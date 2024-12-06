package main

import (
	"fmt"
	"strings"
)

func main() {
	input := sampleInput
	input = realInput

	theMap := loadMap(input)

	escapes := runMap(&theMap)
	fmt.Println("visited at end:", len(theMap.visited), ", escaped:", escapes)

	theMap = loadMap(input)

	validObstaclePosCount := 0
	for x:=0;x<theMap.width;x++ {
		for y:=0;y<theMap.height;y++ {
			p := pos{x,y}
			if !theMap.obstacles[p] && !theMap.visited[p] {
				m := theMap.clone(p)
				escapes := runMap(&m)
				if !escapes {
					fmt.Println("extra obstacle at", p, "works")
					validObstaclePosCount++
				}
			}
		}
	}

	fmt.Println("There are", validObstaclePosCount, "possible places for an obstacle")
}

func runMap(theMap *aMap) (escapes bool) {
	p := theMap.startPos
	dir := theMap.startDir

	for {
		dx, dy := dir.getDelta()
		newPos := pos{p.x + dx, p.y + dy}
		if newPos.x < 0 || newPos.x >= theMap.width || newPos.y < 0 || newPos.y >= theMap.height {
			return true
		}
		if theMap.obstacles[newPos] {
			dir = dir.getNewDir()
			//fmt.Println("turning right at", p, "due to obstacle at", newPos, ", new dir is", dir)
		} else {
			p = newPos
			theMap.visited[p] = true
		}
		pd := posDir{p.x,p.y,dir}
		if theMap.visitedDirs[pd] { return false }
		theMap.visitedDirs[pd] = true
	}
}


func (m aMap) clone(extraObstacle pos) aMap {
	m2 := aMap {
		width: m.width,
		height: m.height,
		visited: map[pos]bool{m.startPos:true},
		visitedDirs: map[posDir]bool{posDir{m.startPos.x, m.startPos.y, m.startDir}:true},
		obstacles: make(map[pos]bool, len(m.obstacles) + 1),
		startPos: m.startPos,
		startDir: m.startDir,
	}
	for o,_ := range m.obstacles {
		m2.obstacles[o] = true
	}
	m2.obstacles[extraObstacle] = true
	return m2
}


type direction rune

func (dir direction) getDelta() (dx int, dy int) {
	switch dir {
	case '^':
		dx = 0
		dy = -1
	case 'v':
		dx = 0
		dy = 1
	case '<':
		dx = -1
		dy = 0
	case '>':
		dx = 1
		dy = 0
	default:
		panic(fmt.Sprintf("invalid direction '%v'", dir))
	}
	return
}

func (dir direction) getNewDir() direction {
	switch dir {
	case '^':
		return '>'
	case 'v':
		return '<'
	case '<':
		return '^'
	case '>':
		return 'v'
	default:
		panic(fmt.Sprintf("invalid direction '%v'", dir))
	}
}

func loadMap(s string) (m aMap) {
	lines := strings.Split(s, "\n")
	m.height = len(lines)
	m.width = len(lines[0])
	m.visited = map[pos]bool{}
	m.obstacles = map[pos]bool{}
	m.visitedDirs = map[posDir]bool{}
	for y, l := range lines {
		for x, c := range l {
			p := pos{x, y}
			if c == '#' {
				m.obstacles[p] = true
			} else if c == '^' {
				m.visited[p] = true
				m.startPos = p
				m.startDir = direction(c)
				m.visitedDirs[posDir{x,y,m.startDir}] = true
			}
		}
	}
	fmt.Println("visited at start:", len(m.visited))
	fmt.Println("obstacles at start:", len(m.obstacles))
	return
}

type aMap struct {
	width, height      int
	visited, obstacles map[pos]bool
	startPos pos
	startDir direction
	visitedDirs map[posDir]bool
}

type pos struct {
	x int
	y int
}

type posDir struct {
	x int
	y int
	dir direction
}
