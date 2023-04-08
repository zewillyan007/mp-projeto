package adapter

import (
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type ChipStatusTypeHandlerRest struct {
	service  *service.ChipStatusTypeService
	resource *resource.ServerResource
}

func NewChipStatusTypeHandlerRest(resource *resource.ServerResource) *ChipStatusTypeHandlerRest {
	return &ChipStatusTypeHandlerRest{
		resource: resource,
	}
}

func (h *ChipStatusTypeHandlerRest) MakeRoutes() {

	h.service = service.NewChipStatusTypeService(NewChipStatusTypeRepository(h.resource.Db))

	router := h.resource.DefaultRouter("/chip-status-types", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
}

func (h *ChipStatusTypeHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var ChipStatusTypeDtoIn *dto.ChipStatusTypeDtoIn
		var ChipStatusTypeDtoOut *dto.ChipStatusTypeDtoOut

		h.resource.Restful.LoadData(w, r)
		ChipStatusTypeDtoIn = dto.NewChipStatusTypeDtoIn()
		h.resource.Restful.BindData(&ChipStatusTypeDtoIn)
		ChipStatusTypeDtoOut, err = h.service.Get(ChipStatusTypeDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, ChipStatusTypeDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *ChipStatusTypeHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ChipStatusTypes := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, ChipStatusTypes)
	})
}
