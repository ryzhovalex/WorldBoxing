package orwynn

import "worldboxing/lib/utils"

type Transport interface {
	GetMaxConnectionSize() int
	GetConnectionSize() int
	GetConnection(id utils.Id) Connection
	Accept() (Connection, *utils.Error)
	Close()
}

type Connection interface {
	Id() utils.Id
	GetTransport() Transport
	Send([]byte) *utils.Error
	Recv() ([]byte, *utils.Error)
	Close()
}
