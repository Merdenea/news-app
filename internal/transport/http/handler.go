package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"news-app/internal/news"
)

const (
	newsPath    = "/news"
	sourcesPath = "/sources"
)

type NewsFetcher interface {
	FetchFromAllSources() ([]news.Source, error)
}

type Service interface {
	GetAllSources() ([]news.Source, error)
}

func NewHandler(fetcher NewsFetcher, port string, svc Service) *Handler {
	return &Handler{
		router:      mux.NewRouter(),
		newsFetcher: fetcher,
		port:        port,
		service:     svc,
	}
}

type Handler struct {
	router      *mux.Router
	newsFetcher NewsFetcher
	service     Service
	port        string
}

func (h *Handler) Serve() {
	h.router.HandleFunc(newsPath, h.GetNews).Methods(http.MethodGet)
	h.router.HandleFunc(sourcesPath, h.GetSources).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(h.port, h.router))
}

func (h *Handler) GetNews(w http.ResponseWriter, r *http.Request) {
	//TODO: read request params for categories, etc.
	log.Println("handling get news request")
	w.Header().Set("Content-Type", "application/json")
	newsArticles, err := h.newsFetcher.FetchFromAllSources()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//TODO: return a response body with an error message
		return
	}
	data, err := json.Marshal(newsArticles)
	if err != nil {
		log.Println("error marshaling response")
		w.WriteHeader(http.StatusInternalServerError)
		//TODO: return a response body with an error message
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) GetSources(w http.ResponseWriter, r *http.Request) {
	log.Println("handling get sources request")
	w.Header().Set("Content-Type", "application/json")
	sources, err := h.service.GetAllSources()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//TODO: return a response body with an error message
		return
	}
	data, err := json.Marshal(sources)
	if err != nil {
		log.Println("error marshaling response")
		w.WriteHeader(http.StatusInternalServerError)
		//TODO: return a response body with an error message
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//TODO: GET GetSource, POST AddSource, etc.
