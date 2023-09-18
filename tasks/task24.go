package main

import (
    "math"
    "fmt"
)

// создаем интерфейс для
// представления общего поведения всех точек
type Movable interface {
    GetCoords() (float64, float64)
}

// создаем тип
type Point struct {
    x float64
    y float64
}

// создаем метод, возвращающий координаты точки
func (p *Point) GetCoords() (float64, float64) {
    return p.x, p.y
}

func NewPoint(x, y float64) *Point {
    return &Point{
        x:          x,
        y:          y,
    }
}

// создаем функцию, котоая принимает две точки
// вычисляет расстояние между ними и возвращает результат
func Distance(a, b Movable) float64 {
    ax, ay := a.GetCoords()
    bx, by := b.GetCoords()
    l := math.Abs(ax - bx)
    h := math.Abs(ay - by)
    return math.Sqrt(math.Pow(l, 2) + math.Pow(h, 2))
}

func main() {
    p_a := NewPoint(2.0, 1.5)
    p_b := NewPoint(-3.5, 0.0)
    dist := Distance(p_a, p_b)
    fmt.Printf("Distance between %+v & %+v = %.2f\n", *p_a, *p_b, dist)
}
