package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorules/zen-go"
	"go-editor/nodes"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/api/health", healthHandler)
	r.POST("/api/simulate", simulateHandler)

	r.Run(":8080")
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

type EvaluateRequest struct {
	Context json.RawMessage `json:"context" binding:"required"`
	Content json.RawMessage `json:"content" binding:"required"`
}

func simulateHandler(c *gin.Context) {
	var request EvaluateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	engine := zen.NewEngine(zen.EngineConfig{
		CustomNodeHandler: nodes.CustomNodeHandler,
	})

	decision, err := engine.CreateDecision(request.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := decision.Evaluate(request.Context)
	if err != nil {
		var res gin.H
		if err := json.Unmarshal([]byte(err.Error()), &res); err == nil {
			c.JSON(http.StatusBadRequest, res)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":      result.Result,
		"trace":       result.Trace,
		"performance": result.Performance,
	})
}
