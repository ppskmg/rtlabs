package mwr

import "github.com/julienschmidt/httprouter"

type Constructor func(httprouter.Handle) httprouter.Handle
type Chain struct {
	constructors []Constructor
}

func Middlewares(constructors ...Constructor) Chain {
	return Chain{append(([]Constructor)(nil), constructors...)}
}
func (c Chain) Then(h httprouter.Handle) httprouter.Handle {
	for i := len(c.constructors) - 1; i >= 0; i-- {
		h = c.constructors[i](h)
	}

	return h
}
func (c Chain) ThenFunc(fn httprouter.Handle) httprouter.Handle {
	if fn == nil {
		return c.Then(nil)
	}
	return c.Then(fn)
}
func (c Chain) Append(constructors ...Constructor) Chain {
	newCons := make([]Constructor, 0, len(c.constructors)+len(constructors))
	newCons = append(newCons, c.constructors...)
	newCons = append(newCons, constructors...)

	return Chain{newCons}
}
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.constructors...)
}
