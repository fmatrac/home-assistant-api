package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type wydarzeniaKalendarzRepo struct {
	db DB
}

func NewWydarzeniaKalendarzRepository(db DB) WydarzeniaKalendarzRepository {
	return &wydarzeniaKalendarzRepo{db: db}
}

func (r *wydarzeniaKalendarzRepo) Create(ctx context.Context, wydarzenie *WydarzenieKalendarz) error {
	query := `
		INSERT INTO home_assistant.wydarzenia_kalendarz 
		(tytul, opis, priorytet, data_startu, data_zakonczenia, miejsce)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, utworzono
	`
	return r.db.QueryRowContext(
		ctx, query,
		wydarzenie.Tytul,
		wydarzenie.Opis,
		wydarzenie.Priorytet,
		wydarzenie.DataStartu,
		wydarzenie.DataZakonczenia,
		wydarzenie.Miejsce,
	).Scan(&wydarzenie.ID, &wydarzenie.Utworzono)
}

func (r *wydarzeniaKalendarzRepo) FindByID(ctx context.Context, id int) (*WydarzenieKalendarz, error) {
	query := `
		SELECT id, tytul, opis, priorytet, data_startu, data_zakonczenia, miejsce, utworzono
		FROM home_assistant.wydarzenia_kalendarz
		WHERE id = $1
	`
	wydarzenie := &WydarzenieKalendarz{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&wydarzenie.ID,
		&wydarzenie.Tytul,
		&wydarzenie.Opis,
		&wydarzenie.Priorytet,
		&wydarzenie.DataStartu,
		&wydarzenie.DataZakonczenia,
		&wydarzenie.Miejsce,
		&wydarzenie.Utworzono,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("wydarzenie not found")
	}
	return wydarzenie, err
}

func (r *wydarzeniaKalendarzRepo) FindAll(ctx context.Context) ([]*WydarzenieKalendarz, error) {
	query := `
		SELECT id, tytul, opis, priorytet, data_startu, data_zakonczenia, miejsce, utworzono
		FROM home_assistant.wydarzenia_kalendarz
		ORDER BY data_startu DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wydarzenia []*WydarzenieKalendarz
	for rows.Next() {
		wydarzenie := &WydarzenieKalendarz{}
		if err := rows.Scan(
			&wydarzenie.ID,
			&wydarzenie.Tytul,
			&wydarzenie.Opis,
			&wydarzenie.Priorytet,
			&wydarzenie.DataStartu,
			&wydarzenie.DataZakonczenia,
			&wydarzenie.Miejsce,
			&wydarzenie.Utworzono,
		); err != nil {
			return nil, err
		}
		wydarzenia = append(wydarzenia, wydarzenie)
	}
	return wydarzenia, rows.Err()
}

func (r *wydarzeniaKalendarzRepo) Update(ctx context.Context, wydarzenie *WydarzenieKalendarz) error {
	query := `
		UPDATE home_assistant.wydarzenia_kalendarz
		SET tytul = $1, opis = $2, priorytet = $3, data_startu = $4, data_zakonczenia = $5, miejsce = $6
		WHERE id = $7
	`
	result, err := r.db.ExecContext(
		ctx, query,
		wydarzenie.Tytul,
		wydarzenie.Opis,
		wydarzenie.Priorytet,
		wydarzenie.DataStartu,
		wydarzenie.DataZakonczenia,
		wydarzenie.Miejsce,
		wydarzenie.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("wydarzenie not found")
	}
	return nil
}

func (r *wydarzeniaKalendarzRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM home_assistant.wydarzenia_kalendarz WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("wydarzenie not found")
	}
	return nil
}
