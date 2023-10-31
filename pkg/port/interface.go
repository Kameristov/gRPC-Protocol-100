package port

type Connecter interface {
	Conect() error
}

type Disconnecter interface {
	Disconect() error
}

type Communicater interface {
	Communicate(data []byte) ([]byte, error)
}

type Sender interface {
	Send(data []byte) error
}

type Recver interface {
	Recv() ([]byte, error)
}

type Porter interface {
	Connecter
	Disconnecter
	Communicater
	Sender
	Recver
}
