package dao

import (
	"errors"
	"github.com/Sfeir/handsongo/model"
)

const (
	// NoPaging used with skip, limit parameters
	NoPaging = -1
)

var (
	// ErrInvalidObjectID is used on invalid ObjectId
	ErrInvalidObjectID = errors.New("invalid input to ObjectIdHex")
)

// SpiritDAO is the DAO interface to work with spirits
type SpiritDAO interface {

	// GetSpiritByID returns a spirit by its ID
	GetSpiritByID(ID string) (*model.Spirit, error)

	// GetAllSpirits returns all spirits with paging capability
	GetAllSpirits(start, end int) ([]model.Spirit, error)

	// GetSpiritsByName returns all spirits by name
	GetSpiritsByName(name string) ([]model.Spirit, error)

	// GetSpiritsByType returns all spirits by type
	GetSpiritsByType(spiritType string) ([]model.Spirit, error)

	// GetSpiritsByTypeAndScore returns all spirits by type and score greater than parameter
	GetSpiritsByTypeAndScore(spiritType string, score uint8) ([]model.Spirit, error)

	// SaveSpirit saves the spirit
	SaveSpirit(spirit *model.Spirit) error

	// UpsertSpirit updates or creates a spirit
	UpsertSpirit(ID string, spirit *model.Spirit) (bool, error)

	// DeleteSpirit deletes a spirits by its ID
	DeleteSpirit(ID string) error
}
