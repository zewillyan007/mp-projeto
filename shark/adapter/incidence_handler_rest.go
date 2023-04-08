package adapter

import (
	adapter_shared "mp-projeto/shared/adapter"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type IncidenceHandlerRest struct {
	service  *service.IncidenceService
	resource *resource.ServerResource
}

func NewIncidenceHandlerRest(resource *resource.ServerResource) *IncidenceHandlerRest {
	return &IncidenceHandlerRest{
		resource: resource,
	}
}

func (h *IncidenceHandlerRest) MakeRoutes() {

	scChip := service.NewChipService(NewChipRepository(h.resource.Db))
	scSharkChip := service.NewSharkChipService(NewSharkChipRepository(h.resource.Db), scChip)
	scShark := service.NewSharkService(NewSharkRepository(h.resource.Db), scSharkChip)

	h.service = service.NewIncidenceService(NewIncidenceRepository(h.resource.Db), scSharkChip, scShark)

	router := h.resource.DefaultRouter("/incidences", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *IncidenceHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var IncidenceDtoIn *dto.IncidenceDtoIn
		var IncidenceDtoOut *dto.IncidenceDtoOut

		h.resource.Restful.LoadData(w, r)
		IncidenceDtoIn = dto.NewIncidenceDtoIn()
		h.resource.Restful.BindData(&IncidenceDtoIn)
		IncidenceDtoOut, err = h.service.Get(IncidenceDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, IncidenceDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *IncidenceHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Incidences := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, Incidences)
	})
}

func (h *IncidenceHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var IncidenceDtoIn *dto.IncidenceDtoIn

		h.resource.Restful.LoadData(w, r)
		IncidenceDtoIn = dto.NewIncidenceDtoIn()
		h.resource.Restful.BindData(&IncidenceDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(IncidenceDtoIn)

		if err != nil {
			transaction.Rollback()
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, err.Error())
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *IncidenceHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var IncidenceDtoIn *dto.IncidenceDtoIn

		h.resource.Restful.LoadData(w, r)
		IncidenceDtoIn = dto.NewIncidenceDtoIn()
		h.resource.Restful.BindData(&IncidenceDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(IncidenceDtoIn)

		if err != nil {
			transaction.Rollback()
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, err.Error())
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *IncidenceHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var IncidenceDtoIn *dto.IncidenceDtoIn

		h.resource.Restful.LoadData(w, r)
		IncidenceDtoIn = dto.NewIncidenceDtoIn()
		h.resource.Restful.BindData(&IncidenceDtoIn)
		err = h.service.Remove(IncidenceDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *IncidenceHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GridConfig := grid.NewGridConfig()
		h.resource.Restful.LoadData(w, r).BindData(&GridConfig)
		dataGrid := h.service.Grid(GridConfig)

		if GridConfig.RowsPage == "0" {
			grid.ResponseDataGrid(w, "csv", dataGrid, "Incidence")
		} else {
			h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
		}
	})
}
