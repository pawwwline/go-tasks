package main

import "testing"

func TestCheckUnique(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},       // все уникальны
		{"abCdefAaf", false}, // повторяются 'a'/'A'
		{"aabcd", false},     // повторение в начале
		{"", true},           // пустая строка
		{"a", true},          // один символ
		{"abcABC", false},    // с учётом приведения к lower
		{"абвг", true},       // юникод кириллица
		{"абва", false},      // юникод повтор
		{"🙂🙃😉", true},        // эмодзи уникальные
		{"🙂🙃🙂", false},       // эмодзи повторяются
	}

	for _, tt := range tests {
		got := checkUnique(tt.input)
		if got != tt.expected {
			t.Errorf("checkUnique(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}
