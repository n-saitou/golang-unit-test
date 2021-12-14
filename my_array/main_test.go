package my_array

import "testing"

func TestSum(t *testing.T) {

	type args struct {
		numbers []int
	}

	myArgs := args{numbers: []int{1, 2, 3, 4, 5}}
	myArgs2 := args{numbers: []int{}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "aa",
			args: myArgs,
			want: 15,
		}, {
			name: "bb",
			args: myArgs2,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
