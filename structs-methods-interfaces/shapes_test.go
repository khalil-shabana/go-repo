package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("wanted %.2f, but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("wanted %g but got %g", want, got)
		}
	}
	t.Run("Rectangles Area", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		// got := Area(rect)

		/* got  := rect.Area()
		want := 100.0 */

		/*if got != want {
			t.Errorf("wanted %.2f, but got %.2f", want, got)
		}*/
		checkArea(t, rect, 100)
	})

	t.Run("Circles Area", func(t *testing.T) {
		circle := Circle{10}
		// got := Area(circle)
		/*got  := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("wanted %g, but got %g", want, got)

		}*/
		checkArea(t, circle, 314.1592653589793)
	})
}

/* Table driven tests */

func TestArea2(t *testing.T) {
	/* We are declaring a slice of structs by using []struct with two fields, the shape and the want.
	   Then we fill the slice with cases.*/
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("wanted %g, but got %g", tt.want, got)
		}
	}
}

func TestArea3(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 4}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v wanted %g but got %g ", tt.shape, tt.hasArea, got)
			}
		})
	}
}
