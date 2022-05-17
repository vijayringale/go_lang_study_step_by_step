package main
import (
	"fmt"
	"time"
)

func main(){

	i := 2
	fmt.Print("write",i,"as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	whatAmI :=func(i interface{}){
		switch t:=i.(type){
		case bool:
			fmt.Println("i am bool")

		case int:
			fmt.Println("I m an int")
		default:
			fmt.Printf("Don't knoe type %T\n ",t)

		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}