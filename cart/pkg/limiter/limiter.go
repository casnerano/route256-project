package limiter

import (
	"net"
	"sync"
)

type Limiter struct {
	mu   sync.Mutex
	conn net.Conn
}

func New(addr string) (*Limiter, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Limiter{conn: conn}, nil
}

func (l *Limiter) Wait() error {
	_, err := l.conn.Write([]byte{1})
	if err != nil {
		return err
	}

	answer := make([]byte, 1)
	_, err = l.conn.Read(answer)
	if err != nil {
		return err
	}

	return nil
}

func (l *Limiter) Close() error {
	return l.conn.Close()
}
