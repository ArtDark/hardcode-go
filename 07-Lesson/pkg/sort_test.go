package sort_test

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestInts(t *testing.T) {
	data := []int{2, 4, 6, 3, 7, 2}
	want := []int{2, 2, 3, 4, 6, 7}
	sort.Ints(data)
	if !reflect.DeepEqual(data, want) {
		t.Errorf("got %v, want %v", data, want)
	}
}

func TestStrings(t *testing.T) {
	type args struct {
		x []string
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "nums test",
			args: args{
				x: []string{"1", "4", "2", "3"},
			},
			want: args{
				x: []string{"1", "2", "3", "4"},
			},
		},
		{
			name: "nil test",
			args: args{
				x: nil,
			},
			want: args{
				x: nil,
			},
		},
		{
			name: "abc test",
			args: args{
				x: []string{"b", "c", "a", "d"},
			},
			want: args{
				x: []string{"a", "b", "c", "d"},
			},
		},
		{
			name: "Words test",
			args: args{
				x: []string{"Alice", "Carol", "Bob", "Mallory"},
			},
			want: args{
				x: []string{"Alice", "Bob", "Carol", "Mallory"},
			},
		},
		{
			name: "Words & nums test",
			args: args{
				x: []string{"Alice", "1", "Bob", "2"},
			},
			want: args{
				x: []string{"1", "2", "Alice", "Bob"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.args.x)
			if !reflect.DeepEqual(tt.args.x, tt.want.x) {
				t.Errorf("got %v, want %v", tt.args.x, tt.want.x)
			}
		})
	}
}

func BenchmarkInts(b *testing.B) {
	n := rand.Perm(2_000_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(n)
	}
}

func BenchmarkFloat64(b *testing.B) {
	n := rand.Perm(2_000_000)
	f := func(ar []int) []float64 {
		newArr := make([]float64, len(ar))
		for i, v := range ar {
			newArr[i] = float64(v)
		}
		return newArr
	}(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Float64s(f)
	}
}
