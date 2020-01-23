package refmsg

import (
	"github.com/vhaoran/vchat/common/ytime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MsgHisRef struct {
	ID        primitive.ObjectID `json:"id,omitempty"   bson:"_id,omitempty"`
	CreatedAt ytime.Date         `json:"created_at,omitempty"   bson:"created_at,omitempty"`
	Created   time.Time `json:"created,omitempty"   bson:"created,omitempty"`

	//1: 私信消息 群消息
	//2：
	//3：朋友圈消息
	//4：系统消息(如某用户资料变更)
	MsgType MsgType `json:"msg_type,omitempty"   bson:"msg_type,omitempty"`

	//发送方id
	FromUID         int64  `json:"from_uid,omitempty"   bson:"from_uid,omitempty"`
	FromUserCodeRef string `json:"from_user_code,omitempty"   bson:"from_user_code,omitempty"`
	FromNickRef     string `json:"from_nick,omitempty"   bson:"from_nick,omitempty"`
	FromIconRef     string `json:"from_icon,omitempty"   bson:"from_icon,omitempty"`

	//目标用户ID
	ToUID         int64  `json:"to_uid,omitempty"   bson:"to_uid,omitempty"`
	ToUserCodeRef string `json:"to_user_code,omitempty"   bson:"to_user_code,omitempty"`
	ToNickRef     string `json:"to_nick,omitempty"   bson:"to_nick,omitempty"`
	ToIconRef     string `json:"to_icon,omitempty"   bson:"to_icon,omitempty"`

	ToGID      int64  `json:"to_gid,omitempty"   bson:"to_gid,omitempty"`
	ToGNameRef string `json:"to_gname,omitempty"   bson:"to_gname,omitempty"`
	ToGIconRef string `json:"to_gicon,omitempty"   bson:"to_gicon,omitempty"`
	//自定义的消息内容类型，可选
	BodyType int `json:"body_type,omitempty"   bson:"body_type,omitempty"`

	//消息体
	Body interface{} `json:"body,omitempty"   bson:"body,omitempty"`
}
