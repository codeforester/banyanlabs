package storage

import (
	"context"
	"errors"
	"time"
)

var (
	ErrConflict = errors.New("storage conflict")
	ErrNotFound = errors.New("storage record not found")
)

type User struct {
	ID           int64
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	ModifiedAt   time.Time
}

type CreateUserParams struct {
	Username     string
	Email        string
	PasswordHash string
	Now          time.Time
}

type Session struct {
	ID         int64
	UserID     int64
	TokenHash  string
	CreatedAt  time.Time
	ExpiresAt  time.Time
	LastSeenAt *time.Time
}

type CreateSessionParams struct {
	UserID    int64
	TokenHash string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Store interface {
	CreateUser(context.Context, CreateUserParams) (User, error)
	FindUserByUsername(context.Context, string) (User, error)
	CreateSession(context.Context, CreateSessionParams) (Session, error)
	DeleteSessionByTokenHash(context.Context, string) error
}
