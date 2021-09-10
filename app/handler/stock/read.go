package stock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler)Read(c *gin.Context)  {

	stockList := h.buss.GetStockList()

	c.JSON(http.StatusOK, stockList)
}