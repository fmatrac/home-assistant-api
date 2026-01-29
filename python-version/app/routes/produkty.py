from typing import List, Optional
from fastapi import APIRouter, HTTPException, Query
from app.models import Produkt, ProduktCreate
from app.repositories import ProduktyRepository

router = APIRouter()


@router.get("/produkty", response_model=List[Produkt])
def get_produkty(kategoria: Optional[str] = Query(None)):
    if kategoria:
        return ProduktyRepository.find_by_kategoria(kategoria)
    return ProduktyRepository.find_all()


@router.get("/produkty/get", response_model=Produkt)
def get_produkt(id: int = Query(...)):
    produkt = ProduktyRepository.find_by_id(id)
    if not produkt:
        raise HTTPException(status_code=404, detail="Produkt not found")
    return produkt


@router.get("/produkty/ulubione", response_model=List[Produkt])
def get_produkty_ulubione():
    return ProduktyRepository.find_ulubione()


@router.post("/produkty", response_model=Produkt, status_code=201)
def create_produkt(produkt: ProduktCreate):
    return ProduktyRepository.create(produkt)


@router.put("/produkty", response_model=Produkt)
def update_produkt(id: int = Query(...), produkt: ProduktCreate = None):
    if not ProduktyRepository.update(id, produkt):
        raise HTTPException(status_code=404, detail="Produkt not found")
    return ProduktyRepository.find_by_id(id)


@router.delete("/produkty", status_code=204)
def delete_produkt(id: int = Query(...)):
    if not ProduktyRepository.delete(id):
        raise HTTPException(status_code=404, detail="Produkt not found")
    return None
