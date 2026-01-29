from typing import List, Optional
from app.database import get_db
from app.models import HistoriaStanuZapasow, HistoriaStanuZapasowCreate


class HistoriaZapasowRepository:
    @staticmethod
    def create(historia: HistoriaStanuZapasowCreate) -> HistoriaStanuZapasow:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO home_assistant.historia_stanu_zapasow
                    (id_produktu, stan, notatka)
                    VALUES (%s, %s, %s)
                    RETURNING id, id_produktu, stan, data_zmiany, notatka
                    """,
                    (historia.id_produktu, historia.stan, historia.notatka),
                )
                row = cur.fetchone()
                return HistoriaStanuZapasow(**row)

    @staticmethod
    def find_by_id(id: int) -> Optional[HistoriaStanuZapasow]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_produktu, stan, data_zmiany, notatka
                    FROM home_assistant.historia_stanu_zapasow
                    WHERE id = %s
                    """,
                    (id,),
                )
                row = cur.fetchone()
                if row:
                    return HistoriaStanuZapasow(**row)
                return None

    @staticmethod
    def find_by_produkt(id_produktu: int) -> List[HistoriaStanuZapasow]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_produktu, stan, data_zmiany, notatka
                    FROM home_assistant.historia_stanu_zapasow
                    WHERE id_produktu = %s
                    ORDER BY data_zmiany DESC
                    """,
                    (id_produktu,),
                )
                rows = cur.fetchall()
                return [HistoriaStanuZapasow(**row) for row in rows]

    @staticmethod
    def find_all() -> List[HistoriaStanuZapasow]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, id_produktu, stan, data_zmiany, notatka
                    FROM home_assistant.historia_stanu_zapasow
                    ORDER BY data_zmiany DESC
                    """
                )
                rows = cur.fetchall()
                return [HistoriaStanuZapasow(**row) for row in rows]
