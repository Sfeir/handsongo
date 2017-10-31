package statistics

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"time"
)

const (
	statisticsChannelSize = 1000
)

// Statistics is the worker to persist the request statistics
type Statistics struct {
	statistics    chan uint8
	counter       uint32
	start         time.Time
	loggingPeriod time.Duration
}

// NewStatistics creates a new statistics structure and launches its worker routine
func NewStatistics(loggingPeriod time.Duration) *Statistics {
	sw := Statistics{
		statistics:    make(chan uint8, statisticsChannelSize),
		counter:       0,
		start:         time.Now(),
		loggingPeriod: loggingPeriod,
	}
	go sw.run()
	return &sw
}

// PlusOne is used to add one to the counter
func (sw *Statistics) PlusOne() {
	sw.statistics <- uint8(1)
}

func (sw *Statistics) run() {
	ticker := time.NewTicker(sw.loggingPeriod)
	for {
		select {
		case stat := <-sw.statistics:
			logger.WithField("stat", stat).Debug("new count received")
			sw.counter += uint32(stat)
		case <-ticker.C:
			elapsed := time.Since(sw.start)
			logger.WithField("elapsed time", elapsed).WithField("count", sw.counter).
				WithField("request per second", fmt.Sprintf("%.2f", float64(sw.counter)/elapsed.Seconds())).Warn("request monitoring")
			sw.counter = 0
			sw.start = time.Now()
		}
	}
}
