package main

import (
    "github.com/gin-gonic/gin"
    "github.com/naman1402/eigen-bootcamp-assignment-2/Dummy_API/api"
)

func main() {
    r := gin.Default()

    // Register API routes
    r.GET("/api/github", api.GetGithubProfile)
    r.GET("/api/psn", api.GetPSNProfile)

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
} 