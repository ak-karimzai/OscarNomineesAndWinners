package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nmferoz/db/internal/db"
)

type listMovieRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListMovies(ctx *gin.Context) {
	var req listMovieRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListMoviesParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	movies, err := server.query.ListMovies(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

type getMovieRequest struct {
	ID int64 `uri:"id", binding:"required,min=1"`
}

func (server *Server) GetMovie(ctx *gin.Context) {
	var req getMovieRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	movie, err := server.query.GetMovie(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

type createMovieRequest struct {
	Title       string `json:"title" binding:"required"`
	ReleaseYear int32  `json:"release_year" binding:"required,min=1900,max=2023"`
	Director    string `json:"director" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
}

func (server *Server) CreateMovie(ctx *gin.Context) {
	var req createMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateMovieParams{
		Title:       req.Title,
		ReleaseYear: req.ReleaseYear,
		Director:    req.Director,
		Genre:       req.Genre,
	}
	movie, err := server.query.CreateMovie(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

type updateMovieRequest struct {
	ID          int64  `json:"id" binding:"required,min=1"`
	Title       string `json:"title" binding:"required"`
	ReleaseYear int32  `json:"release_year" binding:"required,min=1900,max=2023"`
	Director    string `json:"director" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
}

func (server *Server) UpdateMovie(ctx *gin.Context) {
	var req updateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.UpdateMovieParams{
		Title:       req.Title,
		ReleaseYear: req.ReleaseYear,
		Director:    req.Director,
		Genre:       req.Genre,
		ID:          req.ID,
	}
	err := server.query.UpdateMovie(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
