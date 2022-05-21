package discord

import (
	"fmt"
	"regexp"
	"time"

	"github.com/SinclearClan/GoCDAP/calendar"
	"github.com/SinclearClan/GoCDAP/config"
	"github.com/apognu/gocal"
	"github.com/hugolgst/rich-go/client"
	"jaytaylor.com/html2text"
)

var (
	now gocal.Event
	nowExists bool
	nowTitle string
	nowDescription string
	next gocal.Event
	nextExists bool
	nextTitle string
	nextDescription string
)

func SetActivity(cfg *config.Config) {

	re := regexp.MustCompile(`[a-zA-Z0-9[:blank:]-–—?!.,+]*$`)

	err := client.Login("836637382509068319")
	if err != nil {
		panic(err)
	}

	cal := calendar.Update(cfg)
	pre := calendar.PreviousEvents(cal.Events)
	cur := calendar.CurrentEvents(cal.Events)
	upc := calendar.UpcomingEvents(cal.Events)

	for _, event := range pre {
		fmt.Println("Previous: " + event.Summary)
	}
	for _, event := range cur {
		fmt.Println("Current: " + event.Summary)
	}
	for _, event := range upc {
		fmt.Println("Upcoming: " + event.Summary)
	}

	if len(cur) > 0 {
		nowExists = true
		now = cur[0]
		nowTitle = now.Summary
		nowDescription, err = html2text.FromString(now.Description, html2text.Options{PrettyTables: true})
		fmt.Printf("nowDescription 1: %s\n", nowDescription)
		if err != nil {
			panic(err)
		}
		for _, match := range re.FindAllString(nowDescription, -1) {
			nowDescription = " " + match
		}
		if len(nowDescription) < 3 {
			nowDescription = ""
		}
		fmt.Printf("nowDescription 2: %s\n", nowDescription)
	} else {
		nowExists = false
		nowTitle = "-"
		nowDescription = ""
		fmt.Printf("nowDescription 3: %s\n", nowDescription)
	}

	if len(upc) > 0 {
		nextExists = true
		next = upc[0]
		nextTitle = next.Summary
		nextDescription, err = html2text.FromString(next.Description, html2text.Options{PrettyTables: true})
		if err != nil {
			panic(err)
		}
		for _, match := range re.FindAllString(nextDescription, -1) {
			nextDescription = " " + match
		}
		if len(nextDescription) < 3 {
			nextDescription = ""
		}
	} else {
		nextExists = false
		nextTitle = "-"
		nextDescription = ""
	}

	if nowExists && nextExists {
		fmt.Printf("nowDescription 4: %s\n", nowDescription)
		err = client.SetActivity(client.Activity{
			Details:    "Aktuell: " + nowTitle + nowDescription,
			State:      "Danach: " + next.Start.Format("15:04") + " – " + nextTitle + nextDescription,
			Timestamps: &client.Timestamps{
				Start: now.Start,
				End:  now.End,
			},
			/*Buttons: []*client.Button{
				{
					Label: "Kalender anzeigen",
					Url:   "https://nextcloud.sinclear.de/index.php/apps/calendar/p/ffe5tLHMEjmcPLxD",
				},
			},*/
		})
		if err != nil {
			panic(err)
		}
	} else if nowExists && !nextExists {
		err = client.SetActivity(client.Activity{
			Details:    "Aktuell: " + nowTitle + nowDescription,
			State:      "Danach: " + nextTitle,
			Timestamps: &client.Timestamps{
				Start: now.Start,
				End:  now.End,
			},
		})
		if err != nil {
			panic(err)
		}
	} else if !nowExists && nextExists {
		err = client.SetActivity(client.Activity{
			Details:    "Aktuell: " + nowTitle,
			State:      "Danach: " + next.Start.Format("15:04") + " – " + nextTitle + nextDescription,
		})
		if err != nil {
			panic(err)
		}
	} else {
		err = client.SetActivity(client.Activity{
			Details:    "Aktuell: " + nowTitle,
			State:      "Danach: " + nextTitle,
		})
		if err != nil {
			panic(err)
		}
	}

	time.Sleep(time.Second * 60)

}
