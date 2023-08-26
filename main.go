package main

import (
    "hacker-rank-lambda/src/controllers"
)

func main() {
    err := controllers.HackerRankController()
    if err != nil {
        panic(err)
    }

}