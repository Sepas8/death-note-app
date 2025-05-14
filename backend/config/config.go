package config

import "time"

type Config struct {
	
	
	Address                     string `json:"address"`
	Database                    string `json:"database"`
	DBHost                      string `json:"db_host"`
	DBPort                      string `json:"db_port"`
	DBUser                      string `json:"db_user"`
	DBPassword                  string `json:"db_password"`
	DBName                      string `json:"db_name"`
	KillDuration                int    `json:"kill_duration"`           // 40 segundos
	KillDurationWithDescription int    `json:"kill_duration_with_desc"` // 6 minutos 40 segundos (400 segundos)
	DefaultKillDuration         time.Duration
	ExtendedKillDuration        time.Duration
}
