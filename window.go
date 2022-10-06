package brownie

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Window struct {
	webpage *WebPage
	config  map[string]string
}

func (w *Window) Update() error {
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {

	if err := w.webpage.calculate(); err != nil {
		log.Fatalf("failed to calculate: %v", err)
	}

	err := w.paint(screen, w.webpage.Layout, 0, 0)
	if err != nil {
		log.Fatalf("failed to paint web page: %v", err)
	}

	//ebitenutil.DebugPrint(screen, fmt.Sprintf("fps: %v", ebiten.ActualFPS()))
	//if err != nil {
	//	log.Fatalf("faield to print fps: %v", err)
	//}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
