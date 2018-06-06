package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/ogataka50/stay-typing/stay_typing"
)

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()

	return ch
}

func main() {
	score := 0
	combo := 0
	maxCombo := 0

	t := stay_typing.Tasks{
		Difficulty: "normal",
	}

	c1 := color.New(color.FgHiRed).Add(color.Underline)
	c2 := color.New(color.FgHiGreen).Add(color.Underline)

	ch := input(os.Stdin)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	for {
		q := t.Question()
		c1.Println(q)
		fmt.Print("> ")

		select {
		case a := <-ch:
			fmt.Println(a)
			if a == q {
				combo++
				score++
				c2.Println("correct! " + strconv.Itoa(combo) + "combo!")
				if maxCombo < combo {
					maxCombo = combo
				}
			} else {
				combo = 0
				fmt.Println("nope")
			}
		case <-ctx.Done():
			c := color.New(color.FgHiBlue).Add(color.Underline)
			c.Println("Result: total score => " + strconv.Itoa(score) + " max combo => " + strconv.Itoa(maxCombo))
			return
		}
	}
}
