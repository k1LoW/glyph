package glyph

import (
	"fmt"
	"strconv"
)

type Point struct {
	X float64
	Y float64
}

type Points map[string]*Point

func (p Points) Get(key string) (*Point, error) {
	v, ok := p[key]
	if !ok {
		return nil, fmt.Errorf("invalid key: %s", key)
	}
	return v, nil
}

const dx = 8.660254
const dy = 5.0

var cPoints = Points{}

func GetPoints() Points {
	if len(cPoints) > 0 {
		return cPoints
	}
	points := Points{}
	px := 0xf
	py := 0x0
	f0x := 55.0
	f0y := 5.0
	maxy := 10

	// f,e,d,c,b,a
	for i := 0; i <= 5; i++ {
		max := maxy - i
		for j := 0; j <= max; j++ {
			key := fmt.Sprintf("%s%x", strconv.FormatInt(int64(px-i), 21), py+j)
			points[key] = &Point{
				X: f0x - float64(i)*dx,
				Y: f0y + float64(i)*dy + float64(j)*dy*2,
			}
		}
	}

	// g,h,i,j,k
	for i := 1; i <= 5; i++ {
		max := maxy - i
		for j := 0; j <= max; j++ {
			key := fmt.Sprintf("%s%x", strconv.FormatInt(int64(px+i), 21), py+j)
			points[key] = &Point{
				X: f0x + float64(i)*dx,
				Y: f0y + float64(i)*dy + float64(j)*dy*2,
			}
		}
	}
	cPoints = points

	return points
}
