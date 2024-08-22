package game

// GameMap представляет игровую карту
type GameMap struct {
	Width  int
	Height int
	Grid   [][]int
}

// NewGameMap создает новую карту с заданными размерами
func NewGameMap(width, height int) *GameMap {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	return &GameMap{
		Width:  width,
		Height: height,
		Grid:   grid,
	}
}

// NewGameMapFromLayout создает новую карту на основе предоставленного макета
func NewGameMapFromLayout(layout [][]int) *GameMap {
	height := len(layout)
	width := len(layout[0])

	return &GameMap{
		Width:  width,
		Height: height,
		Grid:   layout,
	}
}

// IsValidMove проверяет, допустимо ли перемещение на указанные координаты
func (m *GameMap) IsValidMove(x, y int) bool {
	return x >= 0 && y >= 0 && x < m.Width && y < m.Height && m.Grid[y][x] == 0
}

// ToSlice возвращает карту в виде двумерного среза
func (m *GameMap) ToSlice() [][]int {
	return m.Grid
}
