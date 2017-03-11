package Stringview

import (
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
		board          Board
	}{
		{
			alias:          "One Row - One Field - Nothing Special",
			expectedString: "|---|\n|   |\n|   |\n|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]int{0, 0},
							[]Border{WALL, WALL, WALL, WALL},
						},
					},
				},
				[]FieldObject{},
			},
		},
		{
			alias:          "One Row - Two Fields seperated by door",
			expectedString: "|---|---|\n|   |   |\n|       |\n|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]int{0, 0},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]int{0, 1},
							[]Border{DOOR, WALL},
						},
					},
				},
				[]FieldObject{},
			},
		},
		{
			alias:          "Two Rows - Four Fields seperated by doors",
			expectedString: "|---|---|\n|   |   |\n|       |\n|- -|- -|\n|   |   |\n|       |\n|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]int{0, 0},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]int{0, 1},
							[]Border{DOOR, WALL},
						},
					},
					1: []Field{
						&FakeField{
							[]int{1, 0},
							[]Border{WALL, DOOR},
						},
						&FakeField{
							[]int{1, 1},
							[]Border{DOOR, DOOR},
						},
					},
				},
				[]FieldObject{},
			},
		},
		{
			alias:          "Two Rows - Four Fields - Three Rooms",
			expectedString: "|---|---|\n|   |   |\n|   |   |\n|   |- -|\n|   |   |\n|       |\n|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]int{0, 0},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]int{0, 1},
							[]Border{WALL, WALL},
						},
					},
					1: []Field{
						&FakeField{
							[]int{1, 0},
							[]Border{WALL, FREE},
						},
						&FakeField{
							[]int{1, 1},
							[]Border{DOOR, DOOR},
						},
					},
				},
				[]FieldObject{},
			},
		},
		{
			alias:          "Two Rows - Six Fields - Crosswalk",
			expectedString: "|---|---|---|\n|   |       |\n|   |       |\n| = |- -|---|\n|   =   =   |\n|   =   =   |\n|---|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]int{0, 0},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]int{0, 1},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]int{0, 2},
							[]Border{FREE, WALL},
						},
					},
					1: []Field{
						&FakeField{
							[]int{1, 0},
							[]Border{WALL, CROSSWALK},
						},
						&FakeField{
							[]int{1, 1},
							[]Border{CROSSWALK, DOOR},
						},
						&FakeField{
							[]int{1, 2},
							[]Border{CROSSWALK, WALL},
						},
					},
				},
				[]FieldObject{},
			},
		},
		{
			alias:          "One Row - Two Fields - Zombs and Survivors",
			expectedString: "|---|---|\n|ab |zzy|\n|       |\n|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]int{0, 0},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]int{0, 1},
							[]Border{DOOR, WALL},
						},
					},
				},
				[]FieldObject{
					&FakeFieldObject{"a", []int{0, 0}},
					&FakeFieldObject{"b", []int{0, 0}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"y", []int{0, 1}},
				},
			},
		},
		{
			alias:          "Big Map - Lots of Zombs",
			expectedString: "|---|---|---|---|\n|a  |zzz|y  |z  |\n|   |zz+        |\n|- -|---|---| = |\n|   =b  =   =   |\n|   =   =   =   |\n|---|---|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]int{0, 0},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]int{0, 1},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]int{0, 2},
							[]Border{DOOR, WALL},
						},
						&FakeField{
							[]int{0, 3},
							[]Border{DOOR, WALL},
						},
					},
					1: []Field{
						&FakeField{
							[]int{1, 0},
							[]Border{WALL, DOOR},
						},
						&FakeField{
							[]int{1, 1},
							[]Border{CROSSWALK, WALL},
						},
						&FakeField{
							[]int{1, 2},
							[]Border{CROSSWALK, WALL},
						},
						&FakeField{
							[]int{1, 3},
							[]Border{CROSSWALK, CROSSWALK},
						},
					},
				},
				[]FieldObject{
					&FakeFieldObject{"a", []int{0, 0}},
					&FakeFieldObject{"b", []int{1, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"z", []int{0, 1}},
					&FakeFieldObject{"y", []int{0, 2}},
					&FakeFieldObject{"z", []int{0, 3}},
				},
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
