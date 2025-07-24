package out

type Pageable[T any] interface {
	Content() []*T
	Total() uint64
	Pages() uint32
	Page() uint32
}

type PageRequest interface {
	Page() uint32
	PageSize() uint32
	SortBy() string
	Direction() bool
}

type Repository[Id any, Entity any] interface {
	Save(*Entity) error
	GetById(Id) (*Entity, error)
	GetAll(offset, limit uint32) (Pageable[Entity], error)
}
