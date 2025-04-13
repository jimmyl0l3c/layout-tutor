package layout

type LayoutCourse struct {
	Name   string
	Detail string

	Levels []Level
}

func (c LayoutCourse) Title() string       { return c.Name }
func (c LayoutCourse) Description() string { return c.Detail }
func (c LayoutCourse) FilterValue() string { return c.Name }

type Level struct {
	Name   string
	Detail string

	Words []string
}

func (c Level) Title() string       { return c.Name }
func (c Level) Description() string { return c.Detail }
func (c Level) FilterValue() string { return c.Name }
