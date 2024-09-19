package search

import (
	"golang.org/x/net/html"
)

type Searcher interface {
	SearchFirstNode(key Key, node *html.Node) (ValueAndNode, bool)
	SearchSecondNode(key Key, prevSibling *html.Node, depth int) (ValueAndParent, bool)
	GetNextSiblingValue(prevParent ValueAndParent) ValueAndParent
}
