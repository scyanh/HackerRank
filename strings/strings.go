package strings

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Palindrome
func isPalindrome(s string) bool {
	array := strings.Split(s, "")

	for i, e := range array {
		if e != array[len(array)-1-i] {
			return false
		}
	}

	return true
}
func isPalindrome2(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

// Shift with string split
func shiftWord(positions int) string {
	alfabet := "abcdefghijklmnopqrstuvwxyz"
	str := "xyz"
	strShifted := ""
	strArray := strings.Split(str, "")
	newPosition := 0
	for _, e := range strArray {
		for i, char := range alfabet {
			if string(char) == e {
				shift := i + positions
				if shift <= len(alfabet)-1 {
					newPosition = shift
				} else {
					newPosition = shift % len(alfabet)
				}
				strShifted += string(alfabet[newPosition])
			}
		}
	}

	return strShifted
}

// Palindromic
func LongestPalindromicSubstring(s string) string {
	palindrom := ""
	maxSize := 0

	if len(s) == 1 {
		return s
	}

	for i, e := range s {
		s2 := ""
		for j := i + 1; j < len(s); j++ {
			s2 += string(s[j])
			if isPalindrome(string(e) + s2) {
				palindromeSize := 1 + len(s2)
				if palindromeSize > maxSize {
					maxSize = palindromeSize
					palindrom = string(e) + s2
				}
			}
		}
	}

	return palindrom
}

// Palindromic
func GroupAnagrams(words []string) [][]string {
	result := [][]string{}
	memo := make(map[string]bool)

	for i, e := range words {
		array := []string{}
		if !memo[e] {
			memo[e] = true
			array = append(array, e)
		}
		if i < len(words)-1 {
			for j := i + 1; j < len(words); j++ {
				if orderedString(e) == orderedString(words[j]) &&
					!memo[words[j]] {
					memo[words[j]] = true
					array = append(array, words[j])
				}
			}
		}
		if len(array) > 0 {
			result = append(result, array)
		}
	}

	return result
}

func orderedString(s string) string {
	array := strings.Split(s, "")
	sort.Strings(array)
	s = strings.Join(array, "")
	return s
}
func SplitWord(s string, n int) {

	s1 := s[:n]
	s2 := s[1:n]
	s3 := s[n:]
	s4 := s[n+1:]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

func CamelCaseCount(s string) int {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetArray := strings.Split(alphabet, "")
	sum := 0
	array := strings.Split(s, "")

	for _, e := range array {
		for _, e2 := range alphabetArray {
			if e == e2 {
				sum += 1
			}
		}
	}

	return sum + 1
}
func StrongPassword(s string) int {
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"
	sum := 0

	haveNumber := false
	haveLowerCase := false
	haveUpperCase := false
	haveSpecialChar := false

	//array := strings.Split(s, "")
	for _,e:=range s{
		//numbers
		for _,e2:=range numbers{
			if e==e2{
				haveNumber=true
			}
		}
		//lower case
		for _,e2:=range lower_case{
			if e==e2{
				haveLowerCase=true
			}
		}
		//upper case
		for _,e2:=range upper_case{
			if e==e2{
				haveUpperCase=true
			}
		}
		//upper case
		for _,e2:=range special_characters{
			if e==e2{
				haveSpecialChar=true
			}
		}
	}

	if !haveNumber {
		sum += 1
		s += "1"
	}
	if !haveLowerCase {
		sum += 1
		s += "a"
	}
	if !haveUpperCase {
		sum += 1
		s += "A"
	}
	if !haveSpecialChar {
		sum += 1
		s += "+"
	}
	if len(s) < 6 {
		sum += 6 - len(s)
	}

	return sum
}

func ProgressBar(s string) int{
	array := strings.Split(s, "")

	size:=len(s)
	percent:=float64(0)
	progress:=0
	for i,e:=range array{
		if i<len(array)-1{
			percent=float64(i+1)*100/float64(size)
		}else{
			percent=100
		}

		if i<len(array)-1 && e!=array[i+1]{
			progress=int(percent)
			return progress
		}else {
			progress=int(percent)

		}
	}
	return progress
}

//# Input : K = 3 # Output : 000, 001, 010, 100, 101
//  Input : K = 4 # Output :0000, 0001, 0010, 0100, 0101, 1000, 1001, 1010
func Creator(input int) []string{
	var output []string

	for i:=0.0; i<math.Pow(2.0, float64(input)); i++{
		var sb strings.Builder
		val:=int(i)
		s:=fmt.Sprintf("%b",val)
		// eliminate the consecutives ones
		if strings.Contains(s,"11"){
			continue
		}

		// append with 0's
		total:=len(s)
		for j:=total; j<input; j++{
			sb.WriteString("0")
		}
		sb.WriteString(s)
		output=append(output, sb.String())
	}

	fmt.Println("output=",output)

	return output
}

func StartThis() {
	//fmt.Println(shiftWord(2))
	//fmt.Println(LongestPalindromicSubstring("a"))

	//words := []string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}
	//fmt.Println(GroupAnagrams(words))

	//SplitWord("abcdefgh",3)

	//fmt.Println(CamelCaseCount("saveChangesInTheEditor"))
	//fmt.Println(StrongPassword("Ab1"))
	fmt.Println(ProgressBar("0....."))
	fmt.Println(ProgressBar("000000"))
}


