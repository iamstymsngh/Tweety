package service

import (
	"database/sql"
	"github.com/hako/branca"
)

// Service handles business logic
type Service struct {
	db *sql.DB
	codec *branca.Branca
}