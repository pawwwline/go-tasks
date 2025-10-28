package main

import (
	"reflect"
	"testing"
)

func TestSearchAngs(t *testing.T) {
	tests := []struct {
		input []string
		want  map[string][]string
	}{
		{[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}, map[string][]string{
			"листок": []string{"листок", "слиток", "столик"}, "пятак": []string{"пятак", "пятка", "тяпка"},
		}}, {[]string{"яблоко", "банан", "вишня"}, map[string][]string{}},
		{[]string{"кот", "ток", "окт", "дом"}, map[string][]string{
			"кот": {"кот", "окт", "ток"},
		}},
	}
	for _, tt := range tests {
		got := searchAnagrams(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("searchAnagrams(%v) = %v, want %v", tt.input, got, tt.want)
		}

	}
}
