package vote

import (
	"testing"
)

func TestVote(t *testing.T) *Vote {
	return &Vote{
		WorkID:    12313,
		UserEmail: "test@test.ru",
		Contest:   "Кулинария",
	}
}
