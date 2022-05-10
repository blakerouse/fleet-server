// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package config

import (
	"time"
)

type Limit struct {
	Interval time.Duration `config:"interval"`
	Burst    int           `config:"burst"`
	Max      int64         `config:"max"`
	MaxBody  int64         `config:"max_body_byte_size"`
}

type WebsocketLimit struct {
	Limit `config:",inline"`

	ReadConcurrency int `config:"websocket_read_concurrency"`
	WriteConcurrency int `config:"websocket_write_concurrency"`
}

type ServerLimits struct {
	MaxAgents         int           `config:"max_agents"`
	PolicyThrottle    time.Duration `config:"policy_throttle"`
	MaxHeaderByteSize int           `config:"max_header_byte_size"`
	MaxConnections    int           `config:"max_connections"`

	CheckinLimit  WebsocketLimit `config:"checkin_limit"`
	ArtifactLimit Limit `config:"artifact_limit"`
	EnrollLimit   Limit `config:"enroll_limit"`
	AckLimit      Limit `config:"ack_limit"`
	StatusLimit   Limit `config:"status_limit"`
}

// InitDefaults initializes the defaults for the configuration.
func (c *ServerLimits) InitDefaults() {
	c.LoadLimits(loadLimits(0))
}

func (c *ServerLimits) LoadLimits(limits *envLimits) {
	l := limits.Server

	c.MaxHeaderByteSize = 8192 // 8k
	c.MaxConnections = l.MaxConnections
	c.PolicyThrottle = l.PolicyThrottle

	c.CheckinLimit = WebsocketLimit{
		Limit: Limit{
			Interval: l.CheckinLimit.Interval,
			Burst:    l.CheckinLimit.Burst,
			Max:      l.CheckinLimit.Max,
			MaxBody:  l.CheckinLimit.MaxBody,
		},
		ReadConcurrency: l.CheckinLimit.ReadConcurrency,
		WriteConcurrency: l.CheckinLimit.WriteConcurrency,
	}
	c.ArtifactLimit = Limit{
		Interval: l.ArtifactLimit.Interval,
		Burst:    l.ArtifactLimit.Burst,
		Max:      l.ArtifactLimit.Max,
		MaxBody:  l.ArtifactLimit.MaxBody,
	}
	c.EnrollLimit = Limit{
		Interval: l.EnrollLimit.Interval,
		Burst:    l.EnrollLimit.Burst,
		Max:      l.EnrollLimit.Max,
		MaxBody:  l.EnrollLimit.MaxBody,
	}
	c.AckLimit = Limit{
		Interval: l.AckLimit.Interval,
		Burst:    l.AckLimit.Burst,
		Max:      l.AckLimit.Max,
		MaxBody:  l.AckLimit.MaxBody,
	}
	c.StatusLimit = Limit{
		Interval: l.StatusLimit.Interval,
		Burst:    l.StatusLimit.Burst,
		Max:      l.StatusLimit.Max,
		MaxBody:  l.StatusLimit.MaxBody,
	}
}
