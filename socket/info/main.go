package info

type info struct {
	List []int
}

// Infos public
var Infos = info{}

func (i *info) Add(id int) {
	i.List = append(i.List, id)
}

func (i *info) Del(id int) {
	for index, idAccount := range i.List {
		if idAccount == id {
			i.List = append(i.List[:index], i.List[index+1:]...)
			break
		}
	}
}

func (i *info) Alive(id int) bool {
	for _, idAccount := range i.List {
		if idAccount == id {
			return true
		}
	}
	return false
}

func (i *info) NbAlive() int {
	return len(i.List)
}
