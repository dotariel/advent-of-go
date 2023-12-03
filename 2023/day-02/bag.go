package main

type Bag struct {
	cubeSet CubeSet
}

func NewBag(cubeSet CubeSet) Bag {
	return Bag{cubeSet: cubeSet}
}

func (b Bag) Validate(game Game) bool {
	for _, pull := range game.CubeSets {
		for color, count := range pull {
			if count > b.cubeSet[color] {
				return false
			}
		}
	}

	return true
}
