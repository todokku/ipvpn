package vpn

import (
	"net"

	"github.com/xaionaro-go/errors"
)

type udpWriter struct {
	conn *net.UDPConn
	addr *net.UDPAddr
}

func newUDPWriter(conn *net.UDPConn, addr *net.UDPAddr) *udpWriter {
	return &udpWriter{
		conn: conn,
		addr: addr,
	}
}

func (w *udpWriter) Write(b []byte) (size int, err error) {
	size, err = w.conn.WriteToUDP(b, w.addr)
	if err != nil {
		err = errors.Wrap(err)
	}
	return
}
