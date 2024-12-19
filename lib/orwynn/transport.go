package orwynn

import "worldboxing/lib/utils"

type Transport interface {
	GetMaxConnectionSize() int
	GetConnectionSize() int
	GetConnection(id utils.Id) *Connection
	Accept() *Connection
}

type Connection interface {
	Id() utils.Id
	GetTransport() *Transport
	Send(any) *utils.Error
	Recv() (any, *utils.Error)
	Close()
}
