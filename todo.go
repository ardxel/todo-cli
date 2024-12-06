package todo

type Todo struct {
	Status string
	Id     string
	Text   string
}

func New() *Todo {
	return &Todo{}
}

func List(status string) []*Todo {
	return []*Todo{}
}

func Done(id string) bool {
	return true
}

func Delete(id string) {}
