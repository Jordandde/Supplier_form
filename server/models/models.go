package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SupplierList struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Supplier string             `json:"supplier,omitempty"`
	Address  string             `json:"address,omitempty"`
	Logo     string             `json:"logo,omitempty"`
}
