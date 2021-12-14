package my_struct

import "testing"

func TestRectangle_Perimeter(t1 *testing.T) {
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "1",
			fields: fields{10, 20},
			want:   60,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := t.Perimeter(); got != tt.want {
				t1.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t1 *testing.T) {
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "2",
			fields: fields{10, 20},
			want:   200,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := t.Area(); got != tt.want {
				t1.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	type fields struct {
		Radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			if got := c.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
