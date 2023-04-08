package adapter

import (
	adapter_shared "mp-projeto/shared/adapter"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type SharkChipHandlerRest struct {
	service  *service.SharkChipService
	resource *resource.ServerResource
}

func NewSharkChipHandlerRest(resource *resource.ServerResource) *SharkChipHandlerRest {
	return &SharkChipHandlerRest{
		resource: resource,
	}
}

func (h *SharkChipHandlerRest) MakeRoutes() {

	h.service = service.NewSharkChipService(NewSharkChipRepository(h.resource.Db))

	router := h.resource.DefaultRouter("/shark-chips", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *SharkChipHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkChipDtoIn *dto.SharkChipDtoIn
		var SharkChipDtoOut *dto.SharkChipDtoOut

		h.resource.Restful.LoadData(w, r)
		SharkChipDtoIn = dto.NewSharkChipDtoIn()
		h.resource.Restful.BindData(&SharkChipDtoIn)
		SharkChipDtoOut, err = h.service.Get(SharkChipDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, SharkChipDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *SharkChipHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		SharkChips := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, SharkChips)
	})
}

func (h *SharkChipHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkChipDtoIn *dto.SharkChipDtoIn

		h.resource.Restful.LoadData(w, r)
		SharkChipDtoIn = dto.NewSharkChipDtoIn()
		h.resource.Restful.BindData(&SharkChipDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(SharkChipDtoIn)

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

func (h *SharkChipHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkChipDtoIn *dto.SharkChipDtoIn

		h.resource.Restful.LoadData(w, r)
		SharkChipDtoIn = dto.NewSharkChipDtoIn()
		h.resource.Restful.BindData(&SharkChipDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(SharkChipDtoIn)

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

func (h *SharkChipHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkChipDtoIn *dto.SharkChipDtoIn

		h.resource.Restful.LoadData(w, r)
		SharkChipDtoIn = dto.NewSharkChipDtoIn()
		h.resource.Restful.BindData(&SharkChipDtoIn)
		err = h.service.Remove(SharkChipDtoIn)

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

func (h *SharkChipHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GridConfig := grid.NewGridConfig()
		h.resource.Restful.LoadData(w, r).BindData(&GridConfig)
		dataGrid := h.service.Grid(GridConfig)

		if GridConfig.RowsPage == "0" {
			grid.ResponseDataGrid(w, "csv", dataGrid, "SharkChip")
		} else {
			h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
		}
	})
}
