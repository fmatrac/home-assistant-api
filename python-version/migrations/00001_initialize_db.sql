-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS home_assistant;

GRANT ALL ON SCHEMA home_assistant TO home_assistant_user;

-- ENUMS
CREATE TYPE home_assistant.priorytet AS ENUM (
    'maly',
    'sredni',
    'duzy'
);

CREATE TYPE home_assistant.status_przypomnienia AS ENUM (
    'aktywne',
    'zrealizowane',
    'zapauzowane'
);

CREATE TYPE home_assistant.status_listy_zakupow AS ENUM (
    'otwarta',
    'zamknieta'
);

CREATE TYPE home_assistant.stan_zapasow AS ENUM (
    'ok',
    'malo',
    'brak'
);

CREATE TYPE home_assistant.rola_uzytkownika AS ENUM (
    'user',
    'admin'
);

-- TABLES

-- Tabela użytkowników (musi być pierwsza ze względu na FK)
CREATE TABLE home_assistant.uzytkownicy (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    haslo_hash VARCHAR(255) NOT NULL,
    imie VARCHAR(100),
    rola home_assistant.rola_uzytkownika DEFAULT 'user',
    aktywny BOOLEAN DEFAULT TRUE,
    utworzono TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE home_assistant.wydarzenia_kalendarz (
    id SERIAL PRIMARY KEY,
    uzytkownik_id INT REFERENCES home_assistant.uzytkownicy(id) ON DELETE CASCADE,
    tytul VARCHAR(255) NOT NULL,
    opis VARCHAR(255) NOT NULL,
    priorytet home_assistant.priorytet NOT NULL,
    data_startu TIMESTAMP NOT NULL,
    data_zakonczenia TIMESTAMP NOT NULL,
    miejsce VARCHAR(255) NOT NULL,
    utworzono TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE home_assistant.przypomnienia (
    id                    SERIAL PRIMARY KEY,
    uzytkownik_id         INT REFERENCES home_assistant.uzytkownicy(id) ON DELETE CASCADE,
    id_wydarzenia         INT REFERENCES home_assistant.wydarzenia_kalendarz (id) ON DELETE CASCADE,
    tytul                 VARCHAR(255) NOT NULL,
    opis                  VARCHAR(255) NOT NULL,
    status                home_assistant.status_przypomnienia NOT NULL,
    nastepne_uruchomienie TIMESTAMP NOT NULL,
    utworzono             TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE home_assistant.produkty (
    id            SERIAL PRIMARY KEY,
    uzytkownik_id INT REFERENCES home_assistant.uzytkownicy(id) ON DELETE CASCADE,
    nazwa         VARCHAR(255) NOT NULL,
    kategoria     VARCHAR(100),
    ulubiony      BOOLEAN DEFAULT FALSE,
    link_do_kupna VARCHAR(500),
    utworzono     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE home_assistant.listy_zakupow (
    id            SERIAL PRIMARY KEY,
    uzytkownik_id INT REFERENCES home_assistant.uzytkownicy(id) ON DELETE CASCADE,
    nazwa         VARCHAR(255) NOT NULL,
    status        home_assistant.status_listy_zakupow NOT NULL,
    utworzono     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE home_assistant.pozycje_listy_zakupow (
    id SERIAL PRIMARY KEY,
    id_listy_zakupow INT REFERENCES home_assistant.listy_zakupow (id) ON DELETE CASCADE,
    id_produktu INT NOT NULL REFERENCES home_assistant.produkty (id) ON DELETE CASCADE,
    ilosc INT NOT NULL,
    notatka VARCHAR(255),
    czy_kupione BOOLEAN DEFAULT FALSE
);

CREATE TABLE home_assistant.historia_stanu_zapasow (
    id            SERIAL PRIMARY KEY,
    uzytkownik_id INT REFERENCES home_assistant.uzytkownicy(id) ON DELETE CASCADE,
    id_produktu   INT REFERENCES home_assistant.produkty (id) ON DELETE CASCADE,
    stan          INT NOT NULL,
    data_zmiany   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notatka       VARCHAR(255)
);

CREATE TABLE home_assistant.stany_magazynowe (
    id                  SERIAL PRIMARY KEY,
    uzytkownik_id       INT REFERENCES home_assistant.uzytkownicy(id) ON DELETE CASCADE,
    id_produktu         INT REFERENCES home_assistant.produkty (id) ON DELETE CASCADE,
    stan                home_assistant.stan_zapasow NOT NULL,
    ostatnio_sprawdzono TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indeksy dla szybszego wyszukiwania po użytkowniku
CREATE INDEX idx_wydarzenia_uzytkownik ON home_assistant.wydarzenia_kalendarz(uzytkownik_id);
CREATE INDEX idx_przypomnienia_uzytkownik ON home_assistant.przypomnienia(uzytkownik_id);
CREATE INDEX idx_produkty_uzytkownik ON home_assistant.produkty(uzytkownik_id);
CREATE INDEX idx_listy_zakupow_uzytkownik ON home_assistant.listy_zakupow(uzytkownik_id);
CREATE INDEX idx_stany_magazynowe_uzytkownik ON home_assistant.stany_magazynowe(uzytkownik_id);
CREATE INDEX idx_historia_zapasow_uzytkownik ON home_assistant.historia_stanu_zapasow(uzytkownik_id);

-- Przykładowi użytkownicy (hasło: test123)
INSERT INTO home_assistant.uzytkownicy (email, haslo_hash, imie, rola)
VALUES
    ('admin@example.com', '$2b$12$BhOKTltgSJ1yPIBfgZQf7O9vdaZjMvbjMlQOPwZVu8DcIdR/J5QJC', 'Administrator', 'admin'),
    ('user@example.com', '$2b$12$BhOKTltgSJ1yPIBfgZQf7O9vdaZjMvbjMlQOPwZVu8DcIdR/J5QJC', 'Zwykly User', 'user');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS home_assistant.idx_historia_zapasow_uzytkownik;
DROP INDEX IF EXISTS home_assistant.idx_stany_magazynowe_uzytkownik;
DROP INDEX IF EXISTS home_assistant.idx_listy_zakupow_uzytkownik;
DROP INDEX IF EXISTS home_assistant.idx_produkty_uzytkownik;
DROP INDEX IF EXISTS home_assistant.idx_przypomnienia_uzytkownik;
DROP INDEX IF EXISTS home_assistant.idx_wydarzenia_uzytkownik;

DROP TABLE IF EXISTS home_assistant.stany_magazynowe;
DROP TABLE IF EXISTS home_assistant.historia_stanu_zapasow;
DROP TABLE IF EXISTS home_assistant.pozycje_listy_zakupow;
DROP TABLE IF EXISTS home_assistant.listy_zakupow;
DROP TABLE IF EXISTS home_assistant.produkty;
DROP TABLE IF EXISTS home_assistant.przypomnienia;
DROP TABLE IF EXISTS home_assistant.wydarzenia_kalendarz;
DROP TABLE IF EXISTS home_assistant.uzytkownicy;

DROP TYPE IF EXISTS home_assistant.rola_uzytkownika;
DROP TYPE IF EXISTS home_assistant.stan_zapasow;
DROP TYPE IF EXISTS home_assistant.status_listy_zakupow;
DROP TYPE IF EXISTS home_assistant.status_przypomnienia;
DROP TYPE IF EXISTS home_assistant.priorytet;

DROP SCHEMA IF EXISTS home_assistant;
-- +goose StatementEnd
