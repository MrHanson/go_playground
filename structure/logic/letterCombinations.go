package logic

import (
	"fmt"
)

// 数字到字母表的映射：
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// 实现一个函数，对于给定的一串数字，例如"1", "233", 返回一个包含所有可能的组合的字符串列表
// 233 -> ["dff", "dgg"]
func letterCombinations(digits string) []string {
	combinations := make([]string, 0)
	if len(digits) == 0 {
		return combinations
	}
	backtrack(digits, 0, "", &combinations)

	return combinations
}

func backtrack(digits string, index int, combination string, combinations *[]string) {
	if index == len(digits) {
		*combinations = append(*combinations, combination)
		return
	}

	digit := string(digits[index])
	letters := phoneMap[digit]
	for _, letter := range letters {
		backtrack(digits, index+1, combination+string(letter), combinations)
	}
}

func LetterCombinationsRun() {
	ret := letterCombinations("233")
	fmt.Println("letterCombinations 233:", ret)
}
