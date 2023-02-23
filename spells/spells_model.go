package spells

import "go.mongodb.org/mongo-driver/bson/primitive"

type Spell struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name,omitempty" validate:"required"`
	Level           int                `json:"level,omitempty" validate:"required"`
	School          string             `json:"school,omitempty" validate:"required"`
	ConjurationTime string             `json:"conjurationTime,omitempty" validate:"required"`
	SpellRange      string             `json:"spellRange,omitempty" validate:"required"`
	Components      []string           `json:"components,omitempty" validate:"required"`
	Duration        string             `json:"duration,omitempty" validate:"required"`
	Description     string             `json:"description,omitempty" validate:"required"`
	Source          string             `json:"source,omitempty" validate:"required"`
}
