package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	UseCase OrderUseCase
}

func NewOrderController(useCase OrderUseCase) *OrderController {
	return &OrderController{UseCase: useCase}
}

func (s *OrderController) GetAllOrders(c *gin.Context) {
	// return s.OrderUseCase.ListOrders(c)
}

func (o *OrderController) Create(ctx *gin.Context) {
	var order Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := o.UseCase.Create(ctx.Request.Context(), &order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (u *OrderController) GetById(c *gin.Context) {
	panic("unimplemented")
	// userID := c.GetString("x-user-id")

	// tasks, err := u.OrderUseCase.FetchByCustomerID(c, userID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, tasks)
}
