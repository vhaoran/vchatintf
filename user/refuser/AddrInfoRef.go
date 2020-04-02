package refuser

type AddrInfoRef struct {
	//省
	AddrProvince string `json:"addr_province" gorm:"type:varchar(50)"`
	//城市
	AddrCity string `json:"addr_city" gorm:"type:varchar(50)"`
	//县区
	AddrArea   string `json:"addr_area" gorm:"type:varchar(50)"`
	AddrStreet string `json:"addr_street"`
	//
	AddrLinkman string `json:"addr_link_man" gorm:"type:varchar(50)"`
	//
	AddrMobile string `json:"addr_mobile" gorm:"type:varchar(50)"`
}
