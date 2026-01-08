package tui

import (
	"PrayerTime/prayer"
	"fmt"
	"strings"
)

func (m Model) View() string {
	if m.Err != nil {
		return fmt.Sprintf("Error: %v\n\nPress q to quit", m.Err)
	}

	switch m.Step {
	case stepState:
		return renderList("Select State", m.States, m.Cursor)

	case stepCity:
		return renderList(
			fmt.Sprintf("State: %s\nSelect City", m.SelectedState),
			m.Cities,
			m.Cursor,
		)

	case stepResult:
		ws := prayer.NewPrayerTime(
			m.Zone.Lat,
			m.Zone.Lon,
			m.Now.Year(),
			int(m.Now.Month()),
			m.Now.Day(),
			m.App.Timezone,
		)

		return fmt.Sprintf(
			`Coordinate: %f %f
Location: %s, %s


Imsak   : %s
Fajr    : %s
Syuruk  : %s
Zuhr    : %s
Asar    : %s
Maghrib : %s
Isyak   : %s
			`,
			m.Zone.Lat,
			m.Zone.Lon,
			m.Zone.City,
			m.Zone.State,
			ws.Imsak(),
			ws.Fajr(),
			ws.Shuruq(),
			ws.Dhuhr(),
			ws.Asr(),
			ws.Maghrib(),
			ws.Isya(),
		)
	}
	return ""
}

func renderList(title string, items []string, cursor int) string {
	var b strings.Builder

	b.WriteString(title + "\n\n")

	for i, item := range items {
		prefix := "  "
		if i == cursor {
			prefix = "> "
		}
		b.WriteString(prefix + item + "\n")
	}

	b.WriteString("\n↑↓ navigate • enter select • q quit")
	return b.String()
}
