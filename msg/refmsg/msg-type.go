package refmsg

type MsgType int

const (
	MsgType_Common   MsgType = 1
	MsgType_bulletin MsgType = 2
	MsgType_circle   MsgType = 3
	MsgType_sys      MsgType = 4
)

func GetMsgTypeTitle(i MsgType) string {
	l := []string{"聊天消息", "公众号消息", "朋友圈消息", "系统消息"}
	//
	if i >= 0 && i < 4 {
		return l[i-1]
	}
	return "类型错误"
}
