package test

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"

	"github.com/bluewon/testing/utils"
	"github.com/stretchr/testify/require"
)

// Simple test struct
func TestSquare(t *testing.T) {
	table := []struct {
		TestCase string //tên test case
		Request  int    //Object chứa tất cả parameter của func cần test
		Response int    //Object chứa response của func cần test
	}{
		{TestCase: "with 0", Request: 0, Response: 0},
		{TestCase: "with 1", Request: 1, Response: 1},
	}

	for _, row := range table {
		t.Run(row.TestCase, func(t *testing.T) {
			n := utils.Square(row.Request)

			//Compare result
			require.Equalf(t, row.Response, n, "\nExpect response %d but func `Square` response %d", row.Response, n)
		})
	}
}

func TestSquareWithCSV(t *testing.T) {
	csvFile, err := os.Open("data.csv")
	if err != nil {
		t.Errorf("open file error")
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		t.Errorf("read file error")
	}
	table := []struct {
		TestCase string //tên test case
		Request  int    //Object chứa tất cả parameter của func cần test
		Response int    //Object chứa response của func cần test
	}{}
	for _, line := range csvLines {
		a, err := strconv.Atoi(line[0])
		b, errResp := strconv.Atoi(line[1])
		if err != nil || errResp != nil {
			t.Errorf("parse error")
		}
		table = append(table, struct {
			TestCase string
			Request  int
			Response int
		}{
			TestCase: "square with " + line[0] + " and result " + line[1],
			Request:  a,
			Response: b,
		})
	}

	for _, row := range table {
		t.Run(row.TestCase, func(t *testing.T) {
			n := utils.Square(row.Request)

			//Compare result
			require.Equalf(t, row.Response, n, "\nExpect response %d but func `Square` response %d", row.Response, n)
		})
	}
}
