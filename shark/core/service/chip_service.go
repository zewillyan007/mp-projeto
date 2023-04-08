package service

import (
	"fmt"
	"mp-projeto/shared/grid"
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

	dtoOut := dto.NewChipDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Chip.Id)
	dtoOut.Number = Chip.Number

	if Chip.CreationDateTime != nil {
		dtoOut.CreationDateTime = Chip.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
	}

	if Chip.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = Chip.ChangeDateTime.Format("2006-01-02 15:04:05 -0700")
	}

	return dtoOut, nil
}

func (o *ChipService) GetAll(conditions ...interface{}) []*dto.ChipDtoOut {

	var arrayChipDto []*dto.ChipDtoOut

	arrayChip := o.ucGetAll.Execute(conditions...)

	for _, Chip := range arrayChip {

		dtoOut := dto.NewChipDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Chip.Id)
		dtoOut.Number = Chip.Number

		if Chip.CreationDateTime != nil {
			dtoOut.CreationDateTime = Chip.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
		}

		if Chip.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = Chip.ChangeDateTime.Format("2006-01-02 15:04:05 -0700")
		}

		arrayChipDto = append(arrayChipDto, dtoOut)
	}

	return arrayChipDto
}

func (o *ChipService) Save(dtoIn *dto.ChipDtoIn) error {

	Chip := FactoryChip()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Chip.Id = int64(id)
	}

	Chip.Number = dtoIn.Number
	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Chip.Id == 0 {
			Chip.CreationDateTime = &now
		} else {
			ChipCurrent, _ := o.ucGet.Execute(Chip.Id)
			Chip.CreationDateTime = ChipCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05 -0700", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Chip.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Chip.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05 -0700", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Chip.ChangeDateTime = &ChangeDateTime
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
