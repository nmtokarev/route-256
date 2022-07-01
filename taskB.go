package taskB

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	const maxCapacity = 2 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	in.Buffer(buf, maxCapacity)
	in.Scan()
	line := in.Text()
	n, _ := strconv.Atoi(strings.TrimSuffix(line, "\n"))
	ans := []string{}
	for i := 0; i < n; i++ {
		in.Scan()
		dct := map[int]int{}
		in.Scan()
		line := in.Text()
		line = strings.TrimSuffix(line, "\n") // trim
		for _, key := range strings.Fields(line) {
			tmp, _ := strconv.Atoi(key)
			dct[tmp]++ // counter of number
		}
		s := 0 // evaluate sum
		for key := range dct {
			s += (dct[key] - dct[key]/3) * key
		}
		ans = append(ans, fmt.Sprintf("%d", s))
	}
	fmt.Print(strings.Join(ans, "\n"))
}
