package optional

type Optional[T any] struct {
	value T
	ok    bool
}

func Of[T any](v T) Optional[T] {
	return Optional[T]{value: v, ok: true}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) HasData() bool { 
	return o.ok 
}

func (o Optional[T]) Get() (T, bool) { 
	return o.value, o.ok 
}

// If optional does not has any value it will return the
// `def` as default value.
func (o Optional[T]) OrElse(def T) T {
	if o.ok {
		return o.value
	}
	return def
}

func (o *Optional[T]) Set(v T) { 
	o.value, o.ok = v, true 
}

func (o *Optional[T]) Clear()  { 
	var zero T
	o.value, o.ok = zero, false 
}
