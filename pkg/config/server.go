package config

import "time"

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))

	return err
}


type Server struct {
	ShutdownTimeout Duration `toml:"shutdown_timeout" yaml:"shutdown_timeout"`
	ReadTimeout     Duration `toml:"read_timeout" yaml:"read_timeout"`
	WriteTimeout    Duration `toml:"write_timeout" yaml:"write_timeout"`
	IdleTimeout     Duration `toml:"idle_timeout" yaml:"idle_timeout"`
}
