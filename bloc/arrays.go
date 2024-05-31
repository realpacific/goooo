package bloc

import "strings"

func MultiplyBy100(arr [7]int32) {
	for index, element := range arr {
		arr[index] = element * 100
	}
}

func MultiplyBy100_(arr *[7]int32) {
	for index, element := range *arr {
		(*arr)[index] = element * 100
	}
}

func MatrixMultiply(a [][]int, b [][]int) [][]int {
	result := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = make([]int, len(b[0]))
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

const LOREM_IPSUM = `Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?`

func Take(from, to int) []string {
	return strings.Split(LOREM_IPSUM, " ")[from:to]
}

func InsertInto(array []int, index int, replaceWith []int) []int {
	before := array[:index]
	after := array[index:]
	result := make([]int, 0, len(before)+len(replaceWith)+len(after))

	result = append(result, before...)
	result = append(result, replaceWith...)
	result = append(result, after...)

	return result
}
