package service

import (
	"fmt"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/helper"
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/dto"
	shark_err "mp-projeto/shark/core/err"
	"mp-projeto/shark/core/port"
	"mp-projeto/shark/core/usecase"
	"strconv"
	"time"
)

type SharkChipService struct {
	Repository        port.SharkChipIRepository
	ucGet             *usecase.SharkChipUseCaseGet
	ucSave            *usecase.SharkChipUseCaseSave
	ucGrid            *usecase.SharkChipUseCaseGrid
	ucGetAll          *usecase.SharkChipUseCaseGetAll
	ucRemove          *usecase.SharkChipUseCaseRemove
	ucRemoveByIdShark *usecase.SharkChipUseCaseRemoveByIdShark

	//SERVICES
	scChip *ChipService
}

func NewSharkChipService(repository port.SharkChipIRepository, scChip *ChipService) *SharkChipService {

	return &SharkChipService{
		Repository:        repository,
		ucGet:             usecase.NewSharkChipUseCaseGet(repository),
		ucSave:            usecase.NewSharkChipUseCaseSave(repository),
		ucGrid:            usecase.NewSharkChipUseCaseGrid(repository),
		ucGetAll:          usecase.NewSharkChipUseCaseGetAll(repository),
		ucRemove:          usecase.NewSharkChipUseCaseRemove(repository),
		ucRemoveByIdShark: usecase.NewSharkChipUseCaseRemoveByIdShark(repository),
		scChip:            scChip,
	}
}

func (o *SharkChipService) WithTransaction(transaction port_shared.ITransaction) *SharkChipService {

	return &SharkChipService{
		ucGet:             o.ucGet,
		ucSave:            o.ucSave.WithTransaction(transaction),
		ucGrid:            o.ucGrid,
		ucGetAll:          o.ucGetAll,
		ucRemove:          o.ucRemove.WithTransaction(transaction),
		ucRemoveByIdShark: o.ucRemoveByIdShark.WithTransaction(transaction),
		scChip:            o.scChip.WithTransaction(transaction),
	}
}

func (o *SharkChipService) Get(dtoIn *dto.SharkChipDtoIn) (*dto.SharkChipDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	SharkChip, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewSharkChipDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", SharkChip.Id)
	dtoOut.IdShark = fmt.Sprintf("%d", SharkChip.IdShark)
	dtoOut.IdChip = fmt.Sprintf("%d", SharkChip.IdChip)
	dtoOut.ChipNumber = SharkChip.ChipNumber
	dtoOut.Status = SharkChip.Status

	if SharkChip.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *SharkChip.CreationDateTime)
	}

	return dtoOut, nil
}

func (o *SharkChipService) GetAll(conditions ...interface{}) []*dto.SharkChipDtoOut {

	var arraySharkChipDto []*dto.SharkChipDtoOut

	arraySharkChip := o.ucGetAll.Execute(conditions...)

	for _, SharkChip := range arraySharkChip {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewSharkChipDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", SharkChip.Id)
		dtoOut.IdShark = fmt.Sprintf("%d", SharkChip.IdShark)
		dtoOut.IdChip = fmt.Sprintf("%d", SharkChip.IdChip)
		dtoOut.ChipNumber = SharkChip.ChipNumber
		dtoOut.Status = SharkChip.Status

		if SharkChip.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *SharkChip.CreationDateTime)
		}

		arraySharkChipDto = append(arraySharkChipDto, dtoOut)
	}

	return arraySharkChipDto
}

func (o *SharkChipService) Save(dtoIn *dto.SharkChipDtoIn) error {

	DateHelper := helper.NewDateHelper()
	arrayChips := o.scChip.GetAll("id = ?", dtoIn.IdChip)

	for _, chip := range arrayChips {
		if chip.Status != "NEW" && len(dtoIn.Id) == 0 {
			return shark_err.SharkChipErrorNewLinked
		} else {
			chipDtoIn := dto.NewChipDtoIn()
			chipDtoIn.Id = chip.Id
			chipDtoIn.Number = chip.Number
			chipDtoIn.Status = dtoIn.Status

			err := o.scChip.Save(chipDtoIn)

			if err != nil {
				return err
			}
		}
	}

	SharkChip := FactorySharkChip()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		SharkChip.Id = int64(id)
	}

	SharkChip.IdShark, _ = strconv.ParseInt(dtoIn.IdShark, 10, 64)
	SharkChip.IdChip, _ = strconv.ParseInt(dtoIn.IdChip, 10, 64)
	SharkChip.ChipNumber = dtoIn.ChipNumber
	SharkChip.Status = dtoIn.Status
	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if SharkChip.Id == 0 {
			SharkChip.CreationDateTime = &now
		} else {
			SharkChipCurrent, _ := o.ucGet.Execute(SharkChip.Id)
			SharkChip.CreationDateTime = SharkChipCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		SharkChip.CreationDateTime = CreationDateTime
	}

	_, err := o.ucSave.Execute(SharkChip)
	if err != nil {
		return err
	}

	return nil
}

func (o *SharkChipService) Remove(dtoIn *dto.SharkChipDtoIn) error {

	SharkChip := FactorySharkChip()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		SharkChip.Id = int64(id)
	}
	err := o.ucRemove.Execute(SharkChip)
	if err != nil {
		return err
	}

	return nil
}

func (o *SharkChipService) RemoveAllByIdShark(IdShark int64) error {

	return o.ucRemoveByIdShark.Execute(IdShark)
}

func (o *SharkChipService) Grid(GridConfig *grid.GridConfig) map[string]interface{} {

	return o.ucGrid.Execute(GridConfig)
}
