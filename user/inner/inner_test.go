package inner

import (
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"
	"log"
	"testing"
)

func init() {
	_, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: true,
		LoadEtcd:         true,
		LoadPg:           true,
		LoadRedis:        true,
		LoadMongo:        true,
		LoadMq:           true,
		LoadRabbitMq:     true,
		LoadJwt:          true,
	})
	if err != nil {
		panic(err.Error())
	}
}

func Test_GetUserFriendsInner(t *testing.T) {
	in := &GetUserFriendsInnerIn{
		UID: 1,
	}

	r, err := new(GetUserFriendsInnerH).Call(in)
	log.Println("----------", "err", err, "------------")
	ylog.DebugDump("r", r)
}

func Test_GetUserInfoInnerHandler_Call(t *testing.T) {
	in := GetUserInfoInnerIn{
		UID: 1,
	}

	r, err := new(GetUserInfoInnerH).Call(in)
	log.Println("----------", "err", err, "------------")
	ylog.DebugDump("r", r)
}

func Test_GetGroupMembersInnerHandler_Call(t *testing.T) {
	in := &GetGroupMembersInnerIn{
		GID: 1,
	}

	r, err := new(GetGroupMembersInnerH).Call(in)
	log.Println("----------", "err", err, "------------")
	ylog.DebugDump("r", r)
}

func Test_GetGroupInfoInnerHandler_Call(t *testing.T) {
	in := &GetGroupInfoInnerIn{
		GID: 1,
	}

	r, err := new(GetGroupInfoInnerH).Call(in)
	log.Println("----------", "err", err, "------------")
	ylog.DebugDump("r", r)
}

func Test_GetBulletinSubsInnerHandler_Call(t *testing.T) {
	in := &GetBuSubsInnerIn{
		BID: 1,
	}

	r, err := new(GetBuSubsInnerH).Call(in)
	log.Println("----------", "err", err, "------------")
	ylog.DebugDump("r", r)
}

func Test_GetBulletinInfoInnerHandler_Call(t *testing.T) {
	in := &GetBuInfoInnerIn{
		BID: 1,
	}

	r, err := new(GetBuInfoInnerH).Call(in)
	log.Println("----------", "err", err, "------------")
	ylog.DebugDump("r", r)
}
