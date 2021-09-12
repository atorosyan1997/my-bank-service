package handlers

import (
	"github.com/gin-gonic/gin"
	"my-bank-service/internal/config"
	data2 "my-bank-service/internal/data"
	"my-bank-service/internal/reposytory"
	"my-bank-service/pkg/logging"
	"net/http"
)

type AccountHandler struct {
	auth   *AuthHandler
	logger logging.Logger
	repo   reposytory.AccountRepository
}

// NewAccountHandler returns a new AccountHandler instance
func NewAccountHandler(auth *AuthHandler, l logging.Logger, r reposytory.AccountRepository) *AccountHandler {
	return &AccountHandler{
		auth:   auth,
		logger: l,
		repo:   r,
	}
}

func (a *AccountHandler) Routes(engine *gin.Engine) {
	accessTok := engine.Group(config.GroupPath)
	{
		accessTok.Use(a.auth.MiddlewareValidateAccessToken)
		accessTok.GET(config.PaymentPath, a.Payment)
	}
}

// Payment handles payment request
// @Summary payment
// @Tags Payment
// @Security ApiKeyAuth
// @Description balance payment
// @ID balance-payment
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 500 {integer} integer 2
// @Router /payment/ [get]
func (a *AccountHandler) Payment(ctx *gin.Context) {
	ctx.Set("Content-Type", "application/json")
	authUser := ctx.Request.Context().Value(UserAuthKey{}).(data2.AuthUser)
	balance, err := a.repo.Withdraw(authUser.User.ID, config.DefSum)
	if err != nil {
		a.logger.Error("unable to payment", " error ", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		_ = data2.ToJSON(&GenericResponse{Status: false, Message: "Unable to Payment.Insufficient funds."}, ctx.Writer)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
	_ = data2.ToJSON(&GenericResponse{
		Status:  true,
		Message: "Successfully payment",
		Data: struct {
			Balance float64
		}{
			Balance: balance,
		},
	}, ctx.Writer)
}
