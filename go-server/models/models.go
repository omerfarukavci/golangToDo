package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoList struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task         string             `json:"task,omitempty"`
	Status       bool               `json:"status,omitempty"`
	Description  string             `json:"description,omitempty"`
	Category     string             `json:"category,omitempty"`
	Progress     string             `json:"progress,omitempty"`
	Deadline     time.Time          `json:"deadline,omitempty"`
	CreatedAt    time.Time          `json:"createdat,omitempty"`
	UpdatedAt    time.Time          `json:"updatedat,omitempty"`
	RemainingDay uint               `json:"remainingday,omitempty"`
}
