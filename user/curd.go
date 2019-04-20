package user

import (
	"context"

	"github.com/lexkong/log"
)

type User struct {
}

func (t User) Create(ctx context.Context, req *CreateRequest, res *CreateResponse) error {

	log.Infof("--------------------------------------------------------------------------create name:%s  passwd:%s", req.Password, req.Username)
	return nil
}

func (t User) Delete(ctx context.Context, req *DeleteRequest, res *DeleteResponse) error {

	log.Infof("------------------------------------------------------------------------- Delete ID: %d", req.UserID)
	return nil
}

func (t User) Get(ctx context.Context, req *GetUserRequest, res *GetUserResponse) error {

	log.Infof("--------------------------------------------------------------------------Get: %d", req.Username)
	return nil
}

func (t User) List(ctx context.Context, req *ListRequest, res *ListResponse) error {

	log.Infof("--------------------------------------------------------------------------List: %d", req.Username)
	return nil
}

func (t User) Update(ctx context.Context, req *UpdateRequest, res *UpdateResponse) error {

	log.Infof("--------------------------------------------------------------------------update: %d", req.Username)
	return nil
}
