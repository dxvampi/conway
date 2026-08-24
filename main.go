package main

import (
	"fmt"
	"strings"
	"time"
)

type Grid struct {
	w, h int
	cell []bool
}

func NewGrid(w, h int) *Grid {
	return &Grid{w: w, h: h, cell: make([]bool, w*h)}
}

func (g *Grid) at(x, y int) bool {
	x = (x + g.w) % g.w
	y = (y + g.h) % g.h
	return g.cell[y*g.w+x]
}

func (g *Grid) set(x, y int, alive bool) {
	g.cell[y*g.w+x] = alive
}

func (g *Grid) neighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if g.at(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

func (g *Grid) Step() *Grid {
	next := NewGrid(g.w, g.h)
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			n := g.neighbors(x, y)
			alive := g.at(x, y)
			switch {
			case alive && (n == 2 || n == 3):
				next.set(x, y, true)
			
			case !alive && n == 3:
				next.set(x, y, true)

			default:
			  next.set(x, y, false)
			}
		}
	}
	return next
}

func (g *Grid) String() string {
	var b strings.Builder
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			if g.at(x, y) {
				b.WriteString("O")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func spawnRPentonimo(g *Grid, ox, oy int) {
	pattern := [][2]int{
		{1, 0}, {2, 0},
		{0, 1}, {1, 1},
		{1, 2},
	}

	for _, p := range pattern {
		g.set(ox+p[0], oy+p[1], true)
	}
}

func (g *Grid) Equal(other *Grid) bool {
	if g.w != other.w || g.h != other.h {
		return false
	}
	for i := range g.cell {
		if g.cell[i] != other.cell[i] {
			return false
		}
	}
	return true
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	const width, heigth = 80, 40
	grid := NewGrid(width, heigth)
	
	spawnRPentonimo(grid, 19, 9)
	generation := 0

	var history []*Grid

	for {
		clearScreen()
		fmt.Print(grid.String())
		fmt.Printf("\nGeneration: %d\n", generation)
		time.Sleep(5 * time.Millisecond)

		matchIndex := -1
		for i, past := range history {
			if grid.Equal(past) {
				matchIndex = i
				break
			}
		}

		if matchIndex != -1 {
			stepsAgo := len(history) - matchIndex
			if stepsAgo == 1 {
				fmt.Printf("\nStill life in generation %d\n", generation)
			} else {
				fmt.Printf("\nLoop detected, period %d (last seen %d generations ago)\n", stepsAgo, stepsAgo)
			}
			break
		}

		history = append(history, grid)
		if len(history) > 8 {
			history = history[1:]
		}

		grid = grid.Step()
		generation++
	}
}
