package keysandvalues

import "golang.org/x/exp/maps"

func Get[M ~map[K]V, K comparable, V any](m M) ([]K, []V) {
	return maps.Keys(m), maps.Values(m)
}
