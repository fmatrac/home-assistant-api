from typing import List, Optional
from fastapi import APIRouter, HTTPException, Query
from app.models import Przypomnienie, PrzypomnienieCreate
from app.repositories import PrzypomnieniRepository

router = APIRouter()


@router.get("/przypomnienia", response_model=List[Przypomnienie])
def get_przypomnienia(id_wydarzenia: Optional[int] = Query(None)):
    if id_wydarzenia is not None:
        return PrzypomnieniRepository.find_by_wydarzenie(id_wydarzenia)
    return PrzypomnieniRepository.find_all()


@router.get("/przypomnienia/get", response_model=Przypomnienie)
def get_przypomnienie(id: int = Query(...)):
    przypomnienie = PrzypomnieniRepository.find_by_id(id)
    if not przypomnienie:
        raise HTTPException(status_code=404, detail="Przypomnienie not found")
    return przypomnienie


@router.post("/przypomnienia", response_model=Przypomnienie, status_code=201)
def create_przypomnienie(przypomnienie: PrzypomnienieCreate):
    return PrzypomnieniRepository.create(przypomnienie)


@router.put("/przypomnienia", response_model=Przypomnienie)
def update_przypomnienie(id: int = Query(...), przypomnienie: PrzypomnienieCreate = None):
    if not PrzypomnieniRepository.update(id, przypomnienie):
        raise HTTPException(status_code=404, detail="Przypomnienie not found")
    return PrzypomnieniRepository.find_by_id(id)


@router.delete("/przypomnienia", status_code=204)
def delete_przypomnienie(id: int = Query(...)):
    if not PrzypomnieniRepository.delete(id):
        raise HTTPException(status_code=404, detail="Przypomnienie not found")
    return None
