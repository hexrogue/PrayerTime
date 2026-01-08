package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			m.Cursor++

		case "enter":
			return m.handleEnter()
		}
	}

	return m, nil
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	switch m.Step {
	case stepState:
		if len(m.States) == 0 {
			return m, nil
		}
		m.SelectedState = m.States[m.Cursor]
		m.Cursor = 0

		cities, err := m.Repo.GetCitiesByState(m.SelectedState)
		if err != nil {
			m.Err = err
			return m, nil
		}

		m.Cities = cities
		m.Step = stepCity

	case stepCity:
		if len(m.Cities) == 0 {
			return m, nil
		}
		m.SelectedCity = m.Cities[m.Cursor]

		z, err := m.Repo.GetZone(m.SelectedCity, m.SelectedState)
		if err != nil {
			m.Err = err
			return m, nil
		}

		m.Zone = z
		m.Step = stepResult
		m.Cursor = 0
	}

	return m, nil
}
