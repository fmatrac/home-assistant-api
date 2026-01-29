-- Create the home_assistant schema
CREATE SCHEMA IF NOT EXISTS home_assistant;

-- Grant privileges to the user
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

-- TABLES
CREATE TABLE home_assistant.wydarzenia_kalendarz (
    id SERIAL PRIMARY KEY,
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
    id_wydarzenia         INT REFERENCES home_assistant.wydarzenia_kalendarz (id) ON DELETE CASCADE,
    tytul                 VARCHAR(255) NOT NULL,
    opis                  VARCHAR(255) NOT NULL,
    status                home_assistant.status_przypomnienia NOT NULL,
    nastepne_uruchomienie TIMESTAMP NOT NULL,
    utworzono             TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE home_assistant.produkty (
    id        SERIAL PRIMARY KEY,
    nazwa     VARCHAR(255) NOT NULL,
    kategoria VARCHAR(100),
    ulubiony  BOOLEAN DEFAULT FALSE,
    link_do_kupna VARCHAR(500),
    utworzono TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE home_assistant.listy_zakupow (
    id SERIAL PRIMARY KEY,
    nazwa VARCHAR(255) NOT NULL,
    status home_assistant.status_listy_zakupow NOT NULL,
    utworzono TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
    id SERIAL PRIMARY KEY,
    id_produktu INT REFERENCES home_assistant.produkty (id) ON DELETE CASCADE,
    stan INT NOT NULL,
    data_zmiany TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notatka VARCHAR(255)
);

CREATE TABLE home_assistant.stany_magazynowe (
    id SERIAL PRIMARY KEY,
    id_produktu         INT REFERENCES home_assistant.produkty (id) ON DELETE CASCADE,
    stan                home_assistant.stan_zapasow NOT NULL,
    ostatnio_sprawdzono TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
