package premium

import (
	"net/http"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	"github.com/rahadianir/swiper/internal/pkg/xerrors"
	"github.com/rahadianir/swiper/internal/pkg/xhttp"
)

type PremiumHandler struct {
	*common.Dependencies
	PremiumLogic PremiumLogicInterface
}

func NewPremiumHandler(deps *common.Dependencies, premiumLogic PremiumLogicInterface) *PremiumHandler {
	return &PremiumHandler{
		Dependencies: deps,
		PremiumLogic: premiumLogic,
	}
}

func (handler *PremiumHandler) EnablePremium(w http.ResponseWriter, r *http.Request) {
	var payload models.EnablePremiumRequest
	err := xhttp.BindJSONRequest(r, &payload)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to bind request body",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	err = handler.Validator.Struct(payload)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to validate request body",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	err = handler.PremiumLogic.EnablePremium(r.Context(), payload.ID)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to enable premium user",
			Data:    nil,
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, models.BaseResponse{
		Message: "premium user enabled successfully",
		Data: models.User{
			ID: payload.ID,
		},
	}, http.StatusOK)
}
