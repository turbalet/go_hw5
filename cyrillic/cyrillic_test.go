package cyrillic

import (
	"reflect"
	"testing"
)

type SubUser struct {
	ChildName string
}

type User struct {
	Id   int64
	Name string
	Sub  SubUser
	Sptr *string
}

func TestListFields(t *testing.T) {
	s1 := "покакириллицаBye"
	s2 := "ПhРeИlВlЕoТ"
	s1Res := "Bye"
	s2Res := "hello"
	testTable := []struct {
		data     User
		expected User
	}{
		{
			data:     User{1, "ЮзернеймUsername", SubUser{"СабЮзерНеймSubUserName"}, &s1},
			expected: User{1, "Username", SubUser{"SubUserName"}, &s1Res},
		},
		{
			data:     User{2, "пHрEиLвLеOт", SubUser{"ПhРeИlВlЕoТ"}, &s2},
			expected: User{2, "HELLO", SubUser{"hello"}, &s2Res},
		},
	}

	for _, testCase := range testTable {
		ListFields(&testCase.data)
		if !reflect.DeepEqual(testCase.data, testCase.expected) {

			t.Errorf("Incorrect result. Expect %v got %v", testCase.expected, testCase.data)
		}
	}
}
