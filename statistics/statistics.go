package statistics

import (
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
	// TODO build a new Statistics instance

	// TODO start the "run" loop in a separate go routine

	// TODO return the started instance
	return nil
}

// PlusOne is used to add one to the counter
func (sw *Statistics) PlusOne() {
	sw.statistics <- uint8(1)
}

func (sw *Statistics) run() {
	// TODO build a new time.NewTicker instance of sw.loggingPeriod duration

	// infinite loop select for two channel
	for {
		select {

		// process the incoming statistics coming from the channel
		case stat := <-sw.statistics:
			logger.WithField("stat", stat).Debug("new count received")
			sw.counter += uint32(stat)

			// TODO add a case to listen to the ticker channel

			// TODO compute the time elapsed since the start date
			// TODO log the elapsed time, the hit count since last reset
			// TODO bonus compute the request/sec rate

			// TODO reset the counter and the start date
		}
	}
}
