package main
import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("WeekDAY!")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before Noon")
	default:
		fmt.Println("after Noon")
	}

	whatIAm := func(i interface{}) {
		switch t := i.(type){
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatIAm(true)
	whatIAm(2)
	whatIAm("hey")
}