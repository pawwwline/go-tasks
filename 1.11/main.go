package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.
//
//Пример:
//A = {1,2,3}
//B = {2,3,4}
//Пересечение = {2,3}

// Intersect сложность O(n + m) память O(n) где n - это длина первого слайса
func Intersect(a, b []int) []int {
	hashMap := make(map[int]struct{})

	ans := []int{}

	for i := 0; i < len(a); i++ {
		hashMap[a[i]] = struct{}{}
	}

	for i := 0; i < len(b); i++ {
		_, ok := hashMap[b[i]]
		if ok {
			ans = append(ans, b[i])
			delete(hashMap, b[i])
		}
	}

	return ans
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	fmt.Println(Intersect(A, B))
}
