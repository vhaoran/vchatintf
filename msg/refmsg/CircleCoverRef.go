package refmsg

type CircleCoverRef struct {
	//用户id
	UID         int64
	UserCodeRef string `json:"user_code"   bson:"user_code"`
	NickRef     string `json:"nick"   bson:"nick"`
	IconRef     string `json:"icon"   bson:"icon"`

	CoverPath string `json:"cover_path"   bson:"cover_path"`
}
