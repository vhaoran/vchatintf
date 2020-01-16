package refmsg

import (
	"github.com/vhaoran/vchat/common/ytime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MsgHisRef struct {
	ID        primitive.ObjectID `json:"id,omitempty"   bson:"_id,omitempty"`
	CreatedAt ytime.Date         `json:"created_at,omitempty"   bson:"created_at,omitempty"`
	//0: 私信消息
	//1：群消息
	//2：朋友圈消息
	//3：系统消息(如某用户资料变更)
	MsgType int `json:"msg_type,omitempty"   bson:"msg_type,omitempty"`

	//发送方id
	FromUID int64 `json:"from_uid,omitempty"   bson:"from_uid,omitempty"`

	//目标用户ID
	ToUID         int64  `json:"to_uid,omitempty"   bson:"to_uid,omitempty"`
	ToUserCodeRef string `json:"to_user_code_ref,omitempty"   bson:"to_user_code_ref,omitempty"`
	ToNickRef     string `json:"to_nic_ref,omitempty"   bson:"to_nic_ref,omitempty"`
	ToIconRef     string `json:"to_icon_ref,omitempty"   bson:"to_icon_ref,omitempty"`

	ToGID      int64 `json:"to_gid,omitempty"   bson:"to_gid,omitempty"`
	ToGNameRef string `json:"to_gname_ref,omitempty"   bson:"to_gname_ref,omitempty"`
	ToGIconRef string `json:"to_gicon_ref,omitempty"   bson:"to_gicon_ref,omitempty"`
	//自定义的消息内容类型，可选
	BodyType int `json:"body_type,omitempty"   bson:"body_type,omitempty"`

	//消息体
	Body interface{} `json:"body,omitempty"   bson:"body,omitempty"`
}
