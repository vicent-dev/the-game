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
	currentUserOpponent bool
}

func NewMatch() *Match {
	return &Match{
		Player:              &PlayerMatch{},
		Opponent:            &PlayerMatch{},
		Ball:                &BallMatch{},
		currentUserOpponent: false,
	}
}

func (m *Match) JoinMatch() {
	alog.Info("connecting to match")
	m.connectMatch()
}

func (m *Match) Ready() bool {
	return m.Player.Id != "" && m.Opponent.Id != ""
}

func (m *Match) Sync(gameSync func(data string)) {
	// if users is currently playing as opponent we swap entities
	if m.currentUserOpponent {
		player := m.Player
		opponent := m.Opponent

		m.Opponent = player
		m.Player = opponent
	}

	matchJson, err := json.Marshal(m)

	if err != nil {
		alog.Error(err.Error())
		return
	}

	sendServer(matchJson, gameSync)
}

func (match *Match) connectMatch() {
	conn := connectUdpServer()
	matchJson, err := json.Marshal(match)

	playerIdIsAlreadyAssigned := match.Player.Id != ""

	if err != nil {
		alog.Error(err.Error())
		return
	}

	_, err = conn.Write([]byte(string(matchJson)))

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal([]byte(data), match)
	alog.Debug(match.Player.Id, match.Opponent.Id)

	if playerIdIsAlreadyAssigned && match.Player.Id != "" && match.Opponent.Id != "" {
		match.currentUserOpponent = true
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}
