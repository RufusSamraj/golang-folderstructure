package mapper

import (
	"main.go/dto"
	"main.go/entity"
)

func StudentEntityToDto(reqEntity entity.StudentEntity) dto.Student {
	return dto.Student{
		RollNo: reqEntity.RollNo,
		Name:   reqEntity.Name,
	}
}

func StudentDtoToEntity(reqDto dto.Student) entity.StudentEntity {
	return entity.StudentEntity{
		RollNo: reqDto.RollNo,
		Name:   reqDto.Name,
	}
}
