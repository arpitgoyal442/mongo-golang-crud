package model

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Netflix struct{

// 	ID  primitive.ObjectID  `json:"_id" bson:"_id"`
// 	Movie string              `json:"movie"`
// 	Watched bool              `json:"watched"`

// }

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty bson:"watched,omitempty"`
}