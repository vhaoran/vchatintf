package refuser

type BulletinInfoRef struct {
	//
	ID int64
	//公众号名称
	Name string `json:"name,omitempty" gorm:"type:varchar(100);unique_index;"`
	//system type 0: common 1:system
	Attr int

	//图标
	Icon string `json:"icon,omitempty" gorm:"type:varchar(100)"`
	//公众号所有人
	OwnerID int64 `json:"owner_id,omitempty" gorm:"index"`
	//公众号主体
	OwnerCorp string `json:"owner_corp,omitempty"`
	//简介
	Brief string `json:"brief,omitempty"`
}
