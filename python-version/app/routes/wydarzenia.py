from typing import List, Optional
from fastapi import APIRouter, HTTPException, Query
from app.models import WydarzenieKalendarz, WydarzenieKalendarzCreate
from app.repositories import WydarzeniaRepository

router = APIRouter()


@router.get("/wydarzenia", response_model=List[WydarzenieKalendarz])
def get_wydarzenia():
    return WydarzeniaRepository.find_all()


@router.get("/wydarzenia/get", response_model=WydarzenieKalendarz)
def get_wydarzenie(id: int = Query(...)):
    wydarzenie = WydarzeniaRepository.find_by_id(id)
    if not wydarzenie:
        raise HTTPException(status_code=404, detail="Wydarzenie not found")
    return wydarzenie


@router.post("/wydarzenia", response_model=WydarzenieKalendarz, status_code=201)
def create_wydarzenie(wydarzenie: WydarzenieKalendarzCreate):
    return WydarzeniaRepository.create(wydarzenie)


@router.put("/wydarzenia", response_model=WydarzenieKalendarz)
def update_wydarzenie(id: int = Query(...), wydarzenie: WydarzenieKalendarzCreate = None):
    if not WydarzeniaRepository.update(id, wydarzenie):
        raise HTTPException(status_code=404, detail="Wydarzenie not found")
    return WydarzeniaRepository.find_by_id(id)


@router.delete("/wydarzenia", status_code=204)
def delete_wydarzenie(id: int = Query(...)):
    if not WydarzeniaRepository.delete(id):
        raise HTTPException(status_code=404, detail="Wydarzenie not found")
    return None
