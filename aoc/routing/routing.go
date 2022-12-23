package routing

type Graph[T comparable] interface {
	Neighbours(p T) []T
}
