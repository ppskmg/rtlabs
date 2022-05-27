package store

type Store interface {
	Vote() VoteRepository
}

type Tables struct {
	Votes  string
	Winner string
}

var (
	STables = &Tables{
		Votes:  "work_votes",
		Winner: "winner",
	}
)
