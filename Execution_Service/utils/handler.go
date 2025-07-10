package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/naman1402/eigen-bootcamp-assignment-2/Execution_Service/config"

	"github.com/gin-gonic/gin"
)

type TaskRequest struct {
	TaskDefinitionId int    `json:"taskDefinitionId"`
	Address          string `json:"address"`
	AchievementType  string `json:"achievementType"`
}

func ExecuteTask(c *gin.Context) {
	log.Println("Executing Task for AVS")

	if c.Request.Method == http.MethodPost {
		var req TaskRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Call DummyAPI based on achievement type
		var apiURL string
		if req.AchievementType == "github" {
			apiURL = fmt.Sprintf("%s/api/github?address=%s", config.DUMMY_API_URL, req.Address)
		} else if req.AchievementType == "psn" {
			apiURL = fmt.Sprintf("%s/api/psn?address=%s", config.DUMMY_API_URL, req.Address)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported achievement type"})
			return
		}

		resp, err := http.Get(apiURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch off-chain data"})
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		// Generate a simple proof: hash of the returned data
		hash := sha256.Sum256(body)
		proofOfTask := fmt.Sprintf("%x", hash[:])

		var data interface{}
		json.Unmarshal(body, &data)

		response := gin.H{
			"proofOfTask":      proofOfTask,
			"data":             data,
			"taskDefinitionId": req.TaskDefinitionId,
		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid method"})
	}
}