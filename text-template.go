import "text/template"

func do() []byte{
	type someType struct {
		Count int
	}
	var report someType
	report.Count = 42

	t := template.New("template") //create a new template with some name
	t = t.Delims("[[", "]]")
	t, err := t.Parse(`something [[.Count]] something`)
	if err != nil {
		log.Println(err)
	}
	var text bytes.Buffer

	t.Execute(&text, report)
	return text.Bytes()
}
