package postgres

import (
	"database/sql"
	"fmt"
	"squid/postgres-stress-test/domain"
	"squid/postgres-stress-test/logger"
	"time"

	"github.com/google/uuid"
)

type incomingRepository struct {
	db *sql.DB
}

// NewEventRepository creates new events repository.
func NewIncomingRepository(db *sql.DB) logger.IncomingRepository {
	return &incomingRepository{
		db: db,
	}
}

func (r *incomingRepository) Store(log domain.IncomingLog) (*uuid.UUID, error) {
	query := `
	INSERT INTO public.incoming_logs(
		request_id, request, response, ip, provider, general_provider, request_at, response_at, response_time, added_at, events, fallback_reason, row_number, test_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	RETURNING
		request_id
	`

	result := r.db.QueryRow(
		query,
		uuid.New(),
		log.Request,
		log.Response,
		log.IP,
		log.FallbackProvider,
		log.GeneralProvider,
		log.RequestAt,
		log.ResponseAt,
		log.ResponseTime,
		log.ResponseAt,
		log.Events,
		log.FallbackReason,
		log.RowNumber,
		log.TestID,
	)

	var id uuid.UUID
	if err := result.Scan(&id); err != nil {
		return nil, fmt.Errorf("error storing data: %v", err)
	}
	//fmt.Printf("%v: %v\n", i, id)
	return &id, nil
}

func (r *incomingRepository) GetStats(testID string) (*time.Time, *time.Time, error) {
	query := `
	SELECT min(log_at), max(log_at) test_id
	FROM public.incoming_logs
	where test_id = $1;
	`
	result := r.db.QueryRow(query, testID)
	var minTime, maxTime time.Time
	if err := result.Scan(&minTime, &maxTime); err != nil {
		return nil, nil, err
	}
	return &minTime, &maxTime, nil
}

func (r *incomingRepository) FindAll() (*[]domain.IncomingLog, error) {
	query := `
	SELECT request, response, ip, provider, general_provider, request_at, response_at, response_time, added_at, fallback_reason/*, events*/
	FROM public.incoming_logs;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error query performing, %v", err)
	}

	var items []domain.IncomingLog
	for rows.Next() {
		log := domain.IncomingLog{}
		err = rows.Scan(&log.Request, &log.Response, &log.IP, &log.FallbackProvider, &log.GeneralProvider,
			&log.RequestAt, &log.ResponseAt, &log.ResponseTime, &log.AddedAt, &log.FallbackReason /*, &log.Events*/)
		if err != nil {
			return nil, err
		}
		items = append(items, log)
	}
	rows.Close()

	return &items, nil
}
