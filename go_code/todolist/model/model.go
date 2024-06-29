package model

type Model struct {
	Id     int
	Name   string
	Status string
}

var task = []*Model{&Model{1, "tom", "ing..."}, &Model{2, "mom ", "ing..."}}

func GetModel() []*Model {
	return task
}
func AddModel(name string) {
	task = append(task, &Model{len(task), name, name})
}
