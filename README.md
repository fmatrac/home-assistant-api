# Home Assistant API

Link do aplikacji: https://ti.matra.cc

Link do repozytorium: https://github.com/fmatrac/home-assistant-api

Aplikacja do zarządzania domem i życiem osobistym napisana w języku Go z wbudowanym interfejsem webowym. System umożliwia zarządzanie kalendarzem wydarzeń, przypomnieniami, produktami, listami zakupów oraz stanami magazynowymi.

> **Uwaga:** Cała inicjalizacja bazy danych (schemat, tabele, typy) znajduje się w folderze `migrations/`.

---

## Spis treści

1. [Wprowadzanie danych](#1-wprowadzanie-danych)
2. [Dokumentacja użytkownika](#2-dokumentacja-użytkownika)
3. [Dokumentacja techniczna](#3-dokumentacja-techniczna)

---

## 1. Wprowadzanie danych

System obsługuje następujące metody wprowadzania danych:

### 1.1 Wprowadzanie ręczne (główna metoda)

Podstawowym sposobem wprowadzania danych jest interfejs webowy aplikacji. Użytkownik może ręcznie dodawać, edytować i usuwać dane poprzez formularze modalne dostępne w każdej sekcji aplikacji:

- **Wydarzenia kalendarzowe** - formularze z polami: tytuł, opis, priorytet, data rozpoczęcia/zakończenia, miejsce
- **Przypomnienia** - formularze z polami: tytuł, opis, status, data uruchomienia, powiązanie z wydarzeniem
- **Produkty** - formularze z polami: nazwa, kategoria, ulubiony, link do zakupu
- **Listy zakupów** - formularze z polami: nazwa, status listy
- **Pozycje list zakupów** - formularze z polami: produkt, ilość, notatka
- **Stany magazynowe** - formularze z polami: produkt, stan (ok/mało/brak)

### 1.2 Wprowadzanie przez API (automatyczne/programowe)

System udostępnia REST API umożliwiające automatyczne wprowadzanie danych przez zewnętrzne aplikacje lub skrypty:

```bash
# Przykład dodania produktu przez API
curl -X POST http://localhost:8000/api/produkty \
  -H "Content-Type: application/json" \
  -d '{"nazwa": "Mleko", "kategoria": "Nabiał", "ulubiony": true}'
```

### 1.3 Import danych (SQL)

Dane można importować bezpośrednio do bazy PostgreSQL za pomocą skryptów SQL. Przykładowe skrypty znajdują się w katalogu `scripts/`:

```bash
# Import danych z pliku SQL
psql -h localhost -U home_assistant_user -d home_assistant -f scripts/seed_produkty.sql
```

### 1.4 Podsumowanie metod wprowadzania danych

| Metoda | Typ | Opis |
|--------|-----|------|
| Interfejs webowy | Ręczne | Formularze modalne w przeglądarce |
| REST API | Automatyczne | Żądania HTTP JSON |
| Skrypty SQL | Import | Bezpośredni import do bazy danych |

---

## 2. Dokumentacja użytkownika

### 2.1 Wymagania systemowe

- Docker i Docker Compose (zalecane)
- Lub: Go 1.24+, PostgreSQL 18+

### 2.2 Instalacja i uruchomienie

#### Metoda 1: Docker Compose (zalecana)

```bash
# Klonowanie repozytorium
git clone <url-repozytorium>
cd home-assistant-api

# Uruchomienie aplikacji
docker compose up -d

# Aplikacja dostępna pod adresem:
# http://localhost:8000
```

#### Metoda 2: Uruchomienie lokalne

```bash
# Instalacja zależności
go mod download

# Konfiguracja bazy danych (PostgreSQL)
# Edytuj config/config.yaml z danymi połączenia

# Uruchomienie migracji
goose -dir migrations postgres "postgres://user:pass@localhost:5432/home_assistant" up

# Uruchomienie aplikacji
go run cmd/main.go
```

---

## 3. Dokumentacja techniczna

### 3.1 Architektura systemu

Aplikacja wykorzystuje architekturę warstwową:

```
┌─────────────────────────────────────────┐
│           Frontend (SPA)                │
│    HTML + CSS + Vanilla JavaScript      │
├─────────────────────────────────────────┤
│          HTTP Server (api.go)           │
│         Routing i middleware            │
├─────────────────────────────────────────┤
│      Application (application.go)       │
│      Handlery HTTP, logika biznesowa    │
├─────────────────────────────────────────┤
│         Repositories (*_repository.go)  │
│         Warstwa dostępu do danych       │
├─────────────────────────────────────────┤
│            PostgreSQL 18                │
│         Schemat: home_assistant         │
└─────────────────────────────────────────┘
```

### 3.2 Stos technologiczny

| Komponent | Technologia | Wersja |
|-----------|-------------|--------|
| Backend | Go | 1.24.5 |
| Baza danych | PostgreSQL | 18 |
| Frontend | HTML/CSS/JavaScript | - |
| Logowanie | Logrus | 1.9.3 |
| Konfiguracja | Viper | 1.21.0 |
| Migracje | Goose | - |
| Konteneryzacja | Docker | - |

### 3.3 Struktura projektu

```
home-assistant-api/
├── cmd/
│   └── main.go                 # Punkt wejścia aplikacji
├── pkg/
│   ├── api/
│   │   └── api.go              # Serwer HTTP, rejestracja tras
│   ├── application/
│   │   └── application.go      # Handlery HTTP, logika biznesowa
│   ├── postgres/
│   │   ├── models.go           # Modele danych (struktury Go)
│   │   ├── types.go            # Interfejs DB i enumy
│   │   ├── repository_interface.go  # Interfejsy repozytoriów
│   │   ├── produkty_repository.go
│   │   ├── listy_zakupow_repository.go
│   │   ├── pozycje_listy_zakupow_repository.go
│   │   ├── stany_magazynowe_repository.go
│   │   ├── historia_stanu_zapasow_repository.go
│   │   ├── przypomnienia_repository.go
│   │   └── wydarzenia_kalendarz_repository.go
│   └── logger/
│       └── logger.go           # Wrapper Logrus
├── web/
│   ├── embed.go                # Deklaracja embedded FS
│   ├── index.html              # Główny interfejs HTML
│   ├── js/
│   │   ├── api.js              # Warstwa klienta API
│   │   └── app.js              # Logika frontendu
│   └── css/
│       └── style.css           # Style CSS
├── config/
│   ├── config.go               # Ładowanie konfiguracji
│   └── config.yaml             # Plik konfiguracyjny
├── migrations/
│   └── 00001_initialize_db.sql # Schemat bazy danych
├── db-init/
│   └── 01-init-schema.sql      # Inicjalizacja schematu (Docker)
├── compose.yaml                # Konfiguracja Docker Compose
├── Dockerfile                  # Multi-stage build
├── go.mod                      # Zależności Go
└── Makefile                    # Komendy build/clean
```

### 3.4 Model danych

#### 3.4.1 Diagram ERD

```
wydarzenia_kalendarz          przypomnienia
┌──────────────────┐         ┌──────────────────────┐
│ id (PK)          │◄───────┤│ id (PK)              │
│ tytul            │         │ id_wydarzenia (FK)   │
│ opis             │         │ tytul                │
│ priorytet        │         │ opis                 │
│ data_startu      │         │ status               │
│ data_zakonczenia │         │ nastepne_uruchomienie│
│ miejsce          │         │ utworzono            │
│ utworzono        │         └──────────────────────┘
└──────────────────┘

produkty                      listy_zakupow
┌──────────────────┐         ┌──────────────────┐
│ id (PK)          │         │ id (PK)          │
│ nazwa            │         │ nazwa            │
│ kategoria        │         │ status           │
│ ulubiony         │         │ utworzono        │
│ link_do_kupna    │         └────────┬─────────┘
│ utworzono        │                  │
└────────┬─────────┘                  │
         │                            │
         │    pozycje_listy_zakupow   │
         │   ┌──────────────────────┐ │
         └──►│ id (PK)              │◄┘
             │ id_listy_zakupow (FK)│
             │ id_produktu (FK)     │
             │ ilosc                │
             │ notatka              │
             │ czy_kupione          │
             └──────────────────────┘

produkty                     stany_magazynowe
┌──────────────────┐        ┌─────────────────────┐
│ id (PK)          │◄──────┤│ id (PK)             │
└──────────────────┘        │ id_produktu (FK)    │
         │                  │ stan                │
         │                  │ ostatnio_sprawdzono │
         │                  └─────────────────────┘
         │
         │   historia_stanu_zapasow
         │  ┌──────────────────┐
         └─►│ id (PK)          │
            │ id_produktu (FK) │
            │ stan             │
            │ data_zmiany      │
            │ notatka          │
            └──────────────────┘
```

#### 3.4.2 Typy wyliczeniowe (ENUM)

```sql
-- Priorytet wydarzenia
priorytet: 'maly', 'sredni', 'duzy'

-- Status przypomnienia
status_przypomnienia: 'aktywne', 'zrealizowane', 'zapauzowane'

-- Status listy zakupów
status_listy_zakupow: 'otwarta', 'zamknieta'

-- Stan zapasów
stan_zapasow: 'ok', 'malo', 'brak'
```

### 3.5 API REST - Dokumentacja endpointów

Serwer nasłuchuje na porcie **8000** (konfigurowalny w `config.yaml`).

#### 3.5.1 Health Check

| Metoda | Endpoint | Opis |
|--------|----------|------|
| GET | `/health` | Sprawdzenie stanu aplikacji |

**Odpowiedź:** `{"status": "OK"}`

#### 3.5.2 Wydarzenia (wydarzenia_kalendarz)

| Metoda | Endpoint | Opis |
|--------|----------|------|
| GET | `/api/wydarzenia` | Pobierz wszystkie wydarzenia |
| GET | `/api/wydarzenia/get?id={id}` | Pobierz pojedyncze wydarzenie |
| POST | `/api/wydarzenia` | Utwórz wydarzenie |
| PUT | `/api/wydarzenia?id={id}` | Aktualizuj wydarzenie |
| DELETE | `/api/wydarzenia?id={id}` | Usuń wydarzenie |

**Przykładowe body (POST/PUT):**
```json
{
  "tytul": "Spotkanie",
  "opis": "Opis spotkania",
  "priorytet": "sredni",
  "data_startu": "2025-01-15T10:00:00Z",
  "data_zakonczenia": "2025-01-15T11:00:00Z",
  "miejsce": "Biuro"
}
```

#### 3.5.3 Przypomnienia

| Metoda | Endpoint | Opis |
|--------|----------|------|
| GET | `/api/przypomnienia` | Pobierz wszystkie przypomnienia |
| GET | `/api/przypomnienia?id_wydarzenia={id}` | Filtruj po wydarzeniu |
| GET | `/api/przypomnienia/get?id={id}` | Pobierz pojedyncze przypomnienie |
| POST | `/api/przypomnienia` | Utwórz przypomnienie |
| PUT | `/api/przypomnienia?id={id}` | Aktualizuj przypomnienie |
| DELETE | `/api/przypomnienia?id={id}` | Usuń przypomnienie |

**Przykładowe body (POST/PUT):**
```json
{
  "tytul": "Pamiętaj o spotkaniu",
  "opis": "Przypomnienie",
  "status": "aktywne",
  "nastepne_uruchomienie": "2025-01-15T09:00:00Z",
  "id_wydarzenia": 1
}
```

#### 3.5.4 Produkty

| Metoda | Endpoint | Opis |
|--------|----------|------|
| GET | `/api/produkty` | Pobierz wszystkie produkty |
| GET | `/api/produkty?kategoria={nazwa}` | Filtruj po kategorii |
| GET | `/api/produkty/ulubione` | Pobierz ulubione produkty |
| GET | `/api/produkty/get?id={id}` | Pobierz pojedynczy produkt |
| POST | `/api/produkty` | Utwórz produkt |
| PUT | `/api/produkty?id={id}` | Aktualizuj produkt |
| DELETE | `/api/produkty?id={id}` | Usuń produkt |

**Przykładowe body (POST/PUT):**
```json
{
  "nazwa": "Mleko",
  "kategoria": "Nabiał",
  "ulubiony": true,
  "link_do_kupna": "https://sklep.pl/mleko"
}
```

#### 3.5.5 Listy zakupów

| Metoda | Endpoint | Opis |
|--------|----------|------|
| GET | `/api/listy-zakupow` | Pobierz wszystkie listy |
| GET | `/api/listy-zakupow?status={status}` | Filtruj po statusie |
| GET | `/api/listy-zakupow/get?id={id}` | Pobierz pojedynczą listę |
| POST | `/api/listy-zakupow` | Utwórz listę |
| PUT | `/api/listy-zakupow?id={id}` | Aktualizuj listę |
| DELETE | `/api/listy-zakupow?id={id}` | Usuń listę |

**Przykładowe body (POST/PUT):**
```json
{
  "nazwa": "Zakupy tygodniowe",
  "status": "otwarta"
}
```

#### 3.5.6 Pozycje list zakupów

| Metoda | Endpoint | Opis |
|--------|----------|------|
| GET | `/api/pozycje-listy?id_listy={id}` | Pobierz pozycje listy |
| GET | `/api/pozycje-listy/get?id={id}` | Pobierz pojedynczą pozycję |
| POST | `/api/pozycje-listy` | Dodaj pozycję do listy |
| PUT | `/api/pozycje-listy?id={id}` | Aktualizuj pozycję |
| DELETE | `/api/pozycje-listy?id={id}` | Usuń pozycję |
| POST | `/api/pozycje-listy/kupione?id={id}` | Oznacz jako kupione |

**Przykładowe body (POST/PUT):**
```json
{
  "id_listy_zakupow": 1,
  "id_produktu": 5,
  "ilosc": 2,
  "notatka": "Bio"
}
```

#### 3.5.7 Stany magazynowe

| Metoda | Endpoint | Opis |
|--------|----------|------|
| GET | `/api/stany-magazynowe` | Pobierz wszystkie stany |
| GET | `/api/stany-magazynowe?stan={stan}` | Filtruj po stanie |
| GET | `/api/stany-magazynowe/get?id={id}` | Pobierz pojedynczy stan |
| POST | `/api/stany-magazynowe` | Utwórz stan |
| PUT | `/api/stany-magazynowe?id={id}` | Aktualizuj stan |
| DELETE | `/api/stany-magazynowe?id={id}` | Usuń stan |

**Przykładowe body (POST/PUT):**
```json
{
  "id_produktu": 5,
  "stan": "malo"
}
```

#### 3.5.8 Historia zapasów

| Metoda | Endpoint | Opis |
|--------|----------|------|
| GET | `/api/historia-zapasow` | Pobierz całą historię |
| GET | `/api/historia-zapasow?id_produktu={id}` | Historia produktu |
| POST | `/api/historia-zapasow` | Dodaj wpis historii |

### 3.6 Wzorce projektowe

#### 3.6.1 Repository Pattern

Każda encja posiada własny interfejs repozytorium i implementację:

```go
// Interfejs repozytorium (repository_interface.go)
type ProduktyRepository interface {
    GetAll(ctx context.Context) ([]Produkt, error)
    GetByID(ctx context.Context, id int) (*Produkt, error)
    GetByKategoria(ctx context.Context, kategoria string) ([]Produkt, error)
    GetUlubione(ctx context.Context) ([]Produkt, error)
    Create(ctx context.Context, p *Produkt) error
    Update(ctx context.Context, p *Produkt) error
    Delete(ctx context.Context, id int) error
}

// Implementacja (produkty_repository.go)
type produktyRepository struct {
    db DB
}

func NewProduktyRepository(db DB) ProduktyRepository {
    return &produktyRepository{db: db}
}
```

#### 3.6.2 Dependency Injection

Repozytoria są wstrzykiwane do struktury Application:

```go
// application.go
type Application struct {
    wydarzenia       postgres.WydarzeniaKalendarzRepository
    przypomnienia    postgres.PrzypomieniaRepository
    produkty         postgres.ProduktyRepository
    listyZakupow     postgres.ListyZakupowRepository
    pozycjeListy     postgres.PozycjeListyZakupowRepository
    stanyMagazynowe  postgres.StanyMagazynoweRepository
    historiaZapasow  postgres.HistoriaStanuZapasowRepository
}
```

#### 3.6.3 Embedded Frontend

Frontend jest wbudowany w binarkę Go:

```go
// web/embed.go
//go:embed index.html css js
var WebFS embed.FS
```

### 3.7 Konfiguracja

Plik `config/config.yaml`:

```yaml
server:
  port: 8000                    # Port serwera HTTP

postgres:
  conn: "postgres://home_assistant_user:home_assistant_password@db:5432/home_assistant?sslmode=disable"

logging:
  level: "info"                 # Poziom logowania: debug, info, warn, error
```
