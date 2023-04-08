package service

import (
	"fmt"
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

	dtoOut := dto.NewSexDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Sex.Id)
	dtoOut.Name = Sex.Name
	dtoOut.Mnemonic = Sex.Mnemonic
	dtoOut.Hint = Sex.Hint

	if Sex.CreationDateTime != nil {
		dtoOut.CreationDateTime = Sex.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
	}

	if Sex.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = Sex.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
	}

	if Sex.DisableDateTime != nil {
		dtoOut.DisableDateTime = Sex.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
	}

	return dtoOut, nil
}

func (o *SexService) GetAll(conditions ...interface{}) []*dto.SexDtoOut {

	var arraySexDto []*dto.SexDtoOut

	arraySex := o.ucGetAll.Execute(conditions...)

	for _, Sex := range arraySex {

		dtoOut := dto.NewSexDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Sex.Id)
		dtoOut.Name = Sex.Name
		dtoOut.Mnemonic = Sex.Mnemonic
		dtoOut.Hint = Sex.Hint

		if Sex.CreationDateTime != nil {
			dtoOut.CreationDateTime = Sex.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
		}

		if Sex.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = Sex.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
		}

		if Sex.DisableDateTime != nil {
			dtoOut.DisableDateTime = Sex.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
		}

		arraySexDto = append(arraySexDto, dtoOut)
	}
	return arraySexDto
}
