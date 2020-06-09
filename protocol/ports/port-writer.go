package ports

import (
	"github.com/tarm/serial"
)

//Writer .
type Writer struct {
	name string
	sw   *serial.Port
}

//NewWriter .
func NewWriter(name string) (Writer, error) {
	w := &serial.Config{Name: name, Baud: 115200}
	sw, err := serial.OpenPort(w)
	if err != nil {
		return Writer{}, err
	}
	return Writer{name, sw}, nil
}

//GetName .
func (w *Writer) GetName() string {
	return w.name
}

//GetSerialPort .
func (w *Writer) GetSerialPort() *serial.Port {
	return w.sw
}

//Writing .
func (w *Writer) Writing(messages []string) {
	(w.sw).Write([]byte(unite(putLimiter(messages))))
}

func putLimiter(m []string) []string {
	var messages []string
	for _, message := range m {
		messages = append(messages, "B1"+message+"O75")
	}
	return messages
}
func unite(m []string) string {
	var messages string
	for _, message := range m {
		messages += message
	}
	return messages
}
