package square

import (
	"fmt"
)

/* "0", "-", "|", " " */

func SquareA(column int, row int) {
	/* Iterate over each row (x) */
	for i := 1; i <= row; i++ {

		/* Iterate over each column (y) */
		for j := 1; j <= column; j++ {

			isFirstLastRow := i == 1 || i == row
			isFirstLastCol := j == 1 || j == column

			/* Check if it's First and Last row and First and last Column */
			if isFirstLastRow && isFirstLastCol {
				fmt.Printf("o")
			}

			if !isFirstLastRow && isFirstLastCol {
				fmt.Printf("|")
			}

			if isFirstLastRow && !isFirstLastCol {
				fmt.Printf("-")
			}

			if !isFirstLastCol && !isFirstLastRow {
				fmt.Printf(" ")
			}

		}

		/* Move to the next line */
		fmt.Printf("\n")
	}
}
