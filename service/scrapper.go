package service

import (
	"golang.org/x/net/html"
)

type Searcher interface {
	searchFirstNode(key string, node *html.Node) (ValueAndNode, bool)
	searchSecondNode(key string, prevSibling *html.Node, depth int) (ValueAndParent, bool)
	getNextSiblingValue(prevParent ValueAndParent) ValueAndParent
}

type ScrapperService struct {
	Searcher Searcher
}

type Found struct {
	Values []string
}

func (service *ScrapperService) Scrap(keys []string, source *html.Node) []Found {
	var results []Found
	firstValueToNodes := service.searchFirstNodes(keys, source)
	results = append(results, Found{firstValueToNodes.mapToStrings()})

	secondValueToNodes := service.searchSecondNodes(firstValueToNodes)
	results = append(results, Found{secondValueToNodes.mapToStrings()})

	currentValueToNodes := make(ValuesAndParents, len(secondValueToNodes))
	copy(currentValueToNodes, secondValueToNodes)
	for !currentValueToNodes.isEmpty() {
		nextValueToNodes := service.searchNextNodes(currentValueToNodes)
		if !nextValueToNodes.isEmpty() {
			results = append(results, Found{Values: nextValueToNodes.mapToStrings()})
		}
		currentValueToNodes = nextValueToNodes
	}
	return results
}

func (service *ScrapperService) searchFirstNodes(keys []string, source *html.Node) ValuesAndNodes {
	searcher := service.Searcher
	firstValueToNodes := make(ValuesAndNodes, 0, len(keys))
	for _, key := range keys {
		valueToNode, _ := searcher.searchFirstNode(key, source)
		firstValueToNodes = append(firstValueToNodes, valueToNode)
	}
	return firstValueToNodes
}

func (service *ScrapperService) searchSecondNodes(firstValueToNodes ValuesAndNodes) ValuesAndParents {
	searcher := service.Searcher
	secondValueToNodes := make(ValuesAndParents, 0, len(firstValueToNodes))
	for _, first := range firstValueToNodes {
		second, _ := searcher.searchSecondNode(first.Key, first.Node, 0)
		secondValueToNodes = append(secondValueToNodes, second)
	}
	return secondValueToNodes
}

func (service *ScrapperService) searchNextNodes(currentValueToNodes ValuesAndParents) ValuesAndParents {
	searcher := service.Searcher
	nextValueToNodes := ValuesAndParents{}
	for _, value := range currentValueToNodes {
		nextValueToNodes = append(nextValueToNodes, searcher.getNextSiblingValue(value))
	}
	return nextValueToNodes
}
