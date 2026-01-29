from typing import List, Optional
from app.database import get_db
from app.models import StanMagazynowy, StanMagazynowyCreate, StanZapasow


class StanyMagazynoweRepository:
    @staticmethod
    def create(stan: StanMagazynowyCreate) -> StanMagazynowy:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO home_assistant.stany_magazynowe
                    (id_produktu, stan)
                    VALUES (%s, %s)
                    RETURNING id, id_produktu, stan, ostatnio_sprawdzono
                    """,
                    (stan.id_produktu, stan.stan.value),
                )
                row = cur.fetchone()
                return StanMagazynowy(**row)

    @staticmethod
    def find_by_id(id: int) -> Optional[StanMagazynowy]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_produktu, stan, ostatnio_sprawdzono
                    FROM home_assistant.stany_magazynowe
                    WHERE id = %s
                    """,
                    (id,),
                )
                row = cur.fetchone()
                if row:
                    return StanMagazynowy(**row)
                return None

    @staticmethod
    def find_by_produkt(id_produktu: int) -> Optional[StanMagazynowy]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_produktu, stan, ostatnio_sprawdzono
                    FROM home_assistant.stany_magazynowe
                    WHERE id_produktu = %s
                    """,
                    (id_produktu,),
                )
                row = cur.fetchone()
                if row:
                    return StanMagazynowy(**row)
                return None

    @staticmethod
    def find_all() -> List[StanMagazynowy]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_produktu, stan, ostatnio_sprawdzono
                    FROM home_assistant.stany_magazynowe
                    ORDER BY ostatnio_sprawdzono DESC
                    """
                )
                rows = cur.fetchall()
                return [StanMagazynowy(**row) for row in rows]

    @staticmethod
    def find_by_stan_zapasow(stan_zapasow: StanZapasow) -> List[StanMagazynowy]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_produktu, stan, ostatnio_sprawdzono
                    FROM home_assistant.stany_magazynowe
                    WHERE stan = %s
                    ORDER BY ostatnio_sprawdzono DESC
                    """,
                    (stan_zapasow.value,),
                )
                rows = cur.fetchall()
                return [StanMagazynowy(**row) for row in rows]

    @staticmethod
    def update(id: int, stan: StanMagazynowyCreate) -> Optional[StanMagazynowy]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    UPDATE home_assistant.stany_magazynowe
                    SET id_produktu = %s, stan = %s, ostatnio_sprawdzono = CURRENT_TIMESTAMP
                    WHERE id = %s
                    RETURNING id, id_produktu, stan, ostatnio_sprawdzono
                    """,
                    (stan.id_produktu, stan.stan.value, id),
                )
                row = cur.fetchone()
                if row:
                    return StanMagazynowy(**row)
                return None

    @staticmethod
    def delete(id: int) -> bool:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    "DELETE FROM home_assistant.stany_magazynowe WHERE id = %s",
                    (id,),
                )
                return cur.rowcount > 0
