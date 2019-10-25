package socket

type info struct {
	List []int
}

// Infos public
var Infos = info{}

func (i *info) add(id int) {
	i.List = append(i.List, id)
}

func (i *info) del(id int) {
	for index, idAccount := range i.List {
		if idAccount == id {
			i.List = append(i.List[:index], i.List[index+1:]...)
			break
		}
	}
}

func (i *info) alive(id int) bool {
	for _, idAccount := range i.List {
		if idAccount == id {
			return true
		}
	}
	return false
}

func (i *info) nbAlive() int {
	return len(i.List)
}
