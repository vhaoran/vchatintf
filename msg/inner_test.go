package msg

import (
	"log"
	"testing"

	"github.com/vhaoran/vchat/common/ytime"
	"github.com/vhaoran/vchat/lib"

	"github.com/vhaoran/vchatintf/msg/refmsg"
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

func Test_NotifyMsgInner(t *testing.T) {
	in := &NotifyMsgInnerIn{
		MsgHisRef: refmsg.MsgHisRef{
			CreatedAt:       ytime.OfNow(),
			Created:         ytime.OfNow().Time.UnixNano(),
			MsgType:         1,
			FromUID:         1,
			FromUserCodeRef: "",
			FromNickRef:     "",
			FromIconRef:     "",
			ToUID:           2,
			ToUserCodeRef:   "",
			ToNickRef:       "",
			ToIconRef:       "",
			ToGID:           0,
			ToGNameRef:      "",
			ToGIconRef:      "",
			BodyType:        0,
			Body:            "hello,inner send",
		},
	}

	r, err := new(NotifyMsgInnerH).Call(in)
	log.Println("----------", "", "------------")
	log.Println("----------", "r", r, "------------")
	log.Println("----------", "err", err, "------------")

}
