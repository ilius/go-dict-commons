package search_utils

import (
	"reflect"
	"sort"
	"testing"
	"time"

	commons "github.com/ilius/go-dict-commons"
)

var timeout = 10 * time.Millisecond

func TestRunWorkers1(t *testing.T) {
	results := RunWorkers(
		10,
		1,
		timeout,
		func(i1 int, i2 int) []*commons.SearchResultLow {
			return []*commons.SearchResultLow{
				{F_Terms: []string{"a"}, F_Score: uint8(i1)},
			}
		},
	)
	sort.Slice(results, func(i, j int) bool {
		return results[i].F_Score < results[j].F_Score
	})
	expectedResults := []*commons.SearchResultLow{
		{F_Terms: []string{"a"}, F_Score: uint8(0)},
	}
	if !reflect.DeepEqual(results, expectedResults) {
		t.Errorf("results = %#v", results)
		t.Fail()
	}
}

func TestRunWorkers2(t *testing.T) {
	results := RunWorkers(
		3,
		2,
		timeout,
		func(i1 int, i2 int) []*commons.SearchResultLow {
			return []*commons.SearchResultLow{
				{F_Terms: []string{"a"}, F_Score: uint8(i1)},
			}
		},
	)
	sort.Slice(results, func(i, j int) bool {
		return results[i].F_Score < results[j].F_Score
	})
	expectedResults := []*commons.SearchResultLow{
		{F_Terms: []string{"a"}, F_Score: uint8(0)},
	}
	if !reflect.DeepEqual(results, expectedResults) {
		t.Errorf("results = %#v", results)
		t.Fail()
	}
}

func TestRunWorkers3(t *testing.T) {
	results := RunWorkers(
		10,
		2,
		timeout,
		func(i1 int, i2 int) []*commons.SearchResultLow {
			return []*commons.SearchResultLow{
				{F_Terms: []string{"a"}, F_Score: uint8(i1)},
			}
		},
	)
	sort.Slice(results, func(i, j int) bool {
		return results[i].F_Score < results[j].F_Score
	})
	expectedResults := []*commons.SearchResultLow{
		{F_Terms: []string{"a"}, F_Score: uint8(0)},
		{F_Terms: []string{"a"}, F_Score: uint8(5)},
	}
	if !reflect.DeepEqual(results, expectedResults) {
		t.Errorf("results = %#v", results)
		t.Fail()
	}
}

func TestRunWorkersTimeout(t *testing.T) {
	results := RunWorkers(
		10,
		2,
		time.Millisecond,
		func(i1 int, i2 int) []*commons.SearchResultLow {
			time.Sleep(2 * time.Millisecond)
			return []*commons.SearchResultLow{
				{F_Terms: []string{"a"}, F_Score: uint8(i1)},
			}
		},
	)
	if len(results) > 0 {
		t.Errorf("results = %#v", results)
		t.Fail()
	}
}
