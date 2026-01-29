from typing import List, Optional
from app.database import get_db
from app.models import ListaZakupow, ListaZakupowCreate, StatusListyZakupow


class ListyZakupowRepository:
    @staticmethod
    def create(lista: ListaZakupowCreate) -> ListaZakupow:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO home_assistant.listy_zakupow
                    (nazwa, status)
                    VALUES (%s, %s)
                    RETURNING id, nazwa, status, utworzono
                    """,
                    (lista.nazwa, lista.status.value),
                )
                row = cur.fetchone()
                return ListaZakupow(**row)

    @staticmethod
    def find_by_id(id: int) -> Optional[ListaZakupow]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, nazwa, status, utworzono
                    FROM home_assistant.listy_zakupow
                    WHERE id = %s
                    """,
                    (id,),
                )
                row = cur.fetchone()
                if row:
                    return ListaZakupow(**row)
                return None

    @staticmethod
    def find_all() -> List[ListaZakupow]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, nazwa, status, utworzono
                    FROM home_assistant.listy_zakupow
                    ORDER BY utworzono DESC
                    """
                )
                rows = cur.fetchall()
                return [ListaZakupow(**row) for row in rows]

    @staticmethod
    def find_by_status(status: StatusListyZakupow) -> List[ListaZakupow]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, nazwa, status, utworzono
                    FROM home_assistant.listy_zakupow
                    WHERE status = %s
                    ORDER BY utworzono DESC
                    """,
                    (status.value,),
                )
                rows = cur.fetchall()
                return [ListaZakupow(**row) for row in rows]

    @staticmethod
    def update(id: int, lista: ListaZakupowCreate) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    UPDATE home_assistant.listy_zakupow
                    SET nazwa = %s, status = %s
                    WHERE id = %s
                    """,
                    (lista.nazwa, lista.status.value, id),
                )
                return cur.rowcount > 0

    @staticmethod
    def delete(id: int) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    "DELETE FROM home_assistant.listy_zakupow WHERE id = %s",
                    (id,),
                )
                return cur.rowcount > 0
