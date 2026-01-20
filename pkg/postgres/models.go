package postgres

import "time"

// Enums
type Priorytet string

const (
	PriorytetMaly   Priorytet = "maly"
	PriorytetSredni Priorytet = "sredni"
	PriorytetDuzy   Priorytet = "duzy"
)

type StatusPrzipomnienia string

const (
	StatusAktywne      StatusPrzipomnienia = "aktywne"
	StatusZrealizowane StatusPrzipomnienia = "zrealizowane"
	StatusZapauzowane  StatusPrzipomnienia = "zapauzowane"
)

type StatusListyZakupow string

const (
	StatusOtwarta   StatusListyZakupow = "otwarta"
	StatusZamknieta StatusListyZakupow = "zamknieta"
)

type StanZapasow string

const (
	StanOk   StanZapasow = "ok"
	StanMalo StanZapasow = "malo"
	StanBrak StanZapasow = "brak"
)

// Models
type WydarzenieKalendarz struct {
	ID              int       `db:"id" json:"id"`
	Tytul           string    `db:"tytul" json:"tytul"`
	Opis            string    `db:"opis" json:"opis"`
	Priorytet       Priorytet `db:"priorytet" json:"priorytet"`
	DataStartu      time.Time `db:"data_startu" json:"data_startu"`
	DataZakonczenia time.Time `db:"data_zakonczenia" json:"data_zakonczenia"`
	Miejsce         string    `db:"miejsce" json:"miejsce"`
	Utworzono       time.Time `db:"utworzono" json:"utworzono"`
}

type Przypomnienie struct {
	ID                   int                 `db:"id" json:"id"`
	IDWydarzenia         *int                `db:"id_wydarzenia" json:"id_wydarzenia"`
	Tytul                string              `db:"tytul" json:"tytul"`
	Opis                 string              `db:"opis" json:"opis"`
	Status               StatusPrzipomnienia `db:"status" json:"status"`
	NastepneUruchomienie time.Time           `db:"nastepne_uruchomienie" json:"nastepne_uruchomienie"`
	Utworzono            time.Time           `db:"utworzono" json:"utworzono"`
}

type Produkt struct {
	ID          int       `db:"id" json:"id"`
	Nazwa       string    `db:"nazwa" json:"nazwa"`
	Kategoria   *string   `db:"kategoria" json:"kategoria,omitempty"`
	Ulubiony    bool      `db:"ulubiony" json:"ulubiony"`
	LinkDoKupna *string   `db:"link_do_kupna" json:"link_do_kupna,omitempty"`
	Utworzono   time.Time `db:"utworzono" json:"utworzono"`
}

type ListaZakupow struct {
	ID        int                `db:"id" json:"id"`
	Nazwa     string             `db:"nazwa" json:"nazwa"`
	Status    StatusListyZakupow `db:"status" json:"status"`
	Utworzono time.Time          `db:"utworzono" json:"utworzono"`
}

type PozycjaListyZakupow struct {
	ID             int     `db:"id" json:"id"`
	IDListyZakupow int     `db:"id_listy_zakupow" json:"id_listy_zakupow"`
	IDProduktu     int     `db:"id_produktu" json:"id_produktu"`
	Ilosc          int     `db:"ilosc" json:"ilosc"`
	Notatka        *string `db:"notatka" json:"notatka,omitempty"`
	CzyKupione     bool    `db:"czy_kupione" json:"czy_kupione"`
}

type HistoriaStanuZapasow struct {
	ID         int       `db:"id" json:"id"`
	IDProduktu int       `db:"id_produktu" json:"id_produktu"`
	Stan       int       `db:"stan" json:"stan"`
	DataZmiany time.Time `db:"data_zmiany" json:"data_zmiany"`
	Notatka    *string   `db:"notatka" json:"notatka,omitempty"`
}

type StanMagazynowy struct {
	ID                 int         `db:"id" json:"id"`
	IDProduktu         int         `db:"id_produktu" json:"id_produktu"`
	Stan               StanZapasow `db:"stan" json:"stan"`
	OstatnioSprawdzono time.Time   `db:"ostatnio_sprawdzono" json:"ostatnio_sprawdzono"`
}
