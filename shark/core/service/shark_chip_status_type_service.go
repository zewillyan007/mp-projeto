package service

import (
	"fmt"
	"mp-projeto/shared/helper"
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/port"
	"mp-projeto/shark/core/usecase"
	"strconv"
)

type SharkChipStatusTypeService struct {
	Repository port.SharkChipStatusTypeIRepository
	ucGet      *usecase.SharkChipStatusTypeUseCaseGet
	ucGetAll   *usecase.SharkChipStatusTypeUseCaseGetAll
}

func NewSharkChipStatusTypeService(repository port.SharkChipStatusTypeIRepository) *SharkChipStatusTypeService {

	return &SharkChipStatusTypeService{
		Repository: repository,
		ucGet:      usecase.NewSharkChipStatusTypeUseCaseGet(repository),
		ucGetAll:   usecase.NewSharkChipStatusTypeUseCaseGetAll(repository),
	}
}

func (o *SharkChipStatusTypeService) WithTransaction(transaction port_shared.ITransaction) *SharkChipStatusTypeService {

	return &SharkChipStatusTypeService{
		ucGet:    o.ucGet,
		ucGetAll: o.ucGetAll,
	}
}

func (o *SharkChipStatusTypeService) Get(dtoIn *dto.SharkChipStatusTypeDtoIn) (*dto.SharkChipStatusTypeDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	SharkChipStatusType, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewSharkChipStatusTypeDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", SharkChipStatusType.Id)
	dtoOut.Name = SharkChipStatusType.Name
	dtoOut.Mnemonic = SharkChipStatusType.Mnemonic
	dtoOut.Hint = SharkChipStatusType.Hint

	if SharkChipStatusType.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *SharkChipStatusType.CreationDateTime)
	}

	if SharkChipStatusType.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *SharkChipStatusType.ChangeDateTime)
	}

	if SharkChipStatusType.DisableDateTime != nil {
		dtoOut.DisableDateTime = DateHelper.Format("2006-01-02 15:04:05", *SharkChipStatusType.DisableDateTime)
	}

	return dtoOut, nil
}

func (o *SharkChipStatusTypeService) GetAll(conditions ...interface{}) []*dto.SharkChipStatusTypeDtoOut {

	var arraySharkChipStatusTypeDto []*dto.SharkChipStatusTypeDtoOut

	arraySharkChipStatusType := o.ucGetAll.Execute(conditions...)

	for _, SharkChipStatusType := range arraySharkChipStatusType {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewSharkChipStatusTypeDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", SharkChipStatusType.Id)
		dtoOut.Name = SharkChipStatusType.Name
		dtoOut.Mnemonic = SharkChipStatusType.Mnemonic
		dtoOut.Hint = SharkChipStatusType.Hint

		if SharkChipStatusType.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *SharkChipStatusType.CreationDateTime)
		}

		if SharkChipStatusType.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *SharkChipStatusType.ChangeDateTime)
		}

		if SharkChipStatusType.DisableDateTime != nil {
			dtoOut.DisableDateTime = DateHelper.Format("2006-01-02 15:04:05", *SharkChipStatusType.DisableDateTime)
		}

		arraySharkChipStatusTypeDto = append(arraySharkChipStatusTypeDto, dtoOut)
	}
	return arraySharkChipStatusTypeDto
}
