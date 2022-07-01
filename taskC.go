package taskC

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func printmat(mat [][]int) {
	for _, row := range mat {
		s := []string{}
		for _, x := range row {
			s = append(s, fmt.Sprintf("%d", x))
		}
		fmt.Print(strings.Join(s, " ") + "\n")
	}
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	line := in.Text()
	n, _ := strconv.Atoi(strings.TrimSuffix(line, "\n"))

	ans := [][][]int{}
	for i := 0; i < n; i++ {
		in.Scan()
		in.Scan()
		line = in.Text()
		line = strings.TrimSuffix(line, "\n")
		rows, _ := strconv.Atoi(strings.Fields(line)[0])
		mat := [][]int{}
		for row := 0; row < rows; row++ {
			tmp := []int{}
			in.Scan()
			line = in.Text()
			line = strings.TrimSuffix(line, "\n")
			for _, el := range strings.Fields(line) {
				x, _ := strconv.Atoi(el)
				tmp = append(tmp, x)
			}
			mat = append(mat, tmp)
		}
		in.Scan()
		in.Scan()
		line = in.Text()
		line = strings.TrimSuffix(line, "\n")

		order := []int{}
		for _, el := range strings.Fields(line) {
			x, _ := strconv.Atoi(el)
			order = append(order, x)
		}
		for _, column := range order {
			sort.SliceStable(mat, func(i, j int) bool {
				return mat[i][column-1] < mat[j][column-1]
			})
		}
		ans = append(ans, mat)
	}
	for _, mat := range ans {
		printmat(mat)
		fmt.Print("\n")
	}
}
