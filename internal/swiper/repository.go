package swiper

import (
	"context"
	"log/slog"

	"github.com/huandu/go-sqlbuilder"
	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
)

type SwipeRepo struct {
	*common.Dependencies
}

func NewSwipeRepo(deps *common.Dependencies) *SwipeRepo {
	return &SwipeRepo{
		Dependencies: deps,
	}
}

func (r *SwipeRepo) GetUserLikedUserIDs(ctx context.Context, userID int, params models.LikedUserParams) ([]int, error) {
	q := sqlbuilder.Select(`target_id`).From(`swipes s`)
	q.Where(q.Equal(`s.user_id`, userID))

	if params.IsMatched {
		q.Where(q.Equal(`s.is_matched`, true))
	}

	if params.FromDate != "" {
		q.Where(q.GreaterEqualThan(`s.created_at`, params.FromDate))
	}

	if params.ToDate != "" {
		q.Where(q.LessEqualThan(`s.created_at`, params.ToDate))
	}

	query, args := q.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := r.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		return []int{}, err
	}

	var result []int
	var temp int
	for rows.Next() {
		err := rows.Scan(&temp)
		if err != nil {
			slog.WarnContext(ctx, "failed to scan target id", slog.Any("error", err))
		}
		result = append(result, temp)
	}

	return result, nil
}

func (r *SwipeRepo) StoreUserLike(ctx context.Context, userID int, targetID int) error {
	q := `INSERT INTO swipes (user_id, target_id, created_at) VALUES ($1, $2, now())`

	_, err := r.DB.ExecContext(ctx, q, userID, targetID)
	if err != nil {
		return err
	}

	return nil
}

func (r *SwipeRepo) UpdateMatchStatus(ctx context.Context, userID int, targetID int, status bool) error {
	q := `UPDATE swipes
			SET
				is_matched = $1,
				updated_at = now()
			WHERE
				(user_id = $2 and target_id = $3)
				OR
				(user_id = $3 and target_id = $2);`
	_, err := r.DB.Exec(q, status, userID, targetID)
	if err != nil {
		return err
	}

	return nil
}
