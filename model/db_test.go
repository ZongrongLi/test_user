package model

import (
	"testing"

	"github.com/golang/glog"

	"github.com/tiancai110a/test_user/config"
)

func TestDb(t *testing.T) {
	if err := config.Init("/root/go/src/github.com/tiancai110a/test_user/conf/config.yaml"); err != nil {
		panic(err)
	}

	DB.Init()
	defer DB.Close()

	u := UserModel{
		Username: "xiaoming",
		Password: "123",
	}

	if err := u.Create(); err != nil {
		glog.Info("user exist or other error")
		return
	}

	u1, err := GetUser("xiaoming")
	if err != nil {
		t.Error("get user failed")
		return
	}
	glog.Info("================================ add a user", u1)

	if u1.Username != u.Username || u1.Password != u.Password {
		t.Error("get failed")
		return
	}

	u1.Password = "567"
	u1.Update()

	u11, err := GetUser("xiaoming")
	if err != nil {
		t.Error("get user failed")
		return
	}

	if u11.Password != "567" {
		t.Error("get failed")
		return
	}

	if err := DeleteUser(uint64(u1.Id)); err != nil {
		t.Error("delete failed")
		return
	}

	u2, err := GetUser("xiaoming")
	if err == nil {
		t.Error("delete user failed,  user still exist")
		return
	}
	glog.Info("================================ delete a user", u2)

	info, totalcount, err := ListUser("tiancai", 0, 10)

	if err != nil {
		t.Error("ListUser  failed")
		return
	}

	glog.Info("================================", info, totalcount)

}
