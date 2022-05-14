package discord

import (
	"regexp"
	"time"

	"github.com/SinclearClan/GoCDAP/calendar"
	"github.com/SinclearClan/GoCDAP/config"
	"github.com/hugolgst/rich-go/client"
	"jaytaylor.com/html2text"
)

func SetActivity(cfg *config.Config) {
	err := client.Login("836637382509068319")
	if err != nil {
		panic(err)
	}

	cal := calendar.Update(cfg)

	if len(cal.Events) > 0 { 																	// ONE OR MORE EVENTS
		first := cal.Events[0]

		re := regexp.MustCompile(`[a-zA-Z0-9[:blank:]–-—?!.,+]*$`)
		firstDescription, err := html2text.FromString(first.Description, html2text.Options{PrettyTables: true})
		if err != nil {
			panic(err)
		}
		for _, match := range re.FindAllString(firstDescription, -1) {
			firstDescription = " " + match
		}
		if len(firstDescription) < 3 {
			firstDescription = ""
		}

		if calendar.InTimeSpan(first.Start, first.End) {										// ONE OR MORE EVENTS, FIRST ENTRY IS RIGHT NOW

			if len(cal.Events) > 1 {															// ONE OR MORE EVENTS, FIRST ENTRY IS RIGHT NOW, A SECOND EVENT EXISTS
				second := cal.Events[1]

				secondDescription, err := html2text.FromString(second.Description, html2text.Options{PrettyTables: true})
				if err != nil {
					panic(err)
				}
				for _, match := range re.FindAllString(secondDescription, -1) {
					secondDescription = " " + match
				}
				if len(secondDescription) < 3 {
					secondDescription = ""
				}

				err = client.SetActivity(client.Activity{
					Details:    "Aktuell: " + first.Summary + firstDescription,
					State:      "Danach: " + second.Start.Format("15:04") + " – " + second.Summary + secondDescription,
					Timestamps: &client.Timestamps{
						Start: first.Start,
						End:  first.End,
					},
					Buttons: []*client.Button{
						{
							Label: "Kalender anzeigen",
							Url:   "https://nextcloud.sinclear.de/index.php/apps/calendar/p/ffe5tLHMEjmcPLxD",
						},
					},
				})
				if err != nil {
					panic(err)
				}

			} else {																			// ONE OR MORE EVENTS, FIRST ENTRY IS RIGHT NOW, NO SECOND EVENT EXISTS

				err = client.SetActivity(client.Activity{
					Details:    "Aktuell: " + first.Summary + firstDescription,
					State:      "Danach: – ",
					Timestamps: &client.Timestamps{
						Start: first.Start,
						End:  first.End,
					},
					Buttons: []*client.Button{
						{
							Label: "Kalender anzeigen",
							Url:   "https://nextcloud.sinclear.de/index.php/apps/calendar/p/ffe5tLHMEjmcPLxD",
						},
					},
				})
				if err != nil {
					panic(err)
				}

			}

		} else {																				// ONE OR MORE EVENTS, FIRST ENTRY IS NOT RIGHT NOW

			err = client.SetActivity(client.Activity{
				Details:    "Aktuell: – ",
				State:      "Danach: " + first.Start.Format("15:04") + " – " + first.Summary + firstDescription,
				Buttons: []*client.Button{
					{
						Label: "Kalender anzeigen",
						Url:   "https://nextcloud.sinclear.de/index.php/apps/calendar/p/ffe5tLHMEjmcPLxD",
					},
				},
			})
			if err != nil {
				panic(err)
			}

		}
		
	} else {																					// NO EVENTS

		err = client.SetActivity(client.Activity{
			Details:    "Keine Einträge im Kalender",
			Buttons: []*client.Button{
				{
					Label: "Kalender anzeigen",
					Url:   "https://nextcloud.sinclear.de/index.php/apps/calendar/p/ffe5tLHMEjmcPLxD",
				},
			},
		})
		if err != nil {
			panic(err)
		}

	}

	time.Sleep(time.Second * 60)
}
