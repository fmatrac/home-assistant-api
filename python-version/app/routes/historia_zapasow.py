from typing import List, Optional
from fastapi import APIRouter, Query
from app.models import HistoriaStanuZapasow, HistoriaStanuZapasowCreate
from app.repositories import HistoriaZapasowRepository

router = APIRouter()


@router.get("/historia-zapasow", response_model=List[HistoriaStanuZapasow])
def get_historia_zapasow(id_produktu: Optional[int] = Query(None)):
    if id_produktu is not None:
        return HistoriaZapasowRepository.find_by_produkt(id_produktu)
    return HistoriaZapasowRepository.find_all()


@router.post("/historia-zapasow", response_model=HistoriaStanuZapasow, status_code=201)
def create_historia_zapasow(historia: HistoriaStanuZapasowCreate):
    return HistoriaZapasowRepository.create(historia)
