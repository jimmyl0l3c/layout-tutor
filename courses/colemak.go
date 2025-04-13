package courses

import "github.com/jimmyl0l3c/layout-tutor/layout"

var Colemak = layout.LayoutCourse{
	Name: "Colemak",

	Detail: "Standard Colemak",

	Levels: []layout.Level{
		{
			Name:   "Sire",
			Detail: "Words with characters 's', 'i', 'r', 'e'",

			Words: []string{"sire", "re", "si", "iri", "siri", "sir"},
		},
		{
			Name:   "Home row",
			Detail: "Words using characters on home row",

			Words: []string{"sons", "seas", "tree", "stories", "inns"},
		},
	},
}
