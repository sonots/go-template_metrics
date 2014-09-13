package template_metrics

import (
	"html/template" // ToDo: how can I change this from outside?
	"time"
)

// print infomation on each request
var Verbose = false

// Set Enable = false if you want to turn off the instrumentation
var Enable = true

// a set of proxies
var proxyRegistry = make(map[string](*Template))

//Wrap  instrument template
func WrapTemplate(name string, template *template.Template) *Template {
	proxy := newTemplate(name, template)
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
			for _, proxy := range proxyRegistry {
				proxy.printMetrics(duration)
			}
			elapsedTime := time.Now().Sub(startTime)
			time.Sleep(timeDuration*time.Second - elapsedTime)
		}
	}()
}
