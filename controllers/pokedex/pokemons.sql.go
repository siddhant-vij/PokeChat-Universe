// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: pokemons.sql

package pokedex

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const deletePokemonByID = `-- name: DeletePokemonByID :exec
DELETE FROM pokemons
WHERE id = $1
`

func (q *Queries) DeletePokemonByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePokemonByID, id)
	return err
}

const getOnePokemonAfterCollection = `-- name: GetOnePokemonAfterCollection :one
SELECT id, name, picture_url
FROM pokemons
WHERE id > $1
ORDER BY id ASC
LIMIT 1
`

type GetOnePokemonAfterCollectionRow struct {
	ID         int32
	Name       string
	PictureUrl string
}

func (q *Queries) GetOnePokemonAfterCollection(ctx context.Context, id int32) (GetOnePokemonAfterCollectionRow, error) {
	row := q.db.QueryRowContext(ctx, getOnePokemonAfterCollection, id)
	var i GetOnePokemonAfterCollectionRow
	err := row.Scan(&i.ID, &i.Name, &i.PictureUrl)
	return i, err
}

const getPokemonByID = `-- name: GetPokemonByID :one
SELECT id, created_at, updated_at, name, height, weight, picture_url, base_experience, types, hp, attack, defense, special_attack, special_defense, speed FROM pokemons
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPokemonByID(ctx context.Context, id int32) (Pokemon, error) {
	row := q.db.QueryRowContext(ctx, getPokemonByID, id)
	var i Pokemon
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Height,
		&i.Weight,
		&i.PictureUrl,
		&i.BaseExperience,
		pq.Array(&i.Types),
		&i.Hp,
		&i.Attack,
		&i.Defense,
		&i.SpecialAttack,
		&i.SpecialDefense,
		&i.Speed,
	)
	return i, err
}

const getPokemonCount = `-- name: GetPokemonCount :one
SELECT COUNT(*) FROM pokemons
`

func (q *Queries) GetPokemonCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPokemonCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getPokemonDetailsByName = `-- name: GetPokemonDetailsByName :one
SELECT id, created_at, updated_at, name, height, weight, picture_url, base_experience, types, hp, attack, defense, special_attack, special_defense, speed FROM pokemons
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetPokemonDetailsByName(ctx context.Context, name string) (Pokemon, error) {
	row := q.db.QueryRowContext(ctx, getPokemonDetailsByName, name)
	var i Pokemon
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Height,
		&i.Weight,
		&i.PictureUrl,
		&i.BaseExperience,
		pq.Array(&i.Types),
		&i.Hp,
		&i.Attack,
		&i.Defense,
		&i.SpecialAttack,
		&i.SpecialDefense,
		&i.Speed,
	)
	return i, err
}

const getPokemonsSortedByIdAsc = `-- name: GetPokemonsSortedByIdAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
ORDER BY pokemons.id ASC
LIMIT $1 OFFSET $2
`

type GetPokemonsSortedByIdAscParams struct {
	Limit  int32
	Offset int32
}

type GetPokemonsSortedByIdAscRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetPokemonsSortedByIdAsc(ctx context.Context, arg GetPokemonsSortedByIdAscParams) ([]GetPokemonsSortedByIdAscRow, error) {
	rows, err := q.db.QueryContext(ctx, getPokemonsSortedByIdAsc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPokemonsSortedByIdAscRow
	for rows.Next() {
		var i GetPokemonsSortedByIdAscRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.PictureUrl,
			pq.Array(&i.Types),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPokemonsSortedByIdDesc = `-- name: GetPokemonsSortedByIdDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
ORDER BY pokemons.id DESC
LIMIT $1 OFFSET $2
`

type GetPokemonsSortedByIdDescParams struct {
	Limit  int32
	Offset int32
}

type GetPokemonsSortedByIdDescRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetPokemonsSortedByIdDesc(ctx context.Context, arg GetPokemonsSortedByIdDescParams) ([]GetPokemonsSortedByIdDescRow, error) {
	rows, err := q.db.QueryContext(ctx, getPokemonsSortedByIdDesc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPokemonsSortedByIdDescRow
	for rows.Next() {
		var i GetPokemonsSortedByIdDescRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.PictureUrl,
			pq.Array(&i.Types),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPokemonsSortedByNameAsc = `-- name: GetPokemonsSortedByNameAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
ORDER BY pokemons.name ASC
LIMIT $1 OFFSET $2
`

type GetPokemonsSortedByNameAscParams struct {
	Limit  int32
	Offset int32
}

type GetPokemonsSortedByNameAscRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetPokemonsSortedByNameAsc(ctx context.Context, arg GetPokemonsSortedByNameAscParams) ([]GetPokemonsSortedByNameAscRow, error) {
	rows, err := q.db.QueryContext(ctx, getPokemonsSortedByNameAsc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPokemonsSortedByNameAscRow
	for rows.Next() {
		var i GetPokemonsSortedByNameAscRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.PictureUrl,
			pq.Array(&i.Types),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPokemonsSortedByNameDesc = `-- name: GetPokemonsSortedByNameDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
ORDER BY pokemons.name DESC
LIMIT $1 OFFSET $2
`

type GetPokemonsSortedByNameDescParams struct {
	Limit  int32
	Offset int32
}

type GetPokemonsSortedByNameDescRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetPokemonsSortedByNameDesc(ctx context.Context, arg GetPokemonsSortedByNameDescParams) ([]GetPokemonsSortedByNameDescRow, error) {
	rows, err := q.db.QueryContext(ctx, getPokemonsSortedByNameDesc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPokemonsSortedByNameDescRow
	for rows.Next() {
		var i GetPokemonsSortedByNameDescRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.PictureUrl,
			pq.Array(&i.Types),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertPokemon = `-- name: InsertPokemon :exec
INSERT INTO pokemons
  (id, name, height, weight, picture_url, base_experience, types, hp, attack, defense, special_attack, special_defense, speed)
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
`

type InsertPokemonParams struct {
	ID             int32
	Name           string
	Height         int32
	Weight         int32
	PictureUrl     string
	BaseExperience int32
	Types          []string
	Hp             int32
	Attack         int32
	Defense        int32
	SpecialAttack  int32
	SpecialDefense int32
	Speed          int32
}

func (q *Queries) InsertPokemon(ctx context.Context, arg InsertPokemonParams) error {
	_, err := q.db.ExecContext(ctx, insertPokemon,
		arg.ID,
		arg.Name,
		arg.Height,
		arg.Weight,
		arg.PictureUrl,
		arg.BaseExperience,
		pq.Array(arg.Types),
		arg.Hp,
		arg.Attack,
		arg.Defense,
		arg.SpecialAttack,
		arg.SpecialDefense,
		arg.Speed,
	)
	return err
}

const updatePokemonByID = `-- name: UpdatePokemonByID :exec
UPDATE pokemons
SET
  name = $2,
  height = $3,
  weight = $4,
  picture_url = $5,
  base_experience = $6,
  types = $7,
  hp = $8,
  attack = $9,
  defense = $10,
  special_attack = $11,
  special_defense = $12,
  speed = $13,
  created_at = $14,
  updated_at = NOW()
WHERE id = $1
`

type UpdatePokemonByIDParams struct {
	ID             int32
	Name           string
	Height         int32
	Weight         int32
	PictureUrl     string
	BaseExperience int32
	Types          []string
	Hp             int32
	Attack         int32
	Defense        int32
	SpecialAttack  int32
	SpecialDefense int32
	Speed          int32
	CreatedAt      time.Time
}

func (q *Queries) UpdatePokemonByID(ctx context.Context, arg UpdatePokemonByIDParams) error {
	_, err := q.db.ExecContext(ctx, updatePokemonByID,
		arg.ID,
		arg.Name,
		arg.Height,
		arg.Weight,
		arg.PictureUrl,
		arg.BaseExperience,
		pq.Array(arg.Types),
		arg.Hp,
		arg.Attack,
		arg.Defense,
		arg.SpecialAttack,
		arg.SpecialDefense,
		arg.Speed,
		arg.CreatedAt,
	)
	return err
}
