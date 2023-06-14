package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nmferoz/db/internal/db"
)

type performanceResponse struct {
	ID    int64    `json:"id"`
	Actor db.Actor `json:"actor"`
	Movie db.Movie `json:"movie"`
	Year  int32    `json:"year"`
}

type listPerformancesRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListPerformances(ctx *gin.Context) {
	var req listNominationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListPerformancesParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	performances, err := server.query.ListPerformances(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resPerformances, err := performancesToPerformanceResponses(
		ctx, server, performances)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resPerformances)
}

type getPerformanceRequest struct {
	ID int64 `uri:"id", binding:"required,min=1"`
}

func (server *Server) GetPerformance(ctx *gin.Context) {
	var req getPerformanceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	performance, err := server.query.GetPerformance(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resPerformance, err := performanceToPerformanceResponse(
		ctx, server, &performance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resPerformance)
}

type createPerformaceRequest struct {
	ActorID int32 `json:"actor_id"`
	MovieID int32 `json:"movie_id"`
	Year    int32 `json:"year"`
}

func (server *Server) CreatePerformance(ctx *gin.Context) {
	var req createPerformaceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreatePerformanceParams{
		MovieID: req.MovieID,
		ActorID: req.ActorID,
		Year:    req.Year,
	}
	performance, err := server.query.CreatePerformance(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resPerformance, err := performanceToPerformanceResponse(
		ctx, server, &performance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resPerformance)
}

type updatePerformaceRequest struct {
	ActorID int32 `json:"actor_id"`
	MovieID int32 `json:"movie_id"`
	Year    int32 `json:"year"`
	ID      int64 `json:"id"`
}

func (server *Server) UpdatePerformance(ctx *gin.Context) {
	var req updatePerformaceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.UpdatePerformanceParams{
		MovieID: req.MovieID,
		ActorID: req.ActorID,
		Year:    req.Year,
		ID:      req.ID,
	}
	err := server.query.UpdatePerformance(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func performanceToPerformanceResponse(
	ctx context.Context, server *Server, performance *db.Performance) (*performanceResponse, error) {
	actor, err := server.query.GetActor(ctx, int64(performance.ActorID))
	if err != nil {
		return nil, err
	}

	movie, err := server.query.GetMovie(ctx, int64(performance.MovieID))
	if err != nil {
		return nil, err
	}

	return &performanceResponse{
		ID:    performance.ID,
		Actor: actor,
		Movie: movie,
		Year:  performance.Year,
	}, nil
}

func performancesToPerformanceResponses(
	ctx context.Context, server *Server, performances []db.Performance) ([]performanceResponse, error) {
	var resPerofrmances = []performanceResponse{}

	for _, proformance := range performances {
		resPerformance, err := performanceToPerformanceResponse(
			ctx, server, &proformance)
		if err != nil {
			return nil, err
		}
		resPerofrmances = append(resPerofrmances, *resPerformance)
	}

	return resPerofrmances, nil
}
