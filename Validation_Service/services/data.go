package services

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

// FetchUserData fetches user data from DummyAPI
func FetchUserData(address, achievementType string) ([]byte, error) {
	apiURL := fmt.Sprintf("%s/api/%s?address=%s", os.Getenv("DUMMY_API_URL"), achievementType, address)
	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch off-chain data from DummyAPI")
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}