package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	rawDate := "20251228T152000Z"

	// 2006 = année ; 01 = mois ; 02 = jour ; 15 = heure ; 04 = minute ; 05 = seconde
	d, _ := time.Parse("20060102T150405Z", rawDate)

	if d.Before(time.Now()) {
		fmt.Println("Avant !")
	} else {
		fmt.Println("Après !")
	}

	fmt.Println(d)

	resp, err := http.Get("https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=13295,13345&projectId=3&calType=ical&nbWeeks=4&displayConfigId=128")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	lines := strings.Split(string(body), "\n")

	currentlyParsing := false
	tmpObj := map[string]interface{}{}

	for _, line := range lines {
		if strings.HasPrefix(line, "BEGIN:VEVENT") {
			currentlyParsing = true
		} else {
			if currentlyParsing {
				if strings.HasPrefix(line, "END:VEVENT") {
					// Affichage amélioré
					printEvent(tmpObj)
					tmpObj = map[string]interface{}{}
					currentlyParsing = false
				} else {
					if strings.HasPrefix(line, "DTSTAMP:") {
						tmpObj["stamp"], _ = time.Parse("20060102T150405Z", strings.Replace(strings.Replace(line, "DTSTAMP:", "", 1), "\r", "", 1))
					}

					if strings.HasPrefix(line, "DTSTART:") {
						tmpObj["start"], _ = time.Parse("20060102T150405Z", strings.Replace(strings.Replace(line, "DTSTART:", "", 1), "\r", "", 1))
					}
					if strings.HasPrefix(line, "DTEND:") {
						tmpObj["end"], _ = time.Parse("20060102T150405Z", strings.Replace(strings.Replace(line, "DTEND:", "", 1), "\r", "", 1))
					}
					if strings.HasPrefix(line, "SUMMARY:") {
						tmpObj["summary"] = strings.Replace(strings.Replace(line, "SUMMARY:", "", 1), "\r", "", 1)
					}
					if strings.HasPrefix(line, "LOCATION:") {
						tmpObj["location"] = strings.Replace(strings.Replace(line, "LOCATION:", "", 1), "\r", "", 1)
					}
					if strings.HasPrefix(line, "DESCRIPTION:") {
						tmpObj["description"] = strings.Replace(strings.Replace(line, "DESCRIPTION:", "", 1), "\r", "", 1)
					}
					if strings.HasPrefix(line, "LAST-MODIFIED:") {
						tmpObj["last-modified"], _ = time.Parse("20060102T150405Z", strings.Replace(strings.Replace(line, "LAST-MODIFIED:", "", 1), "\r", "", 1))
					}
				}
			} else {
				continue
			}
		}
	}
}

func printEvent(event map[string]interface{}) {
	fmt.Println("=" + strings.Repeat("=", 80))

	// Titre du cours
	if summary, ok := event["summary"].(string); ok {
		fmt.Printf("📚 COURS: %s\n", summary)
	}

	// Dates et heures
	if start, ok := event["start"].(time.Time); ok {
		if end, ok := event["end"].(time.Time); ok {
			fmt.Printf("📅 DATE: %s\n", start.Format("Monday 02/01/2006"))
			fmt.Printf("🕐 HORAIRE: %s - %s\n", start.Format("15:04"), end.Format("15:04"))

			duration := end.Sub(start)
			fmt.Printf("⏱️  DURÉE: %s\n", duration)
		}
	}

	// Lieu
	if location, ok := event["location"].(string); ok && location != "" {
		fmt.Printf("📍 LIEU: %s\n", location)
	}

	// Description (enseignant, groupe)
	if desc, ok := event["description"].(string); ok && desc != "" {
		// Nettoyer la description
		desc = strings.ReplaceAll(desc, "\\n", " ")
		desc = strings.TrimSpace(desc)
		if desc != "" {
			fmt.Printf("👤 INFO: %s\n", desc)
		}
	}

	// Dernière modification
	if lastMod, ok := event["last-modified"].(time.Time); ok {
		fmt.Printf("🔄 Modifié le: %s\n", lastMod.Format("02/01/2006 à 15:04"))
	}

	fmt.Println()
}
