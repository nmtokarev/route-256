package taskD

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func main() {
	in := bufio.NewScanner(os.Stdin)
	re, _ := regexp.Compile(`^[a-zA-Z0-9_][a-zA-Z0-9_-]{1,23}$`)
	n := Scan_Int(in)
	ans := [][]string{}
	for i := 0; i < n; i++ {
		m := Scan_Int(in)
		s := map[string]bool{}
		ans_part := []string{}
		for j := 0; j < m; j++ {
			login := Scan_Str(in)
			matched := re.MatchString(login)
			_, ok := s[strings.ToLower(login)]
			if matched && !ok {
				ans_part = append(ans_part, "YES")
				s[strings.ToLower(login)] = true
			} else {
				ans_part = append(ans_part, "NO")
			}
		}
		ans = append(ans, ans_part)
	}
	for _, a := range ans {
		fmt.Print(strings.Join(a, "\n"))
		fmt.Print("\n\n")
	}
}
