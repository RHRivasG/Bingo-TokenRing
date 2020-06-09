package ports

import (
	"strings"

	"github.com/tarm/serial"
)

//Listener .
type Listener struct {
	name string
	sl   *serial.Port
}

//NewListener .
func NewListener(name string) (Listener, error) {
	l := &serial.Config{Name: name, Baud: 115200}
	sl, err := serial.OpenPort(l)
	if err != nil {
		return Listener{}, err
	}
	return Listener{name, sl}, nil
}

//Listening .
func (l *Listener) Listening() ([]string, error) {
	buf := make([]byte, 128)
	n, err := (l.sl).Read(buf)
	if err != nil {
		return nil, err
	}
	messages := cutLimiter(separate(string(buf[:n])))
	return messages, nil
}

func cutIndexAndSuffix(t string) string {
	return strings.TrimSuffix(strings.TrimPrefix(t, "B1"), "O75")
}
func cutLimiter(m []string) []string {
	var messages []string
	for _, message := range m {
		messages = append(messages, cutIndexAndSuffix(message))
	}
	return messages
}
func separate(t string) []string {
	return strings.SplitAfter(t, "O75")
}
