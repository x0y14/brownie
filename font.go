package brownie

import (
	"bufio"
	"github.com/flopp/go-findfont"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"io"
	"log"
	"os"
)

var (
	DefaultFont *sfnt.Font
	DefaultFace *font.Face // 16px

	DefaultFaceH1 *font.Face // 32px
	DefaultFaceH2 *font.Face // 24px
	DefaultFaceH3 *font.Face // 18.72px
	DefaultFaceH4 *font.Face // 16px
	DefaultFaceH5 *font.Face // 13.28px
	DefaultFaceH6 *font.Face // 10.72px
)

func init() {
	// setup times-font
	fontPath, err := findfont.Find("Times New Roman.ttf")
	if err != nil {
		panic(err)
	}
	fontFile, err := os.Open(fontPath)
	if err != nil {
		log.Fatal(fontFile)
	}
	defer func(fontFile *os.File) {
		_ = fontFile.Close()
	}(fontFile)

	reader := bufio.NewReader(fontFile)
	fontData, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("failed to read font file: %v", err)
	}

	timesFont, err := sfnt.Parse(fontData)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	DefaultFont = timesFont

	const dpi = 72

	// h1
	timesH1, err := opentype.NewFace(timesFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to generate font face: %v", err)
	}
	DefaultFaceH1 = &timesH1

	// h2
	timesH2, err := opentype.NewFace(timesFont, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to generate font face: %v", err)
	}
	DefaultFaceH2 = &timesH2

	// h3
	timesH3, err := opentype.NewFace(timesFont, &opentype.FaceOptions{
		Size:    18.72,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to generate font face: %v", err)
	}
	DefaultFaceH3 = &timesH3

	// h4
	timesH4, err := opentype.NewFace(timesFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to generate font face: %v", err)
	}
	DefaultFaceH4 = &timesH4
	DefaultFace = &timesH4

	// h5
	timesH5, err := opentype.NewFace(timesFont, &opentype.FaceOptions{
		Size:    13.28,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to generate font face: %v", err)
	}
	DefaultFaceH5 = &timesH5

	// h6
	timesH6, err := opentype.NewFace(timesFont, &opentype.FaceOptions{
		Size:    10.72,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to generate font face: %v", err)
	}
	DefaultFaceH6 = &timesH6
}
