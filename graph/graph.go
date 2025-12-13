package graph

type Graph[T comparable] struct {
	EdgeList map[T][]Edge[T]
}

type Edge[T comparable] struct {
	Next T
	Cost int64
}
