from typing import List
from fastapi import APIRouter, HTTPException, Query
from app.models import PozycjaListyZakupow, PozycjaListyZakupowCreate
from app.repositories import PozycjeListyRepository

router = APIRouter()


@router.get("/pozycje-listy", response_model=List[PozycjaListyZakupow])
def get_pozycje_listy(id_listy: int = Query(...)):
    return PozycjeListyRepository.find_by_lista(id_listy)


@router.get("/pozycje-listy/get", response_model=PozycjaListyZakupow)
def get_pozycja(id: int = Query(...)):
    pozycja = PozycjeListyRepository.find_by_id(id)
    if not pozycja:
        raise HTTPException(status_code=404, detail="Pozycja not found")
    return pozycja


@router.post("/pozycje-listy", response_model=PozycjaListyZakupow, status_code=201)
def create_pozycja(pozycja: PozycjaListyZakupowCreate):
    return PozycjeListyRepository.create(pozycja)


@router.put("/pozycje-listy", response_model=PozycjaListyZakupow)
def update_pozycja(id: int = Query(...), pozycja: PozycjaListyZakupowCreate = None):
    if not PozycjeListyRepository.update(id, pozycja):
        raise HTTPException(status_code=404, detail="Pozycja not found")
    return PozycjeListyRepository.find_by_id(id)


@router.delete("/pozycje-listy", status_code=204)
def delete_pozycja(id: int = Query(...)):
    if not PozycjeListyRepository.delete(id):
        raise HTTPException(status_code=404, detail="Pozycja not found")
    return None


@router.post("/pozycje-listy/kupione")
def mark_pozycja_as_bought(id: int = Query(...)):
    if not PozycjeListyRepository.mark_as_bought(id):
        raise HTTPException(status_code=404, detail="Pozycja not found")
    return {"message": "Pozycja marked as bought"}
