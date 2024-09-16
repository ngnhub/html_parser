package defaultsearcher

import (
	"github.com/ngnhub/html_scrapper/internal/service/searcher"
	"golang.org/x/net/html"
	"strings"
)

type DefaultSearcher struct{}

func (d DefaultSearcher) SearchFirstNode(key string, node *html.Node) (searcher.ValueAndNode, bool) {
	if node == nil {
		return searcher.EmptyValueAndNode(key), false
	}
	if isMatch(key, node) {
		return searcher.ValueAndNode{Key: key, Value: extractValue(node.FirstChild), Node: node}, true
	}
	val, ok := d.SearchFirstNode(key, firstChildWithoutATrash(node))
	if !ok {
		val, ok = d.SearchFirstNode(key, nextSiblingWithoutATrash(node))
	}
	return val, ok
}

func isMatch(search string, node *html.Node) bool {
	attr := node.Attr
	return node.Data == "div" && attr != nil && anyMatch(search, attr) // todo can be not only div
}

func anyMatch(attrName string, attributes []html.Attribute) bool {
	match := false
	for _, attr := range attributes {
		match = attr.Val == attrName
	}
	return match
}

func (d DefaultSearcher) SearchSecondNode(key string, prevSibling *html.Node, depth int) (searcher.ValueAndParent, bool) {
	if prevSibling == nil {
		return searcher.EmptyValueAndParent(key), false
	}
	current := nextSiblingWithoutATrash(prevSibling)
	parent := prevSibling.Parent
	if current == nil {
		return d.SearchSecondNode(key, parent, depth+1)
	}
	child := current
	currDepth := depth
	for child != nil && currDepth != 0 {
		child = firstChildWithoutATrash(child)
		currDepth--
	}
	found, ok := scanSiblingsValues(key, child)
	if ok {
		return searcher.ValueAndParent{Key: key, Value: found, Parent: current, Depth: depth}, true
	}
	return d.SearchSecondNode(key, current, depth)
}

func scanSiblingsValues(key string, current *html.Node) (string, bool) {
	if current == nil {
		return "", false
	}
	if isMatch(key, current) {
		return extractValue(current.FirstChild), true
	}
	return scanSiblingsValues(key, nextSiblingWithoutATrash(current))
}

// todo: backtracking. find first non empty value non element
func extractValue(node *html.Node) string {
	if node == nil {
		return ""
	}
	if node.Type == html.TextNode {
		if node.NextSibling != nil {
			return extractValue(node.NextSibling)
		}
		return strings.TrimSpace(node.Data)
	}
	return extractValue(node.FirstChild)
}

func (d DefaultSearcher) GetNextSiblingValue(prevParent searcher.ValueAndParent) searcher.ValueAndParent {
	current := prevParent.Parent
	key := prevParent.Key
	if current == nil {
		return searcher.EmptyValueAndParent(key)
	}
	current = nextSiblingWithoutATrash(current)
	if current == nil {
		return searcher.EmptyValueAndParent(key)
	}
	child := current
	currDepth := prevParent.Depth
	for child != nil && currDepth != 0 {
		child = firstChildWithoutATrash(child)
		currDepth--
	}
	if child != nil && isMatch(key, child) {
		return searcher.ValueAndParent{Key: key,
			Value:  extractValue(current.FirstChild),
			Parent: current,
			Depth:  prevParent.Depth}
	}
	return searcher.EmptyValueAndParent(key)
}

func nextSiblingWithoutATrash(n *html.Node) *html.Node {
	next := n.NextSibling
	for next != nil && next.Type == html.TextNode {
		next = next.NextSibling
	}
	return next
}

func firstChildWithoutATrash(n *html.Node) *html.Node {
	next := n.FirstChild
	for next != nil && next.Type == html.TextNode {
		next = next.NextSibling
	}
	return next
}
