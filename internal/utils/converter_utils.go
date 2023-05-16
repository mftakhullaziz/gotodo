package utils

import (
	"gotodo/internal/domain/dto"
	"gotodo/internal/persistence/record"
	"time"
)

func ConvertTaskRecordToTaskDto(record record.TaskRecord) dto.TasksDTO {
	return dto.TasksDTO{
		TaskID:      record.TaskID,
		UserID:      record.UserID,
		Title:       record.Title,
		Description: record.Description,
		Completed:   record.Completed,
		TaskStatus:  record.TaskStatus,
		CompletedAt: record.CompletedAt,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func ConvertTaskDtoToTaskRecord(recordDTO dto.TasksDTO) record.TaskRecord {
	return record.TaskRecord{
		TaskID:      recordDTO.TaskID,
		UserID:      recordDTO.UserID,
		Title:       recordDTO.Title,
		Description: recordDTO.Description,
		Completed:   recordDTO.Completed,
		TaskStatus:  recordDTO.TaskStatus,
		CompletedAt: recordDTO.CompletedAt,
		CreatedAt:   recordDTO.CreatedAt,
		UpdatedAt:   recordDTO.UpdatedAt,
	}
}

func RecordToAccountDTO(accountRecord record.AccountRecord) dto.AccountDTO {
	return dto.AccountDTO{
		AccountID: accountRecord.AccountID,
		UserID:    accountRecord.UserID,
		Username:  accountRecord.Username,
		Password:  accountRecord.Password,
		Status:    accountRecord.Status,
		CreatedAt: accountRecord.CreatedAt,
		UpdatedAt: accountRecord.UpdatedAt,
	}
}

func AccountDtoToRecord(accountDTO dto.AccountDTO) record.AccountRecord {
	return record.AccountRecord{
		AccountID: accountDTO.AccountID,
		UserID:    accountDTO.UserID,
		Username:  accountDTO.Username,
		Password:  accountDTO.Password,
		Status:    accountDTO.Status,
		CreatedAt: accountDTO.CreatedAt,
		UpdatedAt: accountDTO.UpdatedAt,
	}
}

func RecordToUserDTO(detailRecord record.UserDetailRecord) dto.UserDetailDTO {
	return dto.UserDetailDTO{
		UserID:      detailRecord.UserID,
		Username:    detailRecord.Username,
		Password:    detailRecord.Password,
		Email:       detailRecord.Email,
		Name:        detailRecord.Name,
		MobilePhone: detailRecord.MobilePhone,
		Address:     detailRecord.Address,
		Status:      detailRecord.Status,
		CreatedAt:   detailRecord.CreatedAt,
		UpdatedAt:   detailRecord.UpdatedAt,
	}
}

func UserDTOToRecord(detailDTO dto.UserDetailDTO) record.UserDetailRecord {
	return record.UserDetailRecord{
		UserID:      detailDTO.UserID,
		Username:    detailDTO.Username,
		Password:    detailDTO.Password,
		Email:       detailDTO.Email,
		Name:        detailDTO.Name,
		MobilePhone: detailDTO.MobilePhone,
		Address:     detailDTO.Address,
		Status:      detailDTO.Status,
		CreatedAt:   detailDTO.CreatedAt,
		UpdatedAt:   detailDTO.UpdatedAt,
	}
}

type NewOptionalColumnParams struct {
	BearerToken string
	TimeNow     time.Time
	TimeIsNull  time.Time
}

func UserAndAccountRecordToAccountLoginHistoryRecord(detailRecord record.UserDetailRecord,
	accountRecord record.AccountRecord, params NewOptionalColumnParams, expireToken time.Time) record.AccountLoginHistoriesRecord {
	return record.AccountLoginHistoriesRecord{
		AccountID:     int(accountRecord.AccountID),
		UserID:        int(detailRecord.UserID),
		Username:      accountRecord.Username,
		Password:      accountRecord.Password,
		Token:         params.BearerToken,
		TokenExpireAt: expireToken,
		LoginAt:       params.TimeNow,
		LoginOutAt:    params.TimeIsNull,
		CreatedAt:     params.TimeNow,
		UpdatedAt:     params.TimeIsNull,
	}
}

func TaskRecordsToTaskDTOs(tasks []record.TaskRecord) []dto.TasksDTO {
	var tasksDto []dto.TasksDTO

	for _, task := range tasks {
		dtoTask := dto.TasksDTO{
			TaskID:      task.TaskID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			TaskStatus:  task.TaskStatus,
			CompletedAt: task.CompletedAt,
			UpdatedAt:   task.UpdatedAt,
			CreatedAt:   task.CreatedAt,
		}
		tasksDto = append(tasksDto, dtoTask)
	}

	return tasksDto
}

func UserDetailRecordToUserDetailDTO(detailRecord record.UserDetailRecord) dto.UserDetailDTO {
	return dto.UserDetailDTO{
		UserID:      detailRecord.UserID,
		Username:    detailRecord.Username,
		Password:    detailRecord.Password,
		Email:       detailRecord.Email,
		Name:        detailRecord.Name,
		MobilePhone: detailRecord.MobilePhone,
		Address:     detailRecord.Address,
		Status:      detailRecord.Status,
		CreatedAt:   detailRecord.CreatedAt,
		UpdatedAt:   detailRecord.UpdatedAt,
	}
}
