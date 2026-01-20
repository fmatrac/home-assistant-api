package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fmatrac/home-assistant-api/config"
	"github.com/fmatrac/home-assistant-api/pkg/postgres"
	"github.com/sirupsen/logrus"
)

type (
	App interface {
		// Default handlers
		HandleDefault(w http.ResponseWriter, r *http.Request)
		HandleHealth(w http.ResponseWriter, r *http.Request)

		// Wydarzenia Kalendarz
		HandleGetWydarzenia(w http.ResponseWriter, r *http.Request)
		HandleGetWydarzenie(w http.ResponseWriter, r *http.Request)
		HandleCreateWydarzenie(w http.ResponseWriter, r *http.Request)
		HandleUpdateWydarzenie(w http.ResponseWriter, r *http.Request)
		HandleDeleteWydarzenie(w http.ResponseWriter, r *http.Request)

		// Przypomnienia
		HandleGetPrzypomnienia(w http.ResponseWriter, r *http.Request)
		HandleGetPrzypomnienie(w http.ResponseWriter, r *http.Request)
		HandleCreatePrzypomnienie(w http.ResponseWriter, r *http.Request)
		HandleUpdatePrzypomnienie(w http.ResponseWriter, r *http.Request)
		HandleDeletePrzypomnienie(w http.ResponseWriter, r *http.Request)

		// Produkty
		HandleGetProdukty(w http.ResponseWriter, r *http.Request)
		HandleGetProdukt(w http.ResponseWriter, r *http.Request)
		HandleCreateProdukt(w http.ResponseWriter, r *http.Request)
		HandleUpdateProdukt(w http.ResponseWriter, r *http.Request)
		HandleDeleteProdukt(w http.ResponseWriter, r *http.Request)
		HandleGetProduktyUlubione(w http.ResponseWriter, r *http.Request)

		// Listy Zakupow
		HandleGetListyZakupow(w http.ResponseWriter, r *http.Request)
		HandleGetListaZakupow(w http.ResponseWriter, r *http.Request)
		HandleCreateListaZakupow(w http.ResponseWriter, r *http.Request)
		HandleUpdateListaZakupow(w http.ResponseWriter, r *http.Request)
		HandleDeleteListaZakupow(w http.ResponseWriter, r *http.Request)

		// Pozycje Listy Zakupow
		HandleGetPozycjeListy(w http.ResponseWriter, r *http.Request)
		HandleGetPozycja(w http.ResponseWriter, r *http.Request)
		HandleCreatePozycja(w http.ResponseWriter, r *http.Request)
		HandleUpdatePozycja(w http.ResponseWriter, r *http.Request)
		HandleDeletePozycja(w http.ResponseWriter, r *http.Request)
		HandleMarkPozycjaAsBought(w http.ResponseWriter, r *http.Request)

		// Stany Magazynowe
		HandleGetStanyMagazynowe(w http.ResponseWriter, r *http.Request)
		HandleGetStanMagazynowy(w http.ResponseWriter, r *http.Request)
		HandleCreateStanMagazynowy(w http.ResponseWriter, r *http.Request)
		HandleUpdateStanMagazynowy(w http.ResponseWriter, r *http.Request)
		HandleDeleteStanMagazynowy(w http.ResponseWriter, r *http.Request)

		// Historia Stanu Zapasow
		HandleGetHistoriaStanuZapasow(w http.ResponseWriter, r *http.Request)
		HandleCreateHistoriaStanuZapasow(w http.ResponseWriter, r *http.Request)
	}

	Application struct {
		cfg                      *config.Config
		logger                   *logrus.Logger
		wydarzeniaRepo           postgres.WydarzeniaKalendarzRepository
		przypomnieniRepo         postgres.PrzypomnieniRepository
		produktyRepo             postgres.ProduktyRepository
		listyZakupowRepo         postgres.ListyZakupowRepository
		pozycjeListyZakupowRepo  postgres.PozycjeListyZakupowRepository
		stanyMagazynoweRepo      postgres.StanyMagazynoweRepository
		historiaStanuZapasowRepo postgres.HistoriaStanuZapasowRepository
	}
)

var _ App = (*Application)(nil)

func New(
	cfg *config.Config,
	logger *logrus.Logger,
	wydarzeniaRepo postgres.WydarzeniaKalendarzRepository,
	przypomnieniRepo postgres.PrzypomnieniRepository,
	produktyRepo postgres.ProduktyRepository,
	listyZakupowRepo postgres.ListyZakupowRepository,
	pozycjeListyZakupowRepo postgres.PozycjeListyZakupowRepository,
	stanyMagazynoweRepo postgres.StanyMagazynoweRepository,
	historiaStanuZapasowRepo postgres.HistoriaStanuZapasowRepository,
) *Application {
	return &Application{
		cfg:                      cfg,
		logger:                   logger,
		wydarzeniaRepo:           wydarzeniaRepo,
		przypomnieniRepo:         przypomnieniRepo,
		produktyRepo:             produktyRepo,
		listyZakupowRepo:         listyZakupowRepo,
		pozycjeListyZakupowRepo:  pozycjeListyZakupowRepo,
		stanyMagazynoweRepo:      stanyMagazynoweRepo,
		historiaStanuZapasowRepo: historiaStanuZapasowRepo,
	}
}

// Helper functions
func (a *Application) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (a *Application) respondError(w http.ResponseWriter, status int, message string) {
	a.respondJSON(w, status, map[string]string{"error": message})
}

func (a *Application) parseID(r *http.Request, param string) (int, error) {
	idStr := r.URL.Query().Get(param)
	if idStr == "" {
		return 0, fmt.Errorf("missing %s parameter", param)
	}
	return strconv.Atoi(idStr)
}

// Default handlers
func (a *Application) HandleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Assistant API")
}

func (a *Application) HandleHealth(w http.ResponseWriter, r *http.Request) {
	a.respondJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

// ==================== WYDARZENIA KALENDARZ ====================

func (a *Application) HandleGetWydarzenia(w http.ResponseWriter, r *http.Request) {
	wydarzenia, err := a.wydarzeniaRepo.FindAll(r.Context())
	if err != nil {
		a.logger.Error("Failed to get wydarzenia: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to get wydarzenia")
		return
	}
	a.respondJSON(w, http.StatusOK, wydarzenia)
}

func (a *Application) HandleGetWydarzenie(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	wydarzenie, err := a.wydarzeniaRepo.FindByID(r.Context(), id)
	if err != nil {
		a.respondError(w, http.StatusNotFound, "Wydarzenie not found")
		return
	}
	a.respondJSON(w, http.StatusOK, wydarzenie)
}

func (a *Application) HandleCreateWydarzenie(w http.ResponseWriter, r *http.Request) {
	var wydarzenie postgres.WydarzenieKalendarz
	if err := json.NewDecoder(r.Body).Decode(&wydarzenie); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := a.wydarzeniaRepo.Create(r.Context(), &wydarzenie); err != nil {
		a.logger.Error("Failed to create wydarzenie: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to create wydarzenie")
		return
	}
	a.respondJSON(w, http.StatusCreated, wydarzenie)
}

func (a *Application) HandleUpdateWydarzenie(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var wydarzenie postgres.WydarzenieKalendarz
	if err := json.NewDecoder(r.Body).Decode(&wydarzenie); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	wydarzenie.ID = id

	if err := a.wydarzeniaRepo.Update(r.Context(), &wydarzenie); err != nil {
		a.respondError(w, http.StatusNotFound, "Wydarzenie not found")
		return
	}
	a.respondJSON(w, http.StatusOK, wydarzenie)
}

func (a *Application) HandleDeleteWydarzenie(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.wydarzeniaRepo.Delete(r.Context(), id); err != nil {
		a.respondError(w, http.StatusNotFound, "Wydarzenie not found")
		return
	}
	a.respondJSON(w, http.StatusNoContent, nil)
}

// ==================== PRZYPOMNIENIA ====================

func (a *Application) HandleGetPrzypomnienia(w http.ResponseWriter, r *http.Request) {
	idWydarzenia := r.URL.Query().Get("id_wydarzenia")

	var przypomnienia []*postgres.Przypomnienie
	var err error

	if idWydarzenia != "" {
		id, parseErr := strconv.Atoi(idWydarzenia)
		if parseErr != nil {
			a.respondError(w, http.StatusBadRequest, "Invalid id_wydarzenia")
			return
		}
		przypomnienia, err = a.przypomnieniRepo.FindByWydarzenie(r.Context(), id)
	} else {
		przypomnienia, err = a.przypomnieniRepo.FindAll(r.Context())
	}

	if err != nil {
		a.logger.Error("Failed to get przypomnienia: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to get przypomnienia")
		return
	}
	a.respondJSON(w, http.StatusOK, przypomnienia)
}

func (a *Application) HandleGetPrzypomnienie(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	przypomnienie, err := a.przypomnieniRepo.FindByID(r.Context(), id)
	if err != nil {
		a.respondError(w, http.StatusNotFound, "Przypomnienie not found")
		return
	}
	a.respondJSON(w, http.StatusOK, przypomnienie)
}

func (a *Application) HandleCreatePrzypomnienie(w http.ResponseWriter, r *http.Request) {
	var przypomnienie postgres.Przypomnienie
	if err := json.NewDecoder(r.Body).Decode(&przypomnienie); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := a.przypomnieniRepo.Create(r.Context(), &przypomnienie); err != nil {
		a.logger.Error("Failed to create przypomnienie: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to create przypomnienie")
		return
	}
	a.respondJSON(w, http.StatusCreated, przypomnienie)
}

func (a *Application) HandleUpdatePrzypomnienie(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var przypomnienie postgres.Przypomnienie
	if err := json.NewDecoder(r.Body).Decode(&przypomnienie); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	przypomnienie.ID = id

	if err := a.przypomnieniRepo.Update(r.Context(), &przypomnienie); err != nil {
		a.respondError(w, http.StatusNotFound, "Przypomnienie not found")
		return
	}
	a.respondJSON(w, http.StatusOK, przypomnienie)
}

func (a *Application) HandleDeletePrzypomnienie(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.przypomnieniRepo.Delete(r.Context(), id); err != nil {
		a.respondError(w, http.StatusNotFound, "Przypomnienie not found")
		return
	}
	a.respondJSON(w, http.StatusNoContent, nil)
}

// ==================== PRODUKTY ====================

func (a *Application) HandleGetProdukty(w http.ResponseWriter, r *http.Request) {
	kategoria := r.URL.Query().Get("kategoria")

	var produkty []*postgres.Produkt
	var err error

	if kategoria != "" {
		produkty, err = a.produktyRepo.FindByKategoria(r.Context(), kategoria)
	} else {
		produkty, err = a.produktyRepo.FindAll(r.Context())
	}

	if err != nil {
		a.logger.Error("Failed to get produkty: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to get produkty")
		return
	}
	a.respondJSON(w, http.StatusOK, produkty)
}

func (a *Application) HandleGetProdukt(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	produkt, err := a.produktyRepo.FindByID(r.Context(), id)
	if err != nil {
		a.respondError(w, http.StatusNotFound, "Produkt not found")
		return
	}
	a.respondJSON(w, http.StatusOK, produkt)
}

func (a *Application) HandleCreateProdukt(w http.ResponseWriter, r *http.Request) {
	var produkt postgres.Produkt
	if err := json.NewDecoder(r.Body).Decode(&produkt); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := a.produktyRepo.Create(r.Context(), &produkt); err != nil {
		a.logger.Error("Failed to create produkt: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to create produkt")
		return
	}
	a.respondJSON(w, http.StatusCreated, produkt)
}

func (a *Application) HandleUpdateProdukt(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var produkt postgres.Produkt
	if err := json.NewDecoder(r.Body).Decode(&produkt); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	produkt.ID = id

	if err := a.produktyRepo.Update(r.Context(), &produkt); err != nil {
		a.respondError(w, http.StatusNotFound, "Produkt not found")
		return
	}
	a.respondJSON(w, http.StatusOK, produkt)
}

func (a *Application) HandleDeleteProdukt(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.produktyRepo.Delete(r.Context(), id); err != nil {
		a.respondError(w, http.StatusNotFound, "Produkt not found")
		return
	}
	a.respondJSON(w, http.StatusNoContent, nil)
}

func (a *Application) HandleGetProduktyUlubione(w http.ResponseWriter, r *http.Request) {
	produkty, err := a.produktyRepo.FindUlubione(r.Context())
	if err != nil {
		a.logger.Error("Failed to get ulubione produkty: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to get ulubione produkty")
		return
	}
	a.respondJSON(w, http.StatusOK, produkty)
}

// ==================== LISTY ZAKUPOW ====================

func (a *Application) HandleGetListyZakupow(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	var listy []*postgres.ListaZakupow
	var err error

	if status != "" {
		listy, err = a.listyZakupowRepo.FindByStatus(r.Context(), postgres.StatusListyZakupow(status))
	} else {
		listy, err = a.listyZakupowRepo.FindAll(r.Context())
	}

	if err != nil {
		a.logger.Error("Failed to get listy zakupow: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to get listy zakupow")
		return
	}
	a.respondJSON(w, http.StatusOK, listy)
}

func (a *Application) HandleGetListaZakupow(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	lista, err := a.listyZakupowRepo.FindByID(r.Context(), id)
	if err != nil {
		a.respondError(w, http.StatusNotFound, "Lista zakupow not found")
		return
	}
	a.respondJSON(w, http.StatusOK, lista)
}

func (a *Application) HandleCreateListaZakupow(w http.ResponseWriter, r *http.Request) {
	var lista postgres.ListaZakupow
	if err := json.NewDecoder(r.Body).Decode(&lista); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := a.listyZakupowRepo.Create(r.Context(), &lista); err != nil {
		a.logger.Error("Failed to create lista zakupow: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to create lista zakupow")
		return
	}
	a.respondJSON(w, http.StatusCreated, lista)
}

func (a *Application) HandleUpdateListaZakupow(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var lista postgres.ListaZakupow
	if err := json.NewDecoder(r.Body).Decode(&lista); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	lista.ID = id

	if err := a.listyZakupowRepo.Update(r.Context(), &lista); err != nil {
		a.respondError(w, http.StatusNotFound, "Lista zakupow not found")
		return
	}
	a.respondJSON(w, http.StatusOK, lista)
}

func (a *Application) HandleDeleteListaZakupow(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.listyZakupowRepo.Delete(r.Context(), id); err != nil {
		a.respondError(w, http.StatusNotFound, "Lista zakupow not found")
		return
	}
	a.respondJSON(w, http.StatusNoContent, nil)
}

// ==================== POZYCJE LISTY ZAKUPOW ====================

func (a *Application) HandleGetPozycjeListy(w http.ResponseWriter, r *http.Request) {
	idListy, err := a.parseID(r, "id_listy")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	pozycje, err := a.pozycjeListyZakupowRepo.FindByLista(r.Context(), idListy)
	if err != nil {
		a.logger.Error("Failed to get pozycje: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to get pozycje")
		return
	}
	a.respondJSON(w, http.StatusOK, pozycje)
}

func (a *Application) HandleGetPozycja(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	pozycja, err := a.pozycjeListyZakupowRepo.FindByID(r.Context(), id)
	if err != nil {
		a.respondError(w, http.StatusNotFound, "Pozycja not found")
		return
	}
	a.respondJSON(w, http.StatusOK, pozycja)
}

func (a *Application) HandleCreatePozycja(w http.ResponseWriter, r *http.Request) {
	var pozycja postgres.PozycjaListyZakupow
	if err := json.NewDecoder(r.Body).Decode(&pozycja); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := a.pozycjeListyZakupowRepo.Create(r.Context(), &pozycja); err != nil {
		a.logger.Error("Failed to create pozycja: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to create pozycja")
		return
	}
	a.respondJSON(w, http.StatusCreated, pozycja)
}

func (a *Application) HandleUpdatePozycja(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var pozycja postgres.PozycjaListyZakupow
	if err := json.NewDecoder(r.Body).Decode(&pozycja); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	pozycja.ID = id

	if err := a.pozycjeListyZakupowRepo.Update(r.Context(), &pozycja); err != nil {
		a.respondError(w, http.StatusNotFound, "Pozycja not found")
		return
	}
	a.respondJSON(w, http.StatusOK, pozycja)
}

func (a *Application) HandleDeletePozycja(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.pozycjeListyZakupowRepo.Delete(r.Context(), id); err != nil {
		a.respondError(w, http.StatusNotFound, "Pozycja not found")
		return
	}
	a.respondJSON(w, http.StatusNoContent, nil)
}

func (a *Application) HandleMarkPozycjaAsBought(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.pozycjeListyZakupowRepo.MarkAsBought(r.Context(), id); err != nil {
		a.respondError(w, http.StatusNotFound, "Pozycja not found")
		return
	}
	a.respondJSON(w, http.StatusOK, map[string]string{"message": "Pozycja marked as bought"})
}

// ==================== STANY MAGAZYNOWE ====================

func (a *Application) HandleGetStanyMagazynowe(w http.ResponseWriter, r *http.Request) {
	stan := r.URL.Query().Get("stan")

	var stany []*postgres.StanMagazynowy
	var err error

	if stan != "" {
		stany, err = a.stanyMagazynoweRepo.FindByStanZapasow(r.Context(), postgres.StanZapasow(stan))
	} else {
		stany, err = a.stanyMagazynoweRepo.FindAll(r.Context())
	}

	if err != nil {
		a.logger.Error("Failed to get stany magazynowe: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to get stany magazynowe")
		return
	}
	a.respondJSON(w, http.StatusOK, stany)
}

func (a *Application) HandleGetStanMagazynowy(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	stan, err := a.stanyMagazynoweRepo.FindByID(r.Context(), id)
	if err != nil {
		a.respondError(w, http.StatusNotFound, "Stan magazynowy not found")
		return
	}
	a.respondJSON(w, http.StatusOK, stan)
}

func (a *Application) HandleCreateStanMagazynowy(w http.ResponseWriter, r *http.Request) {
	var stan postgres.StanMagazynowy
	if err := json.NewDecoder(r.Body).Decode(&stan); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := a.stanyMagazynoweRepo.Create(r.Context(), &stan); err != nil {
		a.logger.Error("Failed to create stan magazynowy: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to create stan magazynowy")
		return
	}
	a.respondJSON(w, http.StatusCreated, stan)
}

func (a *Application) HandleUpdateStanMagazynowy(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var stan postgres.StanMagazynowy
	if err := json.NewDecoder(r.Body).Decode(&stan); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	stan.ID = id

	if err := a.stanyMagazynoweRepo.Update(r.Context(), &stan); err != nil {
		a.respondError(w, http.StatusNotFound, "Stan magazynowy not found")
		return
	}
	a.respondJSON(w, http.StatusOK, stan)
}

func (a *Application) HandleDeleteStanMagazynowy(w http.ResponseWriter, r *http.Request) {
	id, err := a.parseID(r, "id")
	if err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.stanyMagazynoweRepo.Delete(r.Context(), id); err != nil {
		a.respondError(w, http.StatusNotFound, "Stan magazynowy not found")
		return
	}
	a.respondJSON(w, http.StatusNoContent, nil)
}

// ==================== HISTORIA STANU ZAPASOW ====================

func (a *Application) HandleGetHistoriaStanuZapasow(w http.ResponseWriter, r *http.Request) {
	idProduktu := r.URL.Query().Get("id_produktu")

	var historia []*postgres.HistoriaStanuZapasow
	var err error

	if idProduktu != "" {
		id, parseErr := strconv.Atoi(idProduktu)
		if parseErr != nil {
			a.respondError(w, http.StatusBadRequest, "Invalid id_produktu")
			return
		}
		historia, err = a.historiaStanuZapasowRepo.FindByProdukt(r.Context(), id)
	} else {
		historia, err = a.historiaStanuZapasowRepo.FindAll(r.Context())
	}

	if err != nil {
		a.logger.Error("Failed to get historia stanu zapasow: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to get historia stanu zapasow")
		return
	}
	a.respondJSON(w, http.StatusOK, historia)
}

func (a *Application) HandleCreateHistoriaStanuZapasow(w http.ResponseWriter, r *http.Request) {
	var historia postgres.HistoriaStanuZapasow
	if err := json.NewDecoder(r.Body).Decode(&historia); err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := a.historiaStanuZapasowRepo.Create(r.Context(), &historia); err != nil {
		a.logger.Error("Failed to create historia stanu zapasow: ", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to create historia stanu zapasow")
		return
	}
	a.respondJSON(w, http.StatusCreated, historia)
}
