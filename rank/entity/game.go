package entity

import (
	"time"

	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/helper"
)

// Game represents the game entity and its attributes.
type Game struct {
	ID           helper.Identifier `bson:"_id,omitempty" json:"id"`
	Name         string            `bson:"name" json:"name"`
	Platforms    []string          `bson:"platforms" json:"platforms"`
	Categories   []string          `bson:"categories" json:"categories"`
	ReleaseDate  time.Time         `bson:"release_date" json:"release_date"`   // Game's release date
	PublicatedAt time.Time         `bson:"publicated_at" json:"publicated_at"` // Game page's publication date
	Rating       string            `bson:"rating" json:"rating"`
	Score        int               `bson:"score" json:"score"`
	Publisher    string            `bson:"publisher" json:"publisher"`
	CoverImage   string            `bson:"cover_image" json:"cover_image"`
}
