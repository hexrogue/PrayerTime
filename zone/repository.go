package zone

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetStates() ([]string, error) {
	rows, err := r.db.Query(`
	SELECT DISTINCT state
	FROM zones
	ORDER BY state
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var states []string
	for rows.Next() {
		var state string
		if err := rows.Scan(&state); err != nil {
			return nil, err
		}

		states = append(states, state)
	}
	return states, nil
}

func (r *Repository) GetCitiesByState(state string) ([]string, error) {
	rows, err := r.db.Query(`
	SELECT city
	FROM zones
	WHERE state = ?
	ORDER BY city
	`, state)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []string
	for rows.Next() {
		var city string
		if err := rows.Scan(&city); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}

func (r *Repository) GetZone(city, state string) (*Zone, error) {
	row := r.db.QueryRow(`
	SELECT id, city, state, lat, lon
	FROM zones
	WHERE city = ? AND state = ?
	LIMIT 1
	`, city, state)

	var z Zone
	if err := row.Scan(&z.ID, &z.City, &z.State, &z.Lat, &z.Lon); err != nil {
		return nil, err
	}

	return &z, nil
}
