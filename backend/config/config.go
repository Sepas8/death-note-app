package config

type Config struct {
	Address                     string `json:"address"`
	Database                    string `json:"database"`
	KillDuration                int    `json:"kill_duration"`
	KillDurationWithDescription int    `json:"kill_duration_with_desc"`
}
