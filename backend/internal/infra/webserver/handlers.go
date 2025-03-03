package webserver

import (
	"encoding/json"
	"github.com/tonnytg/makemoneytarget/internal/domain/target"
	"github.com/tonnytg/makemoneytarget/internal/infra/database"
	"io"
	"log"
	"net/http"
)

func ErrorHandling(w http.ResponseWriter, err error, requestID string) {

	errorJsonFormat := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorJsonFormat)
	log.Printf("[%s] - error: %s\n", requestID, err)

}

func createTarget(w http.ResponseWriter, r *http.Request) {

	requestID := r.Header.Get("X-Request-ID")

	// Create a new target
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ErrorHandling(w, err, requestID)
		return
	}
	defer r.Body.Close()

	dtoTarget := TargetDTO{}
	if err := json.Unmarshal(body, &dtoTarget); err != nil {
		ErrorHandling(w, err, requestID)
		return
	}

	t := target.Target{
		UUID:          dtoTarget.ID,
		Name:          dtoTarget.Name,
		Description:   dtoTarget.Description,
		Status:        dtoTarget.Status,
		TargetAmount:  dtoTarget.TargetAmount,
		CurrentAmount: dtoTarget.CurrentAmount,
		StartDate:     dtoTarget.StartDate,
		EndDate:       dtoTarget.EndDate,
		Members:       dtoTarget.Members,
	}

	repo := database.NewTargetRepositorySqlite3()
	serv := target.NewService(repo)

	targetCreated, err := serv.Save(&t)
	if err != nil {
		ErrorHandling(w, err, requestID)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(targetCreated)
}

func handleTarget(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.Method == http.MethodPost:
		createTarget(w, r)
	case r.Method == http.MethodGet:
		// TODO: Implement get all targets
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
