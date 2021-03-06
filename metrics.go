package template_metrics

import (
	"fmt"
	metrics "github.com/sonots/go-metrics" // max,mean,min,stddev,percentile
	"sync"
	"time"
)

type Metrics struct {
	name   string
	timers map[string]metrics.Timer
}

func newMetrics(name string) *Metrics {
	return &Metrics{
		name:   name,
		timers: map[string]metrics.Timer{},
	}
}

var mutex = sync.Mutex{}

//print the elapsed time on each request if Verbose flag is true
func (proxy *Metrics) printVerbose(elapsedTime time.Duration, base string) {
	fmt.Printf("time:%v\ttemplate:%s\tbase:%s\telapsed:%f\n",
		time.Now(),
		proxy.name,
		base,
		elapsedTime.Seconds(),
	)
}

func (proxy *Metrics) printMetrics(duration int) {
	for base, timer := range proxy.timers {
		count := timer.Count()
		if count > 0 {
			fmt.Printf(
				"time:%v\ttemplate:%s\tbase:%s\tcount:%d\tmax:%f\tmean:%f\tmin:%f\tpercentile95:%f\tsum:%f\tduration:%d\n",
				time.Now(),
				proxy.name,
				base,
				timer.Count(),
				float64(timer.Max())/float64(time.Second),
				timer.Mean()/float64(time.Second),
				float64(timer.Min())/float64(time.Second),
				timer.Percentile(0.95)/float64(time.Second),
				float64(timer.Sum())/float64(time.Second),
				duration,
			)
			proxy.timers[base] = metrics.NewTimer()
		}
	}
}

//measure the time
func (proxy *Metrics) measure(startTime time.Time, base string) {
	elapsedTime := time.Now().Sub(startTime)
	if Summary {
		if proxy.timers[base] == nil {
			mutex.Lock()
			if proxy.timers[base] == nil {
				proxy.timers[base] = metrics.NewTimer()
			}
			mutex.Unlock()
		}
		proxy.timers[base].Update(elapsedTime)
	}
	if Verbose {
		proxy.printVerbose(elapsedTime, base)
	}
}
