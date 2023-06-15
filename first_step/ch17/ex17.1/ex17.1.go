package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func isCorrect(answer int) int {
	var number int
	var stdin = bufio.NewReader(os.Stdin)
	var cnt int = 0
rightAnswer:
	for {
		fmt.Print("숫자 입력:")
		_, err := fmt.Scanln(&number)
		if err != nil {
			stdin.ReadString('\n')
			fmt.Println("0 또는 자연수만 입력해 주세요!!")
			continue
		}
		switch {
		case number >= 100:
			println("100미만의 수를 입력해 주세요")
		case number < 100 && number > answer:
			println("down")
			cnt++
		case number < answer && number >= 0:
			println("up")
			cnt++
		case number == answer:
			println("정답입니다")
			break rightAnswer
		default:
			println("0 이상의 수를 입력해 주세요")
		}
	}
	return cnt
}

func main() {
	rand.Seed(time.Now().UnixNano())

	answer := rand.Intn(100)
	// fmt.Println(n)
	var c int = isCorrect(answer)
	fmt.Print("시도 횟수", c)
}
