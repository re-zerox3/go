package  main

import(
				"fmt"
				"os"
				"strconv"
				"log"
				"sort"
)


func errExit(err error){
				if err != nil{
								log.Fatal(err)
				}
}


func fibonacci(nVal float64, fibMap map[float64]float64)float64{
				val, ok :=fibMap[nVal]
				if(nVal ==0){
								return 0
				}else if nVal==1{
								fibMap[nVal] = 1
								return 1
				}else if ok == true{
								return val

				}else{
								value := fibonacci(nVal-1, fibMap) + fibonacci(nVal-2, fibMap)
								fibMap[nVal] = value
								return value
				}
}

func main(){
				args := os.Args
				if len(args)<2{
								fmt.Println("Usage:... ./fibonacci N")
								os.Exit(3)
				}
				
				nValue, err := strconv.ParseInt(args[1], 10, 64)
				errExit(err)
				fibMap := make(map[float64]float64)
				fmt.Printf("fibonacci(%d) = %.0f\n",nValue, fibonacci(float64(nValue),fibMap))
				keys := make([]float64, 0, len(fibMap))
				for k := range fibMap{
								keys = append(keys, k)
				}
				sort.Float64s(keys)
			
				for _, k := range keys{
								if float64(nValue) != k{
												fmt.Printf("fibonacci(%d) = %.0f\n",int(k), fibMap[k])
								}
				}

}
