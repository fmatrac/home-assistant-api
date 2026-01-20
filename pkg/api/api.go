package api

import (
	"net/http"

	"github.com/fmatrac/home-assistant-api/config"
	a "github.com/fmatrac/home-assistant-api/pkg/application"
	l "github.com/sirupsen/logrus"
)

type Server struct {
	mux         *http.ServeMux
	cfg         *config.Config
	logger      *l.Logger
	application a.App
	webFS       http.FileSystem
}

func InitServer(cfg *config.Config, logger *l.Logger, application a.App, webFS http.FileSystem) {
	server := &Server{
		mux:         http.NewServeMux(),
		cfg:         cfg,
		logger:      logger,
		application: application,
		webFS:       webFS,
	}
	server.registerRoutes()
	server.listenAndServe()
}

func (s *Server) registerRoutes() {
	// Serve embedded frontend files
	s.mux.Handle("/", http.FileServer(s.webFS))

	// Health check
	s.mux.HandleFunc("/health", s.application.HandleHealth)

	// Wydarzenia Kalendarz
	s.mux.HandleFunc("GET /api/wydarzenia", s.application.HandleGetWydarzenia)
	s.mux.HandleFunc("GET /api/wydarzenia/get", s.application.HandleGetWydarzenie)
	s.mux.HandleFunc("POST /api/wydarzenia", s.application.HandleCreateWydarzenie)
	s.mux.HandleFunc("PUT /api/wydarzenia", s.application.HandleUpdateWydarzenie)
	s.mux.HandleFunc("DELETE /api/wydarzenia", s.application.HandleDeleteWydarzenie)

	// Przypomnienia
	s.mux.HandleFunc("GET /api/przypomnienia", s.application.HandleGetPrzypomnienia)
	s.mux.HandleFunc("GET /api/przypomnienia/get", s.application.HandleGetPrzypomnienie)
	s.mux.HandleFunc("POST /api/przypomnienia", s.application.HandleCreatePrzypomnienie)
	s.mux.HandleFunc("PUT /api/przypomnienia", s.application.HandleUpdatePrzypomnienie)
	s.mux.HandleFunc("DELETE /api/przypomnienia", s.application.HandleDeletePrzypomnienie)

	// Produkty
	s.mux.HandleFunc("GET /api/produkty", s.application.HandleGetProdukty)
	s.mux.HandleFunc("GET /api/produkty/get", s.application.HandleGetProdukt)
	s.mux.HandleFunc("GET /api/produkty/ulubione", s.application.HandleGetProduktyUlubione)
	s.mux.HandleFunc("POST /api/produkty", s.application.HandleCreateProdukt)
	s.mux.HandleFunc("PUT /api/produkty", s.application.HandleUpdateProdukt)
	s.mux.HandleFunc("DELETE /api/produkty", s.application.HandleDeleteProdukt)

	// Listy Zakupow
	s.mux.HandleFunc("GET /api/listy-zakupow", s.application.HandleGetListyZakupow)
	s.mux.HandleFunc("GET /api/listy-zakupow/get", s.application.HandleGetListaZakupow)
	s.mux.HandleFunc("POST /api/listy-zakupow", s.application.HandleCreateListaZakupow)
	s.mux.HandleFunc("PUT /api/listy-zakupow", s.application.HandleUpdateListaZakupow)
	s.mux.HandleFunc("DELETE /api/listy-zakupow", s.application.HandleDeleteListaZakupow)

	// Pozycje Listy Zakupow
	s.mux.HandleFunc("GET /api/pozycje-listy", s.application.HandleGetPozycjeListy)
	s.mux.HandleFunc("GET /api/pozycje-listy/get", s.application.HandleGetPozycja)
	s.mux.HandleFunc("POST /api/pozycje-listy", s.application.HandleCreatePozycja)
	s.mux.HandleFunc("PUT /api/pozycje-listy", s.application.HandleUpdatePozycja)
	s.mux.HandleFunc("DELETE /api/pozycje-listy", s.application.HandleDeletePozycja)
	s.mux.HandleFunc("POST /api/pozycje-listy/kupione", s.application.HandleMarkPozycjaAsBought)

	// Stany Magazynowe
	s.mux.HandleFunc("GET /api/stany-magazynowe", s.application.HandleGetStanyMagazynowe)
	s.mux.HandleFunc("GET /api/stany-magazynowe/get", s.application.HandleGetStanMagazynowy)
	s.mux.HandleFunc("POST /api/stany-magazynowe", s.application.HandleCreateStanMagazynowy)
	s.mux.HandleFunc("PUT /api/stany-magazynowe", s.application.HandleUpdateStanMagazynowy)
	s.mux.HandleFunc("DELETE /api/stany-magazynowe", s.application.HandleDeleteStanMagazynowy)

	// Historia Stanu Zapasow
	s.mux.HandleFunc("GET /api/historia-zapasow", s.application.HandleGetHistoriaStanuZapasow)
	s.mux.HandleFunc("POST /api/historia-zapasow", s.application.HandleCreateHistoriaStanuZapasow)
}

func (s *Server) listenAndServe() {
	s.logger.Info("Starting server on port ", s.cfg.Server.Port)
	err := http.ListenAndServe(":"+s.cfg.Server.Port, s.mux)
	if err != nil {
		s.logger.Fatal(err)
	}
}
