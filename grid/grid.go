// Package grid functions to work with 2d grids
package grid

import "fmt"

// Grid is a 2d grid of cells.
type Grid[T cellType] struct {
	Cells  []T
	Width  int64
	Height int64
}

type cellType interface {
	rune | int
}

// IdxToPos converts slice index to 2d coordinate.
func (g *Grid[T]) IdxToPos(idx int64) Coord {
	return Coord{X: idx % g.Width, Y: idx / g.Width}
}

// PosToIdx converts 2d coordinate to slice index.
func (g *Grid[T]) PosToIdx(pos Coord) int64 {
	return pos.X + pos.Y*g.Width
}

// Set modifies the element at the given coordinate.
func (g *Grid[T]) Set(pos Coord, e T) {
	g.Cells[g.PosToIdx(pos)] = e
}

// Get returns the element at the given coordinate.
func (g *Grid[T]) Get(pos Coord) T {
	return g.Cells[g.PosToIdx(pos)]
}

// Count counts occurrences of a given element.
func (g *Grid[T]) Count(e T) int {
	initialCount := 0
	for _, r := range g.Cells {
		if r == e {
			initialCount++
		}
	}

	return initialCount
}

// Inbounds checks if a coordinate is in the grid.
func (g *Grid[T]) Inbounds(pos Coord) bool {
	return 0 <= pos.X && pos.X < g.Width && 0 <= pos.Y && pos.Y < g.Height
}

// Mvt move on the grid, return false if out of bounds.
func (g *Grid[T]) Mvt(pos Coord, mvt Coord) (Coord, bool) {
	newPos := Coord{X: pos.X + mvt.X, Y: pos.Y + mvt.Y}
	if g.Inbounds(newPos) {
		return newPos, true
	}

	return newPos, false
}

// Direction is one of 8 2d directions (straight and diagonals).
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	UpLeft
	UpRight
	DownLeft
	DownRight
)

// Value returns the 2d coordinate for a move in the given direction.
func (d Direction) Value() Coord {
	switch d {
	case Up:
		return Coord{Y: -1, X: 0}
	case Down:
		return Coord{Y: 1, X: 0}
	case Left:
		return Coord{Y: 0, X: -1}
	case Right:
		return Coord{Y: 0, X: 1}
	case UpLeft:
		return Coord{Y: -1, X: -1}
	case UpRight:
		return Coord{Y: -1, X: 1}
	case DownLeft:
		return Coord{Y: 1, X: -1}
	case DownRight:
		return Coord{Y: 1, X: 1}
	default:
		panic(fmt.Sprintf("Unknown direction %v", d))
	}
}

// Clockwise returns clockwise rotation of the direction.
func (d Direction) Clockwise() Direction {
	v := d.Value()
	newDirV := Coord{X: -v.Y, Y: v.X}

	return fromValue(newDirV)
}

// CounterClockwise returns counter-clockwise rotation of the direction.
func (d Direction) CounterClockwise() Direction {
	v := d.Value()
	newDirV := Coord{X: v.Y, Y: -v.X}

	return fromValue(newDirV)
}

func fromValue(c Coord) Direction {
	switch c {
	case Coord{Y: -1, X: 0}:
		return Up
	case Coord{Y: 1, X: 0}:
		return Down
	case Coord{Y: 0, X: -1}:
		return Left
	case Coord{Y: 0, X: 1}:
		return Right
	case Coord{Y: -1, X: -1}:
		return UpLeft
	case Coord{Y: -1, X: 1}:
		return UpRight
	case Coord{Y: 1, X: -1}:
		return DownLeft
	case Coord{Y: 1, X: 1}:
		return DownRight
	default:
		panic(fmt.Sprintf("Unknown direction %v", c))
	}
}

// Add adds two Coords together.
func (c Coord) Add(c2 Coord) Coord {
	return Coord{X: c.X + c2.X, Y: c.Y + c2.Y}
}

// CoordFrom returns a coord from a slice of int64 (size needs to be at least 2).
func CoordFrom(a []int64) Coord {
	return Coord{X: a[0], Y: a[1]}
}

//     pub fn counter_clockwise(self) -> Direction {
//         let v = self.value() * Complex { im: -1, re: 0 };
//         Direction::from_value(v)
//     }

//     pub fn clockwise(self) -> Direction {
//         let v = self.value() * Complex { im: 1, re: 0 };
//         Direction::from_value(v)
//     }

// IterDirs returns an array of all 8 directions.
func IterDirs() [8]Direction {
	return [8]Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}
}

// Coord represents a 2d coordinate.
type Coord struct {
	X int64
	Y int64
}
