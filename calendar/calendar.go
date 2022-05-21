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

	start, end := time.Now().Add(-12*time.Hour), time.Now().Add(24*time.Hour) //only look at the next 12 hours to avoid confusion
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

func Before(t *time.Time) bool {
	now := time.Now()
	return now.Before(*t)
}

func After(t *time.Time) bool {
	now := time.Now()
	return now.After(*t)
}

// sort events by start time from earliest to latest
func sortEvents(events []gocal.Event) []gocal.Event {
	for i := 0; i < len(events); i++ {
		for j := i + 1; j < len(events); j++ {
			if events[i].Start.After(*events[j].Start) {
				events[i], events[j] = events[j], events[i]
			}
		}
	}
	return events
}

func CurrentEvents(events []gocal.Event) []gocal.Event {
	var currentEvents []gocal.Event
	for _, event := range events {
		if InTimeSpan(event.Start, event.End) {
			currentEvents = append(currentEvents, event)
		}
	}
	currentEvents = sortEvents(currentEvents)
	return currentEvents
}

func PreviousEvents(events []gocal.Event) []gocal.Event {
	var prevEvents []gocal.Event
	for _, event := range events {
		if After(event.End) {
			prevEvents = append(prevEvents, event)
		}
	}
	prevEvents = sortEvents(prevEvents)
	return prevEvents
}

func UpcomingEvents(events []gocal.Event) []gocal.Event {
	var upcomingEvents []gocal.Event
	for _, event := range events {
		if Before(event.Start) {
			upcomingEvents = append(upcomingEvents, event)
		}
	}
	upcomingEvents = sortEvents(upcomingEvents)
	return upcomingEvents
}
