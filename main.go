package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n int64
	var str string
	slice := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Cantidad de cadenas a ingresar: ")
	fmt.Scanln(&n)

	for i := 0; i < int(n); i++ {
		scanner.Scan()
		str = scanner.Text()
		slice = append(slice, str)
	}

	sort.Sort(sort.StringSlice(slice))
	slice2 := []string{strings.Join(slice, "\n")}

	file, err := os.Create("ascendente.text")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, e := range slice2 {
		file.WriteString(e)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	slice3 := []string{strings.Join(slice, "\n")}

	file2, err := os.Create("descendente.text")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	for _, e := range slice3 {
		file2.WriteString(e)
	}
}
