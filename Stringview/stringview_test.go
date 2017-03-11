package Stringview

import (
	c "../Constants"

	gb "../Gameboard"
	gf "../Gamefields"
	fo "../Gameobjects"

	"reflect"
	"testing"
)

// Test helper to wrap up test invokation and separate it from test cases definition
// Also, since this IS golang 1.7, this wrapper will be very usefull  with test filtering
func runTest(testAlias string, t *testing.T, testFunction func(t *testing.T)) {
	// For golang 1.7+
	t.Run(testAlias, testFunction)
}

func TestCompleteBoard(t *testing.T) {
	testCases := []struct {
		alias          string
		expectedString string
		board          gb.Board
	}{
		{
			alias:          "One Row - One gf.Field - Nothing Special",
			expectedString: "|---|\n|   |\n|   |\n|---|",
			board: &gb.FakeBoard{
				map[int][]gf.Field{
					0: []gf.Field{
						&gf.FakeField{
							[2]int{0, 0},
							[]c.Border{c.WALL, c.WALL, c.WALL, c.WALL},
						},
					},
				},
				[][]fo.FieldObject{},
				-1,
			},
		},
		{
			alias:          "One Row - Two Fields seperated by c.DOOR",
			expectedString: "|---|---|\n|   |   |\n|       |\n|---|---|",
			board: &gb.FakeBoard{
				map[int][]gf.Field{
					0: []gf.Field{
						&gf.FakeField{
							[2]int{0, 0},
							[]c.Border{c.WALL, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 1},
							[]c.Border{c.DOOR, c.WALL},
						},
					},
				},
				[][]fo.FieldObject{},
				-1,
			},
		},
		{
			alias:          "Two Rows - Four Fields seperated by c.DOORs",
			expectedString: "|---|---|\n|   |   |\n|       |\n|- -|- -|\n|   |   |\n|       |\n|---|---|",
			board: &gb.FakeBoard{
				map[int][]gf.Field{
					0: []gf.Field{
						&gf.FakeField{
							[2]int{0, 0},
							[]c.Border{c.WALL, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 1},
							[]c.Border{c.DOOR, c.WALL},
						},
					},
					1: []gf.Field{
						&gf.FakeField{
							[2]int{1, 0},
							[]c.Border{c.WALL, c.DOOR},
						},
						&gf.FakeField{
							[2]int{1, 1},
							[]c.Border{c.DOOR, c.DOOR},
						},
					},
				},
				[][]fo.FieldObject{},
				-1,
			},
		},
		{
			alias:          "Two Rows - Four Fields - Three Rooms",
			expectedString: "|---|---|\n|   |   |\n|   |   |\n|   |- -|\n|   |   |\n|       |\n|---|---|",
			board: &gb.FakeBoard{
				map[int][]gf.Field{
					0: []gf.Field{
						&gf.FakeField{
							[2]int{0, 0},
							[]c.Border{c.WALL, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 1},
							[]c.Border{c.WALL, c.WALL},
						},
					},
					1: []gf.Field{
						&gf.FakeField{
							[2]int{1, 0},
							[]c.Border{c.WALL, c.FREE},
						},
						&gf.FakeField{
							[2]int{1, 1},
							[]c.Border{c.DOOR, c.DOOR},
						},
					},
				},
				[][]fo.FieldObject{},
				-1,
			},
		},
		{
			alias:          "Two Rows - Six Fields - c.CROSSWALK",
			expectedString: "|---|---|---|\n|   |       |\n|   |       |\n| = |- -|---|\n|   =   =   |\n|   =   =   |\n|---|---|---|",
			board: &gb.FakeBoard{
				map[int][]gf.Field{
					0: []gf.Field{
						&gf.FakeField{
							[2]int{0, 0},
							[]c.Border{c.WALL, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 1},
							[]c.Border{c.WALL, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 2},
							[]c.Border{c.FREE, c.WALL},
						},
					},
					1: []gf.Field{
						&gf.FakeField{
							[2]int{1, 0},
							[]c.Border{c.WALL, c.CROSSWALK},
						},
						&gf.FakeField{
							[2]int{1, 1},
							[]c.Border{c.CROSSWALK, c.DOOR},
						},
						&gf.FakeField{
							[2]int{1, 2},
							[]c.Border{c.CROSSWALK, c.WALL},
						},
					},
				},
				[][]fo.FieldObject{},
				-1,
			},
		},
		{
			alias:          "One Row - Two Fields - Zombs and Survivors",
			expectedString: "|---|---|\n|ab |zzy|\n|       |\n|---|---|",
			board: &gb.FakeBoard{
				map[int][]gf.Field{
					0: []gf.Field{
						&gf.FakeField{
							[2]int{0, 0},
							[]c.Border{c.WALL, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 1},
							[]c.Border{c.DOOR, c.WALL},
						},
					},
				},
				[][]fo.FieldObject{
					[]fo.FieldObject{
						&fo.FakeFieldObject{Identifier: "a"},
						&fo.FakeFieldObject{Identifier: "b"},
					},
					[]fo.FieldObject{
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "y"},
					},
				},
				-1,
			},
		},
		{
			alias:          "Big Map - Lots of Zombs",
			expectedString: "|---|---|---|---|\n|a  |zzz|y  |z  |\n|   |zz+        |\n|- -|---|---| = |\n|   =b  =   =   |\n|   =   =   =   |\n|---|---|---|---|",
			board: &gb.FakeBoard{
				map[int][]gf.Field{
					0: []gf.Field{
						&gf.FakeField{
							[2]int{0, 0},
							[]c.Border{c.WALL, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 1},
							[]c.Border{c.WALL, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 2},
							[]c.Border{c.DOOR, c.WALL},
						},
						&gf.FakeField{
							[2]int{0, 3},
							[]c.Border{c.DOOR, c.WALL},
						},
					},
					1: []gf.Field{
						&gf.FakeField{
							[2]int{1, 0},
							[]c.Border{c.WALL, c.DOOR},
						},
						&gf.FakeField{
							[2]int{1, 1},
							[]c.Border{c.CROSSWALK, c.WALL},
						},
						&gf.FakeField{
							[2]int{1, 2},
							[]c.Border{c.CROSSWALK, c.WALL},
						},
						&gf.FakeField{
							[2]int{1, 3},
							[]c.Border{c.CROSSWALK, c.CROSSWALK},
						},
					},
				},
				[][]fo.FieldObject{
					[]fo.FieldObject{
						&fo.FakeFieldObject{Identifier: "a"},
					},
					[]fo.FieldObject{
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
						&fo.FakeFieldObject{Identifier: "z"},
					},
					[]fo.FieldObject{
						&fo.FakeFieldObject{Identifier: "y"},
					},
					[]fo.FieldObject{
						&fo.FakeFieldObject{Identifier: "z"},
					},
					[]fo.FieldObject{},
					[]fo.FieldObject{
						&fo.FakeFieldObject{Identifier: "b"},
					},
				},
				-1,
			},
		},
	}

	view := NewStringView()

	for _, tc := range testCases {
		alias := tc.alias
		expectedString := tc.expectedString
		board := tc.board

		fn := func(t *testing.T) {

			actualString := view.CompleteBoard(board)

			if !reflect.DeepEqual(actualString, expectedString) {
				t.Errorf("\r\nTestCase: %s failed.\r\n view.ToString(...)\r\n returned \r\n%+v\r\n while expected \r\n%v\r\n", alias, actualString, expectedString)
			}
		}

		runTest(alias, t, fn)
	}
}
