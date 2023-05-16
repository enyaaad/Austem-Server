package handlers

import (
	"AustemServer/models"
	"AustemServer/package/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	STATUS_OK    = "ok"
	STATUS_ERROR = "error"
)

type response struct {
	Status string `json:"status"`
	Msg    string `json:"message,omitempty"`
}

type handler struct {
	useCase auth.UseCase
}

func newHandler(useCase auth.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

type signInResponse struct {
	*response
	Token string `json:"token,omitempty"`
}

func newSignInResponse(status, msg, token string) *signInResponse {
	return &signInResponse{
		&response{Status: status, Msg: msg},
		token,
	}
}
func (h *handler) signIn(c *gin.Context) {
	inp := new(models.Users)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), inp)

	if err != nil {
		if err == auth.ErrInvalidAccessToken {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(STATUS_ERROR, err.Error(), ""))
			return
		}

		if err == auth.ErrUserDoesNotExist {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(STATUS_ERROR, err.Error(), ""))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, newSignInResponse(STATUS_ERROR, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, newSignInResponse(STATUS_OK, "", token))
}

func PostHandler() {

}
func EditHandler() {

}
func GetHandler() {

}
func DeleteHandler() {

}
