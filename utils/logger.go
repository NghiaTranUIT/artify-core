package utils

import (
	"fmt"
	"time"
)

func LogInfo(a ...interface{}) {
	info := fmt.Sprint(a)
	fmt.Println("[INFO] ", time.Now(), info)
}

func LogError(a ...interface{}) {
	error := fmt.Sprint(a)
	fmt.Println("[ERROR] ❌", time.Now(), error)
}

func LogWarning(a ...interface{}) {
	warning := fmt.Sprint(a)
	fmt.Println("[WARNING] 🔥", time.Now(), warning)
}
