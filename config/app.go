package config

type AppConfig struct {
	AppName  string
	Version  string
	Timezone int
	DBPath   string
}

func Load() AppConfig {
	return AppConfig{
		AppName:  "PrayerTimeTUI",
		Version:  "0.1.0",
		Timezone: 8, // GMT+8
		DBPath:   "zones.db",
	}
}
