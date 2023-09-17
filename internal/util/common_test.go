package util

import "testing"

func TestGetOrderByString(t *testing.T) {
	testCases := []struct {
		param string
		want  string
	}{
		{param: "+col1", want: "col1 ASC"},
		{param: "-col1", want: "col1 DESC"},
		{param: "+col1,-col2", want: "col1 ASC,col2 DESC"},
	}

	for _, testCase := range testCases {
		result := GetOrderByFromString(testCase.param)
		if result != testCase.want {
			t.Errorf("get unexpected result - Expected: %s, Got: %s", testCase.want, result)
		}
	}
}
