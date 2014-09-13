package template_metrics

import (
	"fmt"
	metrics "github.com/yvasiyarov/go-metrics" // max,mean,min,stddev,percentile
	"html/template"                            // ToDo: how can I change this from outside?
	"io"
	"time"
)

// print infomation on each request
var Verbose = false

// Set Enable = false if you want to turn off the instrumentation
var Enable = true

// a set of proxies
var proxyRegistry = make(map[string](*proxyTemplate))

type proxyTemplate struct {
	name     string
	Original *template.Template
	timers   map[string]metrics.Timer
}

func newProxyTemplate(name string, template *template.Template) *proxyTemplate {
	return &proxyTemplate{
		name:     name,
		Original: template,
		timers:   map[string]metrics.Timer{},
	}
}

//print the elapsed time on each request if Verbose flag is true
func (proxy *proxyTemplate) printVerbose(elapsedTime time.Duration, base string) {
	fmt.Printf("time:%v\ttemplate:%s\tbase:%s\telapsed:%f\n",
		time.Now(),
		proxy.name,
		base,
		elapsedTime.Seconds(),
	)
}

//measure the time
func (proxy *proxyTemplate) measure(startTime time.Time, base string) {
	elapsedTime := time.Now().Sub(startTime)
	if proxy.timers[base] == nil {
		proxy.timers[base] = metrics.NewTimer()
	}
	proxy.timers[base].Update(elapsedTime)
	if Enable && Verbose {
		proxy.printVerbose(elapsedTime, base)
	}
}

////// instrument functions

// instrument template.Execute
func (proxy *proxyTemplate) Execute(wr io.Writer, data interface{}) error {
	var startTime time.Time
	if Enable {
		startTime = time.Now()
	}
	error := proxy.Original.Execute(wr, data)
	if Enable {
		// treat as no base name
		defer proxy.measure(startTime, "")
	}
	return error
}

// instrucment template.ExecuteTemplate
func (proxy *proxyTemplate) ExecuteTemplate(wr io.Writer, base string, data interface{}) error {
	var startTime time.Time
	if Enable {
		startTime = time.Now()
	}
	error := proxy.Original.ExecuteTemplate(wr, base, data)
	if Enable {
		defer proxy.measure(startTime, base)
	}
	return error
}

///// package functions

//Wrap  instrument template
func WrapTemplate(name string, template *template.Template) *proxyTemplate {
	proxy := newProxyTemplate(name, template)
	proxyRegistry[name] = proxy
	return proxy
}

//Print  print the metrics in each second
func Print(duration int) {
	timeDuration := time.Duration(duration)
	go func() {
		time.Sleep(timeDuration * time.Second)
		for {
			startTime := time.Now()
			for name, proxy := range proxyRegistry {
				for base, timer := range proxy.timers {
					count := timer.Count()
					if count > 0 {
						fmt.Printf(
							"time:%v\ttemplate:%s\tbase:%s\tcount:%d\tmax:%f\tmean:%f\tmin:%f\tpercentile95:%f\tduration:%d\n",
							time.Now(),
							name,
							base,
							timer.Count(),
							float64(timer.Max())/float64(time.Second),
							timer.Mean()/float64(time.Second),
							float64(timer.Min())/float64(time.Second),
							timer.Percentile(0.95)/float64(time.Second),
							duration,
						)
						proxy.timers[base] = metrics.NewTimer()
					}
				}
			}
			elapsedTime := time.Now().Sub(startTime)
			time.Sleep(timeDuration*time.Second - elapsedTime)
		}
	}()
}
