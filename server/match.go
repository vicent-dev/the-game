package main

import "github.com/google/uuid"

type match struct {
	playerId1    string
	playerId2    string
	playerScore1 int
	playerScore2 int
}

func (m *match) joinMatch() string {
	playerId := uuid.New().String()

	if m.playerId1 == "" {
		m.playerId1 = playerId
		return playerId
	}

	m.playerId2 = playerId
	return playerId
}
