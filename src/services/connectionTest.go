package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	config "hacker-rank-lambda/src"

	_ "github.com/lib/pq"
)

func getTableInfo() {
    // This code is only for testing, don't use in the project
    db, err := sql.Open("postgres", config.DB_CONNECTION)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'hard_skills'")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var tableName string
        err := rows.Scan(&tableName)
        if err != nil {
            panic(err)
        }

        fmt.Printf("Tabla: %s\n", tableName)

        var count int
        err = db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM hard_skills.%s", tableName)).Scan(&count)
        if err != nil {
            panic(err)
        }

        fmt.Printf("  Total de registros: %d\n", count)

        tableRows, err := db.Query(fmt.Sprintf("SELECT * FROM hard_skills.%s", tableName))
        if err != nil {
            panic(err)
        }
        defer tableRows.Close()

        for tableRows.Next() {
            columns, err := tableRows.Columns()
            if err != nil {
                panic(err)
            }

            values := make([]interface{}, len(columns))

            for i := range columns {
                values[i] = new(interface{})
            }

            err = tableRows.Scan(values...)
            if err != nil {
                panic(err)
            }

            for i, column := range columns {
                if b, ok := (*(values[i].(*interface{}))).([]byte); ok {
                    if len(b) > 0 && (b[0] == '{' || b[0] == '[') {
                        var s interface{}
                        err := json.Unmarshal(b, &s)
                        if err != nil {
                            panic(err)
                        }
                        fmt.Printf("  %s: %v\n", column, s)
                    } else {
                        fmt.Printf("  %s: %s\n", column, string(b))
                    }
                } else {
                    fmt.Printf("  %s: %v\n", column, *(values[i].(*interface{})))
                }
            }
        }
    }
}