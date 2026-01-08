# PrayerTime TUI

A terminal-based prayer time calculator for Malaysia, built with Go and the Bubble Tea framework. This project demonstrates astronomical coordinate-based calculations and a clean Terminal User Interface (TUI).

---

## Overview

This application allows users to find prayer times for various districts in Malaysia without needing to manually input coordinates. It uses a local SQLite database to store pre-defined coordinates for Malaysian states and cities.

---

## Key Features

- **Offline First**: No internet connection required; calculations are done locally.

- **Interactive TUI**: Navigate through states and cities using your keyboard.

- **Coordinate-Based**: Uses mathematical formulas to derive prayer times based on Latitude and Longitude.

- **Built-in Zone Data**: Includes a `zones.db` containing coordinates for districts across Malaysia.

---

## Calculation Logic

The core astronomical calculations (Solar Declination, Equation of Time, and Julian Days) are based on the methodology described in:

***Source***: [Mohamoud, 2017 - Prayer Time Calculation](https://astronomycenter.net/pdf/mohamoud_2017.pdf)

To ensure reliability during development, the calculation logic was validated against AI models and cross-referenced with peer results.

---

## Disclaimer

**Not for Production Use**. This project was created for **portfolio and educational purposes**. While the mathematical formulas are robust, subtle differences in atmospheric pressure, elevation, and regional criteria might result in slight discrepancies compared to official sources.

For accurate and official prayer times in Malaysia, please refer to:

- **Official Website**: [e-solat.gov.my](https://www.e-solat.gov.my)
- **API**: For production-grade applications, it is highly recommended to use the API provided by JAKIM at the e-solat portal.

---

## Database Setup

This application requires a SQLite database named `zones.db` containing coordinates for Malaysian districts.

1. **Source**: You can obtain or generate the database from this repository: [hexrogue/city2coordinates](https://github.com/hexrogue/city2coordinates).

2. **Setup**: Copy the generated zones.db file and paste it into the root directory of this project.

```bash
# Example
cp path/to/city2coordinates/zones.db ./PrayerTime
```

---

## Tech Stack

- **Language**: Go (Golang)

- **TUI Framework**: [Charmbracelet Bubble Tea](https://github.com/charmbracelet/bubbletea)

- **Database**: SQLite3 (for zone and coordinate storage)

- **Math**: Standard `math` library for trigonometric astronomical functions.

---

## Project Structure

```
.
├── main.go             # Application entry point
├── zones.db            # SQLite database containing Malaysia coordinates
├── config/             # App configuration and constants
├── prayer/             # Core astronomical calculation logic
├── tui/                # UI logic (Model-Update-View)
└── zone/               # Database repository and data models
```

---

## Installation and Usage

1. *Clone the repository*

```bash
git clone https://github.com/yourusername/PrayerTime.git
cd PrayerTime
```

2. *Install dependencies*

```bash
go mod tidy
```

3. *Run the application*

```bash
go run main.go
```

### Navigation

- `↑/↓` or `k/j`: Navigate through the list.

- `Enter`: Select State/City.

- `q` or `Ctrl+C`: Quit the application.
