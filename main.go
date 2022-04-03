package main

import (
	"log"
	"os"

	"github.com/c-m-hunt/golf-booking-template/pkg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	params := getParams()

	params.GenerateTemplate(params)

}

func getParams() pkg.Params {
	month := os.Getenv("MONTH")
	day := os.Getenv("DAY")
	time := os.Getenv("TIME")
	players := os.Getenv("PLAYERS")
	course := os.Getenv("COURSE")
	return pkg.Params{
		Month:   month,
		Day:     day,
		Time:    time,
		Players: players,
		Course:  course,
	}
}
