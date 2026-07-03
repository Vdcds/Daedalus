package catalog

type Category struct {
	ID          string
	Name        string
	Description string

	Packages []Package
}
