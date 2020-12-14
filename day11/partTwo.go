package main

func (l *Life) newStep() bool {
	stateChanged := false
	for x, row := range l.a.b {
		for y := range row {
			state, changed := l.a.newState(x, y)
			l.b.b[x][y] = state
			stateChanged = stateChanged || changed
		}
	}

	l.a, l.b = l.b, l.a
	return stateChanged
}

func (f *Field) areVisibleEmpty(x, y int) bool {
	areVisibleEmpty := true
	for i := -1; i <= 1; i++ {
		if x+i <= -1 || x+i >= f.H {
			continue
		}
		for j := -1; j <= 1; j++ {
			if (y+j <= -1 || y+j >= f.W) || (i == 0 && j == 0) {
				continue
			}
			i_d, j_d := i, j
			for f.b[x+i][y+j] == '.' {

				if i < 0 {
					i -= 1
				} else if i > 0 {
					i += 1
				}

				if j < 0 {
					j -= 1
				} else if j > 0 {
					j += 1
				}
				if y+j <= -1 || y+j >= f.W || x+i <= -1 || x+i >= f.H {
					break
				}
			}

			if !(y+j <= -1 || y+j >= f.W || x+i <= -1 || x+i >= f.H) {
				areVisibleEmpty = areVisibleEmpty && f.b[x+i][y+j] != '#'
			}
			i, j = i_d, j_d
		}
	}

	return areVisibleEmpty
}

func (f *Field) countVisibleTaken(x, y int) int {
	taken := 0
	for i := -1; i <= 1; i++ {
		if x+i <= -1 || x+i >= f.H {
			continue
		}
		for j := -1; j <= 1; j++ {
			if (y+j <= -1 || y+j >= f.W) || (j == 0 && i == 0) {
				continue
			}
			i_d, j_d := i, j
			for f.b[x+i][y+j] == '.' {
				if i < 0 {
					i -= 1
				} else if i > 0 {
					i += 1
				}
				if j < 0 {
					j -= 1
				} else if j > 0 {
					j += 1
				}
				if y+j <= -1 || y+j >= f.W || x+i <= -1 || x+i >= f.H {
					break
				}
			}
			if !(y+j <= -1 || y+j >= f.W || x+i <= -1 || x+i >= f.H) {
				if f.b[x+i][y+j] == '#' {
					taken += 1
				}
			}
			i, j = i_d, j_d
		}
	}

	return taken
}

func (f *Field) newState(x, y int) (rune, bool) {
	if f.b[x][y] == 'L' && f.areVisibleEmpty(x, y) {
		return '#', true
	}
	if f.b[x][y] == '#' && f.countVisibleTaken(x, y) >= 5 {
		return 'L', true
	}
	return f.b[x][y], false
}
