package postgresstore

import (
	"database/sql"
)

type Store struct {
	db             *sql.DB
	voteRepository *VoteRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Vote() *VoteRepository {
	if s.voteRepository != nil {
		return s.voteRepository
	}
	s.voteRepository = &VoteRepository{
		store: s,
	}

	return s.voteRepository
}
