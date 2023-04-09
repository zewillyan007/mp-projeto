package service

import (
	"fmt"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/helper"
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
	"mp-projeto/shark/core/usecase"
	"strconv"
	"time"
)

type SharkService struct {
	Repository port.SharkIRepository
	ucGet      *usecase.SharkUseCaseGet
	ucSave     *usecase.SharkUseCaseSave
	ucGrid     *usecase.SharkUseCaseGrid
	ucGetAll   *usecase.SharkUseCaseGetAll
	ucRemove   *usecase.SharkUseCaseRemove

	//SERVICE
	scSharkChipService *SharkChipService
}

func NewSharkService(repository port.SharkIRepository, scSharkChipService *SharkChipService) *SharkService {

	return &SharkService{
		Repository:         repository,
		ucGet:              usecase.NewSharkUseCaseGet(repository),
		ucSave:             usecase.NewSharkUseCaseSave(repository),
		ucGrid:             usecase.NewSharkUseCaseGrid(repository),
		ucGetAll:           usecase.NewSharkUseCaseGetAll(repository),
		ucRemove:           usecase.NewSharkUseCaseRemove(repository),
		scSharkChipService: scSharkChipService,
	}
}

func (o *SharkService) WithTransaction(transaction port_shared.ITransaction) *SharkService {

	return &SharkService{
		ucGet:              o.ucGet,
		ucSave:             o.ucSave.WithTransaction(transaction),
		ucGrid:             o.ucGrid,
		ucGetAll:           o.ucGetAll,
		ucRemove:           o.ucRemove.WithTransaction(transaction),
		scSharkChipService: o.scSharkChipService.WithTransaction(transaction),
	}
}

func (o *SharkService) Get(dtoIn *dto.SharkDtoIn) (*dto.SharkAllDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Shark, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewSharkAllDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Shark.Id)
	dtoOut.Species = Shark.Species
	dtoOut.Length = strconv.FormatFloat(Shark.Length, 'f', -1, 64)
	dtoOut.Weight = strconv.FormatFloat(Shark.Weight, 'f', -1, 64)
	dtoOut.Sex = Shark.Sex

	if Shark.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Shark.CreationDateTime)
	}

	if Shark.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Shark.ChangeDateTime)
	}

	dtoOut.SharkChips = o.scSharkChipService.GetAll("id_shark = ?", int64(id))

	return dtoOut, nil
}

func (o *SharkService) GetAll(conditions ...interface{}) []*dto.SharkDtoOut {

	var arraySharkDto []*dto.SharkDtoOut

	arrayShark := o.ucGetAll.Execute(conditions...)

	for _, Shark := range arrayShark {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewSharkDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Shark.Id)
		dtoOut.Species = Shark.Species
		dtoOut.Length = strconv.FormatFloat(Shark.Length, 'f', -1, 64)
		dtoOut.Weight = strconv.FormatFloat(Shark.Weight, 'f', -1, 64)
		dtoOut.Sex = Shark.Sex

		if Shark.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Shark.CreationDateTime)
		}

		if Shark.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Shark.ChangeDateTime)
		}

		arraySharkDto = append(arraySharkDto, dtoOut)
	}

	return arraySharkDto
}

func (o *SharkService) SaveAll(dtoIn *dto.SharkAllDtoIn) (*entity.Shark, error) {

	var err error

	dtoShark := dto.NewSharkDtoIn()
	dtoShark.Id = dtoIn.Id
	dtoShark.Species = dtoIn.Species
	dtoShark.Length = dtoIn.Length
	dtoShark.Weight = dtoIn.Weight
	dtoShark.Sex = dtoIn.Sex
	dtoShark.CreationDateTime = dtoIn.CreationDateTime
	dtoShark.ChangeDateTime = dtoIn.ChangeDateTime

	entityShark, err := o.Save(dtoShark)

	if err != nil {
		return nil, err
	}

	if len(dtoShark.Id) > 0 {
		IdShark, _ := strconv.Atoi(dtoShark.Id)
		err = o.scSharkChipService.RemoveAllByIdShark(int64(IdShark))
		if err != nil {
			return nil, err
		}
	}

	for _, item := range dtoIn.SharkChips {

		dtoSharkChip := dto.NewSharkChipDtoIn()
		dtoSharkChip.Id = item.Id
		dtoSharkChip.IdShark = fmt.Sprintf("%d", entityShark.Id)
		dtoSharkChip.IdChip = item.IdChip
		dtoSharkChip.ChipNumber = item.ChipNumber
		dtoSharkChip.Status = item.Status
		dtoSharkChip.CreationDateTime = item.CreationDateTime

		err = o.scSharkChipService.Save(dtoSharkChip)

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (o *SharkService) Save(dtoIn *dto.SharkDtoIn) (*entity.Shark, error) {

	DateHelper := helper.NewDateHelper()
	Shark := FactoryShark()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Shark.Id = int64(id)
	}

	Shark.Species = dtoIn.Species
	Shark.Length, _ = strconv.ParseFloat(dtoIn.Length, 64)
	Shark.Weight, _ = strconv.ParseFloat(dtoIn.Weight, 64)
	Shark.Sex = dtoIn.Sex

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Shark.Id == 0 {
			Shark.CreationDateTime = &now
		} else {
			SharkCurrent, _ := o.ucGet.Execute(Shark.Id)
			Shark.CreationDateTime = SharkCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return nil, err
		}
		Shark.CreationDateTime = CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Shark.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return nil, err
		}
		Shark.ChangeDateTime = ChangeDateTime
	}

	entityShark, err := o.ucSave.Execute(Shark)
	if err != nil {
		return nil, err
	}

	return entityShark, nil
}

func (o *SharkService) Remove(dtoIn *dto.SharkDtoIn) error {

	Shark := FactoryShark()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Shark.Id = int64(id)
	}
	err := o.ucRemove.Execute(Shark)
	if err != nil {
		return err
	}

	return nil
}

func (o *SharkService) Grid(GridConfig *grid.GridConfig) map[string]interface{} {

	return o.ucGrid.Execute(GridConfig)
}
