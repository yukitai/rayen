package transpiler

type Lifetime struct {
	Id string
	Begin int
	End int
}

type LifetimeSolver struct {
	lifetime map[string]*Lifetime
	generation int
}

func NewLifetimeSolver() LifetimeSolver {
	return LifetimeSolver{
		lifetime: make(map[string]*Lifetime),
		generation: 0,
	}
}

func (s *LifetimeSolver) Next() {
	s.generation += 1
}

func (s *LifetimeSolver) Record(id string) {
	lifetime, err := s.lifetime[id]
	if err {
		lifetime := Lifetime{
			Id: id,
			Begin: s.generation,
			End: s.generation,
		}
		s.lifetime[id] = &lifetime
	}
	lifetime.End = s.generation
}

func (s *LifetimeSolver) Distribute(np *NamePool) map[string]*Name {
	names := make(map[string]*Name)
	for gen := 0; gen < s.generation; gen++ {
		for id, lifetime := range s.lifetime {
			if lifetime.Begin == gen {
				names[id] = np.Get()
			}
			if lifetime.End == gen {
				np.Drop(names[id])
			}
		}
	}
	return names
}