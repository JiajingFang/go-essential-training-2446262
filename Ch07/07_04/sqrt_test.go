package sqrt

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMany(t *testing.T) {
	// Read test cases from "sqrt_case.csv" and check them
	testcases := ReadCSVfile("sqrt_cases.csv")
	tcases := make([]testingcase, len(testcases))
	for i, testcase := range testcases {
		tcases[i].value, _ = strconv.ParseFloat(testcase[0], 64)
		tcases[i].expected, _ = strconv.ParseFloat(testcase[1], 64)
	}

	for _, tc := range tcases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			val, err := Sqrt(tc.value)
			require.NoError(t, err)
			require.InDelta(t, tc.expected, val, 0.001)
		})
	}

}

func ReadCSVfile(url string) [][]string {
	fd, err := os.Open(url)

	if err != nil {
		fmt.Println(err)
	}

	fileReader := csv.NewReader(fd)
	records, err := fileReader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()
	return records
}

type testingcase struct {
	value    float64
	expected float64
}

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}
