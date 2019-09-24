package scrape

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/javscrape/go-scrape/query"
)

const bp4xJavURL = "https://www.bp4x.com/?q=%s"
const bp4xAmateurURL = "https://www.bp4x.com/?c=amateur&q=%s"
const bp4xIVURL = "https://www.bp4x.com/?c=iv&q=%s"
const bp4xHentaiURL = "https://www.bp4x.com/?c=hentai&q=%s"

// GrabBP4XType ...
type GrabBP4XType int

// BP4XTypeJAV ...
const (
	BP4XTypeJAV GrabBP4XType = iota
	BP4XTypeAMATEUR
	BP4XTypeIV
	BP4XTypeHENTAI
)

var bp4xGrabList = []string{
	BP4XTypeJAV:     bp4xJavURL,
	BP4XTypeAMATEUR: bp4xAmateurURL,
	BP4XTypeIV:      bp4xIVURL,
	BP4XTypeHENTAI:  bp4xHentaiURL,
}

type grabBP4X struct {
	doc      *goquery.Document
	grabType GrabBP4XType
}

// Name ...
func (g *grabBP4X) Name() string {
	return "bp4x"
}

// Decode ...
func (g *grabBP4X) Decode([]*Message) error {
	panic("implement me")
}

// Find ...
func (g *grabBP4X) Find(name string) (IGrab, error) {
	name = strings.ToUpper(name)
	url := bp4xGrabList[g.grabType]
	url = fmt.Sprintf(url, name)
	document, e := query.New(url)
	if e != nil {
		return g, e
	}
	g.doc = document
	//ret, e := document.Html()
	//log.Println(ret)
	return g, nil
}

// NewGrabBP4X ...
func NewGrabBP4X(grabType GrabBP4XType) IGrab {
	return &grabBP4X{
		grabType: grabType,
	}
}
