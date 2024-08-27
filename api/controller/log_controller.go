package controller

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// AdminLogsHandler allows the admin to view application logs
func AdminLogsHandler(c *gin.Context) {
	var logFilePath = filepath.Join("..", "..", "logs", "app.log")

	// Read the log file
	logs, err := os.ReadFile(logFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read logs"})
		return
	}

	// Return the logs as plain text
	c.Data(http.StatusOK, "text/plain; charset=utf-8", logs)
}
