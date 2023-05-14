package request

type UserRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	MobilePhone int    `json:"mobilePhone"`
	Address     string `json:"address"`
	Status      string `json:"status"`
}
