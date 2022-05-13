package calendar

import (
	"log"
	"net/http"
	"time"

	"github.com/SinclearClan/GoCDAP/config"
	"github.com/apognu/gocal"
)

func grabCalendarFromIcal(cfg *config.Config) *gocal.Gocal {
	onlinecal, err := http.Get(cfg.Calendar.Url+cfg.Calendar.Path)
	if err != nil {
		panic(err)
	}
	defer onlinecal.Body.Close()

	start, end := time.Now(), time.Now().Add(12*time.Hour) //only look at the next 12 hours to avoid confusion
	c := gocal.NewParser(onlinecal.Body)
	c.Start, c.End = &start, &end
	c.Parse()
	return c
}

func Update(cfg *config.Config) *gocal.Gocal {
	if cfg.Calendar.Type == "ical" {
		cal := grabCalendarFromIcal(cfg)
		return cal
	} else {
		log.Fatal("Unknown calendar type")
	}
	return nil
}

func InTimeSpan(start, end *time.Time) bool {
	now := time.Now()
	return now.After(*start) && now.Before(*end)
}
