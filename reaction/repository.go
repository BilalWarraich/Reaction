package reaction

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	PostReactionTypes(ctx context.Context, a ReactionTypes) error
	GetReactionTypeID(ctx context.Context, id string) (*ReactionTypes, error)
	GetReactionTypes(ctx context.Context) ([]ReactionTypes, error)
	PostReactions(ctx context.Context, a Reactions) error
	GetUserByPostID(ctx context.Context, postId string, reactType string) ([]Reactions, error)
	GetReactByUserID(ctx context.Context, postId string) ([]Reactions, error)
	DeleteReactionByPostID(ctx context.Context, postId string) error
	DeleteReactionByUserPostIDRequest(ctx context.Context, postId string, userId string) error
	UpdateReactions(ctx context.Context, a Reactions) error
	TotalReactionCount(ctx context.Context, postId string) (string, error)
	ReactionCountByType(ctx context.Context, postId string, reactType string) (string, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) PostReactionTypes(ctx context.Context, a ReactionTypes) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO reactionTypes(id, reacts) VALUES($1, $2)", a.ID, a.Reacts)
	return err
}

func (r *postgresRepository) GetReactionTypeID(ctx context.Context, id string) (*ReactionTypes, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, reacts FROM reactionTypes WHERE id = $1", id)
	a := &ReactionTypes{}
	if err := row.Scan(&a.ID, &a.Reacts); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *postgresRepository) GetReactionTypes(ctx context.Context) ([]ReactionTypes, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, reacts FROM reactionTypes ",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reactionTypes := []ReactionTypes{}
	for rows.Next() {
		a := &ReactionTypes{}
		if err = rows.Scan(&a.ID, &a.Reacts); err == nil {
			reactionTypes = append(reactionTypes, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return reactionTypes, nil
}

func (r *postgresRepository) PostReactions(ctx context.Context, a Reactions) error {

	var id string
	rows := r.db.QueryRow("SELECT id FROM reactions WHERE postId=$1 AND userId=$2 ", a.PostId, a.UserId).Scan(&id)
	if rows != nil && id == "" {
		_, err := r.db.ExecContext(ctx, "INSERT INTO reactions(id, reactionType,postId,userId) VALUES($1, $2, $3, $4)", a.ID, a.ReactType, a.PostId, a.UserId)

		return err
	}

	var RepoErr1 = errors.New("Reaction Already Added")
	return RepoErr1

}

func (r *postgresRepository) GetUserByPostID(ctx context.Context, postId string, reactType string) ([]Reactions, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, userId FROM reactions WHERE postId = $1 AND reactionType = $2", postId, reactType)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reactionTypes := []Reactions{}
	for rows.Next() {
		a := &Reactions{}
		if err = rows.Scan(&a.ID, &a.UserId); err == nil {
			reactionTypes = append(reactionTypes, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return reactionTypes, nil

}

func (r *postgresRepository) GetReactByUserID(ctx context.Context, userId string) ([]Reactions, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, postId, reactionType FROM reactions WHERE userId = $1", userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reactionTypes := []Reactions{}
	for rows.Next() {
		a := &Reactions{}
		if err = rows.Scan(&a.ID, &a.PostId, &a.ReactType); err == nil {
			reactionTypes = append(reactionTypes, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return reactionTypes, nil

}

func (r *postgresRepository) DeleteReactionByPostID(ctx context.Context, postId string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM reactions WHERE postId=$1", postId)
	return err
}

func (r *postgresRepository) DeleteReactionByUserPostIDRequest(ctx context.Context, postId string, userId string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM reactions WHERE postId=$1 AND userId = $2", postId, userId)
	return err
}

func (r *postgresRepository) UpdateReactions(ctx context.Context, a Reactions) error {
	_, err := r.db.ExecContext(ctx, "UPDATE reactions SET reactionType=$1 WHERE postId=$2 AND userId=$3", a.ReactType, a.PostId, a.UserId)
	return err
}

func (r *postgresRepository) TotalReactionCount(ctx context.Context, postId string) (string, error) {
	var count string
	stmt, err := r.db.Prepare("SELECT COUNT(*) as count FROM reactions WHERE postId=$1")
	if err != nil {
		return "", err
	}
	err = stmt.QueryRow(postId).Scan(&count)
	if err != nil {
		return "", err
	}
	stmt.Close()

	return count, nil

}

func (r *postgresRepository) ReactionCountByType(ctx context.Context, postId string, reactType string) (string, error) {
	var count string
	stmt, err := r.db.Prepare("SELECT COUNT(*) as count FROM reactions WHERE postId=$1 AND reactionType=$2")
	if err != nil {
		return "", err
	}
	err = stmt.QueryRow(postId, reactType).Scan(&count)
	if err != nil {
		return "", err
	}
	stmt.Close()

	return count, nil

}
