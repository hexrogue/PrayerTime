package prayer

import (
	"math"
	"time"
)

type PrayerTime struct {
	Latitude  float64
	Longitude float64

	Year  int
	Month int
	Day   int
	Zone  int

	declination float64
	istiwa      float64
}

func NewPrayerTime(lat, lon float64, year, month, day, zone int) *PrayerTime {
	ws := &PrayerTime{
		Latitude:  lat,
		Longitude: lon,

		Year:  year,
		Month: month,
		Day:   day,
		Zone:  zone,
	}
	ws.calculate()
	return ws
}

/*
calculate computes shared astronomical values:
- solar declination
- solar noon (istiwa)
These values are reused by all prayer time calculations.
*/
func (w *PrayerTime) calculate() {
	n := daySinceJan1(w.Year, w.Month, w.Day)
	t := 2 * math.Pi * float64(n-1) / 365

	w.declination =
		0.006918 -
			0.399912*math.Cos(t) +
			0.070257*math.Sin(t) -
			0.006758*math.Cos(2*t) +
			0.000907*math.Sin(2*t) -
			0.002696*math.Cos(3*t) +
			0.00148*math.Sin(3*t)

	eot := equationOfTime(w.Year, w.Month, w.Day, w.Zone)
	// solar noon adjusted for longitude Malaysia (120E reference)
	w.istiwa = (12 + eot/60) + ((120 - w.Longitude) / 15)
}

/*
returns the number of days since January 1st
Used to determine Earth's orbital position.
*/
func daySinceJan1(year, month, day int) int {
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return int(now.Sub(start).Hours() / 24)
}

/*
equationOfTime calculates the difference between
apparent solar time and mean solar time.
*/
func equationOfTime(year, month, day, zone int) float64 {
	jd := julianDay(year, month, day, zone)

	// longitude of the Sun
	L := math.Mod(280.46+0.9856474*jd, 360)
	// anomaly of the Sun (degrees)
	g := math.Mod(357.528+0.9856003*jd, 360)

	// convert anomaly to radians for trigonometric functions
	gRad := deg2rad(g)
	// true ecliptic longitude of the Sun
	lambda := L + 1.915*math.Sin(gRad) + 0.020*math.Sin(2*gRad)

	// obliquity of the ecliptic (Earth axial tilt)
	eps := 23.439 - 0.0000004*jd
	// sun right ascension from ecliptic coordinates
	alpha := rad2deg(math.Atan2(
		math.Cos(deg2rad(eps))*math.Sin(deg2rad(lambda)),
		math.Cos(deg2rad(lambda)),
	))

	// equation of Time in angular form
	E := L - alpha
	if E > 50 {
		E -= 360 // normalize angle to avoid discontinuity
	}
	return -E * 4 // convert degrees to minutes of time
}

/*
converts a date into a Julian Day number
used as the astronomical epoch.
*/
func julianDay(year, month, day, zone int) float64 {
	if month <= 2 {
		year--
		month += 12
	}

	// A represents the century part of the year.
	// Example: 2026 -> A = 20
	A := year / 100

	// B is a correction factor for the Gregorian calendar.
	// It adjusts for leap year rules introduced in Gregorian reform.
	B := 2 - A + A/4

	// 365.25 is used to approximate leap years.
	// Casting to int truncates the decimal part.
	C := int(365.25 * float64(year))

	// D calculates total days from months elapsed.
	// 30.6001 is a magic number used in astronomical formulas
	// to map months to days accurately.
	D := int(30.6001 * float64(month+1))

	// Final Julian Day Number relative to a reference epoch.
	// 730550.5 shifts the epoch to a modern reference date.
	// zone is the timezone offset in hours.
	return float64(B+C+D+day) - 730550.5 - float64(zone)/24
}

// Degree to Radius
func deg2rad(d float64) float64 {
	return d * math.Pi / 180
}

// Radius to Degree
func rad2deg(r float64) float64 {
	return r * 180 / math.Pi
}

/*
calculates the time difference from solar noon
based on the sun's altitude angle.
*/
func hourAngle(w *PrayerTime, angle float64) float64 {
	a := deg2rad(angle)
	lat := deg2rad(w.Latitude)

	return rad2deg(math.Acos(
		(math.Sin(a)-math.Sin(w.declination)*math.Sin(lat))/
			(math.Cos(w.declination)*math.Cos(lat)),
	)) / 15 // hour angle converted from degrees to hours
}

func formatTime(decimal float64) string {
	h := int(decimal)
	m := int((decimal - float64(h)) * 60)
	s := int((((decimal - float64(h)) * 60) - float64(m)) * 60)

	return time.Date(0, 1, 1, h, m, s, 0, time.UTC).Format("15:04:05")
}

func (w *PrayerTime) Imsak() string {
	return formatTime(w.istiwa - hourAngle(w, -18) - 10.0/60)
}

// Subuh
func (w *PrayerTime) Fajr() string {
	return formatTime(w.istiwa - hourAngle(w, -18))
}

func (w *PrayerTime) Shuruq() string {
	return formatTime(w.istiwa - hourAngle(w, -1))
}

func (w *PrayerTime) Dhuhr() string {
	return formatTime(w.istiwa + 2.0/60)
}

func (w *PrayerTime) Asr() string {
	// Asr uses shadow length ratio of 1
	r := math.Atan(1 / (1 + math.Tan(math.Abs(deg2rad(w.Latitude)-w.declination))))
	return formatTime(w.istiwa + hourAngle(w, rad2deg(r)))
}

func (w *PrayerTime) Maghrib() string {
	return formatTime(w.istiwa + hourAngle(w, -1))
}

func (w *PrayerTime) Isya() string {
	return formatTime(w.istiwa + hourAngle(w, -18))
}
