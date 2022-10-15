package res
//HTTP Response {Response which will be returned to the User}
type Response struct {
	Result  string `json:"result"`
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}