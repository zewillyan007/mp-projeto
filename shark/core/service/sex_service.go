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

type SexService struct {
	Repository port.SexIRepository
	ucGet      *usecase.SexUseCaseGet
	ucGetAll   *usecase.SexUseCaseGetAll
}

func NewSexService(repository port.SexIRepository) *SexService {

	return &SexService{
		Repository: repository,
		ucGet:      usecase.NewSexUseCaseGet(repository),
		ucGetAll:   usecase.NewSexUseCaseGetAll(repository),
	}
}

func (o *SexService) WithTransaction(transaction port_shared.ITransaction) *SexService {

	return &SexService{
		ucGet:    o.ucGet,
		ucGetAll: o.ucGetAll,
	}
}

func (o *SexService) Get(dtoIn *dto.SexDtoIn) (*dto.SexDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Sex, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewSexDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Sex.Id)
	dtoOut.Name = Sex.Name
	dtoOut.Mnemonic = Sex.Mnemonic
	dtoOut.Hint = Sex.Hint

	if Sex.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Sex.CreationDateTime)
	}

	if Sex.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Sex.ChangeDateTime)
	}

	if Sex.DisableDateTime != nil {
		dtoOut.DisableDateTime = DateHelper.Format("2006-01-02 15:04:05", *Sex.DisableDateTime)
	}

	return dtoOut, nil
}

func (o *SexService) GetAll(conditions ...interface{}) []*dto.SexDtoOut {

	var arraySexDto []*dto.SexDtoOut

	arraySex := o.ucGetAll.Execute(conditions...)

	for _, Sex := range arraySex {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewSexDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Sex.Id)
		dtoOut.Name = Sex.Name
		dtoOut.Mnemonic = Sex.Mnemonic
		dtoOut.Hint = Sex.Hint

		if Sex.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Sex.CreationDateTime)
		}

		if Sex.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Sex.ChangeDateTime)
		}

		if Sex.DisableDateTime != nil {
			dtoOut.DisableDateTime = DateHelper.Format("2006-01-02 15:04:05", *Sex.DisableDateTime)
		}

		arraySexDto = append(arraySexDto, dtoOut)
	}
	return arraySexDto
}
