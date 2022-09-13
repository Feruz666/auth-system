// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"

	"github.com/google/uuid"
)

type Layer struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Layer    string `json:"layer"`
}

type Session struct {
	SessionID    uuid.UUID `json:"session_id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	Organization      string    `json:"organization"`
}
