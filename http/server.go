package http

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type loggerService interface {
	RunTest(ctx context.Context, testID string, rows int) (*time.Duration, error)
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
	id := vars["id"]
	rows := r.FormValue("rows")
	ctx := r.Context()

	fmt.Fprintf(w, "test id: %v\nrows: %v\n", vars["id"], rows)

	rowsI, err := strconv.Atoi(rows)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	diff, err := s.logger.RunTest(ctx, id, rowsI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Test done, time: %vms", diff.Milliseconds())
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
