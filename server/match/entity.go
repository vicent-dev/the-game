package match

import (
	"time"

	"github.com/en-vee/alog"
	"github.com/google/uuid"
)

type playerMatch struct {
	Id    string  `json:"id"`
	Score int     `json:"score"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

type ballMatch struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Match struct {
	Player    *playerMatch
	Opponent  *playerMatch
	Ball      *ballMatch
	UpdatedAt time.Time
}

func NewMatch() *Match {
	return &Match{
		Player:   &playerMatch{},
		Opponent: &playerMatch{},
		Ball:     &ballMatch{},
	}
}

func (m *Match) JoinMatch() string {
	playerId := uuid.New().String()

	if m.Player.Id == "" {
		m.Player.Id = playerId
		alog.Info("Player Id ", playerId)
		return playerId
	}

	m.Opponent.Id = playerId
	alog.Info("Opponent Id ", playerId)
	return playerId
}
