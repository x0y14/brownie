package brownie

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type Window struct {
}

func (w *Window) Update(screen *ebiten.Image) error {
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	err := ebitenutil.DebugPrint(screen, fmt.Sprintf("fps: %v", ebiten.CurrentFPS()))
	if err != nil {
		log.Fatalf("faield to print fps: %v", err)
	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
