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

type ChipStatusTypeService struct {
	Repository port.ChipStatusTypeIRepository
	ucGet      *usecase.ChipStatusTypeUseCaseGet
	ucGetAll   *usecase.ChipStatusTypeUseCaseGetAll
}

func NewChipStatusTypeService(repository port.ChipStatusTypeIRepository) *ChipStatusTypeService {

	return &ChipStatusTypeService{
		Repository: repository,
		ucGet:      usecase.NewChipStatusTypeUseCaseGet(repository),
		ucGetAll:   usecase.NewChipStatusTypeUseCaseGetAll(repository),
	}
}

func (o *ChipStatusTypeService) WithTransaction(transaction port_shared.ITransaction) *ChipStatusTypeService {

	return &ChipStatusTypeService{
		ucGet:    o.ucGet,
		ucGetAll: o.ucGetAll,
	}
}

func (o *ChipStatusTypeService) Get(dtoIn *dto.ChipStatusTypeDtoIn) (*dto.ChipStatusTypeDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	ChipStatusType, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewChipStatusTypeDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", ChipStatusType.Id)
	dtoOut.Name = ChipStatusType.Name
	dtoOut.Mnemonic = ChipStatusType.Mnemonic
	dtoOut.Hint = ChipStatusType.Hint

	if ChipStatusType.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *ChipStatusType.CreationDateTime)
	}

	if ChipStatusType.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *ChipStatusType.ChangeDateTime)
	}

	if ChipStatusType.DisableDateTime != nil {
		dtoOut.DisableDateTime = DateHelper.Format("2006-01-02 15:04:05", *ChipStatusType.DisableDateTime)
	}

	return dtoOut, nil
}

func (o *ChipStatusTypeService) GetAll(conditions ...interface{}) []*dto.ChipStatusTypeDtoOut {

	var arrayChipStatusTypeDto []*dto.ChipStatusTypeDtoOut

	arrayChipStatusType := o.ucGetAll.Execute(conditions...)

	for _, ChipStatusType := range arrayChipStatusType {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewChipStatusTypeDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", ChipStatusType.Id)
		dtoOut.Name = ChipStatusType.Name
		dtoOut.Mnemonic = ChipStatusType.Mnemonic
		dtoOut.Hint = ChipStatusType.Hint

		if ChipStatusType.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *ChipStatusType.CreationDateTime)
		}

		if ChipStatusType.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *ChipStatusType.ChangeDateTime)
		}

		if ChipStatusType.DisableDateTime != nil {
			dtoOut.DisableDateTime = DateHelper.Format("2006-01-02 15:04:05", *ChipStatusType.DisableDateTime)
		}

		arrayChipStatusTypeDto = append(arrayChipStatusTypeDto, dtoOut)
	}
	return arrayChipStatusTypeDto
}
