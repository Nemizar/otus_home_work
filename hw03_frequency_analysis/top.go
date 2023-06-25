package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const resultLength = 10

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	if p[j].Value == p[i].Value {
		return p[i].Key < p[j].Key
	}
	return p[j].Value < p[i].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var re = regexp.MustCompile(`([a-zа-яё]-?)+`)

func Top10(str string) []string {
	str = strings.ToLower(str)
	words := re.FindAllString(str, -1)

	scores := make(map[string]int)

	for _, word := range words {
		scores[word]++
	}

	sLen := len(scores)

	pairList := make(PairList, 0, sLen)

	for k, v := range scores {
		pairList = append(pairList, Pair{k, v})
	}

	sort.Sort(pairList)

	res := make([]string, 0, resultLength)

	for _, pair := range pairList {
		if len(res) >= resultLength {
			break
		}

		res = append(res, pair.Key)
	}

	return res
}
