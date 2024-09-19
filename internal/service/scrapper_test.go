package service

import (
	"bytes"
	"github.com/ngnhub/html_scrapper/internal/service/search"
	"github.com/ngnhub/html_scrapper/internal/service/search/default"
	"golang.org/x/net/html"
	"os"
	"reflect"
	"testing"
)

func TestScrap(t *testing.T) {
	simpleCaseFile, _ := os.ReadFile("test_data/scrapper_test_simple_case.html")
	simpleCase, _ := html.Parse(bytes.NewReader(simpleCaseFile))

	childIsMissedCaseFile, _ := os.ReadFile("test_data/scrapper_test_when_child_is_missed.html")
	childIsMissed, _ := html.Parse(bytes.NewReader(childIsMissedCaseFile))

	whenOnlySingleValueCaseFile, _ := os.ReadFile("test_data/scrapper_test_when_only_single_value.html")
	whenOnlySingleValueCase, _ := html.Parse(bytes.NewReader(whenOnlySingleValueCaseFile))

	differentKeysCaseFile, _ := os.ReadFile("test_data/scrapper_test_with_different_keys_case.html")
	differentKeysCase, _ := html.Parse(bytes.NewReader(differentKeysCaseFile))

	defaultKeys := []search.Key{
		{"div", "Test class 1"}, {"div", "Test class 2"},
	}

	searcher := defaultsearcher.DefaultSearcher{}
	scrapperService := PatternDetectScrapperService{Searcher: searcher}

	type args struct {
		keys []search.Key
		node *html.Node
	}
	tests := []struct {
		name    string
		service PatternDetectScrapperService
		args    args
		want    []Found
	}{
		{
			name:    "Simple case",
			service: scrapperService,
			args:    args{keys: defaultKeys, node: simpleCase},
			want: []Found{
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"Some Value", ""}},
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"", "Some Value 2"}},
				{[]string{"", "Some Value 2"}},
			},
		},
		{
			name:    "When child is missed",
			service: scrapperService,
			args:    args{keys: defaultKeys, node: childIsMissed},
			want: []Found{
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"Some Value", ""}},
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"", "Some Value 2"}},
			},
		},
		{
			name:    "When only single value",
			service: scrapperService,
			args:    args{keys: defaultKeys, node: whenOnlySingleValueCase},
			want: []Found{
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"", "Some Value 2"}},
			},
		},
		{
			name:    "With different keys",
			service: scrapperService,
			args: args{keys: []search.Key{{"div", "Div class"},
				{"h1", "h1 class"}}, node: differentKeysCase},
			want: []Found{
				{[]string{"Some Value", "Some P Value"}},
				{[]string{"Some Value", "Some P Value"}},
				{[]string{"Some Value", ""}},
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.service.Scrap(testCase.args.keys, testCase.args.node)
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("Scrap() failed. Want \n%v\n but got \n%v\n", testCase.want, got)
			}
		})
	}
}
