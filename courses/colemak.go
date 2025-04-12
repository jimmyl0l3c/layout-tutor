package courses

import "github.com/jimmyl0l3c/layout-tutor/layout"

var Colemak = layout.LayoutCourse{
	Name: "Colemak",

	Description: "Standard Colemak",

	Levels: []layout.Level{
		{
			Name:        "Sire",
			Description: "Words with characters 's', 'i', 'r', 'e'",

			Words: []string{"sire", "re", "si", "iri", "siri", "sir"},
		},
		{
			Name:        "Home row",
			Description: "Words using characters on home row",

			Words: []string{"sons", "seas", "tree", "stories", "inns"},
		},
	},
}
