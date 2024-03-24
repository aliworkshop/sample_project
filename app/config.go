package app

import "time"

const ServiceName = "sample_project"

type config struct {
	Http struct {
		Address                   string
		GracefullyShutdownTimeout time.Duration
	}
}

func (c *config) Initialize() {
	if c.Http.GracefullyShutdownTimeout == 0 {
		c.Http.GracefullyShutdownTimeout = time.Second * 10
	}
}
