package dto

//func TestCreateAccountDTO(t *testing.T) {
//	t.Run("Test CreateAccountDTO Initialization", func(t *testing.T) {
//		initCreateAccountDTO := dto.AccountDTO{
//			Username: "username_24",
//			Password: "password",
//		}
//		assert.NotEmptyf(t, initCreateAccountDTO, "Init Create Account DTO Not Empty")
//
//		if initCreateAccountDTO.Username != "username_24" {
//			t.Errorf("Expected CreateAccountDTO.Username to be \"username_24\", but got \"%s\"", initCreateAccountDTO.Username)
//		}
//
//		if initCreateAccountDTO.Email != "username@example.com" {
//			t.Errorf("Expected CreateAccountDTO.Email to be \"username@example.com\", but got \"%s\"", initCreateAccountDTO.Email)
//		}
//
//		if initCreateAccountDTO.Password != "password" {
//			t.Errorf("Expected CreateAccountDTO.Password to be \"password\", but got \"%s\"", initCreateAccountDTO.Password)
//		}
//	})
//
//	t.Run("Test CreateAccountDTO JSON Encoding", func(t *testing.T) {
//		createAccountDTO := dto.CreateAccountDTO{
//			Username: "johndoe",
//			Email:    "johndoe@example.com",
//			Password: "password",
//		}
//		assert.NotEmptyf(t, createAccountDTO, "Create Account DTO Not Empty")
//
//		expectedJSON := `{"username":"johndoe","email":"johndoe@example.com","password":"password"}`
//		assert.NotEmptyf(t, expectedJSON, "Should be format as JSON String")
//
//		actualJSON, err := json.Marshal(createAccountDTO)
//		if err != nil {
//			t.Errorf("Unexpected errors marshalling CreateAccountDTO: %v", err)
//		}
//		assert.NoError(t, err)
//		assert.NotEmptyf(t, actualJSON, "actual JSON Not Empty")
//
//		if string(actualJSON) != expectedJSON {
//			t.Errorf("Expected JSON encoding to be %s, but got %s", expectedJSON, string(actualJSON))
//		}
//		assert.JSONEqf(t, string(actualJSON), expectedJSON, "JSON expected is macther")
//	})
//
//	t.Run("Test CreateAccountDTO JSON Encoding", func(t *testing.T) {
//		expectedCreateAccountDTO := dto.CreateAccountDTO{
//			Username: "johndoe",
//			Email:    "johndoe@example.com",
//			Password: "password",
//		}
//		assert.NotEmptyf(t, expectedCreateAccountDTO, "expected Create Account DTO Not Empty")
//
//		jsonStr := `{"username":"johndoe","email":"johndoe@example.com","password":"password"}`
//		assert.NotEmptyf(t, jsonStr, "Should be format as JSON String")
//
//		var actualCreateAccountDTO dto.CreateAccountDTO
//		err := json.Unmarshal([]byte(jsonStr), &actualCreateAccountDTO)
//		if err != nil {
//			t.Errorf("Unexpected errors unmarshalling JSON: %v", err)
//		}
//		assert.NoError(t, err)
//		assert.NotEmptyf(t, actualCreateAccountDTO, "actual Create Account DTO Not Empty")
//
//		if actualCreateAccountDTO != expectedCreateAccountDTO {
//			t.Errorf("Expected unmarshalled CreateAccountDTO to be %v, but got %v", expectedCreateAccountDTO, actualCreateAccountDTO)
//		}
//		assert.Equal(t, actualCreateAccountDTO, expectedCreateAccountDTO)
//	})
//}
