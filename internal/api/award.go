package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nmferoz/db/internal/db"
)

type listAwardRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListAwards(ctx *gin.Context) {
	var req listAwardRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListAwardsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	actors, err := server.query.ListAwards(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, actors)
}

type getAwardRequest struct {
	ID int64 `uri:"id", binding:"required,min=1"`
}

func (server *Server) GetAward(ctx *gin.Context) {
	var req getAwardRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	award, err := server.query.GetAward(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, award)
}

type createAwardRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (server *Server) CreateAward(ctx *gin.Context) {
	var req createAwardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateAwardParams{
		Name:     req.Name,
		Category: req.Category,
	}
	award, err := server.query.CreateAward(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, award)
}

type updateAwardRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	ID       int64  `json:"id"`
}

func (server *Server) UpdateAward(ctx *gin.Context) {
	var req updateAwardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.UpdateAwardParams{
		Name:     req.Name,
		Category: req.Category,
		ID:       req.ID,
	}
	err := server.query.UpdateAward(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
