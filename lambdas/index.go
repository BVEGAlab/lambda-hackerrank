package lambdas

import (
    "context"
    "hacker-rank-lambda/src/controllers"
)

func Handler(ctx context.Context) (string, error) {
    err := controllers.HackerRankController()
    if err != nil {
        return "", err
    }
    return "Success", nil
}