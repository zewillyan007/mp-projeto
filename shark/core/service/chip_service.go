package service

import (
	"fmt"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/helper"
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/port"
	"mp-projeto/shark/core/usecase"
	"strconv"
	"time"
)

type ChipService struct {
	Repository port.ChipIRepository
	ucGet      *usecase.ChipUseCaseGet
	ucSave     *usecase.ChipUseCaseSave
	ucGrid     *usecase.ChipUseCaseGrid
	ucGetAll   *usecase.ChipUseCaseGetAll
	ucRemove   *usecase.ChipUseCaseRemove
}

func NewChipService(repository port.ChipIRepository) *ChipService {

	return &ChipService{
		Repository: repository,
		ucGet:      usecase.NewChipUseCaseGet(repository),
		ucSave:     usecase.NewChipUseCaseSave(repository),
		ucGrid:     usecase.NewChipUseCaseGrid(repository),
		ucGetAll:   usecase.NewChipUseCaseGetAll(repository),
		ucRemove:   usecase.NewChipUseCaseRemove(repository),
	}
}

func (o *ChipService) WithTransaction(transaction port_shared.ITransaction) *ChipService {

	return &ChipService{
		ucGet:    o.ucGet,
		ucSave:   o.ucSave.WithTransaction(transaction),
		ucGrid:   o.ucGrid,
		ucGetAll: o.ucGetAll,
		ucRemove: o.ucRemove.WithTransaction(transaction),
	}
}

func (o *ChipService) Get(dtoIn *dto.ChipDtoIn) (*dto.ChipDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Chip, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewChipDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Chip.Id)
	dtoOut.Number = Chip.Number
	dtoOut.Status = Chip.Status

	if Chip.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Chip.CreationDateTime)
	}

	if Chip.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Chip.ChangeDateTime)
	}

	return dtoOut, nil
}

func (o *ChipService) GetAll(conditions ...interface{}) []*dto.ChipDtoOut {

	var arrayChipDto []*dto.ChipDtoOut

	arrayChip := o.ucGetAll.Execute(conditions...)

	for _, Chip := range arrayChip {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewChipDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Chip.Id)
		dtoOut.Number = Chip.Number
		dtoOut.Status = Chip.Status

		if Chip.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Chip.CreationDateTime)
		}

		if Chip.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Chip.ChangeDateTime)
		}

		arrayChipDto = append(arrayChipDto, dtoOut)
	}

	return arrayChipDto
}

func (o *ChipService) Save(dtoIn *dto.ChipDtoIn) error {

	DateHelper := helper.NewDateHelper()
	Chip := FactoryChip()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Chip.Id = int64(id)
	}

	Chip.Number = dtoIn.Number
	Chip.Status = dtoIn.Status

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Chip.Id == 0 {
			Chip.CreationDateTime = &now
		} else {
			ChipCurrent, _ := o.ucGet.Execute(Chip.Id)
			Chip.CreationDateTime = ChipCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Chip.CreationDateTime = CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Chip.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Chip.ChangeDateTime = ChangeDateTime
	}

	_, err := o.ucSave.Execute(Chip)
	if err != nil {
		return err
	}

	return nil
}

func (o *ChipService) Remove(dtoIn *dto.ChipDtoIn) error {

	Chip := FactoryChip()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Chip.Id = int64(id)
	}
	err := o.ucRemove.Execute(Chip)
	if err != nil {
		return err
	}

	return nil
}

func (o *ChipService) Grid(GridConfig *grid.GridConfig) map[string]interface{} {

	return o.ucGrid.Execute(GridConfig)
}
