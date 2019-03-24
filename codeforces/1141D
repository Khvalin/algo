package main

import "fmt"

func solve(n int, left, right string) map[int]int {
	const QuestionMark = '?'

	strToCountMap := func(str string) map[rune][]int {
		countMap := map[rune][]int{}
		for i, ch := range str {
			_, found := countMap[ch]
			if !found {
				countMap[ch] = []int{}
			}

			countMap[ch] = append(countMap[ch], i)
		}

		return countMap
	}

	res := map[int]int{}

	leftMap, rightMap := strToCountMap(left), strToCountMap(right)

	//fmt.Println(rightMap)

	for i, leftChar := range left {
		if leftChar != QuestionMark {
			for _, ch := range [...]rune{leftChar, QuestionMark} {
				_, found := rightMap[ch]

				if found && len(rightMap[ch]) > 0 {
					res[i] = rightMap[ch][0]

					leftMap[leftChar] = leftMap[leftChar][1:]
					rightMap[ch] = rightMap[ch][1:]

					break
				}
			}
		}
	}

	for _, rightChar := range right {
		if len(rightMap[rightChar]) > 0 {
			//fmt.Println(rightChar, rightMap[rightChar])

			_, found := leftMap[QuestionMark]

			if found && len(leftMap[QuestionMark]) > 0 {
				res[leftMap[QuestionMark][0]] = rightMap[rightChar][0]

				leftMap[QuestionMark] = leftMap[QuestionMark][1:]
				rightMap[rightChar] = rightMap[rightChar][1:]
			}
		}
	}

	return res
}

func main() {
	n := 0
	left, right := "", ""

	fmt.Scan(&n)
	fmt.Scan(&left, &right)

	ans := solve(n, left, right)
	fmt.Println(len(ans))
	for l, r := range ans {
		fmt.Println(l+1, r+1)
	}
}
