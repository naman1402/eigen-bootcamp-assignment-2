package handlers
import (
	"encoding/json"
	"log"
	"net/http"
	"Validation_Service/services"
)

type ValidateRequest struct {
	ProofOfTask     string `json:"proofOfTask"`
	Address         string `json:"address"`
	AchievementType string `json:"achievementType"`
}

// ValidateTask handles the POST request to `/task/validate`
func ValidateTask(w http.ResponseWriter, r *http.Request) {
	var req ValidateRequest
	if r.Body != nil && r.Body != http.NoBody {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			log.Println("Error decoding JSON body:", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	}

	log.Printf("proofOfTask: %v, address: %v, achievementType: %v\n", req.ProofOfTask, req.Address, req.AchievementType)

	result, err := services.Validate(req.ProofOfTask, req.Address, req.AchievementType)
	if err != nil {
		log.Printf("Validation error: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := services.NewCustomError("Error during validation step", nil)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	log.Printf("Vote: %s", func() string {
		if result {
			return "Approve"
		}
		return "Not Approved"
	}())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := services.NewCustomResponse(map[string]interface{}{
		"result":    result,
	}, "Task validated successfully")
	json.NewEncoder(w).Encode(response)
}