package types
//[UserWithdrawals] - Model
type UserWithdrawals struct {
	ID          int    `json:"id"`
	MemberID    int    `json:"member_id"`
	Coin        string `json:"coin"`
	Status      string `json:"status"`
	Amount      string `json:"amount"`
	Fee         string `json:"fee"`
	Address     string `json:"address"`
	Txid        string `json:"txid"`
	SubmitedAt  int    `json:"submited_at"`
	SuccessedAt int    `json:"successed_at"`
	UpdatedAt   int    `json:"updated_at"`
	ChainType   string `json:"chain_type"`
	HashInfo    string `json:"hash_info"`
}