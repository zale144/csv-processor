package user

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type userBatch struct {
	userID     uint
	savedToCRM bool
}

func saveUserBatch(ctx context.Context, tx *sql.Tx, usrs []userBatch) error {

	query := "INSERT INTO public.user_batch (user_id, saved_to_crm) VALUES "

	var values []string
	for _, u := range usrs {
		values = append(values, fmt.Sprintf("(%d, %v)", u.userID, u.savedToCRM))
	}

	query += strings.Join(values, ",") + ` ON CONFLICT DO NOTHING;`

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}

func getUserBatch(ctx context.Context, tx *sql.Tx, fID, lID int64) ([]userBatch, error) {

	stmt, err := tx.Prepare("SELECT user_id, saved_to_crm FROM public.user_batch WHERE user_id BETWEEN $1 AND $2")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, fID, lID)
	if err != nil {
		return nil, err
	}

	var usrb []userBatch

	for rows.Next() {

		us := userBatch{}
		if err = rows.Scan(&us.userID, &us.savedToCRM); err != nil {
			return nil, err
		}

		usrb = append(usrb, us)
	}

	return usrb, nil
}

func updateUserBatchSavedToCRM(ctx context.Context, tx *sql.Tx, usrIds []uint) error {

	if len(usrIds) == 0 {
		return nil
	}

	usrIdsStr := make([]string, len(usrIds))
	for i, id := range usrIds {
		usrIdsStr[i] = fmt.Sprint(id)
	}

	query := fmt.Sprintf("UPDATE public.user_batch SET saved_to_crm = true WHERE user_id IN (%s)", strings.Join(usrIdsStr, ","))
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}
