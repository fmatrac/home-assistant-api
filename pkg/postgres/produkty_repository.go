package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type produktyRepo struct {
	db DB
}

func NewProduktyRepository(db DB) ProduktyRepository {
	return &produktyRepo{db: db}
}

func (r *produktyRepo) Create(ctx context.Context, produkt *Produkt) error {
	query := `
		INSERT INTO home_assistant.produkty 
		(nazwa, kategoria, ulubiony, link_do_kupna)
		VALUES ($1, $2, $3, $4)
		RETURNING id, utworzono
	`
	return r.db.QueryRowContext(
		ctx, query,
		produkt.Nazwa,
		produkt.Kategoria,
		produkt.Ulubiony,
		produkt.LinkDoKupna,
	).Scan(&produkt.ID, &produkt.Utworzono)
}

func (r *produktyRepo) FindByID(ctx context.Context, id int) (*Produkt, error) {
	query := `
		SELECT id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
		FROM home_assistant.produkty
		WHERE id = $1
	`
	produkt := &Produkt{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&produkt.ID,
		&produkt.Nazwa,
		&produkt.Kategoria,
		&produkt.Ulubiony,
		&produkt.LinkDoKupna,
		&produkt.Utworzono,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("produkt not found")
	}
	return produkt, err
}

func (r *produktyRepo) FindAll(ctx context.Context) ([]*Produkt, error) {
	query := `
		SELECT id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
		FROM home_assistant.produkty
		ORDER BY nazwa ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produkty []*Produkt
	for rows.Next() {
		produkt := &Produkt{}
		if err := rows.Scan(
			&produkt.ID,
			&produkt.Nazwa,
			&produkt.Kategoria,
			&produkt.Ulubiony,
			&produkt.LinkDoKupna,
			&produkt.Utworzono,
		); err != nil {
			return nil, err
		}
		produkty = append(produkty, produkt)
	}
	return produkty, rows.Err()
}

func (r *produktyRepo) FindByKategoria(ctx context.Context, kategoria string) ([]*Produkt, error) {
	query := `
		SELECT id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
		FROM home_assistant.produkty
		WHERE kategoria = $1
		ORDER BY nazwa ASC
	`
	rows, err := r.db.QueryContext(ctx, query, kategoria)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produkty []*Produkt
	for rows.Next() {
		produkt := &Produkt{}
		if err := rows.Scan(
			&produkt.ID,
			&produkt.Nazwa,
			&produkt.Kategoria,
			&produkt.Ulubiony,
			&produkt.LinkDoKupna,
			&produkt.Utworzono,
		); err != nil {
			return nil, err
		}
		produkty = append(produkty, produkt)
	}
	return produkty, rows.Err()
}

func (r *produktyRepo) FindUlubione(ctx context.Context) ([]*Produkt, error) {
	query := `
		SELECT id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
		FROM home_assistant.produkty
		WHERE ulubiony = true
		ORDER BY nazwa ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produkty []*Produkt
	for rows.Next() {
		produkt := &Produkt{}
		if err := rows.Scan(
			&produkt.ID,
			&produkt.Nazwa,
			&produkt.Kategoria,
			&produkt.Ulubiony,
			&produkt.LinkDoKupna,
			&produkt.Utworzono,
		); err != nil {
			return nil, err
		}
		produkty = append(produkty, produkt)
	}
	return produkty, rows.Err()
}

func (r *produktyRepo) Update(ctx context.Context, produkt *Produkt) error {
	query := `
		UPDATE home_assistant.produkty
		SET nazwa = $1, kategoria = $2, ulubiony = $3, link_do_kupna = $4
		WHERE id = $5
	`
	result, err := r.db.ExecContext(
		ctx, query,
		produkt.Nazwa,
		produkt.Kategoria,
		produkt.Ulubiony,
		produkt.LinkDoKupna,
		produkt.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("produkt not found")
	}
	return nil
}

func (r *produktyRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM home_assistant.produkty WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("produkt not found")
	}
	return nil
}
