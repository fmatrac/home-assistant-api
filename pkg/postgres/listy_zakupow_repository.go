package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type listyZakupowRepo struct {
	db DB
}

func NewListyZakupowRepository(db DB) ListyZakupowRepository {
	return &listyZakupowRepo{db: db}
}

func (r *listyZakupowRepo) Create(ctx context.Context, lista *ListaZakupow) error {
	query := `
		INSERT INTO home_assistant.listy_zakupow 
		(nazwa, status)
		VALUES ($1, $2)
		RETURNING id, utworzono
	`
	return r.db.QueryRowContext(
		ctx, query,
		lista.Nazwa,
		lista.Status,
	).Scan(&lista.ID, &lista.Utworzono)
}

func (r *listyZakupowRepo) FindByID(ctx context.Context, id int) (*ListaZakupow, error) {
	query := `
		SELECT id, nazwa, status, utworzono
		FROM home_assistant.listy_zakupow
		WHERE id = $1
	`
	lista := &ListaZakupow{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&lista.ID,
		&lista.Nazwa,
		&lista.Status,
		&lista.Utworzono,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("lista zakupow not found")
	}
	return lista, err
}

func (r *listyZakupowRepo) FindAll(ctx context.Context) ([]*ListaZakupow, error) {
	query := `
		SELECT id, nazwa, status, utworzono
		FROM home_assistant.listy_zakupow
		ORDER BY utworzono DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listy []*ListaZakupow
	for rows.Next() {
		lista := &ListaZakupow{}
		if err := rows.Scan(
			&lista.ID,
			&lista.Nazwa,
			&lista.Status,
			&lista.Utworzono,
		); err != nil {
			return nil, err
		}
		listy = append(listy, lista)
	}
	return listy, rows.Err()
}

func (r *listyZakupowRepo) FindByStatus(ctx context.Context, status StatusListyZakupow) ([]*ListaZakupow, error) {
	query := `
		SELECT id, nazwa, status, utworzono
		FROM home_assistant.listy_zakupow
		WHERE status = $1
		ORDER BY utworzono DESC
	`
	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listy []*ListaZakupow
	for rows.Next() {
		lista := &ListaZakupow{}
		if err := rows.Scan(
			&lista.ID,
			&lista.Nazwa,
			&lista.Status,
			&lista.Utworzono,
		); err != nil {
			return nil, err
		}
		listy = append(listy, lista)
	}
	return listy, rows.Err()
}

func (r *listyZakupowRepo) Update(ctx context.Context, lista *ListaZakupow) error {
	query := `
		UPDATE home_assistant.listy_zakupow
		SET nazwa = $1, status = $2
		WHERE id = $3
	`
	result, err := r.db.ExecContext(
		ctx, query,
		lista.Nazwa,
		lista.Status,
		lista.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("lista zakupow not found")
	}
	return nil
}

func (r *listyZakupowRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM home_assistant.listy_zakupow WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("lista zakupow not found")
	}
	return nil
}
