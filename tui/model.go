package tui

import (
	"PrayerTime/config"
	"PrayerTime/zone"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type step int

const (
	stepState step = iota
	stepCity
	stepResult
)

type Model struct {
	App config.AppConfig

	Repo *zone.Repository

	Step step

	States []string
	Cities []string

	SelectedState string
	SelectedCity  string

	Zone *zone.Zone

	Now time.Time

	Cursor int
	Err    error
}

func NewModel(repo *zone.Repository, appConfig *config.AppConfig) Model {
	states, err := repo.GetStates()
	return Model{
		Repo:   repo,
		Step:   stepState,
		States: states,
		Err:    err,
		Now:    time.Now(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
