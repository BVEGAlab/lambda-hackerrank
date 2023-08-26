package utils

import (
	"fmt"
    "strings"
)


func FormatArray(arr []string) string {
    var sb strings.Builder
    sb.WriteString("[")

    for i, s := range arr {
        if i > 0 {
            sb.WriteString(",")
        }
        sb.WriteString(fmt.Sprintf("'%s'", s))
    }

    sb.WriteString("]")
    return sb.String()
}

func Contains(slice []string, str string) bool {
    for _, s := range slice {
        if s == str {
            return true
        }
    }
    return false
}