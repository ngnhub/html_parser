package service

import "golang.org/x/net/html"

type ValueAndNode struct {
	Key   string
	Value string
	Node  *html.Node
}

func emptyValueAndNode(key string) ValueAndNode {
	return ValueAndNode{key, "", nil}
}

type ValueAndParent struct {
	Key    string
	Value  string
	Parent *html.Node
	depth  int
}

func emptyValueAndParent(key string) ValueAndParent {
	return ValueAndParent{key, "", nil, 0}
}

type ValuesAndNodes []ValueAndNode

type ValuesAndParents []ValueAndParent

func (v *ValuesAndParents) isEmpty() bool {
	allNil := true
	for _, v := range *v {
		if v.Parent != nil {
			allNil = false
		}
	}
	return allNil
}

func (v *ValuesAndNodes) mapToStrings() []string {
	res := make([]string, 0, len(*v))
	for _, val := range *v {
		res = append(res, val.Value)
	}
	return res
}

func (v *ValuesAndParents) mapToStrings() []string {
	res := make([]string, 0, len(*v))
	for _, val := range *v {
		res = append(res, val.Value)
	}
	return res
}
