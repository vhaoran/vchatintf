package refmsg

//系统消息中的通知类型
const (
	//"添加好友"
	NOTIFY_ADD_FRIEND = 101
	// "添加好友pass"
	NOTIFY_ADD_FRIEND_ACCEPT = 102
	//"入群邀请"
	NOTIFY_INVITE_JOIN_GROUP = 103
	//"加入群"
	NOTIFY_JOIN_GROUP = 104

	//"朋友圈点赞"
	NOTIFY_CIRCLE_PRIZE = 201
	//"朋友圈评论"
	NOTIFY_CIRCLE_COMMENT = 202
	//"朋友圈评论回复"
	NOTIFY_CIRCLE_REPLY = 203

	//-----资金操作-------------------------
	//充值 成功
	NOTIFY_FINANCE_RECHARGE_OK = 601
	//提现成功
	NOTIFY_FINANCE_DRAWMONEY_OK = 602

	//某人领取红包
	NOTIFY_FINANCE_FETCH_REDPACKET = 603
	//

)
