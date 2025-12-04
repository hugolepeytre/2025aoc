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
	return Coord{x: idx % g.Width, y: idx / g.Width}
}

// PosToIdx converts 2d coordinate to slice index.
func (g *Grid[T]) PosToIdx(pos Coord) int64 {
	return pos.x + pos.y*g.Width
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
	return 0 <= pos.x && pos.x < g.Width && 0 <= pos.y && pos.y < g.Height
}

// Mvt move on the grid, return false if out of bounds.
func (g *Grid[T]) Mvt(pos Coord, mvt Coord) (Coord, bool) {
	newPos := Coord{x: pos.x + mvt.x, y: pos.y + mvt.y}
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
		return Coord{y: -1, x: 0}
	case Down:
		return Coord{y: 1, x: 0}
	case Left:
		return Coord{y: 0, x: -1}
	case Right:
		return Coord{y: 0, x: 1}
	case UpLeft:
		return Coord{y: -1, x: -1}
	case UpRight:
		return Coord{y: -1, x: 1}
	case DownLeft:
		return Coord{y: 1, x: -1}
	case DownRight:
		return Coord{y: 1, x: 1}
	default:
		panic(fmt.Sprintf("Unknown direction %v", d))
	}
}

// IterDirs returns an array of all 8 directions.
func IterDirs() [8]Direction {
	return [8]Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}
}

//     pub fn from_value(value: Complex<i32>) -> Direction {
//         match value {
//             Complex { im: -1, re: 0 } => Direction::Up,
//             Complex { im: 1, re: 0 } => Direction::Down,
//             Complex { im: 0, re: -1 } => Direction::Left,
//             Complex { im: 0, re: 1 } => Direction::Right,
//             Complex { im: -1, re: -1 } => Direction::UpLeft,
//             Complex { im: -1, re: 1 } => Direction::UpRight,
//             Complex { im: 1, re: -1 } => Direction::DownLeft,
//             Complex { im: 1, re: 1 } => Direction::DownRight,
//             _ => panic!("Invalid direction"),
//         }
//     }

//     pub fn counter_clockwise(self) -> Direction {
//         let v = self.value() * Complex { im: -1, re: 0 };
//         Direction::from_value(v)
//     }

//     pub fn clockwise(self) -> Direction {
//         let v = self.value() * Complex { im: 1, re: 0 };
//         Direction::from_value(v)
//     }

//     pub fn opposite(self) -> Direction {
//         let v = self.value() * Complex { im: 0, re: -1 };
//         Direction::from_value(v)
//     }

//     pub fn from_arrow(a: char) -> Direction {
//         match a {
//             '^' => Direction::Up,
//             '<' => Direction::Left,
//             '>' => Direction::Right,
//             'v' => Direction::Down,
//             c => panic!("Not an arrow: {}", c),
//         }
//     }

//     pub fn iter_diags() -> [Direction; 4] {
//         [
//             Direction::UpLeft,
//             Direction::UpRight,
//             Direction::DownLeft,
//             Direction::DownRight,
//         ]
//     }

//     pub fn iter_straight() -> [Direction; 4] {
//         [
//             Direction::Up,
//             Direction::Down,
//             Direction::Left,
//             Direction::Right,
//         ]
//     }

//	    pub fn iter() -> [Direction; 8] {
//	        [
//	            Direction::Up,
//	            Direction::Down,
//	            Direction::Left,
//	            Direction::Right,
//	            Direction::UpLeft,
//	            Direction::UpRight,
//	            Direction::DownLeft,
//	            Direction::DownRight,
//	        ]
//	    }
//	}

// Coord represents a 2d coordinate.
type Coord struct {
	x int64
	y int64
}
