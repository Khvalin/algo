package main

import "fmt"
func maxProfit(k int, prices []int) int {
	type state struct {
		buy   bool
		ind   int
		last  int
		sum   int
		count int
	}

	res := 0
  
  if k > len(prices) >> 1 + 1 {
    k = len(prices) >> 1 + 1
  }

	max := make([][2][]map[int]int, len(prices) + 1)
  for i := range max {
    max[i][0] = make([]map[int]int, k+1)
    max[i][1] = make([]map[int]int, k+1)
  }

	q := []state{{true, 0, -1, 0, 0}}

	for len(q) > 0 {
		st := q[0]

		if st.count > k {
			q = q[1:]
			continue
		}

		i := 0
		if !st.buy {
			i = 1
		}

		key := st.last
		if m, found := max[st.ind][i][st.count][key]; found && st.sum < m {
			q = q[1:]
			continue
		}

    if max[st.ind][i][st.count] == nil {
      max[st.ind][i][st.count] = map[int]int{}
    }
		max[st.ind][i][st.count][key] = st.sum

		if res < st.sum {
			res = st.sum
		}

		if st.ind >= len(prices) || st.count == k {
			q = q[1:]
			continue
		}

		p := prices[st.ind]
		q[0].ind++

		if st.buy {
			if (st.ind == 0) || (st.ind < len(prices)-1 && (prices[st.ind-1] > p) && (prices[st.ind+1] > p)) {
				q = append(q, state{false, st.ind + 1, p, st.sum, st.count})
			}
			continue
		}

		if p > st.last {
			q = append(q, state{true, st.ind + 1, 0, st.sum + p - st.last, st.count + 1})
		}
	}

	return res
}

func main() {
	// fmt.Println(13, maxProfit(2, []int{1,2,4,2,5,7,2,4,9,0}))

	fmt.Println(13, maxProfit(2, []int{3, 5, 3, 4, 1, 4, 5, 0, 7, 8, 5, 6, 9, 4, 1}))
}

