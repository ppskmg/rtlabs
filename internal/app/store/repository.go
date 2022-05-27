package store

import (
	_ "database/sql"
	"rtlabs/internal/app/model/vote"
)

type VoteRepository interface {
	Create(c *vote.Vote) error
	ReadMyVote(userEmail string, contest string) ([]vote.Vote, error)
	Delete(workID int64, userEmail string) (int64, error)
	ReadCountVote() ([]vote.Count, error)
	Winner() ([]vote.Winner, error)
}
