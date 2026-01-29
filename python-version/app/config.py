import os
from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    postgres_conn: str = "postgres://home_assistant_user:home_assistant_password@db:5432/home_assistant"
    server_port: int = 8000
    log_level: str = "info"

    # JWT settings
    jwt_secret: str = "your-secret-key-change-in-production"
    jwt_algorithm: str = "HS256"
    jwt_expire_minutes: int = 60 * 24  # 24 hours

    class Config:
        env_file = ".env"


settings = Settings()
