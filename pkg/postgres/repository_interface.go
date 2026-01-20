package postgres

import "context"

// WydarzeniaKalendarzRepository defines operations for wydarzenia_kalendarz table
type WydarzeniaKalendarzRepository interface {
	Create(ctx context.Context, wydarzenie *WydarzenieKalendarz) error
	FindByID(ctx context.Context, id int) (*WydarzenieKalendarz, error)
	FindAll(ctx context.Context) ([]*WydarzenieKalendarz, error)
	Update(ctx context.Context, wydarzenie *WydarzenieKalendarz) error
	Delete(ctx context.Context, id int) error
}

// PrzypomnieniRepository defines operations for przypomnienia table
type PrzypomnieniRepository interface {
	Create(ctx context.Context, przypomnienie *Przypomnienie) error
	FindByID(ctx context.Context, id int) (*Przypomnienie, error)
	FindAll(ctx context.Context) ([]*Przypomnienie, error)
	FindByWydarzenie(ctx context.Context, idWydarzenia int) ([]*Przypomnienie, error)
	Update(ctx context.Context, przypomnienie *Przypomnienie) error
	Delete(ctx context.Context, id int) error
}

// ProduktyRepository defines operations for produkty table
type ProduktyRepository interface {
	Create(ctx context.Context, produkt *Produkt) error
	FindByID(ctx context.Context, id int) (*Produkt, error)
	FindAll(ctx context.Context) ([]*Produkt, error)
	FindByKategoria(ctx context.Context, kategoria string) ([]*Produkt, error)
	FindUlubione(ctx context.Context) ([]*Produkt, error)
	Update(ctx context.Context, produkt *Produkt) error
	Delete(ctx context.Context, id int) error
}

// ListyZakupowRepository defines operations for listy_zakupow table
type ListyZakupowRepository interface {
	Create(ctx context.Context, lista *ListaZakupow) error
	FindByID(ctx context.Context, id int) (*ListaZakupow, error)
	FindAll(ctx context.Context) ([]*ListaZakupow, error)
	FindByStatus(ctx context.Context, status StatusListyZakupow) ([]*ListaZakupow, error)
	Update(ctx context.Context, lista *ListaZakupow) error
	Delete(ctx context.Context, id int) error
}

// PozycjeListyZakupowRepository defines operations for pozycje_listy_zakupow table
type PozycjeListyZakupowRepository interface {
	Create(ctx context.Context, pozycja *PozycjaListyZakupow) error
	FindByID(ctx context.Context, id int) (*PozycjaListyZakupow, error)
	FindByLista(ctx context.Context, idListy int) ([]*PozycjaListyZakupow, error)
	Update(ctx context.Context, pozycja *PozycjaListyZakupow) error
	Delete(ctx context.Context, id int) error
	MarkAsBought(ctx context.Context, id int) error
}

// HistoriaStanuZapasowRepository defines operations for historia_stanu_zapasow table
type HistoriaStanuZapasowRepository interface {
	Create(ctx context.Context, historia *HistoriaStanuZapasow) error
	FindByID(ctx context.Context, id int) (*HistoriaStanuZapasow, error)
	FindByProdukt(ctx context.Context, idProduktu int) ([]*HistoriaStanuZapasow, error)
	FindAll(ctx context.Context) ([]*HistoriaStanuZapasow, error)
}

// StanyMagazynoweRepository defines operations for stany_magazynowe table
type StanyMagazynoweRepository interface {
	Create(ctx context.Context, stan *StanMagazynowy) error
	FindByID(ctx context.Context, id int) (*StanMagazynowy, error)
	FindByProdukt(ctx context.Context, idProduktu int) (*StanMagazynowy, error)
	FindAll(ctx context.Context) ([]*StanMagazynowy, error)
	FindByStanZapasow(ctx context.Context, stan StanZapasow) ([]*StanMagazynowy, error)
	Update(ctx context.Context, stan *StanMagazynowy) error
	Delete(ctx context.Context, id int) error
}
