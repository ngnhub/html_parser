package service

import (
	"bytes"
	"golang.org/x/net/html"
	"os"
	"reflect"
	"testing"
)

func TestScrap(t *testing.T) {
	// given
	simpleCaseFile, _ := os.ReadFile("test_data/scrapper_test_simple_case.html")
	simpleCase, _ := html.Parse(bytes.NewReader(simpleCaseFile))
	childIsMissedCaseFile, _ := os.ReadFile("test_data/scrapper_test_when_child_is_missed.html")
	childIsMissed, _ := html.Parse(bytes.NewReader(childIsMissedCaseFile))
	whenOnlySingleValueCaseFile, _ := os.ReadFile("test_data/scrapper_test_when_only_single_value.html")
	whenOnlySingleValueCase, _ := html.Parse(bytes.NewReader(whenOnlySingleValueCaseFile))
	keys := []string{
		"Test class 1", "Test class 2",
	}

	type args struct {
		keys []string
		node *html.Node
	}
	searcher := DefaultSearcher{}
	tests := []struct {
		name    string
		service *ScrapperService
		args    args
		want    []Found
	}{
		{name: "Simple case",
			service: &ScrapperService{Searcher: searcher},
			args:    args{keys: keys, node: simpleCase},
			want: []Found{
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"Some Value", ""}},
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"", "Some Value 2"}},
				{[]string{"", "Some Value 2"}},
			},
		},
		{name: "When child is missed",
			service: &ScrapperService{Searcher: searcher},
			args:    args{keys: keys, node: childIsMissed},
			want: []Found{
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"Some Value", ""}},
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"", "Some Value 2"}},
			},
		},
		{name: "When only single value",
			service: &ScrapperService{Searcher: searcher},
			args:    args{keys: keys, node: whenOnlySingleValueCase},
			want: []Found{
				{[]string{"Some Value", "Some Value 2"}},
				{[]string{"", "Some Value 2"}},
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
