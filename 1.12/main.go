package main

import "fmt"

//Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
//
//Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

// сложность O(n), память O(n)
func Set(strs []string) []string {
	hashSet := make(map[string]struct{})
	ans := []string{}
	for _, s := range strs {
		hashSet[s] = struct{}{}
	}
	for s := range hashSet {
		ans = append(ans, s)
	}

	return ans
}
func main() {

	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(Set(strs))

}
