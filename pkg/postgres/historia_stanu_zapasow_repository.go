package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type historiaStanuZapasowRepo struct {
	db DB
}

func NewHistoriaStanuZapasowRepository(db DB) HistoriaStanuZapasowRepository {
	return &historiaStanuZapasowRepo{db: db}
}

func (r *historiaStanuZapasowRepo) Create(ctx context.Context, historia *HistoriaStanuZapasow) error {
	query := `
		INSERT INTO home_assistant.historia_stanu_zapasow 
		(id_produktu, stan, notatka)
		VALUES ($1, $2, $3)
		RETURNING id, data_zmiany
	`
	return r.db.QueryRowContext(
		ctx, query,
		historia.IDProduktu,
		historia.Stan,
		historia.Notatka,
	).Scan(&historia.ID, &historia.DataZmiany)
}

func (r *historiaStanuZapasowRepo) FindByID(ctx context.Context, id int) (*HistoriaStanuZapasow, error) {
	query := `
		SELECT id, id_produktu, stan, data_zmiany, notatka
		FROM home_assistant.historia_stanu_zapasow
		WHERE id = $1
	`
	historia := &HistoriaStanuZapasow{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&historia.ID,
		&historia.IDProduktu,
		&historia.Stan,
		&historia.DataZmiany,
		&historia.Notatka,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("historia not found")
	}
	return historia, err
}

func (r *historiaStanuZapasowRepo) FindByProdukt(ctx context.Context, idProduktu int) ([]*HistoriaStanuZapasow, error) {
	query := `
		SELECT id, id_produktu, stan, data_zmiany, notatka
		FROM home_assistant.historia_stanu_zapasow
		WHERE id_produktu = $1
		ORDER BY data_zmiany DESC
	`
	rows, err := r.db.QueryContext(ctx, query, idProduktu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historie []*HistoriaStanuZapasow
	for rows.Next() {
		historia := &HistoriaStanuZapasow{}
		if err := rows.Scan(
			&historia.ID,
			&historia.IDProduktu,
			&historia.Stan,
			&historia.DataZmiany,
			&historia.Notatka,
		); err != nil {
			return nil, err
		}
		historie = append(historie, historia)
	}
	return historie, rows.Err()
}

func (r *historiaStanuZapasowRepo) FindAll(ctx context.Context) ([]*HistoriaStanuZapasow, error) {
	query := `
		SELECT id, id_produktu, stan, data_zmiany, notatka
		FROM home_assistant.historia_stanu_zapasow
		ORDER BY data_zmiany DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historie []*HistoriaStanuZapasow
	for rows.Next() {
		historia := &HistoriaStanuZapasow{}
		if err := rows.Scan(
			&historia.ID,
			&historia.IDProduktu,
			&historia.Stan,
			&historia.DataZmiany,
			&historia.Notatka,
		); err != nil {
			return nil, err
		}
		historie = append(historie, historia)
	}
	return historie, rows.Err()
}
