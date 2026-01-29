from datetime import timedelta
from fastapi import APIRouter, HTTPException, status, Depends
from fastapi.security import OAuth2PasswordRequestForm
from app.config import settings
from app.models import UzytkownikCreate, Uzytkownik, Token, UzytkownikLogin
from app.repositories.uzytkownicy import UzytkownicyRepository
from app.auth import create_access_token, get_current_user

router = APIRouter(prefix="/auth")


@router.post("/register", response_model=Uzytkownik)
def register(uzytkownik: UzytkownikCreate):
    # Check if email already exists
    existing = UzytkownicyRepository.find_by_email(uzytkownik.email)
    if existing:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Email jest już zarejestrowany",
        )
    return UzytkownicyRepository.create(uzytkownik)


@router.post("/login", response_model=Token)
def login(form_data: OAuth2PasswordRequestForm = Depends()):
    user = UzytkownicyRepository.authenticate(form_data.username, form_data.password)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Nieprawidłowy email lub hasło",
            headers={"WWW-Authenticate": "Bearer"},
        )
    access_token = create_access_token(
        data={"sub": user.id, "email": user.email, "rola": user.rola.value},
        expires_delta=timedelta(minutes=settings.jwt_expire_minutes),
    )
    return Token(access_token=access_token)


@router.post("/login/json", response_model=Token)
def login_json(credentials: UzytkownikLogin):
    """Alternative login endpoint accepting JSON body"""
    user = UzytkownicyRepository.authenticate(credentials.email, credentials.haslo)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Nieprawidłowy email lub hasło",
            headers={"WWW-Authenticate": "Bearer"},
        )
    access_token = create_access_token(
        data={"sub": user.id, "email": user.email, "rola": user.rola.value},
        expires_delta=timedelta(minutes=settings.jwt_expire_minutes),
    )
    return Token(access_token=access_token)


@router.get("/me", response_model=Uzytkownik)
def get_me(current_user: Uzytkownik = Depends(get_current_user)):
    return current_user
