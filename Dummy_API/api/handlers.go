package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/naman1402/eigen-bootcamp-assignment-2/Dummy_API/types"
)

// --- Static data ---

var githubUsers = map[string]types.GithubProfile{
    "0x123": {
        Address:  "0x123",
        Username: "aliceGH",
        Email:    "alice@github.com",
        Repos:    []string{"eigenlayer-demo", "nft-mvp"},
        Badges:   []string{"100-commits", "oss-contributor"},
    },
    "0x456": {
        Address:  "0x456",
        Username: "bobGH",
        Email:    "bob@github.com",
        Repos:    []string{"psn-integration", "dummy-api"},
        Badges:   []string{"first-pr", "starred-repo"},
    },
}

var psnUsers = map[string]types.PSNProfile{
    "0xabc": {
        Address:  "0xabc",
        Username: "alicePSN",
        Email:    "alice@psn.com",
        Games:    []string{"God of War", "Horizon Zero Dawn"},
        Trophies: []string{"Platinum", "Story Complete"},
    },
    "0xdef": {
        Address:  "0xdef",
        Username: "bobPSN",
        Email:    "bob@psn.com",
        Games:    []string{"Spider-Man", "Gran Turismo"},
        Trophies: []string{"Speedster", "Collector"},
    },
}

// --- Handlers ---

func GetGithubProfile(c *gin.Context) {
    address := c.Query("address")
    if profile, ok := githubUsers[address]; ok {
        c.JSON(http.StatusOK, profile)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "GitHub profile not found for this address"})
    }
}

func GetPSNProfile(c *gin.Context) {
    address := c.Query("address")
    if profile, ok := psnUsers[address]; ok {
        c.JSON(http.StatusOK, profile)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "PSN profile not found for this address"})
    }
} 