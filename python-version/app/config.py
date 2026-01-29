import os
from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    postgres_conn: str = "postgres://home_assistant_user:home_assistant_password@db:5432/home_assistant"
    server_port: int = 8000
    log_level: str = "info"

    class Config:
        env_file = ".env"


settings = Settings()
