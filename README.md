# go-template\_metrics

Instrument template rendering

# Usage

```go
import (
  "html/template"
  template_metrics "github.com/sonots/go-template_metrics"
)

func main() {
  // formTmpl := template.Must(template.ParseFiles("views/form.html"))
  formTmpl := template_metrics.WrapTemplate("form", template.Must(template.ParseFiles("views/form.html")))

  formTmpl.Execute(w, struct)
  formTmpl.ExecuteTemplate(w, "base", struct)
}
```

Output Example (LTSV format):

```
time:2014-09-08 05:06:57.99249161 +0900 JST     template:form   count:1 max:0.000301    mean:0.000301   min:0.000301    percentile95:0.000301
```

Verbose Output Example (LTSV format):

```
time:2014-09-08 05:06:57.22659252 +0900 JST     template:form   elapsed:0.000301
```

# ToDo

* Write tests
* Support text/template

# Contribution

* Fork (https://github.com/sonots/go-http_metrics/fork)
* Create a feature branch
* Commit your changes
* Rebase your local changes against the master branch
* Create new Pull Request

# Copyright

* See [LICENSE](./LICENSE)
