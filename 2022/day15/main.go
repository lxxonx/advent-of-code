package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Coordination struct {
	x, y int
}

func (c Coordination) ManhattanDist(o Coordination) int {
	dist := math.Abs(float64(c.x-o.x)) + math.Abs(float64(c.y-o.y))

	return int(dist)
}

func sol1(lines []string) {
	beaconMap := map[Coordination]rune{}

	yPos := 2000000
	for _, line := range lines {
		sx, sy, bx, by := 0, 0, 0, 0
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensor := Coordination{sx, sy}
		beacon := Coordination{bx, by}
		dist := sensor.ManhattanDist(beacon)
		for i := sensor.x - dist; i <= sensor.x+dist; i++ {
			curr := Coordination{i, yPos}
			if curr.ManhattanDist(sensor) <= dist {
				beaconMap[curr] = '#'
			}
		}
		beaconMap[beacon] = 'B'
	}

	count := 0

	for beacon, r := range beaconMap {
		if beacon.y == yPos && r == '#' {
			count++
		}
	}

	fmt.Println(count)
}

const LIMIT = 4000000

func sol2(lines []string) {
	sensorMap := map[Coordination]int{}
	toTry := map[Coordination]bool{}

	for _, line := range lines {
		sx, sy, bx, by := 0, 0, 0, 0
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensor := Coordination{sx, sy}
		beacon := Coordination{bx, by}
		dist := sensor.ManhattanDist(beacon)
		sensorMap[sensor] = dist
		dist++
		for i := 0; i < dist; i++ {
			if sx+i > 0 && sx+i < LIMIT {
				if sy-dist+1+i > 0 && sy-dist+1+i < LIMIT {
					toTry[Coordination{sx + i, sy - dist + 1 + i}] = true
				}
				if sy+dist-1-i > 0 && sy+dist-1-i < LIMIT {
					toTry[Coordination{sx + i, sy + dist - i}] = true
				}
			}
			if sx-i > 0 && sx-i < LIMIT {
				if sy-dist+1+i > 0 && sy-dist+1+i < LIMIT {
					toTry[Coordination{sx - i, sy - dist + 1 + i}] = true
				}
				if sy+dist-1-i > 0 && sy+dist-1-i < LIMIT {
					toTry[Coordination{sx - i, sy + dist - 1 - i}] = true
				}
			}
		}
	}
	for beacon := range toTry {
		newBeacon := true
		for sensor, nearestBeaconDistance := range sensorMap {
			if sensor.ManhattanDist(beacon) <= nearestBeaconDistance {
				newBeacon = false
				break
			}
		}
		if newBeacon {
			fmt.Println(beacon.x*LIMIT + beacon.y)
			break
		}
	}

}

func main() {
	sample, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(sample), "\n")

	// sol1(lines)

	sol2(lines)
}
