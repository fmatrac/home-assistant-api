from typing import List, Optional
from app.database import get_db
from app.models import Produkt, ProduktCreate


class ProduktyRepository:
    @staticmethod
    def create(produkt: ProduktCreate) -> Produkt:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO home_assistant.produkty
                    (nazwa, kategoria, ulubiony, link_do_kupna)
                    VALUES (%s, %s, %s, %s)
                    RETURNING id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
                    """,
                    (
                        produkt.nazwa,
                        produkt.kategoria,
                        produkt.ulubiony,
                        produkt.link_do_kupna,
                    ),
                )
                row = cur.fetchone()
                return Produkt(**row)

    @staticmethod
    def find_by_id(id: int) -> Optional[Produkt]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
                    FROM home_assistant.produkty
                    WHERE id = %s
                    """,
                    (id,),
                )
                row = cur.fetchone()
                if row:
                    return Produkt(**row)
                return None

    @staticmethod
    def find_all() -> List[Produkt]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
                    FROM home_assistant.produkty
                    ORDER BY nazwa ASC
                    """
                )
                rows = cur.fetchall()
                return [Produkt(**row) for row in rows]

    @staticmethod
    def find_by_kategoria(kategoria: str) -> List[Produkt]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
                    FROM home_assistant.produkty
                    WHERE kategoria = %s
                    ORDER BY nazwa ASC
                    """,
                    (kategoria,),
                )
                rows = cur.fetchall()
                return [Produkt(**row) for row in rows]

    @staticmethod
    def find_ulubione() -> List[Produkt]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, nazwa, kategoria, ulubiony, link_do_kupna, utworzono
                    FROM home_assistant.produkty
                    WHERE ulubiony = true
                    ORDER BY nazwa ASC
                    """
                )
                rows = cur.fetchall()
                return [Produkt(**row) for row in rows]

    @staticmethod
    def update(id: int, produkt: ProduktCreate) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    UPDATE home_assistant.produkty
                    SET nazwa = %s, kategoria = %s, ulubiony = %s, link_do_kupna = %s
                    WHERE id = %s
                    """,
                    (
                        produkt.nazwa,
                        produkt.kategoria,
                        produkt.ulubiony,
                        produkt.link_do_kupna,
                        id,
                    ),
                )
                return cur.rowcount > 0

    @staticmethod
    def delete(id: int) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    "DELETE FROM home_assistant.produkty WHERE id = %s",
                    (id,),
                )
                return cur.rowcount > 0
