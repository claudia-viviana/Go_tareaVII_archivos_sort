package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var opc int64
	slice := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println("Selecciona una opción: ")
		fmt.Println("1. Capturar strings")
		fmt.Println("2. Crear Archivo de strings ordenadas Ascendentemente")
		fmt.Println("3. Crear Archivo de strings ordenadas Descendentemente")
		fmt.Println("4. Salir")

		fmt.Scanln(&opc)

		switch {
		// CAPTURAR STRINGS _______________________________________________________________
		case opc == 1:
			var n int64
			var str string

			fmt.Print("Cantidad de strings a ingresar: ")
			fmt.Scanln(&n)

			for i := 0; i < int(n); i++ {
				scanner.Scan()
				str = scanner.Text()
				slice = append(slice, str)
			}

		// ASCENDENTE ________________________________________________________________
		case opc == 2:

			sort.Sort(sort.StringSlice(slice))            // Ordenar slice de strings
			slice2 := []string{strings.Join(slice, "\n")} // Agregar saltos de linea entre cada string

			file, err := os.Create("ascendente.text")

			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			for _, e := range slice2 {
				file.WriteString(e)
			}

		// DESCENDENTE ________________________________________________________________
		case opc == 3:

			sort.Sort(sort.Reverse(sort.StringSlice(slice))) // Ordenar slice de strings
			slice3 := []string{strings.Join(slice, "\n")}    // Agregar saltos de linea entre cada string

			file2, err := os.Create("descendente.text")

			if err != nil {
				fmt.Println(err)
				return
			}
			defer file2.Close()

			for _, e := range slice3 {
				file2.WriteString(e)
			}

		// Salir _______________________________________________________________
		case opc == 4:
			return
		}
	}

}
