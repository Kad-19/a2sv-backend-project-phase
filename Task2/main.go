package main

import ("fmt")

func main () {
	count := Counter("hello")
	for r, c := range count {
		
		fmt.Printf("%c : %d\n", r, c)
	}

	pal := "eye"

	if palindrome_check(pal) {
		fmt.Printf("\n%s is palindrome", pal)

	} else {
		fmt.Printf("\n%s is not palindrome", pal)

	}

}

func Counter(word string) map[rune]int {
	m := make(map[rune]int)

	for _, runeVal := range word {
		e, ok := m[runeVal]
		if ok {
			m[runeVal] = e + 1
		} else {
			m[runeVal] = 1
		}
	}
	return m
}

func palindrome_check(word string) bool {
	reverse := []rune(word)

	for i, runeVal := range word {
		reverse[len(word) - i - 1] = runeVal
	}
	return string(reverse) == word
}