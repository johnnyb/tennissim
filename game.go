package main

import "fmt"

type Game struct {
	Server         Player
	PointsRequired int
	IsTiebreaker   bool
	P1             int
	P2             int
}

func CreateGame(server Player, isTiebreaker bool) *Game {
	points := 4
	if isTiebreaker {
		points = 7
	}
	return &Game{
		Server:         server,
		PointsRequired: points,
		IsTiebreaker:   isTiebreaker,
	}
}

func (g *Game) Play(p1, p2 *PlayerStats) {
	if g.IsTiebreaker {
		g.playTiebreaker(p1, p2)
	} else {
		g.playRegular(p1, p2)
	}
}

func (g *Game) playTiebreaker(p1, p2 *PlayerStats) {
	currentServer := g.Server
	numTimesServed := 1 // only one time for first server
	for !g.HasWinner() {
		g.servePoint(p1, p2, currentServer)
		numTimesServed++
		if numTimesServed == 2 {
			numTimesServed = 0
			currentServer = currentServer.OtherPlayer()
		}
	}
}

func (g *Game) playRegular(p1, p2 *PlayerStats) {
	for !g.HasWinner() {
		g.servePoint(p1, p2, g.Server)
	}
}

func (g *Game) servePoint(p1, p2 *PlayerStats, server Player) {
	if server.ServePoint(p1, p2) == PLAYER_1 {
		g.P1++
	} else {
		g.P2++
	}
}

func (g *Game) Winner() Player {
	if g.P1 >= g.PointsRequired && (g.P1-g.P2) >= 2 {
		return PLAYER_1
	}
	if g.P2 >= g.PointsRequired && (g.P2-g.P1) >= 2 {
		return PLAYER_2
	}
	return NONE
}

func (g *Game) HasWinner() bool {
	return g.Winner() != NONE
}

func (g *Game) ScoreString() string {
	return fmt.Sprintf("%d/%d", g.P1, g.P2)
}
