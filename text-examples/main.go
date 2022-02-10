package main

import (
	"fmt"
	"os"
	"text/template"
)

//func main() {
//	tmplTextNoInvoke := `
//{{define "T1"}}ONE{{end}}
//{{define "T2"}}TWO{{end}}
//{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
//`
//
//	tmpl0 := template.New("tmpl1")
//	tmpl0, err := tmpl0.Parse(tmplTextNoInvoke)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("executing full Template")
//	_ = tmpl0.Execute(os.Stdout, "")
//
//	fmt.Println("executing T1 Template")
//	_ = tmpl0.ExecuteTemplate(os.Stdout, "T1", "")
//	fmt.Println()
//
//	fmt.Println("executing T2 Template")
//	_ = tmpl0.ExecuteTemplate(os.Stdout, "T2", "")
//	fmt.Println()
//
//	fmt.Println("executing T3 Template")
//	_ = tmpl0.ExecuteTemplate(os.Stdout, "T3", "")
//	fmt.Println()
//
//	tmplTextInvoke1 := ` endln
//{{define "T1"}}ONE{{end}} endln
//{{define "T2"}}TWO{{end}} endln
//{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}} endln
//{{template "T1"}}
//`
//
//	tmpl1 := template.New("tmpl1")
//	tmpl1, err = tmpl1.Parse(tmplTextInvoke1)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("executing full Template")
//	_ = tmpl1.Execute(os.Stdout, "")
//
//	fmt.Println("executing T1 Template")
//	_ = tmpl1.ExecuteTemplate(os.Stdout, "T1", "")
//	fmt.Println()
//
//	fmt.Println("executing T2 Template")
//	_ = tmpl1.ExecuteTemplate(os.Stdout, "T2", "")
//	fmt.Println()
//
//	fmt.Println("executing T3 Template")
//	_ = tmpl1.ExecuteTemplate(os.Stdout, "T3", "")
//	fmt.Println()
//
//	tmplTextInvoke3 := ` endln
//{{define "T1"}}ONE{{end}} endln
//{{define "T2"}}TWO{{end}} endln
//{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}} endln
//{{template "T3"}}
//`
//
//	tmpl3 := template.New("tmpl3")
//	tmpl3, err = tmpl1.Parse(tmplTextInvoke3)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("executing full Template")
//	_ = tmpl3.Execute(os.Stdout, "")
//
//	fmt.Println("executing T1 Template")
//	_ = tmpl3.ExecuteTemplate(os.Stdout, "T1", "")
//	fmt.Println()
//
//	fmt.Println("executing T2 Template")
//	_ = tmpl3.ExecuteTemplate(os.Stdout, "T2", "")
//	fmt.Println()
//
//	fmt.Println("executing T3 Template")
//	_ = tmpl3.ExecuteTemplate(os.Stdout, "T3", "")
//	fmt.Println()
//}

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
