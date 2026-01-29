from typing import List, Optional
from fastapi import APIRouter, HTTPException, Query
from app.models import StanMagazynowy, StanMagazynowyCreate, StanZapasow
from app.repositories import StanyMagazynoweRepository

router = APIRouter()


@router.get("/stany-magazynowe", response_model=List[StanMagazynowy])
def get_stany_magazynowe(stan: Optional[str] = Query(None)):
    if stan:
        return StanyMagazynoweRepository.find_by_stan_zapasow(StanZapasow(stan))
    return StanyMagazynoweRepository.find_all()


@router.get("/stany-magazynowe/get", response_model=StanMagazynowy)
def get_stan_magazynowy(id: int = Query(...)):
    stan = StanyMagazynoweRepository.find_by_id(id)
    if not stan:
        raise HTTPException(status_code=404, detail="Stan magazynowy not found")
    return stan


@router.post("/stany-magazynowe", response_model=StanMagazynowy, status_code=201)
def create_stan_magazynowy(stan: StanMagazynowyCreate):
    return StanyMagazynoweRepository.create(stan)


@router.put("/stany-magazynowe", response_model=StanMagazynowy)
def update_stan_magazynowy(id: int = Query(...), stan: StanMagazynowyCreate = None):
    result = StanyMagazynoweRepository.update(id, stan)
    if not result:
        raise HTTPException(status_code=404, detail="Stan magazynowy not found")
    return result


@router.delete("/stany-magazynowe", status_code=204)
def delete_stan_magazynowy(id: int = Query(...)):
    if not StanyMagazynoweRepository.delete(id):
        raise HTTPException(status_code=404, detail="Stan magazynowy not found")
    return None
