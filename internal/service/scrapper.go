package service

import (
	"github.com/ngnhub/html_scrapper/internal/service/searcher"
	"golang.org/x/net/html"
)

type ScrapperService struct {
	Searcher searcher.Searcher
}

type Found struct {
	Values []string
}

func (service *ScrapperService) Scrap(keys []string, source *html.Node) []Found {
	var results []Found
	firstValueToNodes := service.searchFirstNodes(keys, source)
	results = append(results, Found{firstValueToNodes.MapToStrings()})

	secondValueToNodes := service.searchSecondNodes(firstValueToNodes)
	results = append(results, Found{secondValueToNodes.MapToStrings()})

	currentValueToNodes := make(searcher.ValuesAndParents, len(secondValueToNodes))
	copy(currentValueToNodes, secondValueToNodes)
	for !currentValueToNodes.IsEmpty() {
		nextValueToNodes := service.searchNextNodes(currentValueToNodes)
		if !nextValueToNodes.IsEmpty() {
			results = append(results, Found{Values: nextValueToNodes.MapToStrings()})
		}
		currentValueToNodes = nextValueToNodes
	}
	return results
}

func (service *ScrapperService) searchFirstNodes(keys []string, source *html.Node) searcher.ValuesAndNodes {
	firstValueToNodes := make(searcher.ValuesAndNodes, 0, len(keys))
	for _, key := range keys {
		valueToNode, _ := service.Searcher.SearchFirstNode(key, source)
		firstValueToNodes = append(firstValueToNodes, valueToNode)
	}
	return firstValueToNodes
}

func (service *ScrapperService) searchSecondNodes(firstValueToNodes searcher.ValuesAndNodes) searcher.ValuesAndParents {
	secondValueToNodes := make(searcher.ValuesAndParents, 0, len(firstValueToNodes))
	for _, first := range firstValueToNodes {
		second, _ := service.Searcher.SearchSecondNode(first.Key, first.Node, 0)
		secondValueToNodes = append(secondValueToNodes, second)
	}
	return secondValueToNodes
}

func (service *ScrapperService) searchNextNodes(currentValueToNodes searcher.ValuesAndParents) searcher.ValuesAndParents {
	nextValueToNodes := searcher.ValuesAndParents{}
	for _, value := range currentValueToNodes {
		nextValueToNodes = append(nextValueToNodes, service.Searcher.GetNextSiblingValue(value))
	}
	return nextValueToNodes
}
