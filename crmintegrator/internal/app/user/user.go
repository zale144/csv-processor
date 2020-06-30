package user

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"crmintegrator/internal/client"
	"crmintegrator/internal/config"
	"crmintegrator/internal/pb"
	db "crmintegrator/pkg/database"
	"strings"
)

type UserService struct {
	cfg    *config.Config
	db     *db.DB
	client client.CRMClient
}

func NewUserService(cfg *config.Config, db *db.DB, crmClient client.CRMClient) UserService {
	return UserService{
		cfg:    cfg,
		db:     db,
		client: crmClient,
	}
}

// SaveUsers batch inserts a number of users to the database using a single query
// and saves information about their sent_to_crm_status (rollback on error)
func (u UserService) SaveUsers(ctx context.Context, users []*pb.User) error {
	return u.db.Transact(ctx, func(tx *sql.Tx) (er error) {

		query := "INSERT INTO public.user (id, first_name, last_name, email, phone) VALUES "

		// concatenate for each batch's insert query string
		var values []string
		for _, u := range users {
			values = append(values, fmt.Sprintf(" (%d, '%s', '%s', '%s', '%s') ", u.Id, u.FirstName, u.LastName, u.Email, u.Phone))
		}
		query += strings.Join(values, ",") + ` ON CONFLICT (id) DO UPDATE 
				SET (first_name, last_name, email, phone) = (EXCLUDED.first_name, EXCLUDED.last_name, EXCLUDED.email, EXCLUDED.phone);`

		if _, err := tx.ExecContext(ctx, query); err != nil {
			return err
		}

		usrs := make([]userBatch, len(users))
		for i, u := range users {
			usrs[i] = userBatch{userID: uint(u.Id)}
		}

		// save the information about user status
		if er = saveUserBatch(ctx, tx, usrs); er != nil {
			return
		}

		return
	})
}

// SendUsersToCRM sends the users to CRM and saves their status locally (rollback on error)
func (u UserService) SendUsersToCRM(ctx context.Context, users []*pb.User) error {
	return u.db.Transact(ctx, func(tx *sql.Tx) (er error) {

		firstID, lastID := users[0].Id, users[len(users)-1].Id
		var userInfo []userBatch
		// get the info about the users that have been sent to CRM
		userInfo, er = getUserBatch(ctx, tx, firstID, lastID)
		if er != nil && er != sql.ErrNoRows {
			return
		}

		var toUpdate []uint

		for _, u := range userInfo {
			if u.savedToCRM {
				// exclude the users that have been saved to CRM already
				users = removeUserByID(int64(u.userID), users)
			} else {
				// prepare to update the ones that will be sent to CRM
				toUpdate = append(toUpdate, u.userID)
			}
		}

		if len(users) == 0 {
			return fmt.Errorf("all users between id %d and %d already exist on CRM", firstID, lastID)
		}

		// send the users to CRM
		jsnUsr, er := json.Marshal(users)
		if er != nil {
			return er
		}

		if er = u.client.SendUsersWithRetry(ctx, ioutil.NopCloser(bytes.NewBuffer(jsnUsr))); er != nil {
			er = errors.New(fmt.Sprintf("could not save users between '%d' and '%d' to CRM", firstID, lastID))
			return
		}

		// update the statuses for users that were sent to CRM - "set saved_to_crm = true" (rollback on error)
		if er = updateUserBatchSavedToCRM(ctx, tx, toUpdate); er != nil {
			return
		}

		return
	})
}

func removeUserByID(id int64, users []*pb.User) []*pb.User {
	for i, u := range users {
		if u.Id == id {
			users[i] = users[len(users)-1]
			return users[:len(users)-1]
		}
	}
	return nil
}
