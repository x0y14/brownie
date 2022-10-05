package brownie

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/x0y14/dragon/html/parser"
	"github.com/x0y14/dragon/html/tokenizer"
	"log"
)

func OpenWebSite(url string) error {
	rawHtml, err := getHtmlFromUrl(url)
	if err != nil {
		return fmt.Errorf("failed to get html: %v", err)
	}

	tokenizer_ := tokenizer.NewTokenizer()
	tokens, err := tokenizer_.Tokenize([]rune(rawHtml))
	if err != nil {
		return fmt.Errorf("failed to tokenize: %v", err)
	}

	parser_ := parser.NewParser()
	_, err = parser_.Parse(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse: %v", err)
	}

	w := &Window{}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("brownie")
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}

	return nil
}
