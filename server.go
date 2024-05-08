package guic

import (
	"net"
)

type Server struct {
	Addr *net.UDPAddr

	Conns   map[string]*net.UDPConn
	quic    *quic
	Handler Handler
}
type Request struct {
}

type ResponseWriter interface {
}

type Handler interface {
	ServeQUIC(r *Request, w ResponseWriter)
}

func NewServer(addr string) (*Server, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		Addr:  udpAddr,
		Conns: make(map[string]*net.UDPConn),
	}, nil
}

func (s *Server) Serve() error {
	conn, err := net.ListenUDP("udp", s.Addr)
	if err != nil {
		return err
	}

	for {
		go func() {
			r, w := s.quic.apply(conn)
			s.Handler.ServeQUIC(r, w)
		}()
	}
}
