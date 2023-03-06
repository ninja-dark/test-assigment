package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ninja-dark/test-assigment/internal/entity"
)

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) entity.Repository {
	return &pgRepository{pool: pool}
}

func (r *pgRepository) GetList(ctx context.Context) ([]entity.Song, error) {
	rows, _ := r.pool.Query(ctx, `SELECT name, duration FROM playlist`)

	pl := make([]entity.Song, 0)
	for rows.Next(){
		var song entity.Song
		if err := rows.Scan(&song.Name, &song.Duration); err != nil {
			return nil, err
		}

		pl = append(pl, song)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return pl, nil
}

func (r *pgRepository) Add(ctx context.Context, s *entity.Song) (*entity.Song,error) {
	newSong := *s
	_, err := r.pool.Exec(
		ctx, `INSERT INTO playlist (
			name, duration
		) VALUES (
			&2, -- name
			&3, -- duration
		)
	`, newSong.Name, newSong.Duration)
	return &newSong, err
}

func (r *pgRepository) Update(ctx context.Context, s *entity.Song) (int64, error){

	row:= r.pool.QueryRow(ctx, `UPDATE playlist SET name = &1, duration = &2 WERE ID = &3 RETURNING id`, s.Name, s.Duration, s.ID)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *pgRepository) Delete(ctx context.Context, id int64) error{
	_, err := r.pool.Exec(ctx, `DELETE FROM playlist WERE id = &1`, id)
	if err != nil {
		return err
	}
	return nil
}
