package refuser

type GroupInfoRef struct {
	ID int64
	//群名称
	Name string `json:"name" gorm:"type:varchar(50);index"`

	//图标
	Icon string `json:"icon" gorm:"type:varchar(100)"`

	//群主
	OwnerID int64 `json:"owner_id" gorm:"index"`
	//群人數
	GMemberCount int64 `json:"g_member_count"`
	//全局禁言
	NoSpeak bool `json:"no_speak,omitempty"`
	//允许互加好友
	ExchangeInfo bool `json:"exchange_info"`
	//允许拉人入群
	Recommend bool `json:"recommend"`
	//入群确认
	JoinConfirm bool `json:"join_confirm"`
}
