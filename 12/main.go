package main

import (
	"fmt"
	"strings"
)

func main() {

	input := sampleInput5
	input = realInput

	m := loadMap(input)

	regions := []region{}

	score := 0
	discountedScore := 0
	for row :=range m.plants {
		for col := range m.plants[0] {
			if !m.visited[row][col] {
				region := m.measureArea(row,col)
				regions = append(regions, region)
				fmt.Println("region of", string(m.plants[row][col]), "at row",row,"col",col,"has area",region.area,"and perimeter",region.perimeter,"and",region.sides,"sides")
				score += region.area * region.perimeter
				discountedScore += region.area * region.sides
			}
		}
	}

	fmt.Println(score, discountedScore)
}

func loadMap(input string) aMap {
	lines := strings.Split(input, "\n")
	m := aMap{}
	for _,l := range lines {
		m.plants = append(m.plants, []rune(l))
		m.visited = append(m.visited, make([]bool, len(l)))
	}
	return m
}

type aMap struct {
	plants [][]rune
	visited [][]bool
}

type region struct {
	perimeter int
	area int
	sides int
}

func (m aMap) draw() {
	for _,row := range m.plants {
		for _,p := range row {
			fmt.Print(string(p))
		}
		fmt.Println()
	}
}

type pos struct {
	row int
	col int
}

func (m aMap) measureArea(startRow, startCol int) region {
	plant := m.plants[startRow][startCol]
	posToVisit := []pos{{startRow, startCol}}
	posIndex := 0

	area := 0
	perimeter := 0
	sides := 0
	topFences := map[pos]bool{}
	bottomFences := map[pos]bool{}
	leftFences := map[pos]bool{}
	rightFences := map[pos]bool{}

	checkNeighbour := func(p pos, dR, dC int) {
		r := p.row + dR
		c := p.col + dC
		if r >= 0 && r < len(m.plants) && c >= 0 && c < len(m.plants[r]) && m.plants[r][c] == plant {
			if !m.visited[r][c] {
				posToVisit = append(posToVisit, pos{r,c})
			}
		} else {
			perimeter++
			if dR == 1 {
				bottomFences[pos{r,c}] = true
			}
			if dR == -1 {
				topFences[pos{p.row,c}] = true
			}
			if dC == 1 {
				rightFences[pos{r,c}] = true
			}
			if dC == -1 {
				leftFences[pos{r,p.col}] = true
			}
		}
	}

	for posIndex < len(posToVisit) {
		p := posToVisit[posIndex]
		posIndex++
		if m.visited[p.row][p.col] { continue }

		area++
		m.visited[p.row][p.col] = true
		checkNeighbour(p, -1, 0)
		checkNeighbour(p, 0, 1)
		checkNeighbour(p, 1, 0)
		checkNeighbour(p, 0, -1)
	}

	for p := range topFences {
		if !topFences[pos{p.row,p.col-1}] { sides++ }
	}
	for p := range bottomFences {
		if !bottomFences[pos{p.row,p.col-1}] { sides++ }
	}
	for p := range leftFences {
		if !leftFences[pos{p.row-1,p.col}] { sides++ }
	}
	for p := range rightFences {
		if !rightFences[pos{p.row-1,p.col}] { sides++ }
	}

	return region{area: area, perimeter: perimeter, sides: sides}
}

func debug(t ...any) {
	fmt.Println(t...)
}

