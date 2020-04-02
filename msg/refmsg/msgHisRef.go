package refmsg

import (
	"github.com/vhaoran/vchat/common/ytime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MsgHisRef struct {
	ID        primitive.ObjectID `json:"id"   bson:"_id"`
	CreatedAt ytime.Date         `json:"created_at"   bson:"created_at"`
	Created   time.Time `json:"created"   bson:"created"`

	//1: 私信消息 群消息
	//2：
	//3：朋友圈消息
	//4：系统消息(如某用户资料变更)
	MsgType MsgType `json:"msg_type"   bson:"msg_type"`

	//发送方id
	FromUID         int64  `json:"from_uid"   bson:"from_uid"`
	FromUserCodeRef string `json:"from_user_code"   bson:"from_user_code"`
	FromNickRef     string `json:"from_nick"   bson:"from_nick"`
	FromIconRef     string `json:"from_icon"   bson:"from_icon"`

	//目标用户ID
	ToUID         int64  `json:"to_uid"   bson:"to_uid"`
	ToUserCodeRef string `json:"to_user_code"   bson:"to_user_code"`
	ToNickRef     string `json:"to_nick"   bson:"to_nick"`
	ToIconRef     string `json:"to_icon"   bson:"to_icon"`

	ToGID      int64  `json:"to_gid"   bson:"to_gid"`
	ToGNameRef string `json:"to_gname"   bson:"to_gname"`
	ToGIconRef string `json:"to_gicon"   bson:"to_gicon"`
	//自定义的消息内容类型，可选
	BodyType int `json:"body_type"   bson:"body_type"`

	//消息体
	Body interface{} `json:"body"   bson:"body"`
}
