package adapter

import (
	adapter_shared "mp-projeto/shared/adapter"
	"mp-projeto/shared/grid"
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type LocationHandlerRest struct {
	service  *service.LocationService
	resource *resource.ServerResource
}

func NewLocationHandlerRest(resource *resource.ServerResource) *LocationHandlerRest {
	return &LocationHandlerRest{
		resource: resource,
	}
}

func (h *LocationHandlerRest) MakeRoutes() {

	h.service = service.NewLocationService(NewLocationRepository(h.resource.Db))

	router := h.resource.DefaultRouter("/locations", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *LocationHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var LocationDtoIn *dto.LocationDtoIn
		var LocationDtoOut *dto.LocationDtoOut

		h.resource.Restful.LoadData(w, r)
		LocationDtoIn = dto.NewLocationDtoIn()
		h.resource.Restful.BindData(&LocationDtoIn)
		LocationDtoOut, err = h.service.Get(LocationDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, LocationDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *LocationHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Locations := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, Locations)
	})
}

func (h *LocationHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var LocationDtoIn *dto.LocationDtoIn

		h.resource.Restful.LoadData(w, r)
		LocationDtoIn = dto.NewLocationDtoIn()
		h.resource.Restful.BindData(&LocationDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(LocationDtoIn)

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

func (h *LocationHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var LocationDtoIn *dto.LocationDtoIn

		h.resource.Restful.LoadData(w, r)
		LocationDtoIn = dto.NewLocationDtoIn()
		h.resource.Restful.BindData(&LocationDtoIn)
		transaction := adapter_shared.BeginTransaction(h.resource.Db)

		err = h.service.WithTransaction(transaction).Save(LocationDtoIn)

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

func (h *LocationHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var LocationDtoIn *dto.LocationDtoIn

		h.resource.Restful.LoadData(w, r)
		LocationDtoIn = dto.NewLocationDtoIn()
		h.resource.Restful.BindData(&LocationDtoIn)
		err = h.service.Remove(LocationDtoIn)

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

func (h *LocationHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GridConfig := grid.NewGridConfig()
		h.resource.Restful.LoadData(w, r).BindData(&GridConfig)
		dataGrid := h.service.Grid(GridConfig)

		if GridConfig.RowsPage == "0" {
			grid.ResponseDataGrid(w, "csv", dataGrid, "Location")
		} else {
			h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
		}
	})
}
