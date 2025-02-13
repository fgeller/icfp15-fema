package main

import "testing"

func TestMoveToLowerRight(t *testing.T) {
	b := NewBoard(2, 3, []Cell{})
	atom := Unit{Members: []Cell{Cell{X: 1, Y: 0}}, Pivot: Cell{X: 1, Y: 0}}
	target := Unit{Members: []Cell{Cell{X: 2, Y: 1}}, Pivot: Cell{X: 2, Y: 1}}

	actual := b.MoveSequence(atom, target)
	expected := []Move{E, SE, SE}

	if len(actual) != len(expected) {
		t.Errorf("Not the same amount of moves: %v expected %v", actual, expected)
		return
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Failed to move to lower right, got moves: %v expected %v", actual, expected)
		}
	}
}

func TestMoveToLowerRightWithObstacle(t *testing.T) {
	b := NewBoard(2, 3, []Cell{Cell{X: 2, Y: 0}})
	atom := Unit{Members: []Cell{Cell{X: 1, Y: 0}}, Pivot: Cell{X: 1, Y: 0}}
	target := Unit{Members: []Cell{Cell{X: 2, Y: 1}}, Pivot: Cell{X: 2, Y: 1}}

	actual := b.MoveSequence(atom, target)
	expected := []Move{SE, E, SE}

	if len(actual) != len(expected) {
		t.Errorf("Not the same amount of moves: %v expected %v", actual, expected)
		return
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Failed to move to lower right, got moves: %v expected %v", actual, expected)
		}
	}
}

func TestNoSequencePossible(t *testing.T) {
	b := NewBoard(3, 2, []Cell{Cell{X: 0, Y: 1}, Cell{X: 1, Y: 1}})
	atom := Unit{Members: []Cell{Cell{X: 0, Y: 0}}, Pivot: Cell{X: 0, Y: 0}}
	target := Unit{Members: []Cell{Cell{X: 0, Y: 2}}, Pivot: Cell{X: 0, Y: 2}}

	actual := b.MoveSequence(atom, target)
	expected := []Move{}

	if len(actual) != len(expected) {
		t.Errorf("Not the same amount of moves: %v expected %v", actual, expected)
		return
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Failed to move to lower right, got moves: %v expected %v", actual, expected)
		}
	}
}

func TestMoveToLowerLeft(t *testing.T) {
	b := NewBoard(2, 3, []Cell{})
	atom := Unit{Members: []Cell{Cell{X: 1, Y: 0}}, Pivot: Cell{X: 1, Y: 0}}
	target := Unit{Members: []Cell{Cell{X: 0, Y: 1}}, Pivot: Cell{X: 0, Y: 1}}

	actual := b.MoveSequence(atom, target)
	expected := []Move{W, SE, SE}

	if len(actual) != len(expected) {
		t.Errorf("Not the same amount of moves: %v expected %v", actual, expected)
		return
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Failed to move to lower right, got moves: %v expected %v", actual, expected)
		}
	}
}

func TestMoveFurtherToLowerLeft(t *testing.T) {
	b := NewBoard(5, 3, []Cell{})
	atom := Unit{Members: []Cell{Cell{X: 1, Y: 0}}, Pivot: Cell{X: 1, Y: 0}}
	target := Unit{Members: []Cell{Cell{X: 0, Y: 4}}, Pivot: Cell{X: 0, Y: 4}}

	actual := b.MoveSequence(atom, target)
	expected := []Move{W, SE, SE, W, SE, SE, W, SE}

	if len(actual) != len(expected) {
		t.Errorf("Not the same amount of moves: %v expected %v", actual, expected)
		return
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Failed to move to further lower left, got moves: %v expected %v", actual, expected)
		}
	}
}

func TestUnitWidth(t *testing.T) {
	atoms := []struct {
		atom     Unit
		expected int
	}{
		{
			atom:     Unit{Members: []Cell{Cell{0, 0}}, Pivot: Cell{0, 0}},
			expected: 1,
		},
		{
			atom:     Unit{Members: []Cell{Cell{0, 0}, Cell{2, 0}}, Pivot: Cell{1, 0}},
			expected: 3,
		},
	}

	for _, data := range atoms {
		actual := data.atom.Width()

		if actual != data.expected {
			t.Errorf("Failed identify width: %v expected %v", actual, data.expected)
		}
	}

}

func TestUnitHeight(t *testing.T) {
	atoms := []struct {
		atom     Unit
		expected int
	}{
		{
			atom:     Unit{Members: []Cell{Cell{0, 0}}, Pivot: Cell{0, 0}},
			expected: 1,
		},
		{
			atom:     Unit{Members: []Cell{Cell{0, 0}, Cell{2, 2}}, Pivot: Cell{1, 1}},
			expected: 3,
		},
	}

	for _, data := range atoms {
		actual := data.atom.Height()

		if actual != data.expected {
			t.Errorf("Failed identify height: %v expected %v", actual, data.expected)
		}
	}

}

func TestStartLocation(t *testing.T) {
	atoms := []struct {
		board    Board
		atom     Unit
		expected Unit
	}{
		{
			board:    NewBoard(2, 2, []Cell{}),
			atom:     Unit{Members: []Cell{Cell{0, 0}}, Pivot: Cell{0, 0}},
			expected: Unit{Members: []Cell{Cell{0, 0}}, Pivot: Cell{0, 0}},
		},
		{
			board:    NewBoard(2, 3, []Cell{}),
			atom:     Unit{Members: []Cell{Cell{0, 0}}, Pivot: Cell{0, 0}},
			expected: Unit{Members: []Cell{Cell{1, 0}}, Pivot: Cell{1, 0}},
		},
		{
			board:    NewBoard(2, 5, []Cell{}),
			atom:     Unit{Members: []Cell{Cell{0, 1}, Cell{1, 1}}, Pivot: Cell{0, 0}},
			expected: Unit{Members: []Cell{Cell{1, 0}, Cell{2, 0}}, Pivot: Cell{1, -1}},
		},
	}

	for _, data := range atoms {
		actual := data.board.StartLocation(data.atom)

		if len(actual.Members) != len(data.expected.Members) {
			t.Errorf("Not the same number of Members: %v expected %v", actual, data.expected)
			return
		}

		for i, member := range data.expected.Members {
			if actual.Members[i] != member {
				t.Errorf("Failed identify height: %v expected %v", actual, data.expected)
			}
		}
	}

}

func TestReadProgram(t *testing.T) {
	sample := `{"id": 23, "units": [], "width": 5, "height": 5, "filled": [], "sourceLength": 2, "sourceSeeds": []}`
	actual := ReadProgram([]byte(sample))
	if actual.Id != 23 || actual.Width != 5 || actual.Height != 5 || actual.SourceLength != 2 {
		t.Errorf("Failed to read program got: %v", actual)
	}

	sample = `{"id": 21, "units": [], "width": 10, "height": 5, "filled": [], "sourceLength": 2, "sourceSeeds": []}`
	actual = ReadProgram([]byte(sample))
	if actual.Id != 21 || actual.Width != 10 || actual.Height != 5 || actual.SourceLength != 2 {
		t.Errorf("Failed to read program got: %v", actual)
	}

	sample = `{"id": 21, "units": [{"members": [{"x": 1, "y": 1}], "pivot": {"x": 1, "y": 1}}], "width": 10, "height": 5, "filled": [{"x": 1, "y": 2}], "sourceLength": 2, "sourceSeeds": []}`
	actual = ReadProgram([]byte(sample))
	if len(actual.Units) != 1 || len(actual.Units[0].Members) != 1 || len(actual.Filled) != 1 {
		t.Errorf("Failed to read program got: %v", actual)
	}

}

func TestFillBoard(t *testing.T) {
	b := NewBoard(2, 2, []Cell{Cell{X: 1, Y: 1}})
	actual := b.FillCells([]Cell{Cell{X: 0, Y: 0}})

	if !actual[0][0] || !actual[1][1] || actual[0][1] || actual[1][0] {
		t.Errorf("Failed to read fill board got: %v", actual)
	}

}

func TestMoveCell(t *testing.T) {
	data := []struct {
		c        Move
		s        Cell
		expected Cell
	}{
		{
			c:        E,
			s:        Cell{0, 0},
			expected: Cell{1, 0},
		},
		{
			c:        W,
			s:        Cell{1, 0},
			expected: Cell{0, 0},
		},
		{
			c:        SE,
			s:        Cell{0, 0},
			expected: Cell{0, 1},
		},
		{
			c:        SW,
			s:        Cell{0, 0},
			expected: Cell{-1, 1},
		},
		{
			c:        SW,
			s:        Cell{1, 0},
			expected: Cell{0, 1},
		},
		{
			c:        SE,
			s:        Cell{1, 0},
			expected: Cell{1, 1},
		},
		{
			c:        SE,
			s:        Cell{1, 1},
			expected: Cell{2, 2},
		},
	}

	for _, d := range data {
		if actual := d.s.Move(d.c, Cell{}); actual != d.expected {
			t.Errorf("incorrect move: actual %v expected %v", actual, d.expected)
		}
	}

}

func TestMoveCellBiggerBoard(t *testing.T) {
	data := []struct {
		c        Move
		s        Cell
		expected Cell
	}{
		{
			c:        SE,
			s:        Cell{0, 0},
			expected: Cell{0, 1},
		},
		{
			c:        SE,
			s:        Cell{0, 1},
			expected: Cell{1, 2},
		},
		{
			c:        SW,
			s:        Cell{0, 0},
			expected: Cell{-1, 1},
		},
		{
			c:        SW,
			s:        Cell{1, 0},
			expected: Cell{0, 1},
		},
		{
			c:        SW,
			s:        Cell{1, 1},
			expected: Cell{1, 2},
		},
	}

	for _, d := range data {
		if actual := d.s.Move(d.c, Cell{}); actual != d.expected {
			t.Errorf("incorrect move: actual %v expected %v", actual, d.expected)
		}
	}

}

func TestBoardString(t *testing.T) {
	data := []struct {
		board    Board
		expected string
	}{
		{
			board: NewBoard(2, 2, []Cell{}),
			expected: `⬡ ⬡
 ⬡ ⬡`,
		},
		{
			board: NewBoard(2, 2, []Cell{Cell{0, 0}, Cell{1, 1}}),
			expected: `⬢ ⬡
 ⬡ ⬢`,
		},
		{
			board: NewBoard(3, 3, []Cell{Cell{0, 0}, Cell{1, 1}}),
			expected: `⬢ ⬡ ⬡
 ⬡ ⬢ ⬡
⬡ ⬡ ⬡`,
		},
	}

	for _, d := range data {
		if actual := d.board.String(); d.expected != actual {
			t.Errorf("weird string for board, actual\n%v\nexpected:\n%v\n", actual, d.expected)
		}
	}
}

func TestUnitMove(t *testing.T) {
	// east
	u := Unit{Members: []Cell{Cell{0, 0}, Cell{0, 1}, Cell{0, 2}}, Pivot: Cell{0, 0}}
	actual := u.Move(E)
	expected := Unit{Members: []Cell{Cell{1, 0}, Cell{1, 1}, Cell{1, 2}}, Pivot: Cell{1, 0}}

	if actual.Pivot != expected.Pivot {
		t.Errorf("wrong pivot: %v expected %v", actual.Pivot, expected.Pivot)
	}
	for mi, m := range expected.Members {
		if m != actual.Members[mi] {
			t.Errorf("wrong member: %v expected %v", actual.Members[mi], m)
		}
	}

	// west
	u = Unit{Members: []Cell{Cell{1, 0}, Cell{1, 1}, Cell{1, 2}}, Pivot: Cell{1, 0}}
	actual = u.Move(W)
	expected = Unit{Members: []Cell{Cell{0, 0}, Cell{0, 1}, Cell{0, 2}}, Pivot: Cell{0, 0}}

	if actual.Pivot != expected.Pivot {
		t.Errorf("wrong pivot: %v expected %v", actual.Pivot, expected.Pivot)
	}
	for mi, m := range expected.Members {
		if m != actual.Members[mi] {
			t.Errorf("wrong member: %v expected %v", actual.Members[mi], m)
		}
	}

	// southeast
	u = Unit{Members: []Cell{Cell{0, 0}, Cell{0, 1}, Cell{0, 2}}, Pivot: Cell{0, 0}}
	actual = u.Move(SE)
	expected = Unit{Members: []Cell{Cell{0, 1}, Cell{1, 2}, Cell{0, 3}}, Pivot: Cell{0, 1}}

	if actual.Pivot != expected.Pivot {
		t.Errorf("wrong pivot: %v expected %v", actual.Pivot, expected.Pivot)
	}
	for mi, m := range expected.Members {
		if m != actual.Members[mi] {
			t.Errorf("wrong member: %v expected %v", actual.Members[mi], m)
		}
	}

	// southwest
	u = Unit{Members: []Cell{Cell{1, 0}, Cell{1, 1}, Cell{1, 2}}, Pivot: Cell{1, 0}}
	actual = u.Move(SW)
	expected = Unit{Members: []Cell{Cell{0, 1}, Cell{1, 2}, Cell{0, 3}}, Pivot: Cell{0, 1}}

	if actual.Pivot != expected.Pivot {
		t.Errorf("wrong pivot: %v expected %v", actual.Pivot, expected.Pivot)
	}
	for mi, m := range expected.Members {
		if m != actual.Members[mi] {
			t.Errorf("wrong member: %v expected %v", actual.Members[mi], m)
		}
	}
}

func TestMoveTriplet(t *testing.T) {

	// east
	u := Unit{Members: []Cell{Cell{0, 0}, Cell{0, 1}, Cell{0, 2}}, Pivot: Cell{0, 0}}
	actual := u.MoveTo(Cell{1, 0}, u.Pivot)
	expected := Unit{Members: []Cell{Cell{1, 0}, Cell{1, 1}, Cell{1, 2}}, Pivot: Cell{1, 0}}

	if actual.Pivot != expected.Pivot {
		t.Errorf("wrong pivot: %v expected %v", actual.Pivot, expected.Pivot)
	}
	for mi, m := range expected.Members {
		if m != actual.Members[mi] {
			t.Errorf("wrong member: %v expected %v", actual.Members[mi], m)
		}
	}

	// west
	actual = expected.MoveTo(Cell{0, 0}, expected.Pivot)
	expected = Unit{Members: []Cell{Cell{0, 0}, Cell{0, 1}, Cell{0, 2}}, Pivot: Cell{0, 0}}

	if actual.Pivot != expected.Pivot {
		t.Errorf("wrong pivot: %v expected %v", actual.Pivot, expected.Pivot)
	}
	for mi, m := range expected.Members {
		if m != actual.Members[mi] {
			t.Errorf("wrong member: %v expected %v", actual.Members[mi], m)
		}
	}

	// southeast
	u = Unit{Members: []Cell{Cell{0, 0}, Cell{0, 1}, Cell{0, 2}}, Pivot: Cell{0, 0}}
	expected = Unit{Members: []Cell{Cell{0, 1}, Cell{1, 2}, Cell{0, 3}}, Pivot: Cell{0, 1}}
	actual = u.MoveTo(Cell{0, 1}, u.Pivot)

	if actual.Pivot != expected.Pivot {
		t.Errorf("wrong pivot: %v expected %v", actual.Pivot, expected.Pivot)
	}
	for mi, m := range expected.Members {
		if m != actual.Members[mi] {
			t.Errorf("wrong member: %v expected %v", actual.Members[mi], m)
		}
	}

	// southwest
	actual = expected.MoveTo(Cell{0, 1}, expected.Pivot)
	expected = Unit{Members: []Cell{Cell{0, 1}, Cell{1, 2}, Cell{0, 3}}, Pivot: Cell{0, 1}}

	if actual.Pivot != expected.Pivot {
		t.Errorf("wrong pivot: %v expected %v", actual.Pivot, expected.Pivot)
	}
	for mi, m := range expected.Members {
		if m != actual.Members[mi] {
			t.Errorf("wrong member: %v expected %v", actual.Members[mi], m)
		}
	}
}

// ⬡ ⬡ ⬡ ⬡ ⬢ ⬢ ⬡ ⬡ ⬡ ⬡
//  ⬢ ⬡ ⬢ ⬡ ⬢ ⬡ ⬢ ⬡ ⬢ ⬡
// ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡
//  ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡
// ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡
//  ⬡ ⬢ ⬢ ⬢ ⬢ ⬢ ⬢ ⬢ ⬢ ⬡
// ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡ ⬡

// func TestPathFindingFailureP10(t *testing.T) {

// 	b := NewBoard(
// 		7,
// 		10,
// 		[]Cell{
// 			Cell{0, 1},
// 			Cell{2, 1},
// 			Cell{4, 1},
// 			Cell{6, 1},
// 			Cell{8, 1},
// 			Cell{1, 5},
// 			Cell{2, 5},
// 			Cell{3, 5},
// 			Cell{4, 5},
// 			Cell{5, 5},
// 			Cell{6, 5},
// 			Cell{7, 5},
// 			Cell{8, 5},
// 		},
// 	)

// 	sl := Unit{Members: []Cell{Cell{2, 0}, Cell{3, 0}}, Pivot: Cell{0, 0}}
// 	fmt.Printf("%v\n", b.FillCells(sl.Members))
// 	tl := Unit{Members: []Cell{Cell{8, 6}, Cell{9, 6}}, Pivot: Cell{6, 6}}

// 	ms := b.MoveSequence(sl, tl)

// 	fmt.Printf("got moves: %v\n", ms)
// 	// expected := []Move{E, E, E, SE, W, SE, SE, W, SE, SW, SE, SE}
// 	// for i, em := range expected {
// 	// 	if ms[i] != em {
// 	// 		t.Errorf("did not find proper path. got %v expected %v", ms, expected)
// 	// 	}
// 	// }
// }

func TestPathFindingFailure(t *testing.T) {

	b := NewBoard(
		10,
		10,
		[]Cell{
			Cell{0, 6},
			Cell{5, 6},
			Cell{9, 6},
			Cell{0, 7},
			Cell{2, 7},
			Cell{3, 7},
			Cell{7, 7},
			Cell{8, 7},
			Cell{1, 8},
			Cell{2, 8},
			Cell{3, 8},
			Cell{5, 8},
			Cell{6, 8},
			Cell{8, 8},
			Cell{9, 8},
			Cell{1, 9},
			Cell{2, 9},
			Cell{3, 9},
			Cell{4, 9},
			Cell{6, 9},
			Cell{7, 9},
			Cell{8, 9},
			Cell{9, 9},
		},
	)

	sl := Unit{Members: []Cell{Cell{4, 0}, Cell{4, 2}}, Pivot: Cell{4, 1}}
	tl := Unit{Members: []Cell{Cell{7, 6}, Cell{7, 8}}, Pivot: Cell{7, 7}}

	ms := b.MoveSequence(sl, tl)

	expected := []Move{E, E, E, SE, W, SE, SE, W, SE, SW, SE, SE}

	for i, em := range expected {
		if ms[i] != em {
			t.Errorf("did not find proper path. got %v expected %v", ms, expected)
		}
	}
}

func TestRotateMove(t *testing.T) {

	//  ⬡ ⬡ ⬡ ⬡
	//   ⬢ · ⬢ ⬡
	//  ⬡ ⬡ ⬡ ⬡

	u := Unit{Members: []Cell{Cell{0, 1}, Cell{2, 1}}, Pivot: Cell{1, 1}}

	actual := u.Move(RC)
	expected := Unit{Members: []Cell{Cell{1, 0}, Cell{2, 2}}, Pivot: Cell{1, 1}}

	if !actual.eq(expected) {
		t.Errorf("incorrect 1st move, got\n%v expected\n%v", actual, expected)
	}

	//  ⬡ ⬢ ⬡ ⬡
	//   ⬡ · ⬡ ⬡
	//  ⬡ ⬡ ⬢ ⬡

	actual = expected.Move(RC)
	expected = Unit{Members: []Cell{Cell{2, 0}, Cell{1, 2}}, Pivot: Cell{1, 1}}

	if !actual.eq(expected) {
		t.Errorf("incorrect 2rd move, got\n%v expected\n%v", actual, expected)
	}

	//  ⬡ ⬡ ⬢ ⬡
	//   ⬡ · ⬡ ⬡
	//  ⬡ ⬢ ⬡ ⬡

	actual = expected.Move(RC)
	expected = Unit{Members: []Cell{Cell{2, 1}, Cell{0, 1}}, Pivot: Cell{1, 1}}

	if !actual.eq(expected) {
		t.Errorf("incorrect 3rd move, got\n%v expected\n%v", actual, expected)
	}

	//  ⬡ ⬡ ⬡ ⬡
	//   ⬢ · ⬢ ⬡
	//  ⬡ ⬡ ⬡ ⬡

	actual = expected.Move(RCC)
	expected = Unit{Members: []Cell{Cell{2, 0}, Cell{1, 2}}, Pivot: Cell{1, 1}}

	if !actual.eq(expected) {
		t.Errorf("incorrect 3rd move, got\n%v expected\n%v", actual, expected)
	}
}

func TestRotateMoveWeirdo(t *testing.T) {

	//  · ⬡ ⬢ ⬢
	//   ⬡ ⬡ ⬡ ⬡
	//  ⬡ ⬡ ⬡ ⬡
	//   ⬡ ⬡ ⬡ ⬡

	//  · ⬡ ⬡ ⬡
	//   ⬡ ⬡ ⬡ ⬡
	//  ⬡ ⬢ ⬡ ⬡
	//   ⬡ ⬢ ⬡ ⬡

	u := Unit{Members: []Cell{Cell{2, 0}, Cell{3, 0}}, Pivot: Cell{0, 0}}

	actual := u.Move(RC)
	expected := Unit{Members: []Cell{Cell{1, 2}, Cell{1, 3}}, Pivot: Cell{0, 0}}

	if !actual.eq(expected) {
		t.Errorf("incorrect 1st move, got\n%v expected\n%v", actual, expected)
	}

}
