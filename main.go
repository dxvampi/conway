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

func SpawnGlider(g *Grid, ox, oy int) {
	pattern := [][2]int{
		{1, 0},
		{2, 1},
		{0, 2}, {1, 2}, {2, 2},
	}
	for _, p := range pattern {
		g.set(ox+p[0], oy+p[1], true)
	}
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

func spawnGosperGun(g *Grid, ox, oy int) {
	pattern := [][2]int{
		{24, 0},
		{22, 1}, {24, 1},
		{12, 2}, {13, 2}, {20, 2}, {21, 2}, {34, 2}, {35, 2},
		{11, 3}, {15, 3}, {20, 3}, {21, 3}, {34, 3}, {35, 3},
		{0, 4}, {1, 4}, {10, 4}, {16, 4}, {20, 4}, {21, 4},
		{0, 5}, {1, 5}, {10, 5}, {14, 5}, {16, 5}, {17, 5}, {22, 5}, {24, 5},
		{10, 6}, {16, 6}, {24, 6},
		{11, 7}, {15, 7},
		{12, 8}, {13, 8},
	}
	for _, p := range pattern {
		g.set(ox+p[0], oy+p[1], true)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	const width, heigth = 80, 40
	grid := NewGrid(width, heigth)
	
	spawnGosperGun(grid, 1, 1)

	generation := 0
	for {
		clearScreen()
		fmt.Print(grid.String())
		fmt.Printf("\nGeneration: %d", generation)
		time.Sleep(50 * time.Millisecond)
		grid = grid.Step()
		generation++
	}
}
