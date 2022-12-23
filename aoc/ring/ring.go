package ring

type Ring[T comparable] struct {
	next, prev *Ring[T]
	Value      T
}

func (r *Ring[T]) init() *Ring[T] {
	r.next = r
	r.prev = r
	return r
}

func (r *Ring[T]) Find(e T) *Ring[T] {
	for i := 0; i < r.Len(); i++ {
		if r.Value == e {
			return r
		}
		r = r.Next()
	}

	return nil
}

func (r *Ring[T]) Next() *Ring[T] {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring[T]) Prev() *Ring[T] {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

func (r *Ring[T]) Move(n int) *Ring[T] {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

func (r *Ring[T]) Shift(n int) *Ring[T] {
	if r.next == nil {
		return r.init()
	}
	if n == 0 {
		return r
	}
	e := &Ring[T]{
		Value: r.Value,
	}
	r.prev.next, r.next.prev = r.next, r.prev
	switch {
	case n < 0:
		for ; n <= 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	e.next, e.prev = r.next, r
	r.next, r.next.prev = e, e
	return r
}

func New[T comparable](n int) *Ring[T] {
	if n <= 0 {
		return nil
	}
	r := new(Ring[T])
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring[T]{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

func NewFromSlice[T comparable](in []T) *Ring[T] {
	if len(in) == 0 {
		return nil
	}
	r := new(Ring[T])
	r.Value = in[0]
	p := r
	for i := 1; i < len(in); i++ {
		p.next = &Ring[T]{prev: p, Value: in[i]}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

func (r *Ring[T]) Link(s *Ring[T]) *Ring[T] {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		// Note: Cannot use multiple assignment because
		// evaluation order of LHS is not specified.
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

func (r *Ring[T]) Unlink(n int) *Ring[T] {
	if n <= 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

func (r *Ring[T]) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

func (r *Ring[T]) Do(f func(T)) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			f(p.Value)
		}
	}
}
