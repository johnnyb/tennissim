package main

import (
	"math/rand"
)

type PlayerStats struct {
	FirstServeRatio  float32
	SecondServeRatio float32
	HitRatio         float32
}

type Player int

const (
	NONE Player = iota
	PLAYER_1
	PLAYER_2
)

func (p Player) OtherPlayer() Player {
	if p == PLAYER_1 {
		return PLAYER_2
	}
	return PLAYER_1
}

func (p Player) WhichPlayer(p1, p2 *PlayerStats) *PlayerStats {
	if p == PLAYER_1 {
		return p1
	}
	return p2
}

func (p Player) ServePoint(p1, p2 *PlayerStats) Player {
	// Serve
	pstats := p.WhichPlayer(p1, p2)
	// First Serve
	if rand.Float32() > pstats.FirstServeRatio {
		// Second Serve
		if rand.Float32() > pstats.SecondServeRatio {
			return p.OtherPlayer()
		}
	}

	// Now play point
	currentPlayer := p.OtherPlayer()
	for {
		// Find the current percentage we are working from
		pstats := currentPlayer.WhichPlayer(p1, p2)
		// Did the player hit it out?
		if rand.Float32() > pstats.HitRatio {
			// Then the other player gets the point
			return currentPlayer.OtherPlayer()
		}
		// Now it is the other player's turn
		currentPlayer = currentPlayer.OtherPlayer()
	}
}
