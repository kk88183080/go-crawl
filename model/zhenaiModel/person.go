package zhenaiModel

/**
男女对象信息
*/
type Person struct {
	Photo     string
	Name      string
	Id        string
	Sex       string
	City      string
	Age       string
	Schoole   string
	Status    string
	Height    string
	Money     string
	Introduce string
}

func (p *Person) String() string {
	return "Photo:" + p.Photo + ",Name:" + p.Name + ",Id:" + p.Id + ",Introduce:" + p.Introduce
}
