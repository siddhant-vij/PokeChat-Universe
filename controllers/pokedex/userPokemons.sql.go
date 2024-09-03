// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: userPokemons.sql

package pokedex

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getOneAvailablePokemonAfterCollectionByIdAsc = `-- name: GetOneAvailablePokemonAfterCollectionByIdAsc :one
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.id > $2
ORDER BY pokemons.id ASC
LIMIT 1
`

type GetOneAvailablePokemonAfterCollectionByIdAscParams struct {
	UserID uuid.UUID
	ID     int32
}

type GetOneAvailablePokemonAfterCollectionByIdAscRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetOneAvailablePokemonAfterCollectionByIdAsc(ctx context.Context, arg GetOneAvailablePokemonAfterCollectionByIdAscParams) (GetOneAvailablePokemonAfterCollectionByIdAscRow, error) {
	row := q.db.QueryRowContext(ctx, getOneAvailablePokemonAfterCollectionByIdAsc, arg.UserID, arg.ID)
	var i GetOneAvailablePokemonAfterCollectionByIdAscRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PictureUrl,
		pq.Array(&i.Types),
	)
	return i, err
}

const getOneAvailablePokemonAfterCollectionByIdDesc = `-- name: GetOneAvailablePokemonAfterCollectionByIdDesc :one
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.id < $2
ORDER BY pokemons.id DESC
LIMIT 1
`

type GetOneAvailablePokemonAfterCollectionByIdDescParams struct {
	UserID uuid.UUID
	ID     int32
}

type GetOneAvailablePokemonAfterCollectionByIdDescRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetOneAvailablePokemonAfterCollectionByIdDesc(ctx context.Context, arg GetOneAvailablePokemonAfterCollectionByIdDescParams) (GetOneAvailablePokemonAfterCollectionByIdDescRow, error) {
	row := q.db.QueryRowContext(ctx, getOneAvailablePokemonAfterCollectionByIdDesc, arg.UserID, arg.ID)
	var i GetOneAvailablePokemonAfterCollectionByIdDescRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PictureUrl,
		pq.Array(&i.Types),
	)
	return i, err
}

const getOneAvailablePokemonAfterCollectionByNameAsc = `-- name: GetOneAvailablePokemonAfterCollectionByNameAsc :one
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.name > $2
ORDER BY pokemons.name ASC
LIMIT 1
`

type GetOneAvailablePokemonAfterCollectionByNameAscParams struct {
	UserID uuid.UUID
	Name   string
}

type GetOneAvailablePokemonAfterCollectionByNameAscRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetOneAvailablePokemonAfterCollectionByNameAsc(ctx context.Context, arg GetOneAvailablePokemonAfterCollectionByNameAscParams) (GetOneAvailablePokemonAfterCollectionByNameAscRow, error) {
	row := q.db.QueryRowContext(ctx, getOneAvailablePokemonAfterCollectionByNameAsc, arg.UserID, arg.Name)
	var i GetOneAvailablePokemonAfterCollectionByNameAscRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PictureUrl,
		pq.Array(&i.Types),
	)
	return i, err
}

const getOneAvailablePokemonAfterCollectionByNameDesc = `-- name: GetOneAvailablePokemonAfterCollectionByNameDesc :one
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.name < $2
ORDER BY pokemons.name DESC
LIMIT 1
`

type GetOneAvailablePokemonAfterCollectionByNameDescParams struct {
	UserID uuid.UUID
	Name   string
}

type GetOneAvailablePokemonAfterCollectionByNameDescRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetOneAvailablePokemonAfterCollectionByNameDesc(ctx context.Context, arg GetOneAvailablePokemonAfterCollectionByNameDescParams) (GetOneAvailablePokemonAfterCollectionByNameDescRow, error) {
	row := q.db.QueryRowContext(ctx, getOneAvailablePokemonAfterCollectionByNameDesc, arg.UserID, arg.Name)
	var i GetOneAvailablePokemonAfterCollectionByNameDescRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PictureUrl,
		pq.Array(&i.Types),
	)
	return i, err
}

const getUserAvailablePokemonsSortedByIdAsc = `-- name: GetUserAvailablePokemonsSortedByIdAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
ORDER BY pokemons.id ASC
LIMIT $2 OFFSET $3
`

type GetUserAvailablePokemonsSortedByIdAscParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetUserAvailablePokemonsSortedByIdAscRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetUserAvailablePokemonsSortedByIdAsc(ctx context.Context, arg GetUserAvailablePokemonsSortedByIdAscParams) ([]GetUserAvailablePokemonsSortedByIdAscRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserAvailablePokemonsSortedByIdAsc, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserAvailablePokemonsSortedByIdAscRow
	for rows.Next() {
		var i GetUserAvailablePokemonsSortedByIdAscRow
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

const getUserAvailablePokemonsSortedByIdDesc = `-- name: GetUserAvailablePokemonsSortedByIdDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
ORDER BY pokemons.id DESC
LIMIT $2 OFFSET $3
`

type GetUserAvailablePokemonsSortedByIdDescParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetUserAvailablePokemonsSortedByIdDescRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetUserAvailablePokemonsSortedByIdDesc(ctx context.Context, arg GetUserAvailablePokemonsSortedByIdDescParams) ([]GetUserAvailablePokemonsSortedByIdDescRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserAvailablePokemonsSortedByIdDesc, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserAvailablePokemonsSortedByIdDescRow
	for rows.Next() {
		var i GetUserAvailablePokemonsSortedByIdDescRow
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

const getUserAvailablePokemonsSortedByNameAsc = `-- name: GetUserAvailablePokemonsSortedByNameAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
ORDER BY pokemons.name ASC
LIMIT $2 OFFSET $3
`

type GetUserAvailablePokemonsSortedByNameAscParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetUserAvailablePokemonsSortedByNameAscRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetUserAvailablePokemonsSortedByNameAsc(ctx context.Context, arg GetUserAvailablePokemonsSortedByNameAscParams) ([]GetUserAvailablePokemonsSortedByNameAscRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserAvailablePokemonsSortedByNameAsc, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserAvailablePokemonsSortedByNameAscRow
	for rows.Next() {
		var i GetUserAvailablePokemonsSortedByNameAscRow
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

const getUserAvailablePokemonsSortedByNameDesc = `-- name: GetUserAvailablePokemonsSortedByNameDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
ORDER BY pokemons.name DESC
LIMIT $2 OFFSET $3
`

type GetUserAvailablePokemonsSortedByNameDescParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetUserAvailablePokemonsSortedByNameDescRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetUserAvailablePokemonsSortedByNameDesc(ctx context.Context, arg GetUserAvailablePokemonsSortedByNameDescParams) ([]GetUserAvailablePokemonsSortedByNameDescRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserAvailablePokemonsSortedByNameDesc, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserAvailablePokemonsSortedByNameDescRow
	for rows.Next() {
		var i GetUserAvailablePokemonsSortedByNameDescRow
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

const getUserCollectedPokemonsSortedByIdAsc = `-- name: GetUserCollectedPokemonsSortedByIdAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NOT NULL
ORDER BY pokemons.id ASC
LIMIT $2 OFFSET $3
`

type GetUserCollectedPokemonsSortedByIdAscParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetUserCollectedPokemonsSortedByIdAscRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetUserCollectedPokemonsSortedByIdAsc(ctx context.Context, arg GetUserCollectedPokemonsSortedByIdAscParams) ([]GetUserCollectedPokemonsSortedByIdAscRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCollectedPokemonsSortedByIdAsc, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCollectedPokemonsSortedByIdAscRow
	for rows.Next() {
		var i GetUserCollectedPokemonsSortedByIdAscRow
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

const getUserCollectedPokemonsSortedByIdDesc = `-- name: GetUserCollectedPokemonsSortedByIdDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NOT NULL
ORDER BY pokemons.id DESC
LIMIT $2 OFFSET $3
`

type GetUserCollectedPokemonsSortedByIdDescParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetUserCollectedPokemonsSortedByIdDescRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetUserCollectedPokemonsSortedByIdDesc(ctx context.Context, arg GetUserCollectedPokemonsSortedByIdDescParams) ([]GetUserCollectedPokemonsSortedByIdDescRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCollectedPokemonsSortedByIdDesc, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCollectedPokemonsSortedByIdDescRow
	for rows.Next() {
		var i GetUserCollectedPokemonsSortedByIdDescRow
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

const getUserCollectedPokemonsSortedByNameAsc = `-- name: GetUserCollectedPokemonsSortedByNameAsc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NOT NULL
ORDER BY pokemons.name ASC
LIMIT $2 OFFSET $3
`

type GetUserCollectedPokemonsSortedByNameAscParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetUserCollectedPokemonsSortedByNameAscRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetUserCollectedPokemonsSortedByNameAsc(ctx context.Context, arg GetUserCollectedPokemonsSortedByNameAscParams) ([]GetUserCollectedPokemonsSortedByNameAscRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCollectedPokemonsSortedByNameAsc, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCollectedPokemonsSortedByNameAscRow
	for rows.Next() {
		var i GetUserCollectedPokemonsSortedByNameAscRow
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

const getUserCollectedPokemonsSortedByNameDesc = `-- name: GetUserCollectedPokemonsSortedByNameDesc :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NOT NULL
ORDER BY pokemons.name DESC
LIMIT $2 OFFSET $3
`

type GetUserCollectedPokemonsSortedByNameDescParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetUserCollectedPokemonsSortedByNameDescRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) GetUserCollectedPokemonsSortedByNameDesc(ctx context.Context, arg GetUserCollectedPokemonsSortedByNameDescParams) ([]GetUserCollectedPokemonsSortedByNameDescRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCollectedPokemonsSortedByNameDesc, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCollectedPokemonsSortedByNameDescRow
	for rows.Next() {
		var i GetUserCollectedPokemonsSortedByNameDescRow
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

const insertUserCollectedPokemon = `-- name: InsertUserCollectedPokemon :exec
INSERT INTO user_pokemons
  (id, user_id, pokemon_id)
VALUES
  ($1, $2, $3)
`

type InsertUserCollectedPokemonParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PokemonID int32
}

func (q *Queries) InsertUserCollectedPokemon(ctx context.Context, arg InsertUserCollectedPokemonParams) error {
	_, err := q.db.ExecContext(ctx, insertUserCollectedPokemon, arg.ID, arg.UserID, arg.PokemonID)
	return err
}

const isPokemonCollected = `-- name: IsPokemonCollected :one
SELECT
  EXISTS (
    SELECT 1
    FROM user_pokemons
    WHERE user_id = $1 AND pokemon_id = $2
  )
`

type IsPokemonCollectedParams struct {
	UserID    uuid.UUID
	PokemonID int32
}

func (q *Queries) IsPokemonCollected(ctx context.Context, arg IsPokemonCollectedParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isPokemonCollected, arg.UserID, arg.PokemonID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const searchUserAvailablePokemonsByName = `-- name: SearchUserAvailablePokemonsByName :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
LEFT JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE user_pokemons.id IS NULL
AND pokemons.name ILIKE $2
ORDER BY pokemons.name ASC
LIMIT $3
`

type SearchUserAvailablePokemonsByNameParams struct {
	UserID uuid.UUID
	Name   string
	Limit  int32
}

type SearchUserAvailablePokemonsByNameRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) SearchUserAvailablePokemonsByName(ctx context.Context, arg SearchUserAvailablePokemonsByNameParams) ([]SearchUserAvailablePokemonsByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, searchUserAvailablePokemonsByName, arg.UserID, arg.Name, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchUserAvailablePokemonsByNameRow
	for rows.Next() {
		var i SearchUserAvailablePokemonsByNameRow
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

const searchUserCollectedPokemonsByName = `-- name: SearchUserCollectedPokemonsByName :many
SELECT
  pokemons.id,
  pokemons.name,
  pokemons.picture_url,
  pokemons.types
FROM pokemons
INNER JOIN user_pokemons
ON user_pokemons.pokemon_id = pokemons.id
AND user_pokemons.user_id = $1
WHERE pokemons.name ILIKE $2
ORDER BY pokemons.name ASC
LIMIT $3
`

type SearchUserCollectedPokemonsByNameParams struct {
	UserID uuid.UUID
	Name   string
	Limit  int32
}

type SearchUserCollectedPokemonsByNameRow struct {
	ID         int32
	Name       string
	PictureUrl string
	Types      []string
}

func (q *Queries) SearchUserCollectedPokemonsByName(ctx context.Context, arg SearchUserCollectedPokemonsByNameParams) ([]SearchUserCollectedPokemonsByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, searchUserCollectedPokemonsByName, arg.UserID, arg.Name, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchUserCollectedPokemonsByNameRow
	for rows.Next() {
		var i SearchUserCollectedPokemonsByNameRow
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
