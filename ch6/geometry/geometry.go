package main

import (
  "fmt"
  "math"
)

type point struct {
  x, y float64
}

func (p *point) distance(q *point) float64  {
  xdelta := p.x - q.x
  ydelta := p.y - q.y
  return math.Sqrt(xdelta * xdelta + ydelta * ydelta)
}

type path []point

func (p path) length() (sum float64) {
  for i := range p {
    if i > 0 {
      sum += (&p[i-1]).distance(&p[i])
    }
  }
  return
}

func main()  {
  p := point{0, 3}
  q := point{4, 0}
  fmt.Println((&p).distance(&q))

  perim := path{
    {1, 1},
    {5, 1},
    {5, 4},
    {1, 1},
  }
  fmt.Println(perim.length())
}
