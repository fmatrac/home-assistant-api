from typing import Optional, List
import bcrypt
from app.database import get_db
from app.models import Uzytkownik, UzytkownikCreate, UzytkownikUpdateAdmin


class UzytkownicyRepository:
    @staticmethod
    def hash_password(password: str) -> str:
        password_bytes = password.encode('utf-8')
        salt = bcrypt.gensalt()
        return bcrypt.hashpw(password_bytes, salt).decode('utf-8')

    @staticmethod
    def verify_password(plain_password: str, hashed_password: str) -> bool:
        password_bytes = plain_password.encode('utf-8')
        hashed_bytes = hashed_password.encode('utf-8')
        return bcrypt.checkpw(password_bytes, hashed_bytes)

    @staticmethod
    def create(uzytkownik: UzytkownikCreate) -> Uzytkownik:
        hashed_password = UzytkownicyRepository.hash_password(uzytkownik.haslo)
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO home_assistant.uzytkownicy
                    (email, haslo_hash, imie)
                    VALUES (%s, %s, %s)
                    RETURNING id, email, imie, rola, aktywny, utworzono
                    """,
                    (uzytkownik.email, hashed_password, uzytkownik.imie),
                )
                row = cur.fetchone()
                return Uzytkownik(**row)

    @staticmethod
    def find_by_id(id: int) -> Optional[Uzytkownik]:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, email, imie, rola, aktywny, utworzono
                    FROM home_assistant.uzytkownicy
                    WHERE id = %s
                    """,
                    (id,),
                )
                row = cur.fetchone()
                if row:
                    return Uzytkownik(**row)
                return None

    @staticmethod
    def find_by_email(email: str) -> Optional[tuple[Uzytkownik, str]]:
        """Returns user and password hash for authentication"""
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, email, imie, rola, aktywny, utworzono, haslo_hash
                    FROM home_assistant.uzytkownicy
                    WHERE email = %s
                    """,
                    (email,),
                )
                row = cur.fetchone()
                if row:
                    haslo_hash = row.pop("haslo_hash")
                    return Uzytkownik(**row), haslo_hash
                return None

    @staticmethod
    def authenticate(email: str, password: str) -> Optional[Uzytkownik]:
        result = UzytkownicyRepository.find_by_email(email)
        if not result:
            return None
        user, haslo_hash = result
        if not UzytkownicyRepository.verify_password(password, haslo_hash):
            return None
        if not user.aktywny:
            return None
        return user

    # Admin methods
    @staticmethod
    def find_all() -> List[Uzytkownik]:
        """Get all users - admin only"""
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    """
                    SELECT id, email, imie, rola, aktywny, utworzono
                    FROM home_assistant.uzytkownicy
                    ORDER BY utworzono DESC
                    """
                )
                rows = cur.fetchall()
                return [Uzytkownik(**row) for row in rows]

    @staticmethod
    def update_by_admin(id: int, data: UzytkownikUpdateAdmin) -> Optional[Uzytkownik]:
        """Update user by admin"""
        with get_db() as conn:
            with conn.cursor() as cur:
                # Build dynamic update query
                updates = []
                values = []
                if data.imie is not None:
                    updates.append("imie = %s")
                    values.append(data.imie)
                if data.rola is not None:
                    updates.append("rola = %s")
                    values.append(data.rola.value)
                if data.aktywny is not None:
                    updates.append("aktywny = %s")
                    values.append(data.aktywny)

                if not updates:
                    return UzytkownicyRepository.find_by_id(id)

                values.append(id)
                query = f"""
                    UPDATE home_assistant.uzytkownicy
                    SET {", ".join(updates)}
                    WHERE id = %s
                    RETURNING id, email, imie, rola, aktywny, utworzono
                """
                cur.execute(query, values)
                row = cur.fetchone()
                if row:
                    return Uzytkownik(**row)
                return None

    @staticmethod
    def delete(id: int) -> bool:
        """Delete user - admin only"""
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    "DELETE FROM home_assistant.uzytkownicy WHERE id = %s",
                    (id,),
                )
                return cur.rowcount > 0

    @staticmethod
    def count_all() -> int:
        """Count all users"""
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute("SELECT COUNT(*) as count FROM home_assistant.uzytkownicy")
                row = cur.fetchone()
                return row["count"] if row else 0
