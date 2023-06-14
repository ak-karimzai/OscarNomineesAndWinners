package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nmferoz/db/internal/db"
)

type listActorRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListActors(ctx *gin.Context) {
	var req listActorRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListActorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	actors, err := server.query.ListActors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, actors)
}

type getActorRequest struct {
	ID int64 `uri:"id", binding:"required,min=1"`
}

func (server *Server) GetActor(ctx *gin.Context) {
	var req getActorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	actor, err := server.query.GetActor(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, actor)
}

type createActorRequest struct {
	Name        string `json:"name"`
	BirthYear   int32  `json:"birth_year"`
	Nationality string `json:"nationality"`
}

func (server *Server) CreateActor(ctx *gin.Context) {
	var req createActorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateActorParams{
		Name:        req.Name,
		BirthYear:   req.BirthYear,
		Nationality: req.Nationality,
	}
	actor, err := server.query.CreateActor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, actor)
}

type updateActorRequest struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	BirthYear   int32  `json:"birth_year"`
	Nationality string `json:"nationality"`
}

func (server *Server) UpdateActor(ctx *gin.Context) {
	var req updateActorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.UpdateActorParams{
		Name:        req.Name,
		BirthYear:   req.BirthYear,
		Nationality: req.Nationality,
		ID:          req.ID,
	}
	err := server.query.UpdateActor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
