package brownie

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func OpenWebSite(url string) error {
	//wp, err := GetWebPageFromUrl(url)
	wp, err := GetWebPageFromPath("/Users/x0y14/dev/brownie/test.html")
	if err != nil {
		return fmt.Errorf("failed to get webpage: %v", err)
	}

	w := &Window{webpage: wp}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("brownie")
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}

	return nil
}
