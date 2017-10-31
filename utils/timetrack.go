package utils

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// TimeTrack is used to log execution times for functions
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.WithField("name", name).WithField("elapsed time", elapsed).Debug("time monitoring")
}
