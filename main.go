package main

import (
	"fmt"
	"strconv"
)

var (
	input     int
	b         int
	bitsLimit int
	numA      int
	numB      int
	strA      string
	strB      string
)

func main() {
	fmt.Println("1.type1->2進制表示法", "2.type2->2進制1的補數表示法", "3.type3->2進制2的補數表示法")
	println("選擇模式")
	fmt.Scanln(&input)
	switch input {
	case 1:
		result := dosth()
		fmt.Println(result)
	case 2:
		result := dosth()
		result = onebotwo(result)
		fmt.Println(result)
	case 3:
		result := dosth()
		result = tobotwo(result)
		fmt.Println(result)
	}
}

func reverseString(s string) string {
	runes := []byte(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func onebotwo(n string) string {
	var r []string
	var s string
	for _, v := range n {
		if v == '0' {
			r = append(r, "1")
		} else {
			r = append(r, "0")
		}
	}

	for _, v := range r {
		s += v
	}
	return s
}

func tobotwo(n string) string {
	var k int = 0
	var r []string
	var result string
	var rc []byte
	for i := len(n) - 1; i >= 0; i-- {
		rc = append(rc, n[i])
		if n[i] == '1' {
			k = i
			break
		}
	}
	var str string
	for _, v := range rc {
		s := string(v)
		str += s
	}
	str = reverseString(str)
	fmt.Println(k)
	for j := 0; j < k; j++ {
		if n[j] == '0' {
			r = append(r, "1")
		} else {
			r = append(r, "0")
		}
	}

	for _, v := range r {
		result += v
	}
	return result + str
}

func toTwo(n int) string {
	var str string
	if n == 0 || n == 1 {
		str = strconv.Itoa(n)
		return str
	} else {
		for i := 0; n > 0; i++ {
			str += strconv.Itoa(n % 2)
			n /= 2
		}
	}
	return str
}

func addBinary(str1 string, str2 string) string {
	if len(str1) == 0 {
		return str1
	}
	if len(str2) == 0 {
		return str2
	}

	i := len(str1) - 1
	j := len(str2) - 1
	carry := 0
	var result []byte
	for i >= 0 && j >= 0 {
		str1i, _ := strconv.Atoi(string(str1[i]))
		str2j, _ := strconv.Atoi(string(str2[j]))
		num := str1i + str2j + carry
		carry = num / 2
		num %= 2
		result = append(result, byte(num)+'0')
		i--
		j--
	}
	for i >= 0 {
		str1i, _ := strconv.Atoi(string(str1[i]))
		num := str1i + carry
		carry = num / 2
		num %= 2
		result = append(result, byte(num)+'0')
		i--
	}
	for j >= 0 {
		str2j, _ := strconv.Atoi(string(str2[j]))
		num := str2j + carry
		carry = num / 2
		num %= 2
		result = append(result, byte(num)+'0')
		j--
	}
	if carry > 0 {
		result = append(result, byte(carry)+'0')
	}
	for k := 0; k < len(result)/2; k++ {
		result[k], result[len(result)-1-k] = result[len(result)-1-k], result[k]
	}
	return string(result)
}

func checknum(b, num1, num2 int) bool {
	for num1 > 0 {
		a := num1 % 10
		if a > b {
			return false
		}
		num1 /= 10
	}

	for num2 > 0 {
		a := num2 % 10
		if a > b {
			return false
		}
		num2 /= 10
	}
	return true
}

func dosth() string {
	var sum string
	fmt.Println("輸入最大位元數")
	fmt.Scanln(&bitsLimit)
	fmt.Println("輸入最大位元")
	fmt.Scanln(&b)
	fmt.Println("數入第一個數字")
	fmt.Scanln(&numA)
	fmt.Println("輸入第二個數字")
	fmt.Scanln(&numB)
	c := checknum(b, numA, numB)
	if !c {
		fmt.Println("超過")
	} else {
		if len(strconv.Itoa(numA)) > bitsLimit || len(strconv.Itoa(numB)) > bitsLimit {
			fmt.Println("over bitsLimit")
		} else if numA > 0 && numB > 0 || numA < 0 && numB < 0 {
			sum = toTwo(numA + numB)
			return sum
		} else if numA > 0 && numB < 0 || numA < 0 && numB > 0 {
			if numA < 0 {
				numA = -numA
			}
			if numB < 0 {
				numB = -numB
			}
			strA = toTwo(numA)
			strB = toTwo(numB)
			sum = addBinary(strA, strB)
		}
	}
	return sum
}
