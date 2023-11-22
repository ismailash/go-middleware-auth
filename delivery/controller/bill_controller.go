package controller

import (
	"net/http"

	"enigmacamp.com/be-enigma-laundry/delivery/middleware"
	"enigmacamp.com/be-enigma-laundry/model/dto"
	"enigmacamp.com/be-enigma-laundry/usecase"
	"enigmacamp.com/be-enigma-laundry/utils/common"
	"github.com/gin-gonic/gin"
)

type BillController struct {
	uc             usecase.BillUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (b *BillController) createHandler(ctx *gin.Context) {
	var payload dto.BillRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rspPayload, err := b.uc.RegisterNewBill(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (b *BillController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	rspPayload, err := b.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BillController) Route() {
	br := b.rg.Group("/bills", b.authMiddleware.RequireToken("admin", "employee"))
	br.POST("/", b.createHandler)
	br.GET("/:id", b.getHandler)
}

func NewBillController(uc usecase.BillUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *BillController {
	return &BillController{uc: uc, rg: rg, authMiddleware: authMiddleware}
}
