package handler

import (
	"context"
	"errors"
	"log"
	"crmintegrator/internal/app/user"
	"crmintegrator/internal/pb"
)

type User struct {
	userSrvc user.UserService
}

func NewUser(userSrvc user.UserService) User {
	return User{
		userSrvc: userSrvc,
	}
}

func (u User) ProcessUserBatch(ctx context.Context, in *pb.UserBatchReq) (rsp *pb.UserBatchRsp, err error) {

	rsp = &pb.UserBatchRsp{}

	if len(in.Users) == 0 {
		err = errors.New("no users were provided")
		return
	}

	firstID, lastID := in.Users[0].Id, in.Users[len(in.Users)-1].Id

	log.Printf("received users with id from %d to %d\n", firstID, lastID)

	// save the users locally
	if err = u.userSrvc.SaveUsers(ctx, in.Users); err != nil {
		log.Println(err)
		return
	}

	// send the users to CRM
	if err = u.userSrvc.SendUsersToCRM(ctx, in.Users); err != nil {
		log.Println(err)
		return
	}

	rsp.Success = true
	return
}
