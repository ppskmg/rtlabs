package postgresstore_test

import (
	"fmt"
	"rtlabs/internal/app/model/vote"
	"rtlabs/internal/app/store"
	"rtlabs/internal/app/store/postgresstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVoteRepository_Create(t *testing.T) {
	db, teardown := postgresstore.TestDB(t, databaseURL)
	defer teardown(store.STables.Votes)
	s := postgresstore.New(db)
	ct := vote.TestVote(t)
	err := s.Vote().Create(ct)
	assert.NoError(t, err)
}

func TestVoteRepository_ReadCountVote(t *testing.T) {
	db, teardown := postgresstore.TestDB(t, databaseURL)
	defer teardown(store.STables.Votes)
	s := postgresstore.New(db)
	//ct := vote.TestVote(t)
	c, err := s.Vote().ReadCountVote()
	fmt.Println(c)
	assert.NoError(t, err)
}
