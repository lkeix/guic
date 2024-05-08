package guic

import (
	"net"
)

type quic struct {
}

func (q *quic) apply(conn *net.UDPConn) (r *Request, w ResponseWriter) {

	return nil, nil
}
