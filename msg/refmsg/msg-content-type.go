package refmsg

//只知用于私聊消息及群中聊中的类型
const (
	//普通文本消息
	BT_COMMON = iota + 1
	//

	//2	图片
	BT_IMAGE = 2
	//3	文件
	BT_FILE = 3
	//4	离线语音
	BT_AUDIO_OFFLINE = 4
	//5	语音聊天
	BT_AUDIO_TEL = 5
	//6	视频聊天
	BT_VIDEO_TEL = 6
	//7	位置
	BT_POS = 7
	//8	发红包
	BT_MONEY = 8

	//100及以上	自定义
	//

	//10	人个名片
	BT_CARD = 20
	//11	群名片
	BT_CARD_GROUP = 21
	//BT_CARD_GROUP_TITLE = "群名片"

	//红包领取消息
	BT_RED_FETCH = 100
)
