from typing import List
from fastapi import APIRouter, HTTPException, status, Depends
from app.models import Uzytkownik, UzytkownikUpdateAdmin
from app.repositories.uzytkownicy import UzytkownicyRepository
from app.auth import get_current_admin

router = APIRouter(prefix="/admin")


@router.get("/users", response_model=List[Uzytkownik])
def get_all_users(current_admin: Uzytkownik = Depends(get_current_admin)):
    """Get all users - admin only"""
    return UzytkownicyRepository.find_all()


@router.get("/users/{user_id}", response_model=Uzytkownik)
def get_user(user_id: int, current_admin: Uzytkownik = Depends(get_current_admin)):
    """Get user by ID - admin only"""
    user = UzytkownicyRepository.find_by_id(user_id)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Użytkownik nie znaleziony",
        )
    return user


@router.put("/users/{user_id}", response_model=Uzytkownik)
def update_user(
    user_id: int,
    data: UzytkownikUpdateAdmin,
    current_admin: Uzytkownik = Depends(get_current_admin),
):
    """Update user - admin only"""
    # Prevent admin from deactivating themselves
    if user_id == current_admin.id and data.aktywny is False:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Nie możesz dezaktywować własnego konta",
        )
    # Prevent admin from removing their own admin role
    if user_id == current_admin.id and data.rola and data.rola.value != "admin":
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Nie możesz usunąć sobie roli administratora",
        )

    user = UzytkownicyRepository.update_by_admin(user_id, data)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Użytkownik nie znaleziony",
        )
    return user


@router.delete("/users/{user_id}")
def delete_user(user_id: int, current_admin: Uzytkownik = Depends(get_current_admin)):
    """Delete user - admin only"""
    # Prevent admin from deleting themselves
    if user_id == current_admin.id:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Nie możesz usunąć własnego konta",
        )

    success = UzytkownicyRepository.delete(user_id)
    if not success:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Użytkownik nie znaleziony",
        )
    return {"message": "Użytkownik usunięty"}


@router.get("/stats")
def get_stats(current_admin: Uzytkownik = Depends(get_current_admin)):
    """Get admin statistics"""
    return {
        "total_users": UzytkownicyRepository.count_all(),
    }
