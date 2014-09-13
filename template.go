package template_metrics

import (
	"html/template" // ToDo: how can I change this from outside?
	"io"
	"time"
)

type Template struct {
	Original *template.Template
	metrics  *Metrics
}

func newTemplate(template *template.Template, metrics *Metrics) *Template {
	return &Template{
		template,
		metrics,
	}
}

func (proxy *Template) measure(startTime time.Time, query string) {
	proxy.metrics.measure(startTime, query)
}

// instrument template.Execute
func (proxy *Template) Execute(wr io.Writer, data interface{}) error {
	if Enable {
		startTime := time.Now()
		// treat as no base name
		defer proxy.measure(startTime, "")
	}
	return proxy.Original.Execute(wr, data)
}

// instrucment template.ExecuteTemplate
func (proxy *Template) ExecuteTemplate(wr io.Writer, base string, data interface{}) error {
	if Enable {
		startTime := time.Now()
		defer proxy.measure(startTime, base)
	}
	return proxy.Original.ExecuteTemplate(wr, base, data)
}
