package types
//[UserLogout] - Model
type UserLogOut struct {
	MemberID   int    `json:"member_id"`
	LogoutTime int    `json:"logout_time"`
	LogoutIP   string `json:"logoutip"`
	Platform   string `json:"platform"`
}