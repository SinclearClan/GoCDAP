package discord

import (
	"time"

	"github.com/SinclearClan/GoCDAP/calendar"
	"github.com/SinclearClan/GoCDAP/config"
	"github.com/hugolgst/rich-go/client"
)

func SetActivity(cfg *config.Config) {
	err := client.Login("836637382509068319")
	if err != nil {
		panic(err)
	}

	cal := calendar.Update(cfg)

	if len(cal.Events) > 0 { 																	// ONE OR MORE EVENTS
		first := cal.Events[0]

		if calendar.InTimeSpan(first.Start, first.End) {										// ONE OR MORE EVENTS, FIRST ENTRY IS RIGHT NOW

			if len(cal.Events) > 1 {															// ONE OR MORE EVENTS, FIRST ENTRY IS RIGHT NOW, A SECOND EVENT EXISTS
				second := cal.Events[1]

				err = client.SetActivity(client.Activity{
					Details:    "Aktuell: " + first.Summary,
					State:      "Danach: " + second.Start.Format("15:04") + " – " + second.Summary,
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
					Details:    "Aktuell: " + first.Summary,
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
				State:      "Danach: " + first.Start.Format("15:04") + " – " + first.Summary,
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
