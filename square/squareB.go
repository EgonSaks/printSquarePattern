package square

import "fmt"

/* 5,3 Output  */

/***\
*   *
\***/

func SquareB(column int, row int) {
	if row > 0 && column > 0 {
		for i := 1; i <= row; i++ {

			for j := 1; j <= column; j++ {

				/* Print row (x) */
				isFirstRow := i == 1
				isLastRow := i == row

				/* Print column (y) */
				isFirstColumn := j == 1
				isLastColumn := j == column

				// Print corners /  \
				if isFirstRow && isFirstColumn {
					fmt.Print(string('/'))
				}

				if isFirstColumn && isLastRow && row > 1 {
					fmt.Print(string('\\'))
				}

				if isFirstRow && isLastColumn && column > 1 {
					fmt.Print(string('\\'))
				}

				if isLastRow && isLastColumn && row > 1 && column > 1 {
					fmt.Print(string('/'))
				}

				// Print the center part *    *
				firstLast := (isFirstRow || isLastRow) && (!isFirstColumn && !isLastColumn)
				isMiddle := (!isFirstRow && !isLastRow) && (isLastColumn || isFirstColumn)

				if firstLast || isMiddle {
					fmt.Print(string('*'))
				}

				// Print the empty part ' '
				if !isFirstColumn && !isLastRow && !isFirstRow && !isLastColumn {
					fmt.Print(string(' '))
				}

			}
			fmt.Print(string('\n'))

		}
	}
}
