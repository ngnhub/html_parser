package search

import "golang.org/x/net/html"

type Key struct {
	Elem string
	Name string
}

type ValueAndNode struct {
	Key   Key
	Value string
	Node  *html.Node
}

func EmptyValueAndNode(key Key) ValueAndNode {
	return ValueAndNode{key, "", nil}
}

type ValueAndParent struct {
	Key    Key
	Value  string
	Parent *html.Node
	Depth  int
}

func EmptyValueAndParent(key Key) ValueAndParent {
	return ValueAndParent{key, "", nil, 0}
}

type ValuesAndNodes []ValueAndNode

type ValuesAndParents []ValueAndParent

func (v *ValuesAndParents) IsEmpty() bool {
	allNil := true
	for _, v := range *v {
		if v.Parent != nil {
			allNil = false
		}
	}
	return allNil
}

func (v *ValuesAndNodes) MapToStrings() []string {
	res := make([]string, 0, len(*v))
	for _, val := range *v {
		res = append(res, val.Value)
	}
	return res
}

func (v *ValuesAndParents) MapToStrings() []string {
	res := make([]string, 0, len(*v))
	for _, val := range *v {
		res = append(res, val.Value)
	}
	return res
}
