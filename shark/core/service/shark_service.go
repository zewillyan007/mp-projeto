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

type SharkService struct {
	Repository port.SharkIRepository
	ucGet      *usecase.SharkUseCaseGet
	ucSave     *usecase.SharkUseCaseSave
	ucGrid     *usecase.SharkUseCaseGrid
	ucGetAll   *usecase.SharkUseCaseGetAll
	ucRemove   *usecase.SharkUseCaseRemove
}

func NewSharkService(repository port.SharkIRepository) *SharkService {

	return &SharkService{
		Repository: repository,
		ucGet:      usecase.NewSharkUseCaseGet(repository),
		ucSave:     usecase.NewSharkUseCaseSave(repository),
		ucGrid:     usecase.NewSharkUseCaseGrid(repository),
		ucGetAll:   usecase.NewSharkUseCaseGetAll(repository),
		ucRemove:   usecase.NewSharkUseCaseRemove(repository),
	}
}

func (o *SharkService) WithTransaction(transaction port_shared.ITransaction) *SharkService {

	return &SharkService{
		ucGet:    o.ucGet,
		ucSave:   o.ucSave.WithTransaction(transaction),
		ucGrid:   o.ucGrid,
		ucGetAll: o.ucGetAll,
		ucRemove: o.ucRemove.WithTransaction(transaction),
	}
}

func (o *SharkService) Get(dtoIn *dto.SharkDtoIn) (*dto.SharkDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Shark, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewSharkDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Shark.Id)
	dtoOut.Species = Shark.Species
	dtoOut.Length = strconv.FormatFloat(Shark.Length, 'f', -1, 64)
	dtoOut.Weight = strconv.FormatFloat(Shark.Weight, 'f', -1, 64)
	dtoOut.Sex = Shark.Sex

	if Shark.CreationDateTime != nil {
		dtoOut.CreationDateTime = Shark.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
	}

	if Shark.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = Shark.ChangeDateTime.Format("2006-01-02 15:04:05 -0700")
	}

	return dtoOut, nil
}

func (o *SharkService) GetAll(conditions ...interface{}) []*dto.SharkDtoOut {

	var arraySharkDto []*dto.SharkDtoOut

	arrayShark := o.ucGetAll.Execute(conditions...)

	for _, Shark := range arrayShark {

		dtoOut := dto.NewSharkDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Shark.Id)
		dtoOut.Species = Shark.Species
		dtoOut.Length = strconv.FormatFloat(Shark.Length, 'f', -1, 64)
		dtoOut.Weight = strconv.FormatFloat(Shark.Weight, 'f', -1, 64)
		dtoOut.Sex = Shark.Sex

		if Shark.CreationDateTime != nil {
			dtoOut.CreationDateTime = Shark.CreationDateTime.Format("2006-01-02 15:04:05 -0700")
		}

		if Shark.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = Shark.ChangeDateTime.Format("2006-01-02 15:04:05 -0700")
		}

		arraySharkDto = append(arraySharkDto, dtoOut)
	}

	return arraySharkDto
}

func (o *SharkService) Save(dtoIn *dto.SharkDtoIn) error {

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
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05 -0700", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Shark.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Shark.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05 -0700", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Shark.ChangeDateTime = &ChangeDateTime
	}

	_, err := o.ucSave.Execute(Shark)
	if err != nil {
		return err
	}

	return nil
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
