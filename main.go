package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	// :8080 => 0.0.0.0:8080
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

var indexTmpl = template.Must(template.New("").Parse(`
<!doctype html>
<title>Calculator</title>
<h1>Calculator</h1>

<form method=POST>
    <input name=a type="number">
    <input name=b type="number">
    <button>+</button>
</form>

{{if .}}
<p>Result: {{.A}} + {{.B}} = {{.Result}} </p>
{{end}}
`))

type indexData struct {
	A      int
	B      int
	Result int
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		a, _ := strconv.Atoi(r.PostFormValue("a"))
		b, _ := strconv.Atoi(r.PostFormValue("b"))

		result := a + b
		w.Header().Set("Content-Type", "text/html")
		indexTmpl.Execute(w, indexData{
			A:      a,
			B:      b,
			Result: result,
		})
		return
	}

	w.Header().Set("Content-Type", "text/html")
	indexTmpl.Execute(w, nil)
}
