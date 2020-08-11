package protocol

import (
	port "bingo-tokenring/protocol/ports"
)

//Protocol .
type Protocol struct {
	L port.Listener
	W port.Writer
}

//NewProtocol .
func NewProtocol(listener string, writer string) (Protocol, error) {
	l, err := port.NewListener(listener)
	if err != nil {
		return Protocol{}, err
	}
	w, err := port.NewWriter(writer)
	if err != nil {
		return Protocol{}, err
	}
	return Protocol{l, w}, nil
}

//Listen .
func (p *Protocol) Listen() ([]string, error) {
	return p.L.Listening()
}

//Converse .
func (p *Protocol) Converse(message []string) ([]string, error) {
	p.W.Writing(message)
	return p.L.Listening()
}

//Write .
func (p *Protocol) Write(lastMessage []string) {
	p.W.Writing(lastMessage)
}

//GetWriterName .
func (p *Protocol) GetWriterName() string {
	return p.W.GetName()
}

//Close .
func (p *Protocol) Close() {
	p.L.Close()
	p.W.Close()

}

//Reset .
func (p *Protocol) Reset() {
	p.L.Reset()
	p.W.Reset()
}
