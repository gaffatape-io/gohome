package gohome

import "testing"

type testType struct {
}

type testTypeIdentifiable struct {
}

func (t *testTypeIdentifiable) ID() string {
	return "testTypeIdentifiable"
}

func TestID(t *testing.T) {
	tests := []struct {
		d  interface{}
		id string
	}{
		{&testType{}, "testType"},
		{testType{}, "testType"},
		{&testTypeIdentifiable{}, "testTypeIdentifiable"},
		{testTypeIdentifiable{}, "testTypeIdentifiable"},
	}

	for _, tc := range tests {
		id := ID(tc.d)
		t.Log(id, tc.id)
		if id != tc.id {
			t.Fatal()
		}
	}
}
