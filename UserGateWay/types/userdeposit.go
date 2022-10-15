package types

//[UserDeposit] - Model
type UserDeposit  struct {
	ID              int     `json:"id"`
	Coin            string  `json:"coin"`
	ChainDetail     string  `json:"chain_detail"`
	MemberID        int     `json:"member_id"`
	Amount          float64 `json:"amount"`
	Status          int     `json:"status"`
	RequestTime     int     `json:"request_time"`
	RequestDeadline int     `json:"request_deadline"`
	DepositLabel    int     `json:"deposit_label"`
	UpdatedTime     int     `json:"updated_time"`
}