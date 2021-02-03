package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"squid/postgres-stress-test/domain"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// TrackerRepository tracker repository.
type IncomingRepository interface {
	Store(log domain.IncomingLog) (*uuid.UUID, error)
	GetStats(testID string) (*time.Time, *time.Time, error)
	FindAll() (*[]domain.IncomingLog, error)
}

// Service encapsulates the business details of service.
type Service struct {
	repo IncomingRepository
}

// NewService returns a new service with all dependencies set up.
func NewService(in IncomingRepository) *Service {
	return &Service{
		repo: in,
	}
}

func (s *Service) RunTest(ctx context.Context, testID string, rows int) (*time.Duration, error) {
	data, err := readFile()
	if err != nil {
		return nil, fmt.Errorf("coudln't read json file: %v", err)
	}

	data.TestID = testID

	err = s.store(data, rows)
	if err != nil {
		return nil, fmt.Errorf("error storing data: %v", err)
	}

	min, max, err := s.repo.GetStats(testID)
	if err != nil {
		return nil, fmt.Errorf("error to getStat: %v", err)
	}
	diff := max.Sub(*min)
	return &diff, nil
}

func (s *Service) store(data *domain.IncomingLog, rows int) error {
	var wg sync.WaitGroup
	for i := 1; i <= rows; i++ {
		data.RowNumber = i
		wg.Add(1)
		go func(log domain.IncomingLog, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			res, err := s.repo.Store(log)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(i, ": ", res)
		}(*data, i, &wg)
	}
	wg.Wait()

	return nil
}

//
//	min, max, err := s.repo.GetStats(testID)
//	if err != nil {
//		return fmt.Errorf("error to getStat: %v", err)
//	}
//
//	fmt.Printf("min time: %v", min)
//	fmt.Printf("max time: %v", max)
//
//	return nil
//}

func readFile() (*domain.IncomingLog, error) {
	dir := "mock/"
	jsonFile, err := os.Open(dir + "incoming_log.json")
	if err != nil {
		files := getFilesInFolder(dir)
		fmt.Println(files)
		return nil, fmt.Errorf("couldn't open the JSON file %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var log domain.IncomingLog
	if err = json.Unmarshal(byteValue, &log); err != nil {
		return nil, fmt.Errorf("error decoding JSON file %v", err)
	}
	return &log, nil
}

func getFilesInFolder(dir string) string {
	filesNames := ""
	files, err := ioutil.ReadDir("/mock")

	if err != nil {
		filesNames = err.Error()
		return filesNames
	}
	f := make([]string, 0, 20)
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), "affiliate") {
			f = append(f, file.Name())
		}
	}
	return strings.Join(f, ",")
}
