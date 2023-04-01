package helpers

import (
	"gotodo/internal/domain/dto"
	"gotodo/internal/persistence/record"
)

func TaskRecordToDTO(record record.TaskRecord) dto.TasksDTO {
	toDTO := dto.TasksDTO{
		ID:        record.ID,
		Title:     record.Title,
		Completed: record.Completed,
	}
	return toDTO
}

func TaskDTOToRecord(recordDTO dto.TasksDTO) record.TaskRecord {
	toRecord := record.TaskRecord{
		ID:        recordDTO.ID,
		Title:     recordDTO.Title,
		Completed: recordDTO.Completed,
	}
	return toRecord
}
