package main

import (
	"html/template"
	"os"
)

type Father struct {
	Fname string
}

type Lang struct {
	LangName string
	Version  []string
	Fathers  []*Father
}

func main() {
	// New, Parse, Execcute
	t := template.New("langname example")
	t, _ = t.Parse("hello {{.LangName}}!\n")
	p := Lang{LangName: "golang"}
	t.Execute(os.Stdout, p)

	// range, with
	f1 := Father{Fname: "google"}
	f2 := Father{Fname: "linus"}
	t1 := template.New("langname example")
	t1, _ = t1.Parse(`hello {{.LangName}}!
            {{range .Version}}
                an version {{.}}
            {{end}}
            {{with .Fathers}}
            {{range .}}
                my father name is {{.Fname}}
            {{end}}
            {{end}}
            `)

	l := Lang{LangName: "golang",
		Version: []string{"1.0", "5.0", "6.0"},
		Fathers: []*Father{&f1, &f2}}

	t1.Execute(os.Stdout, l)

	// if-else
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipline if demo: {{if ``}} 不会输出。 {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tValue := template.New("template test")
	tValue = template.Must(tValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容 {{end}}\n"))
	tValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo:{{if ``}} if部分 {{else}} else部分. {{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

	// inner var
	tVar := template.New("template test")
	tVar = template.Must(tVar.Parse("with-var {{with $x := `output`}}{{$x | printf `%q`}}{{end}}\n"))
	tVar.Execute(os.Stdout, nil)

}
