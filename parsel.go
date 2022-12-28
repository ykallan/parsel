package parsel

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var err error

type Select struct {
	Text string
	err  []error

	linkFlag          bool
	XpathNode         *html.Node
	XpathStringResult []string
	XpathNodeResult   []*html.Node
}

func Selector(Text string) Select {
	s := Select{}
	s.Load(Text)
	return s
}

func (s *Select) Help() {
	helpString := `
selector := Selector{Text:text}
xpathObject := selector.Xpath("//xpath string")
string := xpathObject.Get()
[]string := xpathObject.GetAll()
	`
	fmt.Println(helpString)
}

func (s *Select) Load(text string) {
	s.Text = text
	s.XpathNode, err = htmlquery.Parse(strings.NewReader(text))
	if err != nil {
		//panic(err.Error())
		fmt.Println("load text catch err = ", err.Error())
	}
}

func (s *Select) Xpath(xpath string) *Select {
	// fmt.Println("s.linkFlag =", s.linkFlag)
	if s.linkFlag {
		// already get A xpath node, can use link extract
		newXpathNode := htmlquery.Find(s.XpathNode, xpath)
		if len(newXpathNode) > 0 {
			s.XpathNode = newXpathNode[0]
			return s
		} else {
			s.XpathNode = nil
			s.linkFlag = false
		}
	}

	s.XpathNodeResult = htmlquery.Find(s.XpathNode, xpath)
	// if got xpath node more than one, save the first one
	if len(s.XpathNodeResult) > 0 {
		s.XpathNode = s.XpathNodeResult[0]
		s.linkFlag = true

		return s
	}
	return s
}

func (s *Select) Re(reString string) *Select {
	s.XpathStringResult = make([]string, 0)
	compiled := regexp.MustCompile(reString)
	match := compiled.FindAllStringSubmatch(s.Text, -1)
	for _, first := range match {
		for _, second := range first {
			s.XpathStringResult = append(s.XpathStringResult, second)
		}
	}
	return s
}

func (s *Select) Get() (string, bool) {
	return s.XpathNode.Data, len(s.XpathNode.Data) > 0
}

func (s *Select) GetAll() ([]string, bool) {
	newStringSlice := make([]string, 0)
	for _, node := range s.XpathNodeResult {
		//fmt.Println(node.Data)
		newStringSlice = append(newStringSlice, node.Data)
	}

	return newStringSlice, len(newStringSlice) > 0
}

func (s *Select) GetText() string {
	return s.Text
}
