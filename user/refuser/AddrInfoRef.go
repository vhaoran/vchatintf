package refuser

type AddrInfoRef struct {
	//省
	AddrProvince string `json:"addr_province omitempty" gorm:"type:varchar(50)"`
	//城市
	AddrCity string `json:"addr_city omitempty" gorm:"type:varchar(50)"`
	//县区
	AddrArea   string `json:"addr_area omitempty" gorm:"type:varchar(50)"`
	AddrStreet string `json:"addr_street omitempty"`
	//
	AddrLinkman string `json:"addr_link_man omitempty" gorm:"type:varchar(50)"`
	//
	AddrMobile string `json:"addr_mobile omitempty" gorm:"type:varchar(50)"`
}
