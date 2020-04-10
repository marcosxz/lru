package lru

type Chan chan interface{}

func NewChan(max int) *Chan {
	c := make(Chan, max)
	return &c
}

func (c *Chan) Put(elem interface{}) {
retry:
	select {
	case *c <- elem:
	default:
		_ = <-*c
		goto retry
	}
}

func (c *Chan) Get() interface{} {
	select {
	case i := <-*c:
		return i
	default:
		return nil
	}
}

func (c *Chan) Range(fn func(interface{})) {
	for {
		select {
		case elem, ok := <-*c:
			if !ok {
				return
			}
			fn(elem)
		default:
			return
		}
	}
}

func (c *Chan) Len() int {
	return len(*c)
}
