package layout

type LayoutCourse struct {
	Name        string
	Description string

	Levels []Level
}

type Level struct {
	Name        string
	Description string

	Words []string
}
