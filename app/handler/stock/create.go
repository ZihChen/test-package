package stock

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"test-package-02/app/global/structs"
)

func (h *Handler)Create(c *gin.Context) {

	raw := structs.StockDetail{}
	if err := c.ShouldBindJSON(&raw); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(raw); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	h.buss.CreateStock(&raw)
}