package swiper

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (handler *SwiperHandler) SwipeRight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   "invalid target id",
			Message: "failed to parse request",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}
	targetID, err := strconv.Atoi(id)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse request",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	userID, err := xcontext.GetUserID(r.Context())
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse request: invalid user id",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	isMatch, err := handler.SwiperLogic.SwipeRight(r.Context(), userID, targetID)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to swipe right",
			Data:    nil,
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, models.BaseResponse{
		Message: "swiped right",
		Data: map[string]bool{
			"is_matched": isMatch,
		},
	}, http.StatusOK)

}

func (handler *SwiperHandler) SwipeLeft(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   "invalid target id",
			Message: "failed to parse request",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}
	targetID, err := strconv.Atoi(id)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse request",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	userID, err := xcontext.GetUserID(r.Context())
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse request: invalid user id",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	err = handler.SwiperLogic.SwipeLeft(r.Context(), userID, targetID)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to swipe left",
			Data:    nil,
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, models.BaseResponse{
		Message: "swiped left",
	}, http.StatusOK)

}
