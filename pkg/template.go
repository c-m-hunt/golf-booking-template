package pkg

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
)

type TemplateParams struct {
	Title string
	CronDay string
	CronMonth string
	Date string
	Time string
	Course string
	Players string
}

func (p Params) GenerateTemplate(params Params) {
	startDate := getFirstDayOfMonth(params.Month)
	compareMonth := startDate.Month()

	for true {
		if startDate.Weekday().String() == p.Day {
			bookingDate := startDate.AddDate(0, 0, -11)
			fmt.Printf(
				"%s %s\n",
				bookingDate.Format("2006-01-02"),
				startDate.Format("2006-01-02"),
			)
			tp := TemplateParams{
				Title: strings.ToLower(startDate.Format("02-Jan-2006")),
				CronDay: bookingDate.Format("02"),
				CronMonth: bookingDate.Format("01"),
				Date: startDate.Format("02/01/06"),
				Time: p.Time,
				Course: p.Course,
				Players: p.Players,
			}

			file, err := os.Create(fmt.Sprintf("out/%s.yaml", startDate.Format("02-01-2006")))
			if err != nil { panic(err) }
			defer file.Close()

			tmpl, err := template.ParseFiles("template/cron.yaml")
			if err != nil { panic(err) }
			err = tmpl.Execute(file, tp)
			if err != nil { panic(err) }

			fmt.Printf("%s\n", tp)
		}

		startDate = startDate.AddDate(0, 0, 1)
		if startDate.Month() != compareMonth {
			break
		}
	}
}

func getFirstDayOfMonth(monthStr string) time.Time {
	dateString := fmt.Sprintf("01 %s", monthStr)
	month, err := time.Parse("01 Jan,2006", dateString)
	if err != nil {
		panic(err)
	}
	return month.AddDate(0, 0, -month.Day()+1)
}
