package parsel

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	_ "github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

var err error

type Selector struct {
	Text string
	err  []error

	XpathNode         *html.Node
	XpathStringResult []string
	XpathNodeResult   []html.Node
}

func (s *Selector) Help() {
	helpString := `
selector := Selector{}.Load(text string)
xpathObject := selector.Xpath("//xpath string")
string := xpathObject.Get()
[]string := xpathObject.GetAll()
	`
	fmt.Println(helpString)
}

func (s *Selector) Load(text string) {
	s.Text = text
	s.XpathNode, err = htmlquery.Parse(strings.NewReader(text))
	if err != nil {
		panic(err.Error())
	}
}
