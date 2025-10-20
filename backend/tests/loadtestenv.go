package tests

import (
	"log"
	"os"
	"strings"
)

func LoadTestEnv() {
	if err := loadDotEnv(".env.test"); err != nil {
		log.Fatalf("failed to load .env.test: %v", err)
	}
}

// simple loader
func loadDotEnv(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(f), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}
		os.Setenv(kv[0], kv[1])
	}
	return nil
}
