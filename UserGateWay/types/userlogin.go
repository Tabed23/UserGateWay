package types
//[UserLogin] - Model
type UserLogin struct {
	MemberID  int    `json:"member_id"`
	LoginTime int    `json:"login_time"`
	LoginIP   string `json:"loginip"`
	Platform  string `json:"platform"`
}