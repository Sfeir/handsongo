package model

import (
	"gopkg.in/mgo.v2/bson"
)

// SpiritType is the type of spirit as typed enumeration
type SpiritType int

const (
	// TypeRhum is the constant used for rhum spirits = 0
	TypeRhum SpiritType = iota
	// TypeWine is the constant used for wine
	TypeWine
	// TypeWhisky is the constant used for whisky
	TypeWhisky
	// TypeChampagne is the constant used for champagne
	TypeChampagne
	// TypeGin is the constant used for gin
	TypeGin
)

// Spirit is the structure to define a spirit
type Spirit struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty" `
	Name        string        `json:"name" bson:"name"`
	Distiller   string        `json:"distiller" bson:"distiller"`
	Bottler     string        `json:"bottler" bson:"bottler"`
	Country     string        `json:"country" bson:"country"`
	Region      string        `json:"region" bson:"region"`
	Composition string        `json:"composition" bson:"composition"`
	SpiritType  SpiritType    `json:"type" bson:"type"`
	// TODO add an uint8 Age attribute
	// TODO add a time BottlingDate attribute
	Score   float32 `json:"score" bson:"score"`
	Comment string  `json:"comment" bson:"comment"`
}

// GetID returns the ID of an Spirit as a string
func (s *Spirit) GetID() string {
	return s.ID.Hex()
}

// Equal returns true if both spirit are equal, false otherwise
func (s *Spirit) Equal(cmp Spirit) bool {
	equal :=
		s.ID == cmp.ID &&
			s.Name == cmp.Name &&
			s.Distiller == cmp.Distiller &&
			s.Bottler == cmp.Bottler &&
			s.Country == cmp.Country &&
			s.Region == cmp.Region &&
			s.Composition == cmp.Composition &&
			s.SpiritType == cmp.SpiritType &&
			// TODO add Age and BottlingDate Comparison
			s.Score == cmp.Score &&
			s.Comment == cmp.Comment

	return equal
}
