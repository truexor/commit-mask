package utils

import "fmt"

func LogSuccess(msg string) {
	fmt.Println("✅ \033[32m" + msg + "\033[0m")
}

func LogUpdate(msg string) {
	fmt.Println("🔄 \033[34m" + msg + "\033[0m")
}

func LogInfo(msg string) {
	fmt.Println("- \033[90m" + msg + "\033[0m")
}

func LogWarning(msg string) {
	fmt.Println("⚠️ \033[33m" + msg + "\033[0m")
}
