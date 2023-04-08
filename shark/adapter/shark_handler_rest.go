package adapter

import (
	adapter_shared "mp-projeto/shared/adapter"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type SharkHandlerRest struct {
	service  *service.SharkService
	resource *resource.ServerResource
}

func NewSharkHandlerRest(resource *resource.ServerResource) *SharkHandlerRest {
	return &SharkHandlerRest{
		resource: resource,
	}
}

func (h *SharkHandlerRest) MakeRoutes() {

	scSharkChipService := service.NewSharkChipService(NewSharkChipRepository(h.resource.Db))

	h.service = service.NewSharkService(NewSharkRepository(h.resource.Db), scSharkChipService)

	router := h.resource.DefaultRouter("/sharks", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *SharkHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkDtoIn *dto.SharkDtoIn
		var SharkAllDtoOut *dto.SharkAllDtoOut

		h.resource.Restful.LoadData(w, r)
		SharkDtoIn = dto.NewSharkDtoIn()
		h.resource.Restful.BindData(&SharkDtoIn)
		SharkAllDtoOut, err = h.service.Get(SharkDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, SharkAllDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *SharkHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Sharks := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, Sharks)
	})
}

func (h *SharkHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkAllDtoIn *dto.SharkAllDtoIn

		h.resource.Restful.LoadData(w, r)
		SharkAllDtoIn = dto.NewSharkAllDtoIn()
		h.resource.Restful.BindData(&SharkAllDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		_, err = h.service.WithTransaction(transaction).SaveAll(SharkAllDtoIn)

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

func (h *SharkHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkAllDtoIn *dto.SharkAllDtoIn

		h.resource.Restful.LoadData(w, r)
		SharkAllDtoIn = dto.NewSharkAllDtoIn()
		h.resource.Restful.BindData(&SharkAllDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		_, err = h.service.WithTransaction(transaction).SaveAll(SharkAllDtoIn)

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

func (h *SharkHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkDtoIn *dto.SharkDtoIn

		h.resource.Restful.LoadData(w, r)
		SharkDtoIn = dto.NewSharkDtoIn()
		h.resource.Restful.BindData(&SharkDtoIn)
		err = h.service.Remove(SharkDtoIn)

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

func (h *SharkHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GridConfig := grid.NewGridConfig()
		h.resource.Restful.LoadData(w, r).BindData(&GridConfig)
		dataGrid := h.service.Grid(GridConfig)

		if GridConfig.RowsPage == "0" {
			grid.ResponseDataGrid(w, "csv", dataGrid, "Shark")
		} else {
			h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
		}
	})
}
