from typing import List, Optional
from app.database import get_db
from app.models import WydarzenieKalendarz, WydarzenieKalendarzCreate


class WydarzeniaRepository:
    @staticmethod
    def create(wydarzenie: WydarzenieKalendarzCreate) -> WydarzenieKalendarz:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO home_assistant.wydarzenia_kalendarz
                    (tytul, opis, priorytet, data_startu, data_zakonczenia, miejsce)
                    VALUES (%s, %s, %s, %s, %s, %s)
                    RETURNING id, tytul, opis, priorytet, data_startu, data_zakonczenia, miejsce, utworzono
                    """,
                    (
                        wydarzenie.tytul,
                        wydarzenie.opis,
                        wydarzenie.priorytet.value,
                        wydarzenie.data_startu,
                        wydarzenie.data_zakonczenia,
                        wydarzenie.miejsce,
                    ),
                )
                row = cur.fetchone()
                return WydarzenieKalendarz(**row)

    @staticmethod
    def find_by_id(id: int) -> Optional[WydarzenieKalendarz]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, tytul, opis, priorytet, data_startu, data_zakonczenia, miejsce, utworzono
                    FROM home_assistant.wydarzenia_kalendarz
                    WHERE id = %s
                    """,
                    (id,),
                )
                row = cur.fetchone()
                if row:
                    return WydarzenieKalendarz(**row)
                return None

    @staticmethod
    def find_all() -> List[WydarzenieKalendarz]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, tytul, opis, priorytet, data_startu, data_zakonczenia, miejsce, utworzono
                    FROM home_assistant.wydarzenia_kalendarz
                    ORDER BY data_startu DESC
                    """
                )
                rows = cur.fetchall()
                return [WydarzenieKalendarz(**row) for row in rows]

    @staticmethod
    def update(id: int, wydarzenie: WydarzenieKalendarzCreate) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    UPDATE home_assistant.wydarzenia_kalendarz
                    SET tytul = %s, opis = %s, priorytet = %s,
                        data_startu = %s, data_zakonczenia = %s, miejsce = %s
                    WHERE id = %s
                    """,
                    (
                        wydarzenie.tytul,
                        wydarzenie.opis,
                        wydarzenie.priorytet.value,
                        wydarzenie.data_startu,
                        wydarzenie.data_zakonczenia,
                        wydarzenie.miejsce,
                        id,
                    ),
                )
                return cur.rowcount > 0

    @staticmethod
    def delete(id: int) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    "DELETE FROM home_assistant.wydarzenia_kalendarz WHERE id = %s",
                    (id,),
                )
                return cur.rowcount > 0
