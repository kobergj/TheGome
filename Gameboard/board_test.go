package Gameboard

import (
	"reflect"
	"testing"

	gf "../Gamefields"
	fo "../Gameobjects"
)

// Test helper to wrap up test invokation and separate it from test cases definition
// Also, since this IS golang 1.7, this wrapper will be very usefull  with test filtering
func runTest(testAlias string, t *testing.T, testFunction func(t *testing.T)) {
	// For golang 1.7+
	t.Run(testAlias, testFunction)
}

func TestBoardObjects(t *testing.T) {
	testCases := []struct {
		alias            string
		rows             int
		cols             int
		field            gf.Field
		availableObjects []fo.FieldObject
		expectedObjects  []fo.FieldObject
	}{
		{
			"Nothing to show",
			1,
			1,
			&gf.FakeField{Cords: [2]int{0, 0}},
			[]fo.FieldObject{},
			[]fo.FieldObject{},
		},
		{
			"One Object, invisible",
			1,
			1,
			&gf.FakeField{Cords: [2]int{0, 0}},
			[]fo.FieldObject{
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{0, 0}},
				},
			},
			[]fo.FieldObject{},
		},
		{
			"One Object, visible",
			1,
			1,
			&gf.FakeField{Cords: [2]int{0, 0}},
			[]fo.FieldObject{
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{0, 0}},
					IsAlive: true,
				},
			},
			[]fo.FieldObject{
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{0, 0}},
					IsAlive: true,
				},
			},
		},
		{
			"Extended, with dead Objects",
			3,
			5,
			&gf.FakeField{Cords: [2]int{2, 1}},
			[]fo.FieldObject{
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{0, 0}},
					IsAlive: true,
				},
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{2, 1}},
					IsAlive: true,
				},
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{2, 1}},
					IsAlive: false,
				},
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{2, 1}},
					IsAlive: true,
				},
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{3, 2}},
					IsAlive: true,
				},
			},
			[]fo.FieldObject{
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{2, 1}},
					IsAlive: true,
				},
				&fo.FakeFieldObject{
					Current: &gf.FakeField{Cords: [2]int{2, 1}},
					IsAlive: true,
				},
			},
		},
	}

	for _, tc := range testCases {

		alias := tc.alias
		rows := tc.rows
		cols := tc.cols
		field := tc.field
		availableObjects := tc.availableObjects
		expectedObjects := tc.expectedObjects

		board := NewBoard(rows, cols)
		for _, obj := range availableObjects {
			board.AddFieldObject(obj)
		}

		fn := func(t *testing.T) {

			res := board.BoardObjectsByField(field)

			if !reflect.DeepEqual(res, expectedObjects) {
				t.Errorf("\r\nTestCase: %s failed.\r\n board.BoardObjectsByField(...)\r\n returned \r\n%+v\r\n while expected \r\n%v\r\n", alias, res, expectedObjects)
			}
		}

		runTest(alias, t, fn)
	}
}

func TestBoardFields(t *testing.T) {
	testCases := []struct {
		alias          string
		rows           int
		cols           int
		allFields      []gf.Field
		rowToCheck     int
		expectedOk     bool
		expectedFields []gf.Field
	}{
		{
			"One Row simple check",
			1,
			1,
			[]gf.Field{
				&gf.FakeField{Cords: [2]int{-1, -1}},
			},
			0,
			true,
			[]gf.Field{
				&gf.FakeField{Cords: [2]int{0, 0}},
			},
		},
		{
			"More Rows simple check",
			2,
			4,
			[]gf.Field{
				&gf.FakeField{Cords: [2]int{-1, -1}},
				&gf.FakeField{Cords: [2]int{-1, -1}},
				&gf.FakeField{Cords: [2]int{-1, -1}},
				&gf.FakeField{Cords: [2]int{-1, -1}},
				&gf.FakeField{Cords: [2]int{-1, -1}},
				&gf.FakeField{Cords: [2]int{-1, -1}},
			},
			1,
			true,
			[]gf.Field{
				&gf.FakeField{Cords: [2]int{1, 0}},
				&gf.FakeField{Cords: [2]int{1, 1}},
			},
		},
		{
			"Check to big row",
			1,
			1,
			[]gf.Field{
				&gf.FakeField{Cords: [2]int{-1, -1}},
			},
			3,
			false,
			nil,
		},
	}

	for _, tc := range testCases {

		alias := tc.alias
		rows := tc.rows
		cols := tc.cols
		allFields := tc.allFields
		rowToCheck := tc.rowToCheck
		expectedOk := tc.expectedOk
		expectedFields := tc.expectedFields

		board := NewBoard(rows, cols)
		for _, field := range allFields {
			board.AddField(field)
		}

		fn := func(t *testing.T) {

			actualFields, actualOk := board.FieldsByRow(rowToCheck)
			if !reflect.DeepEqual(actualOk, expectedOk) {
				t.Errorf("\r\nTestCase: %s failed.\r\n board.FieldsByRow(...)\r\n returned \r\n%+v\r\n while expected \r\n%v\r\n", alias, actualOk, expectedOk)
			}
			if !reflect.DeepEqual(actualFields, expectedFields) {
				t.Errorf("\r\nTestCase: %s failed.\r\n board.FieldsByRow(...)\r\n returned \r\n%+v\r\n while expected \r\n%v\r\n", alias, actualFields, expectedFields)
			}
			for i, af := range actualFields {
				if !reflect.DeepEqual(af, expectedFields[i]) {
					t.Errorf("\r\nTestCase: %s failed.\r\n board.FieldsByRow(...)\r\n returned \r\n%+v\r\n while expected \r\n%v\r\n", alias, af, expectedFields[i])
				}
			}
		}

		runTest(alias, t, fn)
	}
}
