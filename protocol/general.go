package protocol

import (
	port "bingo-tokenring/protocol/ports"
)

//Protocol .
type Protocol struct {
	l port.Listener
	w port.Writer
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

//Converse .
func (p *Protocol) Converse(message []string) ([]string, error) {
	p.w.Writing(message)
	return p.l.Listening()
}

//EndConversation .
func (p *Protocol) EndConversation(lastMessage []string) {
	p.w.Writing(lastMessage)
}

//GetWriterName .
func (p *Protocol) GetWriterName() string {
	return p.w.GetName()
}
