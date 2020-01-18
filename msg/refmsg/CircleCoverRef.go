package refmsg

type CircleCoverRef struct {
	//用户id
	UID         int64
	UserCodeRef string `json:"user_code,omitempty"   bson:"user_code,omitempty"`
	NickRef     string `json:"nick,omitempty"   bson:"nick,omitempty"`
	IconRef     string `json:"icon,omitempty"   bson:"icon,omitempty"`

	CoverPath string `json:"cover_path,omitempty"   bson:"cover_path,omitempty"`
}
