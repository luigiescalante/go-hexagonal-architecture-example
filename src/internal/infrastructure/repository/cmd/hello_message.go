package cmd

type repository struct {
	listName []string
}

func NewHelloMessageRepo() *repository {
	return &repository{}
}

func (r *repository) Save(name string) error {
	r.listName = append(r.listName, name)
	return nil
}

func (r *repository) GetList() []string {
	return r.listName
}
