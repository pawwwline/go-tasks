package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	angs := readInput()
	ans := searchAnagrams(angs)
	fmt.Println(ans)

}

func searchAnagrams(stringSlice []string) map[string][]string {
	res := make(map[string][]string)

	for _, str := range stringSlice {
		str = strings.ToLower(str)
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		res[string(runes)] = append(res[string(runes)], str)
	}

	return changeMap(res)
}

func changeMap(res map[string][]string) map[string][]string {
	ans := make(map[string][]string)
	for k, v := range res {
		if len(v) == 1 {
			delete(res, k)
			continue
		}
		ans[v[0]] = append(ans[v[0]], v...)
	}

	for _, group := range ans {
		sort.Slice(group, func(i, j int) bool {
			return group[i] < group[j]
		})

	}
	return ans
}

func readInput() []string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(data))
	return words
}
