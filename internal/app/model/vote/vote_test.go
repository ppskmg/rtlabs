package vote_test

import (
	"github.com/stretchr/testify/assert"
	"rtlabs/internal/app/model/vote"
	"testing"
)

func TestSpecialization_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *vote.Vote
		isValid bool
	}{
		{
			name: "valid",
			s: func() *vote.Vote {
				return vote.TestVote(t)
			},
			isValid: true,
		},
		{
			name: "empty workID",
			s: func() *vote.Vote {
				s := vote.TestVote(t)
				s.WorkID = 0
				return s
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}
