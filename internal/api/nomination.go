package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nmferoz/db/internal/db"
)

type nominationResponse struct {
	ID       int64    `json:"id"`
	Movie    db.Movie `json:"movie"`
	Award    db.Award `json:"award"`
	Year     int32    `json:"year"`
	IsWinner bool     `json:"is_winner"`
}

type listNominationRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListNominations(ctx *gin.Context) {
	var req listNominationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListNominationsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	nomintations, err := server.query.ListNominations(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	nominationResps, err := nominationsToResponseNominations(
		ctx, server, nomintations)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nominationResps)
}

type getNominationRequest struct {
	ID int64 `uri:"id", binding:"required,min=1"`
}

func (server *Server) GetNomination(ctx *gin.Context) {
	var req getNominationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	nomination, err := server.query.GetNomination(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	nominationResp, err := nominationToResponseNomination(
		ctx, server, &nomination)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nominationResp)
}

type createNominationRequest struct {
	MovieID  int32 `json:"movie_id"`
	AwardID  int32 `json:"award_id"`
	Year     int32 `json:"year"`
	IsWinner bool  `json:"is_winner"`
}

func (server *Server) CreateNomination(ctx *gin.Context) {
	var req createNominationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateNominationParams{
		MovieID:  req.MovieID,
		AwardID:  req.AwardID,
		Year:     req.Year,
		IsWinner: req.IsWinner,
	}
	nomination, err := server.query.CreateNomination(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	nominationResp, err := nominationToResponseNomination(
		ctx, server, &nomination)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nominationResp)
}

type updateNominationRequest struct {
	ID       int64 `json:"id"`
	MovieID  int32 `json:"movie_id"`
	AwardID  int32 `json:"award_id"`
	Year     int32 `json:"year"`
	IsWinner bool  `json:"is_winner"`
}

func (server *Server) UpdateNomination(ctx *gin.Context) {
	var req updateNominationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.UpdateNominationParams{
		MovieID:  req.MovieID,
		AwardID:  req.AwardID,
		Year:     req.Year,
		IsWinner: req.IsWinner,
		ID:       req.ID,
	}
	err := server.query.UpdateNomination(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func nominationToResponseNomination(ctx context.Context, server *Server,
	nomination *db.Nomination) (*nominationResponse, error) {
	movie, err := server.query.GetMovie(ctx, int64(nomination.MovieID))
	if err != nil {
		return nil, err
	}

	award, err := server.query.GetAward(ctx, int64(nomination.AwardID))
	if err != nil {
		return nil, err
	}

	return &nominationResponse{
		ID:       nomination.ID,
		Movie:    movie,
		Award:    award,
		Year:     nomination.Year,
		IsWinner: nomination.IsWinner,
	}, nil
}

func nominationsToResponseNominations(ctx context.Context, server *Server,
	nominations []db.Nomination) ([]nominationResponse, error) {
	var nominationRes = []nominationResponse{}

	for _, nomination := range nominations {
		nonRes, err := nominationToResponseNomination(
			ctx, server, &nomination)
		if err != nil {
			return nil, err
		}
		nominationRes = append(nominationRes, *nonRes)
	}

	return nominationRes, nil
}
