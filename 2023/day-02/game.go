package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Id       int
	CubeSets []CubeSet
}

func NewGame(input string) Game {
	game := Game{}

	parts := strings.Split(input, ":")

	game.Id = extractId(parts[0])
	game.CubeSets = make([]CubeSet, 0)

	for _, part := range strings.Split(parts[1], ";") {
		cubeSet := make(CubeSet, 0)

		for _, subset := range strings.Split(part, ",") {
			cube := strings.Fields(subset)
			count, _ := strconv.Atoi(cube[0])
			color := cube[1]

			cubeSet[color] = count
		}

		game.CubeSets = append(game.CubeSets, cubeSet)
	}

	return game
}

func (g Game) GetMinimumCubeSet() CubeSet {
	cubeSet := CubeSet{}

	for _, cs := range g.CubeSets {
		for color, count := range cs {
			if count > cubeSet[color] {
				cubeSet[color] = count
			}
		}
	}

	return cubeSet
}

func extractId(s string) int {
	r := regexp.MustCompile(`Game (\d+)`)
	matches := r.FindStringSubmatch(s)

	if len(matches) < 1 {
		return 0
	}

	if id, err := strconv.Atoi(matches[1]); err == nil {
		return id
	}

	return 0
}
