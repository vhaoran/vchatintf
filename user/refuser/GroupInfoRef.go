package refuser

type GroupInfoRef struct {
	ID int64
	//群名称
	Name string `json:"name,omitempty" gorm:"type:varchar(50);index"`

	//图标
	Icon string `json:"icon,omitempty" gorm:"type:varchar(100)"`

	//群主
	OwnerID int64 `json:"owner_id,omitempty" gorm:"index"`
	//群人數
	GMemberCount int64 `json:"g_member_count,omitempty"`
	//全局禁言
	NoSpeak bool
	//允许互加好友
	ExchangeInfo bool
	//允许拉人入群
	Recommend bool
	//入群确认
	JoinConfirm bool
}
