package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -5, 0},
		{"один элемент", 1, 1},
		{"десять элементов", 10, 10},
		{"тысяча элементов", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Errorf("generateRandomElements(%d) вернула слайс длины %d, ожидалось %d",
					tt.size, len(got), tt.want)
			}

			for i, v := range got {
				if v <= 0 {
					t.Errorf("generateRandomElements(%d): элемент [%d] = %d, ожидалось положительное число",
						tt.size, i, v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"nil слайс", nil, 0},
		{"один элемент", []int{42}, 42},
		{"два элемента (возрастание)", []int{1, 2}, 2},
		{"два элемента (убывание)", []int{2, 1}, 2},
		{"все одинаковые", []int{5, 5, 5, 5}, 5},
		{"максимум в начале", []int{100, 1, 2, 3}, 100},
		{"максимум в конце", []int{1, 2, 3, 100}, 100},
		{"максимум в середине", []int{1, 50, 2, 3}, 50},
		{"отрицательные числа", []int{-5, -1, -10}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"nil слайс", nil, 0},
		{"один элемент", []int{7}, 7},
		{"меньше CHUNKS", []int{3, 1, 2}, 3},
		{"ровно CHUNKS", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"больше CHUNKS", []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 100},
		{"все одинаковые", []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksVsMaximum(t *testing.T) {
	data := generateRandomElements(10_000)

	got := maxChunks(data)
	want := maximum(data)

	if got != want {
		t.Errorf("maxChunks = %d, maximum = %d, результаты должны совпадать", got, want)
	}
}
