package multiplayer

import (
	"encoding/json"
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
	Player   *PlayerMatch
	Opponent *PlayerMatch
	Ball     *BallMatch
	UpdatedAt   time.Time
}

func NewMatch() *Match {
	return &Match{
		Player:   &PlayerMatch{},
		Opponent: &PlayerMatch{},
		Ball:     &BallMatch{},
	}
}

func (m *Match) JoinMatch() {

}

func (m *Match) Sync(gameSync func(data string)) {
	matchJson, err := json.Marshal(m)

	if err != nil {
		alog.Error(err.Error())
		return
	}

	sendServer(matchJson, gameSync)
}
