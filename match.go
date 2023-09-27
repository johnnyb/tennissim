package main

import "strings"

type Match struct {
	Player1 *PlayerStats
	Player2 *PlayerStats
	Sets    []*Set
}

func CreateMatch(p1, p2 *PlayerStats) *Match {
	m := &Match{
		Player1: p1,
		Player2: p2,
		Sets:    []*Set{},
	}
	return m
}

func (m *Match) Play() {
	currentPlayer := PLAYER_1
	for !m.HasWinner() {
		set := CreateSet(6)
		currentPlayer = set.Play(m.Player1, m.Player2, currentPlayer)
		m.Sets = append(m.Sets, set)
	}
}

func (m *Match) ScoreString() string {
	scores := []string{}
	for _, s := range m.Sets {
		scores = append(scores, s.ScoreString())
	}
	return strings.Join(scores, "; ")
}

func (m *Match) HasWinner() bool {
	return m.Winner() != NONE
}

func (m *Match) Winner() Player {
	p1 := 0
	p2 := 0
	for _, set := range m.Sets {
		winner := set.Winner()
		switch winner {
		case PLAYER_1:
			p1++
		case PLAYER_2:
			p2++
		default:
			return NONE
		}
	}
	if p1 >= 2 {
		return PLAYER_1
	}
	if p2 >= 2 {
		return PLAYER_2
	}

	return NONE
}
