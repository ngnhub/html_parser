package service

import (
	"golang.org/x/net/html"
	"strings"
)

type Searcher interface {
	searchFirstNode(key string, node *html.Node) (ValueAndNode, bool)
	searchSecondNode(key string, prevSibling *html.Node, depth int) (ValueAndParent, bool)
	getNextSiblingValue(prevParent ValueAndParent) ValueAndParent
}

type DefaultSearcher struct{}

func (d DefaultSearcher) searchFirstNode(key string, node *html.Node) (ValueAndNode, bool) {
	if node == nil {
		return emptyValueAndNode(key), false
	}
	if isMatch(key, node) {
		return ValueAndNode{key, extractValue(node.FirstChild), node}, true
	}
	val, ok := d.searchFirstNode(key, firstChildWithoutATrash(node))
	if !ok {
		val, ok = d.searchFirstNode(key, nextSiblingWithoutATrash(node))
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

func (d DefaultSearcher) searchSecondNode(key string, prevSibling *html.Node, depth int) (ValueAndParent, bool) {
	if prevSibling == nil {
		return emptyValueAndParent(key), false
	}
	current := nextSiblingWithoutATrash(prevSibling)
	parent := prevSibling.Parent
	if current == nil {
		return d.searchSecondNode(key, parent, depth+1)
	}
	child := current
	currDepth := depth
	for child != nil && currDepth != 0 {
		child = firstChildWithoutATrash(child)
		currDepth--
	}
	found, ok := scanSiblingsValues(key, child)
	if ok {
		return ValueAndParent{key, found, current, depth}, true
	}
	return d.searchSecondNode(key, current, depth)
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

func (d DefaultSearcher) getNextSiblingValue(prevParent ValueAndParent) ValueAndParent {
	current := prevParent.Parent
	key := prevParent.Key
	if current == nil {
		return emptyValueAndParent(key)
	}
	current = nextSiblingWithoutATrash(current)
	if current == nil {
		return emptyValueAndParent(key)
	}
	child := current
	currDepth := prevParent.depth
	for child != nil && currDepth != 0 {
		child = firstChildWithoutATrash(child)
		currDepth--
	}
	if child != nil && isMatch(key, child) {
		return ValueAndParent{key, extractValue(current.FirstChild), current, prevParent.depth}
	}
	return emptyValueAndParent(key)
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
