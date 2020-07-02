package refmsg

//只知用于私聊消息及群中聊中的类型
const (
	//普通文本消息
	BT_COMMON = iota + 1
	//

	//2	图片
	BT_IMAGE
	//3	文件
	BT_FILE
	//4	离线语音
	BT_AUDIO_OFFLINE
	//5	语音聊天
	BT_AUDIO_TEL
	//6	视频聊天
	BT_VIDEO_TEL
	//7	位置
	BT_POS
	//8	发红包
	BT_MONEY
	//100及以上	自定义
	//

	//10	人个名片
	BT_CARD = iota + 10
	//11	群名片
	BT_CARD_GROUP
	//BT_CARD_GROUP_TITLE = "群名片"

)
