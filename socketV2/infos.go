package socketv2

type info struct {
	list []int
}

// Infos public
var Infos = info{}

func (i *info) add(id int) {
	i.list = append(i.list, id)
}

func (i *info) del(id int) {
	for index, idAccount := range i.list {
		if idAccount == id {
			i.list = append(i.list[:index], i.list[index+1:]...)
		}
	}
}

func (i *info) alive(id int) bool {
	for _, idAccount := range i.list {
		if idAccount == id {
			return true
		}
	}
	return false
}

func (i *info) nbAlive() int {
	return len(i.list)
}
