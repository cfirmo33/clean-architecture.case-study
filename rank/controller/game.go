package controller

import (
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/entity"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/helper"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/repository"
)

type (
	// Controllers contains the Controllers for each Entity.
	Controllers struct {
		Game repository.Game
	}
	// Game contains the injected Game interface from Repository layer.
	Game struct {
		Repository repository.Game
	}
)

// New creates new Controllers for each Entity.
func New(repo *repository.MongoDB) *Controllers {
	return &Controllers{
		Game: newGameController(repo),
	}
}

// GameController contains methods that must be implemented by the injected layer.
type GameController interface {
	DeleteByID(helper.Identifier) error
	FindAll() ([]*entity.Game, error)
	FindByID(helper.Identifier) (*entity.Game, error)
	Store(*entity.Game) (helper.Identifier, error)
	Update(*entity.Game) error
}

// newGameController creates a new Game Controller.
func newGameController(m *repository.MongoDB) *Game {
	return &Game{
		Repository: m,
	}
}

// DeleteByID requests the Repository layer for a Game to be deleted from the database by its ID.
func (g *Game) DeleteByID(id helper.Identifier) error {
	return g.Repository.DeleteByID(id)
}

// FindAll requests the Repository layer to return all Games from database.
func (g *Game) FindAll() ([]*entity.Game, error) {
	return g.Repository.FindAll()
}

// FindByID requests the Repository layer for a certain Game by its ID.
func (g *Game) FindByID(id helper.Identifier) (*entity.Game, error) {
	return g.Repository.FindByID(id)
}

// Store requests the Repository layer for the insertion of a new Game in the database.
func (g *Game) Store(game *entity.Game) (helper.Identifier, error) {
	return g.Repository.Store(game)
}

// Update requests the Repository layer for a Game to be updated in the database.
func (g *Game) Update(game *entity.Game) error {
	return g.Repository.Update(game)
}
