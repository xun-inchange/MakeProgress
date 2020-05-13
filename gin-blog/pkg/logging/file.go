package logging

import (
	"fmt"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeForMat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	preFilePath := getLogFileFullPath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeForMat), LogFileExt)
	return fmt.Sprintf("%s%s", preFilePath, suffixPath)
}

func openLogFile(filePath string) string {
	return ""
}
