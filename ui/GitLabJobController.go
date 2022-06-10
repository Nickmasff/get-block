package ui

import (
	"get-block/application/balance"
	"github.com/gin-gonic/gin"
)

type BalanceController struct {
	balanceService *balance.Service
}

func NewBalanceController(balanceService *balance.Service) *BalanceController {
	return &BalanceController{balanceService: balanceService}
}

// GetMostChangedBalanceAddress godoc
// @Description Get address whose balance has changed (in any direction) more than the rest in the last hundred blocks
// @Summary Returns address
// @Produce json
// @Success 200 {array} balance.MostChangedBalanceDto
// @Failure 500 {object} gb_swagger.JSONResultError "Server error message"
// @Router /most-changed-balance [get]
func (controller *BalanceController) GetMostChangedBalanceAddress(context *gin.Context) {

	result, err := controller.balanceService.GetMostChangedBalanceAddress()
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	context.JSON(200, result)
}
