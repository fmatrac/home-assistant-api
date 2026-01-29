from fastapi import APIRouter
from app.routes import wydarzenia, przypomnienia, produkty, listy_zakupow, pozycje_listy, stany_magazynowe, historia_zapasow

api_router = APIRouter(prefix="/api")

api_router.include_router(wydarzenia.router, tags=["wydarzenia"])
api_router.include_router(przypomnienia.router, tags=["przypomnienia"])
api_router.include_router(produkty.router, tags=["produkty"])
api_router.include_router(listy_zakupow.router, tags=["listy-zakupow"])
api_router.include_router(pozycje_listy.router, tags=["pozycje-listy"])
api_router.include_router(stany_magazynowe.router, tags=["stany-magazynowe"])
api_router.include_router(historia_zapasow.router, tags=["historia-zapasow"])
