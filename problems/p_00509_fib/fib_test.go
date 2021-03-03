package p_00509_fib

import "testing"

func Test_fib__O__memo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{2},
			want: 1,
		},
		{
			name: "case2",
			args: args{3},
			want: 2,
		},
		{
			name: "case3",
			args: args{4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib__On__memo(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib__O2n(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{2},
			want: 1,
		},
		{
			name: "case2",
			args: args{3},
			want: 2,
		},
		{
			name: "case3",
			args: args{4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib__O2n(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fib(b *testing.B) {
	benchmarks := []struct {
		name string
		f    func(n int) int
	}{
		{name: "fib_rescur", f: fib__O2n},
		{name: "fib_memo", f: fib__On__memo},
		{name: "fib_dp", f: fib__On__dp},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.f(30)
			}
		})
	}
}
