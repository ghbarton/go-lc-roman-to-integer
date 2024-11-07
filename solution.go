package solution

func solution(s string) int {
	conv := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for _, v := range s {
		result += conv[v]
	}
	return result
}
