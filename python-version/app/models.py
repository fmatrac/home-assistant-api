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
