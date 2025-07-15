package services

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Validate(proofOfTask, address, achievementType string) (bool, error) {
	// Fetch user data from DummyAPI
	apiURL := fmt.Sprintf("%s/api/%s?address=%s", os.Getenv("DUMMY_API_URL"), achievementType, address)
	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch off-chain data from DummyAPI")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Hash the fetched data
	hash := sha256.Sum256(body)
	computedProof := fmt.Sprintf("%x", hash[:])

	// Compare with provided proofOfTask
	return computedProof == proofOfTask, nil
}