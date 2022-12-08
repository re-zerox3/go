package main
 
import (
				"fmt"
			"math/rand"
	  	"time"
)


func prime(num int)(int){
				i:=2
				for i<= num/2 {
								if(num % i != 0){
												i++
								}else{
												return 0
								}
				}
				fmt.Printf("%d is PRIME!!!\n", num)
				return 1
}

func factor(num int){
				fmt.Printf("%d is NOT prime. factors = ",num)
				for i := 2; i<=num; i++{
								if(num %i ==0){
												fmt.Printf("%d ", i)
								}
				}
				fmt.Print("\n")
}

func main(){
				rand.Seed(time.Now().UnixNano())
				val :=0
				for val ==0 {
								randN := rand.Intn(10000000-2+1)+2
								val := prime(randN)
								if  val==0 {
										factor(randN)
								}else{
												break
								}
				}
}
























