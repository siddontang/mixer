package router

type Node struct {
	Name string

	DB       string
	User     string
	Password string

	Master string
	Slaves []string
}
