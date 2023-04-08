package adapter

import (
	adapter_shared "mp-projeto/shared/adapter"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type ChipHandlerRest struct {
	service  *service.ChipService
	resource *resource.ServerResource
}

func NewChipHandlerRest(resource *resource.ServerResource) *ChipHandlerRest {
	return &ChipHandlerRest{
		resource: resource,
	}
}

func (h *ChipHandlerRest) MakeRoutes() {

	h.service = service.NewChipService(NewChipRepository(h.resource.Db))

	router := h.resource.DefaultRouter("/chips", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *ChipHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var ChipDtoIn *dto.ChipDtoIn
		var ChipDtoOut *dto.ChipDtoOut

		h.resource.Restful.LoadData(w, r)
		ChipDtoIn = dto.NewChipDtoIn()
		h.resource.Restful.BindData(&ChipDtoIn)
		ChipDtoOut, err = h.service.Get(ChipDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, ChipDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *ChipHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Chips := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, Chips)
	})
}

func (h *ChipHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var ChipDtoIn *dto.ChipDtoIn

		h.resource.Restful.LoadData(w, r)
		ChipDtoIn = dto.NewChipDtoIn()
		h.resource.Restful.BindData(&ChipDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(ChipDtoIn)

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

func (h *ChipHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var ChipDtoIn *dto.ChipDtoIn

		h.resource.Restful.LoadData(w, r)
		ChipDtoIn = dto.NewChipDtoIn()
		h.resource.Restful.BindData(&ChipDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(ChipDtoIn)

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

func (h *ChipHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var ChipDtoIn *dto.ChipDtoIn

		h.resource.Restful.LoadData(w, r)
		ChipDtoIn = dto.NewChipDtoIn()
		h.resource.Restful.BindData(&ChipDtoIn)
		err = h.service.Remove(ChipDtoIn)

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

func (h *ChipHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GridConfig := grid.NewGridConfig()
		h.resource.Restful.LoadData(w, r).BindData(&GridConfig)
		dataGrid := h.service.Grid(GridConfig)

		if GridConfig.RowsPage == "0" {
			grid.ResponseDataGrid(w, "csv", dataGrid, "Chip")
		} else {
			h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
		}
	})
}
