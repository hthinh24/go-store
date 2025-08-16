package log

// Log holds logging configuration
type Log struct {
	Level      string `mapstructure:"level" env:"LOG_LEVEL"`
	Format     string `mapstructure:"format" env:"LOG_FORMAT"`
	Output     string `mapstructure:"output" env:"LOG_OUTPUT"`
	FilePath   string `mapstructure:"file_path" env:"LOG_FILE_PATH"`
	MaxSize    int    `mapstructure:"max_size" env:"LOG_MAX_SIZE"`
	MaxBackups int    `mapstructure:"max_backups" env:"LOG_MAX_BACKUPS"`
	MaxAge     int    `mapstructure:"max_age" env:"LOG_MAX_AGE"`
	Compress   bool   `mapstructure:"compress" env:"LOG_COMPRESS"`
}

// IsValid checks if required log fields are set
func (l *Log) IsValid() bool {
	return l.Level != ""
}

// SetDefaults sets default values for log configuration
func (l *Log) SetDefaults() {
	if l.Level == "" {
		l.Level = "info"
	}
	if l.Format == "" {
		l.Format = "json"
	}
	if l.Output == "" {
		l.Output = "stdout"
	}
	if l.MaxSize == 0 {
		l.MaxSize = 100 // MB
	}
	if l.MaxBackups == 0 {
		l.MaxBackups = 3
	}
	if l.MaxAge == 0 {
		l.MaxAge = 28 // days
	}
}
