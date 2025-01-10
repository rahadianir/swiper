package swiper

import (
	"net/http"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	"github.com/rahadianir/swiper/internal/pkg/xcontext"
	"github.com/rahadianir/swiper/internal/pkg/xerrors"
	"github.com/rahadianir/swiper/internal/pkg/xhttp"
)

type SwiperHandler struct {
	*common.Dependencies
	SwiperLogic SwiperLogicInterface
}

func NewSwiperHandler(deps *common.Dependencies, swiperLogic SwiperLogicInterface) *SwiperHandler {
	return &SwiperHandler{
		Dependencies: deps,
		SwiperLogic:  swiperLogic,
	}
}

func (handler *SwiperHandler) GetTargetProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := xcontext.GetUserID(r.Context())
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse request: invalid user id",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	profile, err := handler.SwiperLogic.GetTargetProfile(r.Context(), userID)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to get profile",
			Data:    nil,
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, models.BaseResponse{
		Message: "profile fetched",
		Data:    profile,
	}, http.StatusOK)

}
