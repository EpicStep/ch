package chpool

import (
	puddle "github.com/jackc/puddle/puddleg"

	"github.com/go-faster/ch"
)

type connResource struct {
	client  *ch.Client
	clients []Client
}

func (cr *connResource) getConn(p *Pool, res *puddle.Resource[*connResource]) *Client {
	if len(cr.clients) == 0 {
		cr.clients = make([]Client, 128)
	}

	c := &cr.clients[len(cr.clients)-1]
	cr.clients = cr.clients[0 : len(cr.clients)-1]

	c.res = res
	c.p = p

	return c
}
