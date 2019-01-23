package repository

import (
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/entity"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/helper"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/middlewares/config"
)

// Game defines the methods must be implemented by injected layer.
type Game interface {
	DeleteByID(helper.Identifier) error
	FindAll() ([]*entity.Game, error)
	FindByID(helper.Identifier) (*entity.Game, error)
	Store(*entity.Game) (helper.Identifier, error)
	Update(*entity.Game) error
}

// DeleteByID deletes a Game by its ID.
func (m *MongoDB) DeleteByID(id helper.Identifier) error {
	return m.pool.Session(nil).DB(m.db).C(config.GameCollection).RemoveId(id)
}

// FindAll returns all Games from the database sorted by ID.
func (m *MongoDB) FindAll() ([]*entity.Game, error) {
	var games []*entity.Game

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GameCollection)
	if err := collection.Find(nil).Sort("id").All(&games); err != nil {
		return nil, err
	}

	return games, nil
}

// FindByID finds a Game by its ID.
func (m *MongoDB) FindByID(id helper.Identifier) (*entity.Game, error) {
	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GameCollection)

	var game *entity.Game

	collection.FindId(id).One(&game)

	return game, nil
}

// Store inserts a new Game in the database.
func (m *MongoDB) Store(game *entity.Game) (helper.Identifier, error) {
	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GameCollection)

	game.ID = helper.NewID()

	collection.Insert(game)

	return game.ID, nil
}

// Update updates an existing Game in the database.
func (m *MongoDB) Update(game *entity.Game) error {
	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GameCollection)

	_, err := collection.UpsertId(game.ID, game) // TODO - avoid null Games
	return err
}
