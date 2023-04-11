package adapter

import (
	adapter_shared "mp-projeto/shared/adapter"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type MicrocontrollerHandlerRest struct {
	service  *service.MicrocontrollerService
	resource *resource.ServerResource
}

func NewMicrocontrollerHandlerRest(resource *resource.ServerResource) *MicrocontrollerHandlerRest {
	return &MicrocontrollerHandlerRest{
		resource: resource,
	}
}

func (h *MicrocontrollerHandlerRest) MakeRoutes() {

	h.service = service.NewMicrocontrollerService(NewMicrocontrollerRepository(h.resource.Db))

	router := h.resource.DefaultRouter("/microcontrollers", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *MicrocontrollerHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var MicrocontrollerDtoIn *dto.MicrocontrollerDtoIn
		var MicrocontrollerDtoOut *dto.MicrocontrollerDtoOut

		h.resource.Restful.LoadData(w, r)
		MicrocontrollerDtoIn = dto.NewMicrocontrollerDtoIn()
		h.resource.Restful.BindData(&MicrocontrollerDtoIn)
		MicrocontrollerDtoOut, err = h.service.Get(MicrocontrollerDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, MicrocontrollerDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *MicrocontrollerHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Microcontrollers := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, Microcontrollers)
	})
}

func (h *MicrocontrollerHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var MicrocontrollerDtoIn *dto.MicrocontrollerDtoIn

		h.resource.Restful.LoadData(w, r)
		MicrocontrollerDtoIn = dto.NewMicrocontrollerDtoIn()
		h.resource.Restful.BindData(&MicrocontrollerDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(MicrocontrollerDtoIn)

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

func (h *MicrocontrollerHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var MicrocontrollerDtoIn *dto.MicrocontrollerDtoIn

		h.resource.Restful.LoadData(w, r)
		MicrocontrollerDtoIn = dto.NewMicrocontrollerDtoIn()
		h.resource.Restful.BindData(&MicrocontrollerDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(MicrocontrollerDtoIn)

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

func (h *MicrocontrollerHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var MicrocontrollerDtoIn *dto.MicrocontrollerDtoIn

		h.resource.Restful.LoadData(w, r)
		MicrocontrollerDtoIn = dto.NewMicrocontrollerDtoIn()
		h.resource.Restful.BindData(&MicrocontrollerDtoIn)
		err = h.service.Remove(MicrocontrollerDtoIn)

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

func (h *MicrocontrollerHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GridConfig := grid.NewGridConfig()
		h.resource.Restful.LoadData(w, r).BindData(&GridConfig)
		dataGrid := h.service.Grid(GridConfig)

		if GridConfig.RowsPage == "0" {
			grid.ResponseDataGrid(w, "csv", dataGrid, "Microcontroller")
		} else {
			h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
		}
	})
}
