package http

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type loggerService interface {
	Store(ctx context.Context, testID string) error
}

// Server will perform operations over http.
type Server interface {
	// Open will setup a new listener on the given port.
	Open() error

	// Close will close the connection if it's open.
	Close()

	// Handler returns a http handler with all routes in place.
	Handler() http.Handler
}

// Server represents an HTTP server.
type server struct {
	listener net.Listener
	addr     string
	db       *sql.DB
	logger   loggerService
}

// Open will setup a tcp listener and serve the http requests.
func (s *server) Open() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	fmt.Sprintf("HTTP server listening %v", s.addr)

	// Save listener so we can decide to close it later.
	s.listener = ln

	// Start HTTP server.
	server := http.Server{
		Handler: s.Handler(),
	}
	return server.Serve(s.listener)
}

// Close will close the socket if it's open.
func (s *server) Close() {
	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
	}
}

func (s *server) Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/test/{id}", s.TestHandler)
	return r
}

func (s *server) TestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]

	rows := r.FormValue("rows")

	fmt.Fprintf(w, "got id: %v rows:%v\n", vars["id"], rows)

	ctx := r.Context()

	err := s.logger.Store(ctx, id)
	if err != nil {
		fmt.Fprintf(w, "something went wrong %v\n", err)
	}
	w.Write([]byte("Done"))
}

// NewServer returns a new instance of Server.
func NewServer(
	addr string,
	db *sql.DB,
	service loggerService,
) Server {
	return &server{
		addr:   addr,
		db:     db,
		logger: service,
	}
}
