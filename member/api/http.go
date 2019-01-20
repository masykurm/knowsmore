package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/robertotambunan/knowsmore/member"
	m "github.com/robertotambunan/knowsmore/member/model"
)

type httpMemberHandler struct {
	mUsecase member.Usecase
}

// NewMemberHTTPtpHandler create object representation from new member api using http
func NewMemberHTTPtpHandler(e *echo.Echo, usecase member.Usecase) {
	handler := &httpMemberHandler{
		mUsecase: usecase,
	}
	e.POST("/autocomplete", handler.GetAutocomplete)
}

// GetAutocomplete Handler function for autocomplete member
func (mh *httpMemberHandler) GetAutocomplete(c echo.Context) error {
	var mReq m.Request
	err := c.Bind(&mReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&mReq); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var memberResp m.Response
	memberResp, err = mh.mUsecase.GetByAutocomplete(ctx, mReq.Keyword)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, m.Response{Error: err.Error()})
	} else {
		memberResp.Status = http.StatusOK
	}
	return c.JSON(http.StatusOK, memberResp)
}

func isRequestValid(m *m.Request) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
