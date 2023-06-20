package api

const (
	// define base api in local var
	baseVersion  = "/api/v1"
	authenticate = baseVersion + "/authenticate"
	user         = baseVersion + "/user"
	task         = baseVersion + "/task"

	// define url api in global to use in router
	Register         = authenticate + "/register"
	Login            = authenticate + "/login"
	Logout           = authenticate + "/logout"
	UserFind         = user + "/find"
	UserEdit         = user + "/edit"
	TaskCreate       = task + "/create"
	TaskUpdate       = task + "/update/:task_id"
	TaskUpdateStatus = task + "/update_status"
	TaskFind         = task + "/find"
	TaskFindById     = task + "/find/:task_id"
	TaskDelete       = task + "/delete"
)
