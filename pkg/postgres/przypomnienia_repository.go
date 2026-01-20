package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type przypomnieniRepo struct {
	db DB
}

func NewPrzypomnieniRepository(db DB) PrzypomnieniRepository {
	return &przypomnieniRepo{db: db}
}

func (r *przypomnieniRepo) Create(ctx context.Context, przypomnienie *Przypomnienie) error {
	query := `
		INSERT INTO home_assistant.przypomnienia 
		(id_wydarzenia, tytul, opis, status, nastepne_uruchomienie)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, utworzono
	`
	return r.db.QueryRowContext(
		ctx, query,
		przypomnienie.IDWydarzenia,
		przypomnienie.Tytul,
		przypomnienie.Opis,
		przypomnienie.Status,
		przypomnienie.NastepneUruchomienie,
	).Scan(&przypomnienie.ID, &przypomnienie.Utworzono)
}

func (r *przypomnieniRepo) FindByID(ctx context.Context, id int) (*Przypomnienie, error) {
	query := `
		SELECT id, id_wydarzenia, tytul, opis, status, nastepne_uruchomienie, utworzono
		FROM home_assistant.przypomnienia
		WHERE id = $1
	`
	przypomnienie := &Przypomnienie{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&przypomnienie.ID,
		&przypomnienie.IDWydarzenia,
		&przypomnienie.Tytul,
		&przypomnienie.Opis,
		&przypomnienie.Status,
		&przypomnienie.NastepneUruchomienie,
		&przypomnienie.Utworzono,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("przypomnienie not found")
	}
	return przypomnienie, err
}

func (r *przypomnieniRepo) FindAll(ctx context.Context) ([]*Przypomnienie, error) {
	query := `
		SELECT id, id_wydarzenia, tytul, opis, status, nastepne_uruchomienie, utworzono
		FROM home_assistant.przypomnienia
		ORDER BY nastepne_uruchomienie ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var przypomnienia []*Przypomnienie
	for rows.Next() {
		przypomnienie := &Przypomnienie{}
		if err := rows.Scan(
			&przypomnienie.ID,
			&przypomnienie.IDWydarzenia,
			&przypomnienie.Tytul,
			&przypomnienie.Opis,
			&przypomnienie.Status,
			&przypomnienie.NastepneUruchomienie,
			&przypomnienie.Utworzono,
		); err != nil {
			return nil, err
		}
		przypomnienia = append(przypomnienia, przypomnienie)
	}
	return przypomnienia, rows.Err()
}

func (r *przypomnieniRepo) FindByWydarzenie(ctx context.Context, idWydarzenia int) ([]*Przypomnienie, error) {
	query := `
		SELECT id, id_wydarzenia, tytul, opis, status, nastepne_uruchomienie, utworzono
		FROM home_assistant.przypomnienia
		WHERE id_wydarzenia = $1
		ORDER BY nastepne_uruchomienie ASC
	`
	rows, err := r.db.QueryContext(ctx, query, idWydarzenia)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var przypomnienia []*Przypomnienie
	for rows.Next() {
		przypomnienie := &Przypomnienie{}
		if err := rows.Scan(
			&przypomnienie.ID,
			&przypomnienie.IDWydarzenia,
			&przypomnienie.Tytul,
			&przypomnienie.Opis,
			&przypomnienie.Status,
			&przypomnienie.NastepneUruchomienie,
			&przypomnienie.Utworzono,
		); err != nil {
			return nil, err
		}
		przypomnienia = append(przypomnienia, przypomnienie)
	}
	return przypomnienia, rows.Err()
}

func (r *przypomnieniRepo) Update(ctx context.Context, przypomnienie *Przypomnienie) error {
	query := `
		UPDATE home_assistant.przypomnienia
		SET id_wydarzenia = $1, tytul = $2, opis = $3, status = $4, nastepne_uruchomienie = $5
		WHERE id = $6
	`
	result, err := r.db.ExecContext(
		ctx, query,
		przypomnienie.IDWydarzenia,
		przypomnienie.Tytul,
		przypomnienie.Opis,
		przypomnienie.Status,
		przypomnienie.NastepneUruchomienie,
		przypomnienie.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("przypomnienie not found")
	}
	return nil
}

func (r *przypomnieniRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM home_assistant.przypomnienia WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("przypomnienie not found")
	}
	return nil
}
