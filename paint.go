package brownie

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

func (w *Window) paint(screen *ebiten.Image, model *BoxModel, x, y float64) error {
	paintX := x
	paintY := y
	if model.position.margin.left != 0 {
		paintX += model.position.margin.left
	}
	if model.position.margin.top != 0 {
		paintY += model.position.margin.top
	}

	for _, content := range model.position.margin.border.padding.contents {
		switch content.contentType {
		case Boxes:
			for _, b := range content.boxModels {
				if err := w.paint(screen, b, paintX, paintY); err != nil {
					panic(err)
				}
			}
		case PlainText:
			b := text.BoundString(*DefaultFace, content.text)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(paintX, paintY)
			op.ColorM.Reset()
			//op.GeoM.Translate(400, 200)
			op.ColorM.ScaleWithColor(color.White)

			text.DrawWithOptions(screen, content.text, *DefaultFace, op)
			paintX += float64(b.Max.X)
			paintY += float64(b.Max.Y)
		}
	}
	return nil
}
