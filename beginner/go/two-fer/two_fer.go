// Package twofer returns a custom name and string

package twofer

// ShareWith returns a name and string

func ShareWith(name string) string {
	switch {
	case name == "name" || name == "":
		return "One for you, one for me."
	default:
		return "One for " + (name) + ", one for me."

	}
}
