package main

import "fmt"

type Set struct {
	GamesRequired int
	Games         []*Game
}

func CreateSet(games int) *Set {
	return &Set{
		GamesRequired: games,
		Games:         []*Game{},
	}
}

func (s *Set) GamesEach() (int, int) {
	p1 := 0
	p2 := 0

	for _, g := range s.Games {
		switch g.Winner() {
		case PLAYER_1:
			p1++
		case PLAYER_2:
			p2++
		default:
			// nothing
		}
	}
	return p1, p2
}

func (s *Set) HasWinner() bool {
	return s.Winner() != NONE
}

func (s *Set) Winner() Player {
	p1, p2 := s.GamesEach()
	winningGames := s.GamesRequired
	tiebreakGames := winningGames + 1

	if p1 == tiebreakGames {
		return PLAYER_1
	}
	if p2 == tiebreakGames {
		return PLAYER_2
	}
	if p1 == winningGames && (p1-p2) >= 2 {
		return PLAYER_1
	}
	if p2 == winningGames && (p2-p1) >= 2 {
		return PLAYER_2
	}
	return NONE
}

func (s *Set) NeedsTiebreaker() bool {
	games := s.GamesRequired
	p1, p2 := s.GamesEach()
	return p1 == games && p2 == games
}

func (s *Set) Play(p1, p2 *PlayerStats, startingPlayer Player) Player {
	player := startingPlayer
	for !s.HasWinner() {
		g := CreateGame(player, s.NeedsTiebreaker())
		g.Play(p1, p2)
		s.Games = append(s.Games, g)
		player = player.OtherPlayer()
	}

	return player
}

func (s *Set) ScoreString() string {
	p1, p2 := s.GamesEach()
	return fmt.Sprintf("%d-%d", p1, p2)
}
