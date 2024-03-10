package transpiler

type Name struct {
	Pos int
	Id string
}

func NewName(pos int) Name {
	return Name{
		Pos: pos,
		Id: Id.Next(),
	}
}

type NamePool struct {
	names []Name
	available []int
}

func NewNamePool() NamePool {
	return NamePool{
		names: make([]Name, 0),
		available: make([]int, 0),
	}
}

func (np *NamePool) Get() *Name {
	if len(np.available) > 0 {
		name := np.available[len(np.available)-1]
		np.available = np.available[0:(len(np.available) - 1)]
		return &np.names[name]
	}
	return np.Extend()
}

func (np *NamePool) Extend() *Name {
	name := NewName(len(np.names))
	np.names = append(np.names, name)
	return &name
}

func (np *NamePool) Drop(name *Name) {
	np.available = append(np.available, name.Pos)
}