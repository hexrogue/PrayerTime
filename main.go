package main

import (
	"PrayerTime/config"
	"PrayerTime/tui"
	"PrayerTime/zone"
	"database/sql"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	app := config.Load()

	db, err := sql.Open("sqlite3", app.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	zoneRepo := zone.NewRepository(db)

	model := tui.NewModel(zoneRepo, &app)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
