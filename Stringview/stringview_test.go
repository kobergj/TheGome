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
			expectedString: "|---|\n|   |\n|   |\n|   |\n|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{&FakeField{
						[]FieldObject{},
						[]FieldObject{},
						[]int{1, 1},
						[]Border{WALL, WALL, WALL, WALL},
					},
					},
				},
			},
		},
		{
			alias:          "One Row - Two Fields seperated by door",
			expectedString: "|---|---|\n|   |   |\n|       |\n|   |   |\n|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 1},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 2},
							[]Border{DOOR, WALL},
						},
					},
				},
			},
		},
		{
			alias:          "Two Rows - Four Fields seperated by doors",
			expectedString: "|---|---|\n|   |   |\n|       |\n|   |   |\n|- -|- -|\n|   |   |\n|       |\n|   |   |\n|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 1},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 2},
							[]Border{DOOR, WALL},
						},
					},
					1: []Field{
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{2, 1},
							[]Border{WALL, DOOR},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{2, 2},
							[]Border{DOOR, DOOR},
						},
					},
				},
			},
		},
		{
			alias:          "Two Rows - Four Fields - Three Rooms",
			expectedString: "|---|---|\n|   |   |\n|   |   |\n|   |   |\n|   |- -|\n|   |   |\n|       |\n|   |   |\n|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 1},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 2},
							[]Border{WALL, WALL},
						},
					},
					1: []Field{
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{2, 1},
							[]Border{WALL, FREE},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{2, 2},
							[]Border{DOOR, DOOR},
						},
					},
				},
			},
		},
		{
			alias:          "Two Rows - Six Fields - Crosswalk",
			expectedString: "|---|---|---|\n|   |       |\n|   |       |\n|   |       |\n| = |- -|---|\n|   =   =   |\n|   =   =   |\n|   =   =   |\n|---|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 1},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 2},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{1, 3},
							[]Border{FREE, WALL},
						},
					},
					1: []Field{
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{2, 1},
							[]Border{WALL, CROSSWALK},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{2, 2},
							[]Border{CROSSWALK, DOOR},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{},
							[]int{2, 3},
							[]Border{CROSSWALK, WALL},
						},
					},
				},
			},
		},
		{
			alias:          "One Row - Two Fields - Zombs and Survivors",
			expectedString: "|---|---|\n|ab |zzZ|\n|       |\n|   |   |\n|---|---|",
			board: &FakeBoard{
				map[int][]Field{
					0: []Field{
						&FakeField{
							[]FieldObject{
								&FakeFieldObject{"a"},
								&FakeFieldObject{"b"},
							},
							[]FieldObject{},
							[]int{1, 1},
							[]Border{WALL, WALL},
						},
						&FakeField{
							[]FieldObject{},
							[]FieldObject{
								&FakeFieldObject{"z"},
								&FakeFieldObject{"z"},
								&FakeFieldObject{"Z"},
							},
							[]int{1, 2},
							[]Border{DOOR, WALL},
						},
					},
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
