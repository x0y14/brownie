package brownie

type BoxModel struct {
	position *Position
}

type Position struct {
	margin *Margin
}

type Margin struct {
	top    float64
	bottom float64
	left   float64
	right  float64

	border *Border
}

type Border struct {
	top    float64
	bottom float64
	left   float64
	right  float64

	padding *Padding
}

type Padding struct {
	top    float64
	bottom float64
	left   float64
	right  float64

	contents []*Content
}

type ContentType int

const (
	_ ContentType = iota
	Boxes
	PlainText
	H1
	H2
	H3
	H4
	H5
	H6
)

type Content struct {
	contentType ContentType

	text string

	boxModels []*BoxModel
}

func NewContentModel(contents []*Content) *BoxModel {
	return &BoxModel{
		position: &Position{
			margin: &Margin{
				border: &Border{
					padding: &Padding{
						contents: contents,
					},
				},
			},
		},
	}
}
