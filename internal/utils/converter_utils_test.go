package utils

import (
	"gotodo/internal/domain/dto"
	"gotodo/internal/persistence/record"
	"reflect"
	"testing"
	"time"
)

func TestAccountDtoToRecord(t *testing.T) {
	type args struct {
		accountDTO dto.AccountDTO
	}
	tests := []struct {
		name string
		args args
		want record.AccountRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AccountDtoToRecord(tt.args.accountDTO); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountDtoToRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertTaskDtoToTaskRecord(t *testing.T) {
	type args struct {
		recordDTO dto.TasksDTO
	}
	tests := []struct {
		name string
		args args
		want record.TaskRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertTaskDtoToTaskRecord(tt.args.recordDTO); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertTaskDtoToTaskRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertTaskRecordToTaskDto(t *testing.T) {
	type args struct {
		record record.TaskRecord
	}
	tests := []struct {
		name string
		args args
		want dto.TasksDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertTaskRecordToTaskDto(tt.args.record); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertTaskRecordToTaskDto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecordToAccountDTO(t *testing.T) {
	type args struct {
		accountRecord record.AccountRecord
	}
	tests := []struct {
		name string
		args args
		want dto.AccountDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecordToAccountDTO(tt.args.accountRecord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecordToAccountDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecordToUserDTO(t *testing.T) {
	type args struct {
		detailRecord record.UserDetailRecord
	}
	tests := []struct {
		name string
		args args
		want dto.UserDetailDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecordToUserDTO(tt.args.detailRecord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecordToUserDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRecordsToTaskDTOs(t *testing.T) {
	type args struct {
		tasks []record.TaskRecord
	}
	tests := []struct {
		name string
		args args
		want []dto.TasksDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TaskRecordsToTaskDTOs(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskRecordsToTaskDTOs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAndAccountRecordToAccountLoginHistoryRecord(t *testing.T) {
	type args struct {
		detailRecord  record.UserDetailRecord
		accountRecord record.AccountRecord
		params        NewOptionalColumnParams
		expireToken   time.Time
	}
	tests := []struct {
		name string
		args args
		want record.AccountLoginHistoriesRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserAndAccountRecordToAccountLoginHistoryRecord(tt.args.detailRecord, tt.args.accountRecord, tt.args.params, tt.args.expireToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserAndAccountRecordToAccountLoginHistoryRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDTOToRecord(t *testing.T) {
	type args struct {
		detailDTO dto.UserDetailDTO
	}
	tests := []struct {
		name string
		args args
		want record.UserDetailRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserDTOToRecord(tt.args.detailDTO); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserDTOToRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDetailRecordToUserDetailDTO(t *testing.T) {
	type args struct {
		detailRecord record.UserDetailRecord
	}
	tests := []struct {
		name string
		args args
		want dto.UserDetailDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserDetailRecordToUserDetailDTO(tt.args.detailRecord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserDetailRecordToUserDetailDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
