package helpers

import (
	"gotodo/internal/domain/dto"
	"gotodo/internal/persistence/record"
)

func TaskRecordToDTO(record record.TaskRecord) dto.TasksDTO {
	toDTO := dto.TasksDTO{
		ID:          record.ID,
		Title:       record.Title,
		Description: record.Description,
		Completed:   record.Completed,
		CompletedAt: record.CompletedAt,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
	return toDTO
}

func TaskDTOToRecord(recordDTO dto.TasksDTO) record.TaskRecord {
	toRecord := record.TaskRecord{
		ID:          recordDTO.ID,
		Title:       recordDTO.Title,
		Description: recordDTO.Description,
		Completed:   recordDTO.Completed,
		CompletedAt: recordDTO.CompletedAt,
		CreatedAt:   recordDTO.CreatedAt,
		UpdatedAt:   recordDTO.UpdatedAt,
	}
	return toRecord
}
