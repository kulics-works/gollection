package gollection

// Add subscripts to the incoming iterators.
func Enumerate[T any](it Iterator[T]) Iterator[Pair[int, T]] {
	return &enumerateStream[T]{-1, it}
}

type enumerateStream[T any] struct {
	index    int
	iterator Iterator[T]
}

func (a *enumerateStream[T]) Next() Option[Pair[int, T]] {
	if v, ok := a.iterator.Next().Get(); ok {
		a.index++
		return Some(PairOf(a.index, v))
	}
	return None[Pair[int, T]]()
}

// Use transform to map an iterator to another iterator.
func Map[T any, R any](transform func(T) R, it Iterator[T]) Iterator[R] {
	return &mapStream[T, R]{transform, it}
}

type mapStream[T any, R any] struct {
	transform func(T) R
	iterator  Iterator[T]
}

func (a *mapStream[T, R]) Next() Option[R] {
	if v, ok := a.iterator.Next().Get(); ok {
		return Some(a.transform(v))
	}
	return None[R]()
}

// Use predicate to filter an iterator to another iterator。
func Filter[T any](predicate func(T) bool, it Iterator[T]) Iterator[T] {
	return &filterStream[T]{predicate, it}
}

type filterStream[T any] struct {
	predicate func(T) bool
	iterator  Iterator[T]
}

func (a *filterStream[T]) Next() Option[T] {
	for v, ok := a.iterator.Next().Get(); ok; v, ok = a.iterator.Next().Get() {
		if a.predicate(v) {
			return Some(v)
		}
	}
	return None[T]()
}

// Convert an iterator to another iterator that limits the maximum number of iterations.
func Limit[T any](count int, it Iterator[T]) Iterator[T] {
	return &limitStream[T]{count, 0, it}
}

type limitStream[T any] struct {
	limit    int
	index    int
	iterator Iterator[T]
}

func (a *limitStream[T]) Next() Option[T] {
	if v, ok := a.iterator.Next().Get(); ok && a.index < a.limit {
		a.index++
		return Some(v)
	}
	return None[T]()
}

// Converts an iterator to another iterator that skips a specified number of times.
func Skip[T any](count int, it Iterator[T]) Iterator[T] {
	return &skipStream[T]{count, 0, it}
}

type skipStream[T any] struct {
	skip     int
	index    int
	iterator Iterator[T]
}

func (a *skipStream[T]) Next() Option[T] {
	for v, ok := a.iterator.Next().Get(); ok; v, ok = a.iterator.Next().Get() {
		if a.index < a.skip {
			a.index++
			continue
		}
		return Some(v)
	}
	return None[T]()
}

// Converts an iterator to another iterator that skips a specified number of times each time.
func Step[T any](count int, it Iterator[T]) Iterator[T] {
	return &stepStream[T]{count, count, it}
}

type stepStream[T any] struct {
	step     int
	index    int
	iterator Iterator[T]
}

func (a *stepStream[T]) Next() Option[T] {
	for v, ok := a.iterator.Next().Get(); ok; v, ok = a.iterator.Next().Get() {
		if a.index < a.step {
			a.index++
			continue
		}
		a.index = 1
		return Some(v)
	}
	return None[T]()
}

// By connecting two iterators in series,
// the new iterator will iterate over the first iterator before continuing with the second iterator.
func Concat[T any](left Iterator[T], right Iterator[T]) Iterator[T] {
	return &concatStream[T]{false, left, right}
}

type concatStream[T any] struct {
	firstNotFinished bool
	first            Iterator[T]
	last             Iterator[T]
}

func (a *concatStream[T]) Next() Option[T] {
	if a.firstNotFinished {
		if v, ok := a.first.Next().Get(); ok {
			return Some(v)
		}
		a.firstNotFinished = false
		return a.Next()
	}
	return a.last.Next()
}
