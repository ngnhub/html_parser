package searcher

import (
	"golang.org/x/net/html"
)

type Searcher interface {
	SearchFirstNode(key string, node *html.Node) (ValueAndNode, bool)
	SearchSecondNode(key string, prevSibling *html.Node, depth int) (ValueAndParent, bool)
	GetNextSiblingValue(prevParent ValueAndParent) ValueAndParent
}
