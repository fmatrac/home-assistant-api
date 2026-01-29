from typing import List, Optional
from app.database import get_db
from app.models import PozycjaListyZakupow, PozycjaListyZakupowCreate


class PozycjeListyRepository:
    @staticmethod
    def create(pozycja: PozycjaListyZakupowCreate) -> PozycjaListyZakupow:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO home_assistant.pozycje_listy_zakupow
                    (id_listy_zakupow, id_produktu, ilosc, notatka, czy_kupione)
                    VALUES (%s, %s, %s, %s, %s)
                    RETURNING id, id_listy_zakupow, id_produktu, ilosc, notatka, czy_kupione
                    """,
                    (
                        pozycja.id_listy_zakupow,
                        pozycja.id_produktu,
                        pozycja.ilosc,
                        pozycja.notatka,
                        pozycja.czy_kupione,
                    ),
                )
                row = cur.fetchone()
                return PozycjaListyZakupow(**row)

    @staticmethod
    def find_by_id(id: int) -> Optional[PozycjaListyZakupow]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_listy_zakupow, id_produktu, ilosc, notatka, czy_kupione
                    FROM home_assistant.pozycje_listy_zakupow
                    WHERE id = %s
                    """,
                    (id,),
                )
                row = cur.fetchone()
                if row:
                    return PozycjaListyZakupow(**row)
                return None

    @staticmethod
    def find_by_lista(id_listy: int) -> List[PozycjaListyZakupow]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_listy_zakupow, id_produktu, ilosc, notatka, czy_kupione
                    FROM home_assistant.pozycje_listy_zakupow
                    WHERE id_listy_zakupow = %s
                    ORDER BY czy_kupione ASC, id ASC
                    """,
                    (id_listy,),
                )
                rows = cur.fetchall()
                return [PozycjaListyZakupow(**row) for row in rows]

    @staticmethod
    def update(id: int, pozycja: PozycjaListyZakupowCreate) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    UPDATE home_assistant.pozycje_listy_zakupow
                    SET id_listy_zakupow = %s, id_produktu = %s,
                        ilosc = %s, notatka = %s, czy_kupione = %s
                    WHERE id = %s
                    """,
                    (
                        pozycja.id_listy_zakupow,
                        pozycja.id_produktu,
                        pozycja.ilosc,
                        pozycja.notatka,
                        pozycja.czy_kupione,
                        id,
                    ),
                )
                return cur.rowcount > 0

    @staticmethod
    def delete(id: int) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    "DELETE FROM home_assistant.pozycje_listy_zakupow WHERE id = %s",
                    (id,),
                )
                return cur.rowcount > 0

    @staticmethod
    def mark_as_bought(id: int) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    UPDATE home_assistant.pozycje_listy_zakupow
                    SET czy_kupione = true
                    WHERE id = %s
                    """,
                    (id,),
                )
                return cur.rowcount > 0
