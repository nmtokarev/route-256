package taskF

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Scan_Int(in *bufio.Scanner) int {
	in.Scan()
	line := in.Text()
	n, _ := strconv.Atoi(strings.TrimSuffix(line, "\n"))
	return n
}

func Scan_Str(in *bufio.Scanner) string {
	in.Scan()
	line := in.Text()
	return strings.TrimSuffix(line, "\n")
}

func Scan_Coupe_K(in *bufio.Scanner) (coupe int, k int) {
	in.Scan()
	line := in.Text()
	line = strings.TrimSuffix(line, "\n")
	coupe, _ = strconv.Atoi(strings.Fields(line)[0])
	k, _ = strconv.Atoi(strings.Fields(line)[1])
	return coupe, k
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func num_coupe(x int) int {
	if x%2 == 0 {
		return x / 2
	}
	return (x + 1) / 2
}

func num_place(x int) int {
	return 1 - x%2
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	n := Scan_Int(in)
	ans := [][]string{}
	for i := 0; i < n; i++ {
		_ = Scan_Str(in)
		coupe, k := Scan_Coupe_K(in)
		dct := map[int][2]bool{}
		vagon := []string{}
		h := &IntHeap{}
		for j := 1; j <= coupe; j++ {
			heap.Push(h, j)
		}
		removed := map[int]bool{}
		for j := 0; j < k; j++ {
			in.Scan()
			line := in.Text()
			line = strings.TrimSuffix(line, "\n")
			if len(strings.Fields(line)) == 2 {
				a, _ := strconv.Atoi(strings.Fields(line)[0])
				b, _ := strconv.Atoi(strings.Fields(line)[1])
				if a == 1 {
					if !dct[num_coupe(b)][num_place(b)] && num_coupe(b) <= coupe {
						vagon = append(vagon, "SUCCESS")
						if !dct[num_coupe(b)][0] && !dct[num_coupe(b)][1] {
							removed[num_coupe(b)] = true
						}
						tmp := dct[num_coupe(b)]
						tmp[num_place(b)] = true
						dct[num_coupe(b)] = tmp
					} else {
						vagon = append(vagon, "FAIL")
					}
				}
				if a == 2 {
					if dct[num_coupe(b)][num_place(b)] && num_coupe(b) <= coupe {
						vagon = append(vagon, "SUCCESS")
						tmp := dct[num_coupe(b)]
						tmp[num_place(b)] = false
						dct[num_coupe(b)] = tmp
						if !dct[num_coupe(b)][0] && !dct[num_coupe(b)][1] {
							heap.Push(h, num_coupe(b))
							delete(removed, num_coupe(b))
						}
					} else {
						vagon = append(vagon, "FAIL")
					}
				}
			} else {
				var x int
				var ok bool
				if len(*h) > 0 {
					x = heap.Pop(h).(int)
					_, ok = removed[x]
				}
				for ok && len(*h) > 0 {
					x = heap.Pop(h).(int)
					_, ok = removed[x]
				}

				if x <= coupe && x > 0 && !ok {
					vagon = append(vagon, fmt.Sprintf("SUCCESS %d-%d", 2*x-1, 2*x))
					dct[x] = [2]bool{true, true}
					removed[x] = true
				} else {
					vagon = append(vagon, "FAIL")
				}
			}
		}
		ans = append(ans, vagon)
	}
	for _, vagon := range ans {
		fmt.Printf("%s\n\n", strings.Join(vagon, "\n"))
	}
}
