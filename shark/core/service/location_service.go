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

type LocationService struct {
	Repository port.LocationIRepository
	ucGet      *usecase.LocationUseCaseGet
	ucSave     *usecase.LocationUseCaseSave
	ucGrid     *usecase.LocationUseCaseGrid
	ucGetAll   *usecase.LocationUseCaseGetAll
	ucRemove   *usecase.LocationUseCaseRemove
}

func NewLocationService(repository port.LocationIRepository) *LocationService {

	return &LocationService{
		Repository: repository,
		ucGet:      usecase.NewLocationUseCaseGet(repository),
		ucSave:     usecase.NewLocationUseCaseSave(repository),
		ucGrid:     usecase.NewLocationUseCaseGrid(repository),
		ucGetAll:   usecase.NewLocationUseCaseGetAll(repository),
		ucRemove:   usecase.NewLocationUseCaseRemove(repository),
	}
}

func (o *LocationService) WithTransaction(transaction port_shared.ITransaction) *LocationService {

	return &LocationService{
		ucGet:    o.ucGet,
		ucSave:   o.ucSave.WithTransaction(transaction),
		ucGrid:   o.ucGrid,
		ucGetAll: o.ucGetAll,
		ucRemove: o.ucRemove.WithTransaction(transaction),
	}
}

func (o *LocationService) Get(dtoIn *dto.LocationDtoIn) (*dto.LocationDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Location, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewLocationDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Location.Id)
	dtoOut.Name = Location.Name

	if Location.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Location.CreationDateTime)
	}

	if Location.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Location.ChangeDateTime)
	}

	return dtoOut, nil
}

func (o *LocationService) GetAll(conditions ...interface{}) []*dto.LocationDtoOut {

	var arrayLocationDto []*dto.LocationDtoOut

	arrayLocation := o.ucGetAll.Execute(conditions...)

	for _, Location := range arrayLocation {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewLocationDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Location.Id)
		dtoOut.Name = Location.Name

		if Location.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Location.CreationDateTime)
		}

		if Location.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Location.ChangeDateTime)
		}

		arrayLocationDto = append(arrayLocationDto, dtoOut)
	}

	return arrayLocationDto
}

func (o *LocationService) Save(dtoIn *dto.LocationDtoIn) error {

	DateHelper := helper.NewDateHelper()
	Location := FactoryLocation()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Location.Id = int64(id)
	}

	Location.Name = dtoIn.Name

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Location.Id == 0 {
			Location.CreationDateTime = &now
		} else {
			LocationCurrent, _ := o.ucGet.Execute(Location.Id)
			Location.CreationDateTime = LocationCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Location.CreationDateTime = CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Location.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Location.ChangeDateTime = ChangeDateTime
	}

	_, err := o.ucSave.Execute(Location)
	if err != nil {
		return err
	}

	return nil
}

func (o *LocationService) Remove(dtoIn *dto.LocationDtoIn) error {

	Location := FactoryLocation()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Location.Id = int64(id)
	}
	err := o.ucRemove.Execute(Location)
	if err != nil {
		return err
	}

	return nil
}

func (o *LocationService) Grid(GridConfig *grid.GridConfig) map[string]interface{} {

	return o.ucGrid.Execute(GridConfig)
}
