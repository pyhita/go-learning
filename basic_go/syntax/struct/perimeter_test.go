package main

import "testing"

func TestPerimeter(t *testing.T) {
	assertMessage := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("test perimeter", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		assertMessage(t, got, want)
	})
}

func TestArea(t *testing.T) {

	t.Run("test area", func(t *testing.T) {
		got := Area(Rectangle{
			Width:  10.0,
			Height: 10.0,
		})
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

// 为不同的类型 定义不同的area 方法
func TestMethodArea(t *testing.T) {

	// 通过接口 实现进一步的封装
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		checkArea(t, &rectangle, 100.0)
	})

	t.Run("cycle area", func(t *testing.T) {
		circle := Circle{
			Radius: 10.0,
		}
		checkArea(t, &circle, 314.1592653589793)
	})
}

// 表格驱动测试
func TestArea2(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			"rectangle", &Rectangle{12.0, 6.0}, 72.0,
		},
		{
			"circle", &Circle{110.0}, 314.1592653589793,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}
