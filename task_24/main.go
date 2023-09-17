package main

import (
    "math"
    "fmt"
)

type Movable interface {
    GetCoords() (float64, float64)
    MovX(dist float64)
    MovY(dist float64)
}

type Point struct {
    x float64
    y float64
}

func (p *Point) GetCoords() (float64, float64) {
    return p.x, p.y
}

func (p *Point) MovX(dist float64) {
    p.x += dist
}

func (p *Point) MovY(dist float64) {
    p.y += dist
}

func NewPoint(x, y float64) *Point {
    return &Point{
        x:          x,
        y:          y,
    }
}

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
    p_a.MovX(-16.7)
    p_b.MovY(3.4)
    dist = Distance(p_a, p_b)
    fmt.Printf("Distance between %+v & %+v = %.2f\n", *p_a, *p_b, dist)
    p_a.MovX(6.7)
    p_b.MovY(8.4)
    p_b.MovX(-8.4)
    dist = Distance(p_a, p_b)
    fmt.Printf("Distance between %+v & %+v = %.2f\n", *p_a, *p_b, dist)
}
