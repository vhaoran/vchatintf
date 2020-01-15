package refuser

type GroupInfoRef struct {
	ID int64
	//群名称
	Name string `json:"name,omitempty" gorm:"type:varchar(50);unique_index"`

	//图标
	Icon string `json:"icon omitempty" gorm:"type:varchar(100)"`

	//群主
	OwnerID int64 `json:"owner_id,omitempty" gorm:"index"`
	//群人數
	GMemberCount int64 `json:"g_member_count,omitempty"`
}
