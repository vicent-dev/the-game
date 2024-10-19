package multiplayer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"time"

	"github.com/en-vee/alog"
)

type PlayerMatch struct {
	Id    string  `json:"id"`
	Score int     `json:"score"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

type BallMatch struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Match struct {
	Player              *PlayerMatch
	Opponent            *PlayerMatch
	Ball                *BallMatch
	UpdatedAt           time.Time
	CurrentUserOpponent bool
}

func NewMatch() *Match {
	return &Match{
		Player:              &PlayerMatch{},
		Opponent:            &PlayerMatch{},
		Ball:                &BallMatch{},
		CurrentUserOpponent: false,
	}
}

func (m *Match) JoinMatch() {
	m.connectMatch()
}

func (m *Match) Ready() bool {
	return m.Player.Id != "" && m.Opponent.Id != ""
}

func (m *Match) Sync(gameSync func(data string)) {
	if m.CurrentUserOpponent {
		player := m.Player
		opponent := m.Opponent

		m.Opponent = player
		m.Player = opponent
	}

	matchJson, err := json.Marshal(m)
	alog.Info("send to sync ", string(matchJson))

	if err != nil {
		alog.Error(err.Error())
		return
	}

	sendServer(matchJson, gameSync)
}

func (m *Match) CopyFromServer(sm *Match) {
	m.Player.Score = sm.Player.Score
	m.Player.X = sm.Player.X
	m.Player.Y = sm.Player.Y

	m.Opponent.Score = sm.Opponent.Score
	m.Opponent.X = sm.Opponent.X
	m.Opponent.Y = sm.Opponent.Y

	m.Ball.X = sm.Ball.X
	m.Ball.Y = sm.Ball.Y
}

func (m *Match) connectMatch() {
	conn := connectUdpServer()

	playerIdIsAlreadyAssigned := m.Player.Id != ""
	matchJson, err := json.Marshal(m)

	if err != nil {
		alog.Error(err.Error())
		return
	}

	_, err = conn.Write([]byte(string(matchJson) + "\n"))

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal([]byte(data), m)

	if playerIdIsAlreadyAssigned && m.Player.Id != "" && m.Opponent.Id != "" {
		m.CurrentUserOpponent = true
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}
