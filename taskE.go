package taskE

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func Scan_Name_Number(in *bufio.Scanner) (name string, number string) {
	in.Scan()
	line := in.Text()
	line = strings.TrimSuffix(line, "\n")
	name = strings.Fields(line)[0]
	number = strings.Fields(line)[1]
	return name, number
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func reverse(input []string) (output []string) {
	output = []string{}
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	n := Scan_Int(in)
	ans := []map[string][]string{}
	for i := 0; i < n; i++ {
		m := Scan_Int(in)
		dct := map[string][]string{}
		for j := 0; j < m; j++ {
			name, number := Scan_Name_Number(in)
			if !contains(dct[name], number) {
				dct[name] = append(dct[name], number)
			} else {
				dct[name] = remove(dct[name], number)
				dct[name] = append(dct[name], number)
			}
			if len(dct[name]) == 6 {
				dct[name] = dct[name][1:]
			}
		}
		ans = append(ans, dct)
	}
	for _, dct := range ans {
		keys := []string{}
		for key := range dct {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, key := range keys {
			tmp := reverse(dct[key])
			fmt.Printf("%s: %d %s\n", key, len(dct[key]), strings.Join(tmp, " "))
		}
		fmt.Print("\n")

	}
}
