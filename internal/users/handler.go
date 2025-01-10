package users

import (
	"net/http"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	"github.com/rahadianir/swiper/internal/pkg/xerrors"
	"github.com/rahadianir/swiper/internal/pkg/xhttp"
)

type UserHandler struct {
	*common.Dependencies
	UserLogic UserLogicInterface
}

func NewUserHandler(deps *common.Dependencies, userLogic UserLogicInterface) *UserHandler {
	return &UserHandler{
		Dependencies: deps,
		UserLogic:    userLogic,
	}
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterRequest
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

	userID, err := handler.UserLogic.Register(r.Context(), payload)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to register user",
			Data:    nil,
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, models.BaseResponse{
		Message: "user registered successfully",
		Data: models.User{
			ID: userID,
		},
	}, http.StatusCreated)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload models.LoginRequest
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

	token, refreshToken, err := handler.UserLogic.Login(r.Context(), payload)
	if err != nil {
		xhttp.SendJSONResponse(w, models.BaseResponse{
			Error:   err.Error(),
			Message: "failed to login",
			Data:    nil,
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, models.BaseResponse{
		Message: "logged in successfully",
		Data: models.LoginResponse{
			Token:        token,
			RefreshToken: refreshToken,
		},
	}, http.StatusCreated)
}
