package server

import (
	"dev11/calendar"
	"log"
	"net/http"
	"time"
)

type Server struct {
	events  *calendar.EventsManager
	routing *http.ServeMux
}

func NewServer() *Server {
	return &Server{calendar.NewEventsManager(), http.NewServeMux()}
}

// Logger MiddleWare
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s, %s, %s\n", r.Method, r.URL, time.Since(start))
	}
}

func (s *Server) Run() {
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: s.routing,
	}
	s.routing.HandleFunc("/create_event", Logger(s.CreateEvent))
	s.routing.HandleFunc("/update_event", Logger(s.UpdateEvent))
	s.routing.HandleFunc("/delete_event", Logger(s.DeleteEvent))
	s.routing.HandleFunc("/events_for_day", Logger(s.EventsForDay))
	s.routing.HandleFunc("/events_for_week", Logger(s.EventsForWeek))
	s.routing.HandleFunc("/events_for_month", Logger(s.EventsForMonth))

	log.Println("Server started")
	log.Fatal(srv.ListenAndServe())
}
