package pool

import "errors"

var (
	//ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("pool is closed")
)
type Conn interface {

}

// Pool 基本方法
type Pool interface {
	Get() (*Conn, error)

	Put(*Conn) error

	Close(*Conn) error

	Release()

	Len() int
}
