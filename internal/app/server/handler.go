package server

import (
	"fmt"
	"io"
	"net/http"
	"simpel-gateway/internal/dto"
	"simpel-gateway/internal/factory"
	"simpel-gateway/pkg/constants"
	"simpel-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) GetServers(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.service.GetListServer(ctx)
	if err != nil {
		response := util.APIResponse("Failed to retrieve user list"+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list servers", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *handler) Store(c *gin.Context) {
	var payload dto.PayloadServer
	if err := c.ShouldBind(&payload); err != nil {
		errorMessage := gin.H{"errors": "Please fill data"}
		if err != io.EOF {
			errors := util.FormatValidationError(err)
			errorMessage = gin.H{"errors": errors}
		}
		response := util.APIResponse("Error validation", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.service.Store(c, payload)

	if err == constants.DuplicateStoreServer {
		response := util.APIResponse(fmt.Sprintf("%s", constants.DuplicateStoreServer), http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.APIResponse("Success store server", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *handler) CheckServerIp(c *gin.Context) {
	var payload dto.PayloadCheckServerIp
	if err := c.ShouldBind(&payload); err != nil {
		errorMessage := gin.H{"errors": "Please fill data"}
		if err != io.EOF {
			errors := util.FormatValidationError(err)
			errorMessage = gin.H{"errors": errors}
		}
		response := util.APIResponse("Error validation", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println(payload)
	data := h.service.CheckServerIp(c, payload)
	response := util.APIResponse("Success check server ip", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
