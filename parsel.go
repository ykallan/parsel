package parsel

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var err error

type Selector struct {
	Text string
	err  []error

	linkFlag          bool
	XpathNode         *html.Node
	XpathStringResult []string
	XpathNodeResult   []*html.Node
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
		//panic(err.Error())
		fmt.Println("load text catch err = ", err.Error())
	}
}

func (s *Selector) Xpath(xpath string) *Selector {
	fmt.Println("s.linkFlag =", s.linkFlag)
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

func (s *Selector) Re(reString string) *Selector {
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

func (s *Selector) Get() string {
	return s.XpathNode.Data
}

func (s *Selector) GetAll() []string {
	newStringSlice := make([]string, 0)
	for _, node := range s.XpathNodeResult {
		//fmt.Println(node.Data)
		newStringSlice = append(newStringSlice, node.Data)
	}

	return newStringSlice
}

func (s *Selector) GetText() string {
	return s.Text
}
