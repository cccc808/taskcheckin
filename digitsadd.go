package main
import "fmt"
import "strconv"

func plusOne(digits []int) []int{
	strdigit := ""
	var digitarray []int
	for _, digit := range digits {
		strdigit += strconv.Itoa(digit)
	}

    finaldigit, err := strconv.Atoi(strdigit)
	if err!= nil {
		fmt.Println("Error converting string to integer")
	} else {
		finaldigit = finaldigit + 1
		fmt.Println("final digit:", finaldigit)
		strdigit = strconv.Itoa(finaldigit)
		for _, digit := range strdigit{
			intdigit, err := strconv.Atoi(string(digit))
			if err!= nil {
				fmt.Println("Error converting string to integer")
			} else
			{
				digitarray = append(digitarray,intdigit)
			}
		
		}
	}
	return digitarray
}

func main() {
	digits := []int{10,2,3}
	fmt.Printf("%d digits add one = %d", digits,plusOne(digits))
}