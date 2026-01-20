package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type pozycjeListyZakupowRepo struct {
	db DB
}

func NewPozycjeListyZakupowRepository(db DB) PozycjeListyZakupowRepository {
	return &pozycjeListyZakupowRepo{db: db}
}

func (r *pozycjeListyZakupowRepo) Create(ctx context.Context, pozycja *PozycjaListyZakupow) error {
	query := `
		INSERT INTO home_assistant.pozycje_listy_zakupow 
		(id_listy_zakupow, id_produktu, ilosc, notatka, czy_kupione)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	return r.db.QueryRowContext(
		ctx, query,
		pozycja.IDListyZakupow,
		pozycja.IDProduktu,
		pozycja.Ilosc,
		pozycja.Notatka,
		pozycja.CzyKupione,
	).Scan(&pozycja.ID)
}

func (r *pozycjeListyZakupowRepo) FindByID(ctx context.Context, id int) (*PozycjaListyZakupow, error) {
	query := `
		SELECT id, id_listy_zakupow, id_produktu, ilosc, notatka, czy_kupione
		FROM home_assistant.pozycje_listy_zakupow
		WHERE id = $1
	`
	pozycja := &PozycjaListyZakupow{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&pozycja.ID,
		&pozycja.IDListyZakupow,
		&pozycja.IDProduktu,
		&pozycja.Ilosc,
		&pozycja.Notatka,
		&pozycja.CzyKupione,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("pozycja not found")
	}
	return pozycja, err
}

func (r *pozycjeListyZakupowRepo) FindByLista(ctx context.Context, idListy int) ([]*PozycjaListyZakupow, error) {
	query := `
		SELECT id, id_listy_zakupow, id_produktu, ilosc, notatka, czy_kupione
		FROM home_assistant.pozycje_listy_zakupow
		WHERE id_listy_zakupow = $1
		ORDER BY czy_kupione ASC, id ASC
	`
	rows, err := r.db.QueryContext(ctx, query, idListy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pozycje []*PozycjaListyZakupow
	for rows.Next() {
		pozycja := &PozycjaListyZakupow{}
		if err := rows.Scan(
			&pozycja.ID,
			&pozycja.IDListyZakupow,
			&pozycja.IDProduktu,
			&pozycja.Ilosc,
			&pozycja.Notatka,
			&pozycja.CzyKupione,
		); err != nil {
			return nil, err
		}
		pozycje = append(pozycje, pozycja)
	}
	return pozycje, rows.Err()
}

func (r *pozycjeListyZakupowRepo) Update(ctx context.Context, pozycja *PozycjaListyZakupow) error {
	query := `
		UPDATE home_assistant.pozycje_listy_zakupow
		SET id_listy_zakupow = $1, id_produktu = $2, ilosc = $3, notatka = $4, czy_kupione = $5
		WHERE id = $6
	`
	result, err := r.db.ExecContext(
		ctx, query,
		pozycja.IDListyZakupow,
		pozycja.IDProduktu,
		pozycja.Ilosc,
		pozycja.Notatka,
		pozycja.CzyKupione,
		pozycja.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("pozycja not found")
	}
	return nil
}

func (r *pozycjeListyZakupowRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM home_assistant.pozycje_listy_zakupow WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("pozycja not found")
	}
	return nil
}

func (r *pozycjeListyZakupowRepo) MarkAsBought(ctx context.Context, id int) error {
	query := `
		UPDATE home_assistant.pozycje_listy_zakupow
		SET czy_kupione = true
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("pozycja not found")
	}
	return nil
}
