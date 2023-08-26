package config

import (
    "fmt"
    "os"
    "strings"
)

var (
    TOKEN           = getEnv("token", "")
    TESTS_TO_SEARCH = getEnvSlice("tests_to_search", []string{""})
    DB_CONNECTION   = getEnv("db_connection", "")
)

func getEnv(name string, fallback string) string {
    if value, exists := os.LookupEnv(name); exists {
        return value
    }

    if fallback != "" {
        return fallback
    }

    panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}

func getEnvSlice(name string, fallback []string) []string {
    if value, exists := os.LookupEnv(name); exists {
        return strings.Split(value, ",")
    }

    if len(fallback) > 0 {
        return fallback
    }

    panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}