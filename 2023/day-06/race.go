package main

import (
	"dotariel/util"
	"regexp"
	"strconv"
	"strings"
)

type Races []Race

type Race struct {
	Time           int
	DistanceRecord int
}

func Parse(input string) Races {
	races := Races{}

	rex := regexp.MustCompile(`Time:\s+([\d\s]+)\nDistance:\s+([\d\s]+)`)

	if matches := rex.FindAllStringSubmatch(input, -1); len(matches) > 0 {
		times := util.ToInts(strings.Fields(matches[0][1]))
		distances := util.ToInts(strings.Fields(matches[0][2]))

		for i := 0; i < len(times); i++ {
			races = append(races, Race{Time: times[i], DistanceRecord: distances[i]})
		}
	}

	return races
}

func ParseSingle(input string) Race {
	rex := regexp.MustCompile(`Time:\s+([\d\s]+)\nDistance:\s+([\d\s]+)`)
	if matches := rex.FindAllStringSubmatch(input, -1); len(matches) > 0 {
		t := strings.Join(strings.Fields(matches[0][1]), "")
		d := strings.Join(strings.Fields(matches[0][2]), "")

		time, _ := strconv.Atoi(t)
		distance, _ := strconv.Atoi(d)

		return Race{Time: time, DistanceRecord: distance}
	}

	return Race{}
}

func (r Race) GetChargeTimes() []int {
	times := []int{}

	for charge := 0; charge < r.Time; charge++ {
		remaining := r.Time - charge
		velocity := charge
		distance := remaining * velocity

		if distance > r.DistanceRecord {
			times = append(times, charge)
		}
	}

	return times
}
