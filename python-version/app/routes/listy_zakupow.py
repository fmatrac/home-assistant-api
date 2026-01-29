from typing import List, Optional
from fastapi import APIRouter, HTTPException, Query
from app.models import ListaZakupow, ListaZakupowCreate, StatusListyZakupow
from app.repositories import ListyZakupowRepository

router = APIRouter()


@router.get("/listy-zakupow", response_model=List[ListaZakupow])
def get_listy_zakupow(status: Optional[str] = Query(None)):
    if status:
        return ListyZakupowRepository.find_by_status(StatusListyZakupow(status))
    return ListyZakupowRepository.find_all()


@router.get("/listy-zakupow/get", response_model=ListaZakupow)
def get_lista_zakupow(id: int = Query(...)):
    lista = ListyZakupowRepository.find_by_id(id)
    if not lista:
        raise HTTPException(status_code=404, detail="Lista zakupow not found")
    return lista


@router.post("/listy-zakupow", response_model=ListaZakupow, status_code=201)
def create_lista_zakupow(lista: ListaZakupowCreate):
    return ListyZakupowRepository.create(lista)


@router.put("/listy-zakupow", response_model=ListaZakupow)
def update_lista_zakupow(id: int = Query(...), lista: ListaZakupowCreate = None):
    if not ListyZakupowRepository.update(id, lista):
        raise HTTPException(status_code=404, detail="Lista zakupow not found")
    return ListyZakupowRepository.find_by_id(id)


@router.delete("/listy-zakupow", status_code=204)
def delete_lista_zakupow(id: int = Query(...)):
    if not ListyZakupowRepository.delete(id):
        raise HTTPException(status_code=404, detail="Lista zakupow not found")
    return None
