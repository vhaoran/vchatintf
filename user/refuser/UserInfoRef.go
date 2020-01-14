package refuser

//用刻信息中常用的三个字段
type UserInfoRef struct {
	UID      int64  `json:"uid,omitempty"`
	UserCode string `json:"user_code omitempty" gorm:"index:user_info_multi_code_mobile;type:varchar(50);not null;unique_index;"`
	//眤称
	Nick string `json:"nick omitempty" gorm:"type:varchar(50)"`
	//头像
	Icon string `json:"icon omitempty" gorm:"type:varchar(100)"`
}
