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

type MicrocontrollerService struct {
	Repository port.MicrocontrollerIRepository
	ucGet      *usecase.MicrocontrollerUseCaseGet
	ucSave     *usecase.MicrocontrollerUseCaseSave
	ucGrid     *usecase.MicrocontrollerUseCaseGrid
	ucGetAll   *usecase.MicrocontrollerUseCaseGetAll
	ucRemove   *usecase.MicrocontrollerUseCaseRemove
}

func NewMicrocontrollerService(repository port.MicrocontrollerIRepository) *MicrocontrollerService {

	return &MicrocontrollerService{
		Repository: repository,
		ucGet:      usecase.NewMicrocontrollerUseCaseGet(repository),
		ucSave:     usecase.NewMicrocontrollerUseCaseSave(repository),
		ucGrid:     usecase.NewMicrocontrollerUseCaseGrid(repository),
		ucGetAll:   usecase.NewMicrocontrollerUseCaseGetAll(repository),
		ucRemove:   usecase.NewMicrocontrollerUseCaseRemove(repository),
	}
}

func (o *MicrocontrollerService) WithTransaction(transaction port_shared.ITransaction) *MicrocontrollerService {

	return &MicrocontrollerService{
		ucGet:    o.ucGet,
		ucSave:   o.ucSave.WithTransaction(transaction),
		ucGrid:   o.ucGrid,
		ucGetAll: o.ucGetAll,
		ucRemove: o.ucRemove.WithTransaction(transaction),
	}
}

func (o *MicrocontrollerService) Get(dtoIn *dto.MicrocontrollerDtoIn) (*dto.MicrocontrollerDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Microcontroller, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	DateHelper := helper.NewDateHelper()
	dtoOut := dto.NewMicrocontrollerDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Microcontroller.Id)
	dtoOut.IdLocation = fmt.Sprintf("%d", Microcontroller.IdLocation)
	dtoOut.SerialNumber = Microcontroller.SerialNumber
	dtoOut.Model = Microcontroller.Model
	dtoOut.Status = Microcontroller.Status

	if Microcontroller.CreationDateTime != nil {
		dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Microcontroller.CreationDateTime)
	}

	if Microcontroller.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Microcontroller.ChangeDateTime)
	}

	return dtoOut, nil
}

func (o *MicrocontrollerService) GetAll(conditions ...interface{}) []*dto.MicrocontrollerDtoOut {

	var arrayMicrocontrollerDto []*dto.MicrocontrollerDtoOut

	arrayMicrocontroller := o.ucGetAll.Execute(conditions...)

	for _, Microcontroller := range arrayMicrocontroller {

		DateHelper := helper.NewDateHelper()
		dtoOut := dto.NewMicrocontrollerDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Microcontroller.Id)
		dtoOut.IdLocation = fmt.Sprintf("%d", Microcontroller.IdLocation)
		dtoOut.SerialNumber = Microcontroller.SerialNumber
		dtoOut.Model = Microcontroller.Model
		dtoOut.Status = Microcontroller.Status

		if Microcontroller.CreationDateTime != nil {
			dtoOut.CreationDateTime = DateHelper.Format("2006-01-02 15:04:05", *Microcontroller.CreationDateTime)
		}

		if Microcontroller.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = DateHelper.Format("2006-01-02 15:04:05", *Microcontroller.ChangeDateTime)
		}

		arrayMicrocontrollerDto = append(arrayMicrocontrollerDto, dtoOut)
	}

	return arrayMicrocontrollerDto
}

func (o *MicrocontrollerService) Save(dtoIn *dto.MicrocontrollerDtoIn) error {

	DateHelper := helper.NewDateHelper()
	Microcontroller := FactoryMicrocontroller()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Microcontroller.Id = int64(id)
	}
	Microcontroller.IdLocation, _ = strconv.ParseInt(dtoIn.IdLocation, 10, 64)
	Microcontroller.SerialNumber = dtoIn.SerialNumber
	Microcontroller.Model = dtoIn.Model
	Microcontroller.Status = dtoIn.Status

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Microcontroller.Id == 0 {
			Microcontroller.CreationDateTime = &now
		} else {
			MicrocontrollerCurrent, _ := o.ucGet.Execute(Microcontroller.Id)
			Microcontroller.CreationDateTime = MicrocontrollerCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Microcontroller.CreationDateTime = CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Microcontroller.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := DateHelper.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Microcontroller.ChangeDateTime = ChangeDateTime
	}

	_, err := o.ucSave.Execute(Microcontroller)
	if err != nil {
		return err
	}

	return nil
}

func (o *MicrocontrollerService) Remove(dtoIn *dto.MicrocontrollerDtoIn) error {

	Microcontroller := FactoryMicrocontroller()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Microcontroller.Id = int64(id)
	}
	err := o.ucRemove.Execute(Microcontroller)
	if err != nil {
		return err
	}

	return nil
}

func (o *MicrocontrollerService) Grid(GridConfig *grid.GridConfig) map[string]interface{} {

	return o.ucGrid.Execute(GridConfig)
}
