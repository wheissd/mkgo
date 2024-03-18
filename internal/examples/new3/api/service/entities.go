package service

type Opt[T any] struct {
	isSet bool
	val   T
}

func (o *Opt[T]) Set(v T) {
	o.val = v
	o.isSet = true
}

func (o *Opt[T]) IsSet() bool {
	return o.isSet
}

func (o *Opt[T]) Get() T {
	return o.val
}
