package items

import (
	"log"
	"strconv"
	"strings"
)

//Message .
type Message struct {
	Ball     string
	Bingo    string
	Finished string
}

//GetMessageBall .
func (m *Message) GetMessageBall() Ball {
	message := m.Ball
	var ball Ball
	var err error
	ball.Letter = string([]rune(message)[0])
	ball.Number, err = strconv.Atoi(strings.Trim(message, ball.Letter))
	if err != nil {
		log.Fatal(err)
	}
	return ball
}

//GetMessageBingo .
func (m *Message) GetMessageBingo() []string {
	message := m.Bingo
	var winners []string
	winners = strings.Split(message, ",")
	return winners
}

//GetMessageFinished .
func (m *Message) GetMessageFinished() bool {
	if m.Finished == "true" {
		return true
	}
	return false
}

//SaveWinner .
func (m *Message) SaveWinner(winner string) {
	if m.Bingo == "null" {
		m.Bingo = winner
	} else {
		m.Bingo += "," + winner
	}
}

//SaveBall .
func (m *Message) SaveBall(ball Ball) {
	m.Ball = ball.Letter + strconv.Itoa(ball.Number)
}

//SaveMessage .
func (m *Message) SaveMessage(res []string) {
	m.Ball = res[0]
	m.Bingo = res[1]
	m.Finished = res[2]
}
