package adapter

import (
	"mp-projeto/shared/resource"
	"mp-projeto/shark/core/domain/dto"
	"mp-projeto/shark/core/service"
	"net/http"
)

type SharkChipStatusTypeHandlerRest struct {
	service  *service.SharkChipStatusTypeService
	resource *resource.ServerResource
}

func NewSharkChipStatusTypeHandlerRest(resource *resource.ServerResource) *SharkChipStatusTypeHandlerRest {
	return &SharkChipStatusTypeHandlerRest{
		resource: resource,
	}
}

func (h *SharkChipStatusTypeHandlerRest) MakeRoutes() {

	h.service = service.NewSharkChipStatusTypeService(NewSharkChipStatusTypeRepository(h.resource.Db))

	router := h.resource.DefaultRouter("/shark-chip-status-types", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
}

func (h *SharkChipStatusTypeHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var SharkChipStatusTypeDtoIn *dto.SharkChipStatusTypeDtoIn
		var SharkChipStatusTypeDtoOut *dto.SharkChipStatusTypeDtoOut

		h.resource.Restful.LoadData(w, r)
		SharkChipStatusTypeDtoIn = dto.NewSharkChipStatusTypeDtoIn()
		h.resource.Restful.BindData(&SharkChipStatusTypeDtoIn)
		SharkChipStatusTypeDtoOut, err = h.service.Get(SharkChipStatusTypeDtoIn)

		if err != nil {
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, err.Error())
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, SharkChipStatusTypeDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *SharkChipStatusTypeHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		SharkChipStatusTypes := h.service.GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, SharkChipStatusTypes)
	})
}
