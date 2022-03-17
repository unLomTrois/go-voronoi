package geometry

import (
	"image/color"
	"template/internal/app/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Point struct {
	Pos   pixel.Vec
	Color color.Color
}

func NewPoint(bounds pixel.Rect, pos pixel.Vec) *Point {
	return &Point{Pos: pos, Color: colornames.White}
}

func (p *Point) Draw(imd *imdraw.IMDraw) {
	utils.DrawCircle(imd, p.Pos, 3, p.Color, 0)
}
