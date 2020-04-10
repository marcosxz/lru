package lru

type lRUChan chan interface{}

func NewChan(max int) *lRUChan {
	c := make(lRUChan, max)
	return &c
}

func (c *lRUChan) Put(elem interface{}) {
retry:
	select {
	case *c <- elem:
	default:
		_ = <-*c
		goto retry
	}
}

func (c *lRUChan) First() interface{} {
	select {
	case i := <-*c:
		return i
	default:
		return nil
	}
}

func (c *lRUChan) Range(fn func(interface{})) {
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

func (c *lRUChan) Len() int {
	return len(*c)
}
