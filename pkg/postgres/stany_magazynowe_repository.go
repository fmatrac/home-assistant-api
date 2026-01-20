package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type stanyMagazynoweRepo struct {
	db DB
}

func NewStanyMagazynoweRepository(db DB) StanyMagazynoweRepository {
	return &stanyMagazynoweRepo{db: db}
}

func (r *stanyMagazynoweRepo) Create(ctx context.Context, stan *StanMagazynowy) error {
	query := `
		INSERT INTO home_assistant.stany_magazynowe 
		(id_produktu, stan)
		VALUES ($1, $2)
		RETURNING id, ostatnio_sprawdzono
	`
	return r.db.QueryRowContext(
		ctx, query,
		stan.IDProduktu,
		stan.Stan,
	).Scan(&stan.ID, &stan.OstatnioSprawdzono)
}

func (r *stanyMagazynoweRepo) FindByID(ctx context.Context, id int) (*StanMagazynowy, error) {
	query := `
		SELECT id, id_produktu, stan, ostatnio_sprawdzono
		FROM home_assistant.stany_magazynowe
		WHERE id = $1
	`
	stan := &StanMagazynowy{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&stan.ID,
		&stan.IDProduktu,
		&stan.Stan,
		&stan.OstatnioSprawdzono,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("stan magazynowy not found")
	}
	return stan, err
}

func (r *stanyMagazynoweRepo) FindByProdukt(ctx context.Context, idProduktu int) (*StanMagazynowy, error) {
	query := `
		SELECT id, id_produktu, stan, ostatnio_sprawdzono
		FROM home_assistant.stany_magazynowe
		WHERE id_produktu = $1
	`
	stan := &StanMagazynowy{}
	err := r.db.QueryRowContext(ctx, query, idProduktu).Scan(
		&stan.ID,
		&stan.IDProduktu,
		&stan.Stan,
		&stan.OstatnioSprawdzono,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("stan magazynowy not found")
	}
	return stan, err
}

func (r *stanyMagazynoweRepo) FindAll(ctx context.Context) ([]*StanMagazynowy, error) {
	query := `
		SELECT id, id_produktu, stan, ostatnio_sprawdzono
		FROM home_assistant.stany_magazynowe
		ORDER BY ostatnio_sprawdzono DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stany []*StanMagazynowy
	for rows.Next() {
		stan := &StanMagazynowy{}
		if err := rows.Scan(
			&stan.ID,
			&stan.IDProduktu,
			&stan.Stan,
			&stan.OstatnioSprawdzono,
		); err != nil {
			return nil, err
		}
		stany = append(stany, stan)
	}
	return stany, rows.Err()
}

func (r *stanyMagazynoweRepo) FindByStanZapasow(ctx context.Context, stanZapasow StanZapasow) ([]*StanMagazynowy, error) {
	query := `
		SELECT id, id_produktu, stan, ostatnio_sprawdzono
		FROM home_assistant.stany_magazynowe
		WHERE stan = $1
		ORDER BY ostatnio_sprawdzono DESC
	`
	rows, err := r.db.QueryContext(ctx, query, stanZapasow)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stany []*StanMagazynowy
	for rows.Next() {
		stan := &StanMagazynowy{}
		if err := rows.Scan(
			&stan.ID,
			&stan.IDProduktu,
			&stan.Stan,
			&stan.OstatnioSprawdzono,
		); err != nil {
			return nil, err
		}
		stany = append(stany, stan)
	}
	return stany, rows.Err()
}

func (r *stanyMagazynoweRepo) Update(ctx context.Context, stan *StanMagazynowy) error {
	query := `
		UPDATE home_assistant.stany_magazynowe
		SET id_produktu = $1, stan = $2, ostatnio_sprawdzono = CURRENT_TIMESTAMP
		WHERE id = $3
		RETURNING ostatnio_sprawdzono
	`
	err := r.db.QueryRowContext(
		ctx, query,
		stan.IDProduktu,
		stan.Stan,
		stan.ID,
	).Scan(&stan.OstatnioSprawdzono)

	if err == sql.ErrNoRows {
		return fmt.Errorf("stan magazynowy not found")
	}
	return err
}

func (r *stanyMagazynoweRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM home_assistant.stany_magazynowe WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("stan magazynowy not found")
	}
	return nil
}
