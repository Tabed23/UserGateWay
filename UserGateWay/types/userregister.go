package types
//[UserRegistration] - Model
type UserRegistration struct {
	MemberID         int    `json:"member_id"`
	Mobile           string `json:"mobile"`
	MobilePrefix     string `json:"mobile_prefix"`
	Email            string `json:"email"`
	Status           string `json:"status"`
	RegisterIP       string `json:"register_ip"`
	AccountStatus    string `json:"account_status"`
	CreatedAt        int    `json:"created_at"`
	RegisterPlatform string `json:"register_platform"`
}