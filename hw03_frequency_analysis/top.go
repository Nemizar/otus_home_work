package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

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

	pairList := make(PairList, sLen)

	pairListIdx := 0
	for k, v := range scores {
		pairList[pairListIdx] = Pair{k, v}
		pairListIdx++
	}

	sort.Sort(pairList)

	res := make([]string, 0)
	for _, pair := range pairList {
		res = append(res, pair.Key)
	}

	if sLen > 10 {
		return res[:10]
	}
	return res
}
