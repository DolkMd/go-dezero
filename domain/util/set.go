package util

type set struct {
	key    []interface{}
	values map[interface{}]struct{}
}

func Set(vs ...interface{}) *set {
	values := map[interface{}]struct{}{}
	key := []interface{}{}
	for _, v := range vs {
		if _, ok := values[v]; !ok {
			key = append(key, v)
			values[v] = struct{}{}
		}
	}
	return &set{
		values: values,
		key:    key,
	}
}

func (s *set) Add(values ...interface{}) {
	for _, v := range values {
		if _, ok := s.values[v]; !ok {
			s.key = append(s.key, v)
			s.values[v] = struct{}{}
		}
	}
}

func (s *set) Get(key interface{}) interface{} {
	return s.values[key]
}

func (s *set) Find(key interface{}) bool {
	_, ok := s.values[key]
	return ok
}
