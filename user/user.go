package user

import (
	"context"

	"github.com/lexkong/log"
	"github.com/tiancai110a/test_user/dl"
	"github.com/tiancai110a/test_user/model"
	"github.com/tiancai110a/test_user/pkg/errno"
)

type User struct {
}

func (t User) Create(ctx context.Context, req *dl.CreateRequest, res *dl.CreateResponse) error {

	log.Infof("--------------------------------------------------------------------------create name:%s  passwd:%s", req.Password, req.Username)
	u := model.UserModel{
		Username: req.Username,
		Password: req.Password,
	}

	if err := u.Create(); err != nil {
		log.Error("create user failed", err)
		res.Err = errno.ErrSQL
		return nil
	}
	res.Username = u.Username

	res.Err = errno.OK

	return nil
}

func (t User) Delete(ctx context.Context, req *dl.DeleteRequest, res *dl.DeleteResponse) error {

	log.Infof("------------------------------------------------------------------------- Delete ID: %d", req.UserID)

	if err := model.DeleteUser(int64(req.UserID)); err != nil {
		log.Error("delete failed", err)
		res.Err = errno.ErrSQL
		return nil
	}

	res.Err = errno.OK
	return nil
}

func (t User) Get(ctx context.Context, req *dl.GetUserRequest, res *dl.GetUserResponse) error {

	log.Infof("--------------------------------------------------------------------------Get: %+v", req)
	u, err := model.GetUser(req.Username)
	if err != nil {
		log.Error("get user failed", err)
		res.Err = errno.ErrSQL
		return nil
	}

	res.Username = u.Username
	res.Id = u.Id
	res.Password = u.Password

	res.Err = errno.OK
	return nil
}

func (t User) List(ctx context.Context, req *dl.ListRequest, res *dl.ListResponse) error {

	log.Infof("--------------------------------------------------------------------------List: %d", req.Username)

	infolist, totalcount, err := model.ListUser(req.Username, req.Offset, req.Limit)

	if err != nil {
		log.Error("ListUser  failed", err)
		res.Err = errno.ErrSQL
		return nil
	}
	res.TotalCount = totalcount
	res.UserList = make([]*dl.UserInfo, 0)

	for _, k := range infolist {

		info := dl.UserInfo{
			Id:       k.Id,
			Username: k.Username,
			Password: k.Password,
		}
		res.UserList = append(res.UserList, &info)
	}

	res.Err = errno.OK
	return nil
}

func (t User) Update(ctx context.Context, req *dl.UpdateRequest, res *dl.UpdateResponse) error {
	log.Infof("--------------------------------------------------------------------------update: %d", req.Username)

	u := model.UserModel{
		Username: req.Username,
		Password: req.Password,
	}
	u.Id = req.Id

	if err := u.Update(); err != nil {
		res.Err = errno.ErrSQL
		return nil
	}

	res.Err = errno.OK
	return nil
}
