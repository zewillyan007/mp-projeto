package adapter

import (
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type SexHandlerRest struct {
	service  *service.SexService
	resource *resource.ServerResource
}

func NewSexHandlerRest(resource *resource.ServerResource) *SexHandlerRest {
	return &SexHandlerRest{
		resource: resource,
	}
}

func (h *SexHandlerRest) MakeRoutes() {

	h.service = service.NewSexService(NewSexRepository(h.resource.Db))

	router := h.resource.DefaultRouter("/sexs", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
}

func (h *SexHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SexDtoIn *dto.SexDtoIn
		var SexDtoOut *dto.SexDtoOut

		h.resource.Restful.LoadData(w, r)
		SexDtoIn = dto.NewSexDtoIn()
		h.resource.Restful.BindData(&SexDtoIn)
		SexDtoOut, err = h.service.Get(SexDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, SexDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *SexHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Sexs := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, Sexs)
	})
}
