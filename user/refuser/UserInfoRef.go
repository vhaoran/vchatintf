package refuser

//用刻信息中常用的三个字段
type UserInfoRef struct {
	ID int64 `json:"id,omitempty"`
	//帐号	登录依据，建议用手机号
	UserCode string `json:"user_code,omitempty" gorm:"index:user_info_multi_code_mobile;type:varchar(50);not null;unique_index;"`
	//眤称
	Nick string `json:"nick,omitempty" gorm:"type:varchar(50)"`
	//头像
	Icon string `json:"icon,omitempty" gorm:"type:varchar(100)"`
	//手机号 登录依据
	Mobile string `json:"mobile,omitempty" gorm:"index:user_info_multi_code_mobile;type:varchar(50);not null;"`
	//缺省手机登录帐号
	MobileLoginDefault bool `json:"mobile_login_default,omitempty" gorm:"index:user_info_multi_code_mobile;default:false;"`

	//姓名
	UserName string `json:"user_name,omitempty" gorm:"type:varchar(50)"`
	//状态	//	锁定时为false
	Enabled bool `json:"enabled,omitempty"`
	//姓别(0,女1田,2保密)
	Sex int32 `json:"sex,omitempty"`
	//出生年
	BirthYear int32 `json:"birth_year,omitempty"`
	//出生月
	BirthMonth int32 `json:"birth_month,omitempty"`
	//出生日
	BirthDay int32 `json:"birth_day,omitempty"`
	//国家
	Country string `json:"country,omitempty" gorm:"type:varchar(100)"`
	//省
	Province string `json:"province,omitempty" gorm:"type:varchar(50)"`
	//城市
	City string `json:"city,omitempty" gorm:"type:varchar(50)"`
	//县区
	Area string `json:"area,omitempty" gorm:"type:varchar(50)"`
	//身份证号
	IdCard string `json:"id_card,omitempty" gorm:"type:varchar(20)"`
	//商城地址
	AddrInfoRef
}
