package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AchievementResponse struct {
	Address   string      `json:"address"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Details   interface{} `json:"details"`
}

// Fetches user achievement/profile data from DummyAPI
func GetAchievement(address, achievementType, dummyApiUrl string) (*AchievementResponse, error) {
	var apiURL string
	if achievementType == "github" {
		apiURL = fmt.Sprintf("%s/api/github?address=%s", dummyApiUrl, address)
	} else if achievementType == "psn" {
		apiURL = fmt.Sprintf("%s/api/psn?address=%s", dummyApiUrl, address)
	} else {
		return nil, fmt.Errorf("unsupported achievement type: %s", achievementType)
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("DummyAPI returned status %d", resp.StatusCode)
	}

	var achievement AchievementResponse
	if err := json.NewDecoder(resp.Body).Decode(&achievement); err != nil {
		return nil, err
	}

	return &achievement, nil
}