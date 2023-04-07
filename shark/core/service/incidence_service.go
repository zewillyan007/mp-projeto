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

type IncidenceService struct {
	Repository port.IncidenceIRepository
	ucGet      *usecase.IncidenceUseCaseGet
	ucSave     *usecase.IncidenceUseCaseSave
	ucGrid     *usecase.IncidenceUseCaseGrid
	ucGetAll   *usecase.IncidenceUseCaseGetAll
	ucRemove   *usecase.IncidenceUseCaseRemove
}

func NewIncidenceService(repository port.IncidenceIRepository) *IncidenceService {

	return &IncidenceService{
		Repository: repository,
		ucGet:      usecase.NewIncidenceUseCaseGet(repository),
		ucSave:     usecase.NewIncidenceUseCaseSave(repository),
		ucGrid:     usecase.NewIncidenceUseCaseGrid(repository),
		ucGetAll:   usecase.NewIncidenceUseCaseGetAll(repository),
		ucRemove:   usecase.NewIncidenceUseCaseRemove(repository),
	}
}

func (o *IncidenceService) WithTransaction(transaction port_shared.ITransaction) *IncidenceService {

	return &IncidenceService{
		ucGet:    o.ucGet,
		ucSave:   o.ucSave.WithTransaction(transaction),
		ucGrid:   o.ucGrid,
		ucGetAll: o.ucGetAll,
		ucRemove: o.ucRemove.WithTransaction(transaction),
	}
}

func (o *IncidenceService) Get(dtoIn *dto.IncidenceDtoIn) (*dto.IncidenceDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	Incidence, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewIncidenceDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", Incidence.Id)
	dtoOut.IdShark = fmt.Sprintf("%d", Incidence.IdShark)
	dtoOut.Name = Incidence.Name

	if Incidence.IncidenceDateTime != nil {
		dtoOut.IncidenceDateTime = Incidence.IncidenceDateTime.Format("2006-01-02 15:04:05 -0700")
	}

	return dtoOut, nil
}

func (o *IncidenceService) GetAll(conditions ...interface{}) []*dto.IncidenceDtoOut {

	var arrayIncidenceDto []*dto.IncidenceDtoOut

	arrayIncidence := o.ucGetAll.Execute(conditions...)

	for _, Incidence := range arrayIncidence {

		dtoOut := dto.NewIncidenceDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", Incidence.Id)
		dtoOut.IdShark = fmt.Sprintf("%d", Incidence.IdShark)
		dtoOut.Name = Incidence.Name

		if Incidence.IncidenceDateTime != nil {
			dtoOut.IncidenceDateTime = Incidence.IncidenceDateTime.Format("2006-01-02 15:04:05 -0700")
		}

		arrayIncidenceDto = append(arrayIncidenceDto, dtoOut)
	}

	return arrayIncidenceDto
}

func (o *IncidenceService) Save(dtoIn *dto.IncidenceDtoIn) error {

	Incidence := FactoryIncidence()

	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Incidence.Id = int64(id)
	}

	Incidence.IdShark, _ = strconv.ParseInt(dtoIn.IdShark, 10, 64)
	Incidence.Name = dtoIn.Name
	now := time.Now()

	if len(dtoIn.IncidenceDateTime) == 0 {
		Incidence.IncidenceDateTime = &now
	} else {
		IncidenceDateTime, err := time.Parse("2006-01-02 15:04:05 -0700", dtoIn.IncidenceDateTime)
		if err != nil {
			return err
		}
		Incidence.IncidenceDateTime = &IncidenceDateTime
	}

	_, err := o.ucSave.Execute(Incidence)
	if err != nil {
		return err
	}
	return nil
}

func (o *IncidenceService) Remove(dtoIn *dto.IncidenceDtoIn) error {

	Incidence := FactoryIncidence()
	if len(dtoIn.Id) > 0 {
		id, _ := strconv.Atoi(dtoIn.Id)
		Incidence.Id = int64(id)
	}
	err := o.ucRemove.Execute(Incidence)
	if err != nil {
		return err
	}

	return nil
}

func (o *IncidenceService) Grid(GridConfig *grid.GridConfig) map[string]interface{} {

	return o.ucGrid.Execute(GridConfig)
}
