// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
)

type Querier interface {
	CreateActor(ctx context.Context, arg CreateActorParams) (int64, error)
	CreateAward(ctx context.Context, arg CreateAwardParams) (int64, error)
	CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error)
	CreateNomination(ctx context.Context, arg CreateNominationParams) (int64, error)
	CreatePerformance(ctx context.Context, arg CreatePerformanceParams) (Performance, error)
	DeleteActor(ctx context.Context, id int64) error
	//
	DeleteAward(ctx context.Context, id int64) error
	DeleteMovie(ctx context.Context, id int64) error
	//
	DeleteNomination(ctx context.Context, id int64) error
	GetActor(ctx context.Context, id int64) (Actor, error)
	//
	GetAward(ctx context.Context, id int64) (Award, error)
	GetMovie(ctx context.Context, id int64) (Movie, error)
	//
	GetNomination(ctx context.Context, id int64) (Nomination, error)
	GetPerformance(ctx context.Context, id int64) (Performance, error)
	ListActors(ctx context.Context, arg ListActorsParams) ([]Actor, error)
	//
	ListAwards(ctx context.Context, arg ListAwardsParams) ([]Award, error)
	ListMovies(ctx context.Context, arg ListMoviesParams) ([]Movie, error)
	//
	ListNominations(ctx context.Context, arg ListNominationsParams) ([]Nomination, error)
	UpdateActor(ctx context.Context, arg UpdateActorParams) error
	//
	UpdateAward(ctx context.Context, arg UpdateAwardParams) error
	UpdateMovie(ctx context.Context, arg UpdateMovieParams) error
	//
	UpdateNomination(ctx context.Context, arg UpdateNominationParams) error
}

var _ Querier = (*Queries)(nil)