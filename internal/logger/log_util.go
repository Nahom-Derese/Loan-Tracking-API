package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var logFilePath = filepath.Join("logs", "app.log")
var logFolder = filepath.Join("logs")

func init() {
	// Ensure the logs directory exists
	if _, err := os.Stat(logFolder); os.IsNotExist(err) {
		err = os.Mkdir(logFolder, 0755)
		if err != nil {
			fmt.Printf("Failed to create logs directory: %v\n", err)
			return
		}
	}
}

func logEvent(event string, details string) error {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer logFile.Close()

	logEntry := fmt.Sprintf("%s - %s: %s\n",
		time.Now().Format(time.RFC3339),
		event,
		details,
	)

	if _, err := logFile.WriteString(logEntry); err != nil {
		return fmt.Errorf("failed to write log entry: %v", err)
	}

	return nil
}

// LogLoginAttempt logs both successful and failed login attempts
func LogLoginAttempt(username string, success bool) error {
	event := "Login Attempt"
	status := "FAILED"
	if success {
		status = "SUCCESSFUL"
	}
	details := fmt.Sprintf("Email: %s, Status: %s", username, status)

	return logEvent(event, details)
}

func ResetPasswordAttempt(username string, success bool) error {
	event := "Password Reset Attempt"
	status := "FAILED"
	if success {
		status = "SUCCESSFUL"
	}
	details := fmt.Sprintf("Email: %s, Status: %s", username, status)

	return logEvent(event, details)
}
