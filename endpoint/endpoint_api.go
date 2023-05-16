package endpoint

type Prefix struct {
	Authenticate string
	Account      string
	Task         string
}

type Endpoint struct {
	AuthenticateRegister string
	AuthenticateLogin    string
	AuthenticateLogout   string
	AccountUserFind      string
	AccountUserEdit      string
	TaskCreate           string
	TaskUpdate           string
	TaskUpdateStatus     string
	TaskFindByID         string
	TaskFind             string
	TaskDelete           string
}

func Rest() Endpoint {
	apiRoute := Prefix{
		Authenticate: "/api/v1/authenticate/",
		Account:      "/api/v1/user/",
		Task:         "/api/v1/task/",
	}

	api := Endpoint{
		AuthenticateRegister: apiRoute.Authenticate + "register",
		AuthenticateLogin:    apiRoute.Authenticate + "login",
		AuthenticateLogout:   apiRoute.Authenticate + "logout",
		AccountUserFind:      apiRoute.Account + "find",
		AccountUserEdit:      apiRoute.Account + "edit",
		TaskCreate:           apiRoute.Task + "create",
		TaskUpdate:           apiRoute.Task + "update/:task_id",
		TaskFindByID:         apiRoute.Task + "find/:task_id",
		TaskFind:             apiRoute.Task + "find",
		TaskDelete:           apiRoute.Task + "delete",
		TaskUpdateStatus:     apiRoute.Task + "update_status",
	}

	return api
}
