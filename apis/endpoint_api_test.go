package apis

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPrefix(t *testing.T) {
	// case 1: test prefix sub router
	t.Run("Test Prefix Sub Router", func(t *testing.T) {
		// initiate struct
		prefix := Prefix{
			Authenticate: "/api/v1/authenticate",
			Account:      "/api/v1/account",
			Task:         "/api/v1/task",
		}
		// case value of struct not nil
		assert.NotNil(t, prefix)
		// case value of struct is match of url
		assert.Equal(t, "/api/v1/authenticate", prefix.Authenticate)
		assert.Equal(t, "/api/v1/account", prefix.Account)
		assert.Equal(t, "/api/v1/task", prefix.Task)
		// case value of struct sub route having minimum 1 character
		assert.NotZero(t, len(prefix.Authenticate))
		assert.NotZero(t, len(prefix.Account))
		assert.NotZero(t, len(prefix.Task))
	})
	// case 2: test prefix sub router data type
	t.Run("Test Prefix Sub Router Data Type", func(t *testing.T) {
		// initiate variable struct prefix
		var prefix Prefix
		// read data type from struct
		authenticateStr := reflect.TypeOf(prefix.Authenticate).Kind().String()
		accountStr := reflect.TypeOf(prefix.Account).Kind().String()
		taskStr := reflect.TypeOf(prefix.Task).Kind().String()
		// case data type is match
		assert.Equal(t, "string", authenticateStr)
		assert.Equal(t, "string", accountStr)
		assert.Equal(t, "string", taskStr)
	})
}

func TestEndpoint(t *testing.T) {
	// case 1: test endpoint url
	t.Run("Test Endpoint URL", func(t *testing.T) {
		// initiate struct
		endpointURL := Endpoint{
			AuthenticateRegister: "register",
			AuthenticateLogin:    "login",
			AuthenticateLogout:   "logout",
			AccountUserFind:      "find",
			AccountUserEdit:      "edit",
			TaskCreate:           "create",
			TaskUpdate:           "update/:task_id",
			TaskUpdateStatus:     "update_status",
			TaskFindByID:         "find/:task_id",
			TaskFind:             "find",
			TaskDelete:           "delete",
		}
		// case value struct not nil
		assert.NotNil(t, endpointURL)
		// case value struct is match
		assert.Equal(t, "register", endpointURL.AuthenticateRegister)
		assert.Equal(t, "login", endpointURL.AuthenticateLogin)
		assert.Equal(t, "logout", endpointURL.AuthenticateLogout)
		assert.Equal(t, "find", endpointURL.AccountUserFind)
		assert.Equal(t, "edit", endpointURL.AccountUserEdit)
		assert.Equal(t, "create", endpointURL.TaskCreate)
		assert.Equal(t, "update/:task_id", endpointURL.TaskUpdate)
		assert.Equal(t, "find/:task_id", endpointURL.TaskFindByID)
		assert.Equal(t, "find", endpointURL.TaskFind)
		assert.Equal(t, "delete", endpointURL.TaskDelete)
		// case value struct of string have is value of character
		assert.NotZero(t, len(endpointURL.AuthenticateRegister))
		assert.NotZero(t, len(endpointURL.AuthenticateLogin))
		assert.NotZero(t, len(endpointURL.AuthenticateLogout))
		assert.NotZero(t, len(endpointURL.AccountUserFind))
		assert.NotZero(t, len(endpointURL.AccountUserEdit))
		assert.NotZero(t, len(endpointURL.TaskCreate))
		assert.NotZero(t, len(endpointURL.TaskUpdate))
		assert.NotZero(t, len(endpointURL.TaskUpdateStatus))
		assert.NotZero(t, len(endpointURL.TaskFind))
		assert.NotZero(t, len(endpointURL.TaskFindByID))
		assert.NotZero(t, len(endpointURL.TaskDelete))
	})
	// case 2: test endpoint url data type
	t.Run("Test Endpoint URL Data Type", func(t *testing.T) {
		// initiate variable struct
		var endpoint Endpoint
		// read data type from struct
		registerStr := reflect.TypeOf(endpoint.AuthenticateRegister).Kind().String()
		loginStr := reflect.TypeOf(endpoint.AuthenticateLogin).Kind().String()
		logoutStr := reflect.TypeOf(endpoint.AuthenticateLogout).Kind().String()
		findUserStr := reflect.TypeOf(endpoint.AccountUserFind).Kind().String()
		editUserStr := reflect.TypeOf(endpoint.AccountUserEdit).Kind().String()
		createTaskStr := reflect.TypeOf(endpoint.TaskCreate).Kind().String()
		updateTaskStr := reflect.TypeOf(endpoint.TaskUpdate).Kind().String()
		updateTaskStatusStr := reflect.TypeOf(endpoint.TaskUpdateStatus).Kind().String()
		findTaskStr := reflect.TypeOf(endpoint.TaskFind).Kind().String()
		findByIdTaskStr := reflect.TypeOf(endpoint.TaskFindByID).Kind().String()
		deleteTaskStr := reflect.TypeOf(endpoint.TaskDelete).Kind().String()
		// case data type is match
		assert.Equal(t, "string", registerStr)
		assert.Equal(t, "string", loginStr)
		assert.Equal(t, "string", logoutStr)
		assert.Equal(t, "string", findUserStr)
		assert.Equal(t, "string", editUserStr)
		assert.Equal(t, "string", createTaskStr)
		assert.Equal(t, "string", updateTaskStr)
		assert.Equal(t, "string", updateTaskStatusStr)
		assert.Equal(t, "string", findTaskStr)
		assert.Equal(t, "string", findByIdTaskStr)
		assert.Equal(t, "string", deleteTaskStr)
	})
}

func TestRest(t *testing.T) {
	// case 1: test api url is matching
	t.Run("Test API URL", func(t *testing.T) {
		// initiate endpoint url
		expected := Endpoint{
			AuthenticateRegister: "/api/v1/authenticate/register",
			AuthenticateLogin:    "/api/v1/authenticate/login",
			AuthenticateLogout:   "/api/v1/authenticate/logout",
			AccountUserFind:      "/api/v1/user/find",
			AccountUserEdit:      "/api/v1/user/edit",
			TaskCreate:           "/api/v1/task/create",
			TaskUpdate:           "/api/v1/task/update/:task_id",
			TaskFindByID:         "/api/v1/task/find/:task_id",
			TaskFind:             "/api/v1/task/find",
			TaskDelete:           "/api/v1/task/delete",
			TaskUpdateStatus:     "/api/v1/task/update_status",
		}
		// mock rest func
		result := Rest()
		// case matching api url
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
