package helpers

import (
	"gotodo/internal/domain/dto"
	"gotodo/internal/persistence/record"
	"time"
)

func TaskRecordToDTO(record record.TaskRecord) dto.TasksDTO {
	return dto.TasksDTO{
		ID:          record.ID,
		Title:       record.Title,
		Description: record.Description,
		Completed:   record.Completed,
		CompletedAt: record.CompletedAt,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func TaskDTOToRecord(recordDTO dto.TasksDTO) record.TaskRecord {
	return record.TaskRecord{
		ID:          recordDTO.ID,
		Title:       recordDTO.Title,
		Description: recordDTO.Description,
		Completed:   recordDTO.Completed,
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

func UserAndAccountRecordToAccountLoginHistoryRecord(
	detailRecord record.UserDetailRecord,
	accountRecord record.AccountRecord,
	loginStatus string,
	timeAt time.Time,
	timeOut time.Time) record.AccountLoginHistoriesRecord {

	return record.AccountLoginHistoriesRecord{
		AccountID:   int(accountRecord.AccountID),
		UserID:      int(detailRecord.UserID),
		Email:       detailRecord.Email,
		Password:    accountRecord.Password,
		LoginStatus: loginStatus,
		LoginAt:     timeAt,
		LoginOutAt:  timeOut,
		CreatedAt:   timeAt,
		UpdatedAt:   time.Time{},
	}
}
