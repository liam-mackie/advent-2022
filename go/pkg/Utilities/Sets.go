package Utilities

type StringSet map[string]bool

func (s StringSet) Has(v string) bool {
	_, ok := s[v]
	return ok
}

func (s StringSet) Keys() []string {
	keys := make([]string, 0, len(s))

	for k := range s {
		keys = append(keys, k)
	}

	return keys
}
