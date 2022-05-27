package postgresstore

import (
	"fmt"
	"rtlabs/internal/app/model/vote"
	"rtlabs/internal/app/store"
)

type VoteRepository struct {
	store *Store
}

func (r *VoteRepository) Create(c *vote.Vote) error {
	if err := c.Validate(); err != nil {

		return err
	}
	if err := r.store.db.QueryRow(
		"INSERT INTO work_votes (work_id, user_email, contest) VALUES ($1, $2, $3) RETURNING id",
		c.WorkID, c.UserEmail, c.Contest).Scan(&c.WorkID); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *VoteRepository) Delete(workID int64, userEmail string) (int64, error) {
	sqlExpert := `DELETE FROM work_votes WHERE work_id = $1 AND user_email = $2 RETURNING work_id`
	result, err := r.store.db.Exec(sqlExpert, workID, userEmail)
	var dr int64 = 0
	if result != nil {
		dr, err = result.RowsAffected()
		if err == nil && dr == 0 {
			err = store.ErrRecordNotFound
		}
	}
	if err != nil {
		return dr, err
	}
	return dr, nil
}

func (r *VoteRepository) ReadCountVote() ([]vote.Count, error) {
	SQL := `SELECT work_id, count(work_id) FROM work_votes GROUP BY work_id`
	rows, err := r.store.db.Query(SQL)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var votes []vote.Count
	for rows.Next() {
		var e vote.Count
		if err := rows.Scan(
			&e.WorkID,
			&e.Count,
		); err != nil {
			return nil, err
		}
		votes = append(votes, e)
	}
	return votes, nil
}

func (r *VoteRepository) ReadMyVote(userEmail string, contest string) ([]vote.Vote, error) {
	limit := 6
	SQL := `SELECT work_id, user_email, contest FROM work_votes WHERE user_email = $2 AND contest = $3 LIMIT $1`
	rows, err := r.store.db.Query(SQL,
		limit,
		userEmail,
		contest,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var experts []vote.Vote
	for rows.Next() {
		var e vote.Vote
		if err := rows.Scan(
			&e.WorkID,
			&e.UserEmail,
			&e.Contest,
		); err != nil {
			return nil, err
		}
		experts = append(experts, e)
	}
	return experts, nil
}
