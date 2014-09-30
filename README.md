# go-template\_metrics

Instrument template rendering

# Usage

```go
import (
  "html/template" // or "text/template"
  "github.com/sonots/go-template_metrics"
)

func main() {
  // formTmpl := template.Must(template.ParseFiles("views/base.html", "views/form.html"))
  formTmpl := template_metrics.WrapTemplate("form", template.Must(template.ParseFiles("views/base.html", "views/form.html")))

  // Use as usual
  formTmpl.Execute(w, data)
  // or
  formTmpl.ExecuteTemplate(w, "base", data)

  template_metrics.Verbose = true // print metrics on each rendering
  template_metrics.Print(1) // print metrics on each 1 second
  // template_metrics.Enable = false // turn off the instrumentation
}
```

Output Example (LTSV format):

```
time:2014-09-08 05:06:57.99249161 +0900 JST     template:form   base:base    count:1 max:0.000301    mean:0.000301   min:0.000301    percentile95:0.000301   duration:1
```

Verbose Output Example (LTSV format):

```
time:2014-09-08 05:06:57.22659252 +0900 JST     template:form   base:base    elapsed:0.000301
```

# API

## Print

Print summarized metrics on each specified second:

```go
template_metrics.Print(60)
```

## Verbose

Print metrics on each rendering:

```go
template_metrics.Verbose = true
```

## Enable

Diable instrumentation as:

```
template_metrics.Enable = false
```

## Flush()

Flush metrics on arbitrary timing by calling `Flush()` as:

```
template_metrics.Flush()
```

# ToDo

* Write tests

# Contribution

* Fork (https://github.com/sonots/go-template_metrics/fork)
* Create a feature branch
* Commit your changes
* Rebase your local changes against the master branch
* Create new Pull Request

# Copyright

* See [LICENSE](./LICENSE)
