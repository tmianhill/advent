package main

import (
	"fmt"
	"strings"
)

func main() {
	input := sampleInput
	input = realInput

	theMap := loadMap(input)

	for t,_ := range theMap.nodesByRune {
		theMap.applyAntinodes(t)
	}

	fmt.Println(len(theMap.antinodes), "antinodes found")

	_ = theMap

	//theMap.draw()
}


type aMap struct {
	width, height int
	nodesByPos map[pos]rune
	nodesByRune map[rune][]pos
	antinodes map[pos]bool
}

type pos struct {
	x int
	y int
}

func loadMap(s string) (m aMap) {
	lines := strings.Split(s, "\n")
	m.height = len(lines)
	m.width = len(lines[0])
	m.nodesByPos = map[pos]rune{}
	m.nodesByRune = map[rune][]pos{}
	for y, l := range lines {
		for x, c := range l {
			p := pos{x, y}
			if c != '.' {
				m.nodesByPos[p] = c
				m.nodesByRune[c] = append(m.nodesByRune[c], p)
			}
		}
	}
	fmt.Println("nodes:", len(m.nodesByPos))
	fmt.Println("node types:", len(m.nodesByRune), m.nodesByRune)
	return
}

func (m *aMap) applyAntinodes(nodeType rune) {
	if m.antinodes == nil {
		m.antinodes = map[pos]bool{}
	}
	nodes := m.nodesByRune[nodeType]
	for i:=0;i<len(nodes);i++ {
		n1 := nodes[i]
		for j:=i+1;j<len(nodes);j++ {
			n2 := nodes[j]
			dx := n2.x - n1.x
			dy := n2.y - n1.y
			for l:=2;l<6;l++ {
				if dx % l == 0 && dy % l == 0 {
					dx /= l
					dy /= l
				}
			}
			for k:=-m.height;k<m.height;k++ {
				an1 := pos{n1.x + k * dx, n1.y + k * dy}
				if an1.x >= 0 && an1.x < m.width && an1.y >= 0 && an1.y < m.width {
					fmt.Println("antinode of type", nodeType, "found at", an1, "from", n1, n2)
					m.antinodes[an1] = true
				}
			}
		} 
	}
}

func (m *aMap) draw() {
	for y:=0;y<m.height;y++ {
		fmt.Println()
		for x:=0;x<m.height;x++ {
			if n,ok := m.nodesByPos[pos{x,y}]; ok {
				fmt.Print(string(n))
			} else if m.antinodes[pos{x,y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()
}