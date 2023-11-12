package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var n, t int
	var s []int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	if n > 0 {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &t)
			s = append(s, t)
		}
		s = qsort(s, 0, n)
		for i := 0; i < n; i++ {
			fmt.Print(s[i], " ")
		}

	}

}

func qsort(s []int, i int, j int) []int {

	el := s[i]
	c := 0
	for k := i; k < j; k++ {
		if s[k] != el {
			c++
		}
	}
	if c == 0 {
		return s
	}
	if i < j-1 {
		r := rand.Intn(j-i-1) + i

		r2 := rand.Intn(2)

		sr := s[r] + r2

		t, s := partition(sr, s, i, j)

		if t != j && t != i {

			s = qsort(s, i, t)

			s = qsort(s, t, j)
		} else {
			if (j - i) > 2 {
				if t == j {
					s = qsort(s, i, t)
					//_, s = partition(s[t], s, t, j)
				}
				if t == i {
					//_, s = partition(s[i], s, i, t)
					s = qsort(s, t, j)
				}
			} else {
				_, s = partition(s[t], s, t, j)
				_, s = partition(s[i], s, i, t)
			}

		}
	}
	return s
}

func partition(k int, s []int, l int, r int) (int, []int) {
	var t, i, j int
	i = l
	j = r
	for i < j {

		if s[i] >= k && s[j-1] < k {

			t = s[i]
			s[i] = s[j-1]
			s[j-1] = t
		} else {
			if s[i] < k {
				i++
			}
			if s[j-1] >= k {
				j--
			}
		}

	}

	return j, s
}
