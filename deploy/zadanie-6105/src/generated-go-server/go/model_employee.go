package openapi

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Username  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Organization struct {
	ID          uuid.UUID
	Name        string
	Description string
	Type        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
