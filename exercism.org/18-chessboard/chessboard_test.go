package chessboard

import (
	"testing"
)

// newChessboard return a *Chessboard for tests
//
//   A B C D E F G H
// 8 # _ _ _ # _ _ # 8
// 7 _ _ _ _ _ _ _ _ 7
// 6 _ _ _ _ # _ _ # 6
// 5 _ # _ _ _ _ _ # 5
// 4 _ _ _ _ _ _ # # 4
// 3 # _ # _ _ _ _ # 3
// 2 _ _ _ _ _ _ _ # 2
// 1 # _ _ _ _ _ _ # 1
//   A B C D E F G H

func newChessboard() Chessboard {
	return Chessboard{
		"A": File{true, false, true, false, false, false, false, true},
		"B": File{false, false, false, false, true, false, false, false},
		"C": File{false, false, true, false, false, false, false, false},
		"D": File{false, false, false, false, false, false, false, false},
		"E": File{false, false, false, false, false, true, false, true},
		"F": File{false, false, false, false, false, false, false, false},
		"G": File{false, false, false, true, false, false, false, false},
		"H": File{true, true, true, true, true, true, false, true},
	}
}

func TestCountInFile(t *testing.T) {
	cb := newChessboard()
	for _, test := range []struct {
		in  string
		out int
	}{
		{"A", 3},
		{"B", 1},
		{"C", 1},
		{"D", 0},
		{"E", 2},
		{"F", 0},
		{"G", 1},
		{"H", 7},
		{"Z", 0},
	} {
		if out := CountInFile(cb, test.in); out != test.out {
			t.Errorf(
				"CountInFile(chessboard, '%v') returned %v while %v was expected\n",
				test.in,
				out,
				test.out,
			)
		}
	}
}

func TestCountInRank(t *testing.T) {
	cb := newChessboard()
	for _, test := range []struct {
		in  int
		out int
	}{
		{1, 2},
		{2, 1},
		{3, 3},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 0},
		{8, 3},
		// cases not between 1 and 8, inclusive
		{100, 0},
		{0, 0},
		{-1, 0},
		{-100, 0},
	} {
		if out := CountInRank(cb, test.in); out != test.out {
			t.Errorf(
				"CountInRank(chessboard, %v) returned %v while %v was expected\n",
				test.in,
				out,
				test.out,
			)
		}
	}
}

func TestCountAll(t *testing.T) {
	cb := newChessboard()
	wanted := 64
	if out := CountAll(cb); out != wanted {
		t.Errorf("CountAll(chessboard) returned %v while %v was expected", out, wanted)
	}
}

func TestCountOccupied(t *testing.T) {
	cb := newChessboard()
	wanted := 15
	if out := CountOccupied(cb); out != wanted {
		t.Errorf("CountOccupied(chessboard) returned %v while %v was expected", out, wanted)
	}
}
