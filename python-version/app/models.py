from datetime import datetime
from typing import Optional
from enum import Enum
from pydantic import BaseModel


# Enums
class Priorytet(str, Enum):
    maly = "maly"
    sredni = "sredni"
    duzy = "duzy"


class StatusPrzypomnienia(str, Enum):
    aktywne = "aktywne"
    zrealizowane = "zrealizowane"
    zapauzowane = "zapauzowane"


class StatusListyZakupow(str, Enum):
    otwarta = "otwarta"
    zamknieta = "zamknieta"


class StanZapasow(str, Enum):
    ok = "ok"
    malo = "malo"
    brak = "brak"


class RolaUzytkownika(str, Enum):
    user = "user"
    admin = "admin"


# Models
class WydarzenieKalendarzBase(BaseModel):
    tytul: str
    opis: str
    priorytet: Priorytet
    data_startu: datetime
    data_zakonczenia: datetime
    miejsce: str


class WydarzenieKalendarzCreate(WydarzenieKalendarzBase):
    pass


class WydarzenieKalendarz(WydarzenieKalendarzBase):
    id: int
    utworzono: Optional[datetime] = None

    class Config:
        from_attributes = True


class PrzypomnienieBase(BaseModel):
    id_wydarzenia: Optional[int] = None
    tytul: str
    opis: str
    status: StatusPrzypomnienia
    nastepne_uruchomienie: datetime


class PrzypomnienieCreate(PrzypomnienieBase):
    pass


class Przypomnienie(PrzypomnienieBase):
    id: int
    utworzono: Optional[datetime] = None

    class Config:
        from_attributes = True


class ProduktBase(BaseModel):
    nazwa: str
    kategoria: Optional[str] = None
    ulubiony: bool = False
    link_do_kupna: Optional[str] = None


class ProduktCreate(ProduktBase):
    pass


class Produkt(ProduktBase):
    id: int
    utworzono: Optional[datetime] = None

    class Config:
        from_attributes = True


class ListaZakupowBase(BaseModel):
    nazwa: str
    status: StatusListyZakupow


class ListaZakupowCreate(ListaZakupowBase):
    pass


class ListaZakupow(ListaZakupowBase):
    id: int
    utworzono: Optional[datetime] = None

    class Config:
        from_attributes = True


class PozycjaListyZakupowBase(BaseModel):
    id_listy_zakupow: int
    id_produktu: int
    ilosc: int
    notatka: Optional[str] = None
    czy_kupione: bool = False


class PozycjaListyZakupowCreate(PozycjaListyZakupowBase):
    pass


class PozycjaListyZakupow(PozycjaListyZakupowBase):
    id: int

    class Config:
        from_attributes = True


class StanMagazynowyBase(BaseModel):
    id_produktu: int
    stan: StanZapasow


class StanMagazynowyCreate(StanMagazynowyBase):
    pass


class StanMagazynowy(StanMagazynowyBase):
    id: int
    ostatnio_sprawdzono: Optional[datetime] = None

    class Config:
        from_attributes = True


class HistoriaStanuZapasowBase(BaseModel):
    id_produktu: int
    stan: int
    notatka: Optional[str] = None


class HistoriaStanuZapasowCreate(HistoriaStanuZapasowBase):
    pass


class HistoriaStanuZapasow(HistoriaStanuZapasowBase):
    id: int
    data_zmiany: Optional[datetime] = None

    class Config:
        from_attributes = True


# User models
class UzytkownikBase(BaseModel):
    email: str
    imie: Optional[str] = None


class UzytkownikCreate(UzytkownikBase):
    haslo: str


class Uzytkownik(UzytkownikBase):
    id: int
    rola: RolaUzytkownika = RolaUzytkownika.user
    aktywny: bool = True
    utworzono: Optional[datetime] = None

    class Config:
        from_attributes = True


class UzytkownikLogin(BaseModel):
    email: str
    haslo: str


class Token(BaseModel):
    access_token: str
    token_type: str = "bearer"


class TokenData(BaseModel):
    user_id: Optional[int] = None
    email: Optional[str] = None
    rola: Optional[str] = None


# Admin models
class UzytkownikAdmin(Uzytkownik):
    """Model użytkownika dla widoku admina - zawiera więcej danych"""
    pass


class UzytkownikUpdateAdmin(BaseModel):
    """Model do aktualizacji użytkownika przez admina"""
    imie: Optional[str] = None
    rola: Optional[RolaUzytkownika] = None
    aktywny: Optional[bool] = None
