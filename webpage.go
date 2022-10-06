package brownie

import (
	"fmt"
	commonParser "github.com/x0y14/dragon/common/parser"
	"github.com/x0y14/dragon/html/parser"
	"github.com/x0y14/dragon/html/tokenizer"
	"os"
)

type WebPage struct {
	Url    string
	Nodes  []*commonParser.Node
	Layout *BoxModel
	Config map[string]string
}

func GetWebPageFromUrl(url string) (*WebPage, error) {
	rawHtml, err := getHtmlFromUrl(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get html: %v", err)
	}

	tokenizer_ := tokenizer.NewTokenizer()
	tokens, err := tokenizer_.Tokenize([]rune(rawHtml))
	if err != nil {
		return nil, fmt.Errorf("failed to tokenize: %v", err)
	}

	parser_ := parser.NewParser()
	nodes, err := parser_.Parse(tokens)
	if err != nil {
		return nil, fmt.Errorf("failed to parse: %v", err)
	}

	return &WebPage{Url: url, Nodes: nodes}, nil
}

func GetWebPageFromPath(path string) (*WebPage, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	rawHtml := string(bytes)

	tokenizer_ := tokenizer.NewTokenizer()
	tokens, err := tokenizer_.Tokenize([]rune(rawHtml))
	if err != nil {
		return nil, fmt.Errorf("failed to tokenize: %v", err)
	}

	parser_ := parser.NewParser()
	nodes, err := parser_.Parse(tokens)
	if err != nil {
		return nil, fmt.Errorf("failed to parse: %v", err)
	}

	return &WebPage{Url: path, Nodes: nodes}, nil
}

func nodeToModel(node *commonParser.Node) (*BoxModel, error) {
	switch node.Kind {
	case parser.Tag:
		switch node.N1.S {
		//case "body":
		//	for _, child := range node.Children {
		//
		//	}
		}
	case parser.Text:
		return NewContentModel([]*Content{
			{contentType: PlainText, text: node.S},
		}), nil
	}
	return nil, nil
}

func (wp *WebPage) calculate() error {

	var boxModels []*BoxModel

	for _, node := range wp.Nodes {
		switch node.Kind {
		case parser.Tag:
			switch node.N1.S {
			case "html":
				for _, child := range node.Children {
					bm, err := nodeToModel(child)
					if err != nil {
						panic(err)
					}
					boxModels = append(boxModels, bm)
				}
			default:
			}
		}
	}

	wp.Layout = &BoxModel{
		position: &Position{
			margin: &Margin{
				top:    8,
				bottom: 8,
				left:   8,
				right:  8,
				border: &Border{
					padding: &Padding{
						contents: []*Content{
							{contentType: Boxes, boxModels: boxModels},
						}}}}},
	}
	return nil
}
