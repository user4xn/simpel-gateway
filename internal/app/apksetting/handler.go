package apksetting

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

func (h *handler) Store(c *gin.Context) {
	var payload dto.PayloadApkSetting
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

	response := util.APIResponse("Success create or update apk setting", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *handler) FindOne(c *gin.Context) {
	data, err := h.service.FindOne(c)
	if err == constants.ErrorNoFoundDataApkSetting {
		response := util.APIResponse(fmt.Sprintf("%s", constants.ErrorNoFoundDataApkSetting), http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.APIResponse("Success get apk setting", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
