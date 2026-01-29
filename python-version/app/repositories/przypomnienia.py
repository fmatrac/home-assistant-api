from typing import List, Optional
from app.database import get_db
from app.models import Przypomnienie, PrzypomnienieCreate


class PrzypomnieniRepository:
    @staticmethod
    def create(przypomnienie: PrzypomnienieCreate) -> Przypomnienie:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO home_assistant.przypomnienia
                    (id_wydarzenia, tytul, opis, status, nastepne_uruchomienie)
                    VALUES (%s, %s, %s, %s, %s)
                    RETURNING id, id_wydarzenia, tytul, opis, status, nastepne_uruchomienie, utworzono
                    """,
                    (
                        przypomnienie.id_wydarzenia,
                        przypomnienie.tytul,
                        przypomnienie.opis,
                        przypomnienie.status.value,
                        przypomnienie.nastepne_uruchomienie,
                    ),
                )
                row = cur.fetchone()
                return Przypomnienie(**row)

    @staticmethod
    def find_by_id(id: int) -> Optional[Przypomnienie]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_wydarzenia, tytul, opis, status, nastepne_uruchomienie, utworzono
                    FROM home_assistant.przypomnienia
                    WHERE id = %s
                    """,
                    (id,),
                )
                row = cur.fetchone()
                if row:
                    return Przypomnienie(**row)
                return None

    @staticmethod
    def find_all() -> List[Przypomnienie]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_wydarzenia, tytul, opis, status, nastepne_uruchomienie, utworzono
                    FROM home_assistant.przypomnienia
                    ORDER BY nastepne_uruchomienie ASC
                    """
                )
                rows = cur.fetchall()
                return [Przypomnienie(**row) for row in rows]

    @staticmethod
    def find_by_wydarzenie(id_wydarzenia: int) -> List[Przypomnienie]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_wydarzenia, tytul, opis, status, nastepne_uruchomienie, utworzono
                    FROM home_assistant.przypomnienia
                    WHERE id_wydarzenia = %s
                    ORDER BY nastepne_uruchomienie ASC
                    """,
                    (id_wydarzenia,),
                )
                rows = cur.fetchall()
                return [Przypomnienie(**row) for row in rows]

    @staticmethod
    def update(id: int, przypomnienie: PrzypomnienieCreate) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    UPDATE home_assistant.przypomnienia
                    SET id_wydarzenia = %s, tytul = %s, opis = %s,
                        status = %s, nastepne_uruchomienie = %s
                    WHERE id = %s
                    """,
                    (
                        przypomnienie.id_wydarzenia,
                        przypomnienie.tytul,
                        przypomnienie.opis,
                        przypomnienie.status.value,
                        przypomnienie.nastepne_uruchomienie,
                        id,
                    ),
                )
                return cur.rowcount > 0

    @staticmethod
    def delete(id: int) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    "DELETE FROM home_assistant.przypomnienia WHERE id = %s",
                    (id,),
                )
                return cur.rowcount > 0
