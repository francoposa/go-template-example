package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// define T1
	t1 := `{{define "T1"}}ONE{{end}}`
	// define T2
	t2 := `{{define "T2"}}TWO{{end}}`
	//// define T3, which invokes T1 & T2
	//t3 := `{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}`
	//// define T3, which invokes T1 & T2, then invokes itself
	//t3 := `{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}{{template "T3"}}`
	// define T3 with block, which invokes T1 & T2, then invokes itself in place
	t3 := `{{block "T3" .}}{{template "T1"}} {{template "T2"}}{{end}}`

	// create new template.Template collection
	tmplEx1 := template.New("tmplEx1")

	// build template collection by parsing the templates,
	// using Must to panic on any errors
	tmplEx1 = template.Must(tmplEx1.Parse(t1))
	tmplEx1 = template.Must(tmplEx1.Parse(t2))
	tmplEx1 = template.Must(tmplEx1.Parse(t3))

	// Print out the names of all templates in the collection
	fmt.Print(tmplEx1.DefinedTemplates())

	fmt.Print("\nexecuting T1 Template: ")
	err := tmplEx1.ExecuteTemplate(os.Stdout, "T1", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\nexecuting T2 Template: ")
	err = tmplEx1.ExecuteTemplate(os.Stdout, "T2", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\nexecuting T3 Template: ")
	err = tmplEx1.ExecuteTemplate(os.Stdout, "T3", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\nexecuting full Template collection: ")
	// Since the root template has name "tmplEx1", this is the same as calling
	// tmplEx1.ExecuteTemplate(os.Stdout, "tmplEx1", "")
	err = tmplEx1.Execute(os.Stdout, "")
	if err != nil {
		fmt.Println(err)
	}

}
