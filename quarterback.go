package main

import (
				"fmt"
				"os"
				"io"
				"bufio"
				"log"
				"strings"
				"strconv"
)
func errExit(err error){
				if err !=nil{
								log.Fatal(err)
				}
}

func createFile(fname string){
				file, err := os.Create(fname)
				errExit(err)
				writer := bufio.NewWriter(file)
				writer.WriteString("hello")
				writer.Flush()
				file.Close()
}

func checkVal(valArray [4]float64)[4]float64{
				for i:=0; i<len(valArray); i++{
				if valArray[i] <=0{
								valArray[i] =0
				}else if valArray[i] >=2.375 {
								valArray[i] = 2.375
				}
}
return valArray
}

func rating(name string, yds float64, att float64, comp float64, td float64, intr float64){
				a := ((comp/att) -0.3) * 5
				b :=((yds/att) -3) * 0.25
				c := (td/att)*20
				d := 2.375 - ((intr/att) *25)
				var valArr = [4]float64{a,b,c,d}
				valArr = checkVal(valArr)
				total := ((valArr[0]+valArr[1]+valArr[2]+valArr[3])/6)*100
				fmt.Printf("%-20s: %.1f\n",name, total)
}

func readFile(fname string){
				file, err := os.Open(fname)
				errExit(err)
				row1, err:= bufio.NewReader(file).ReadSlice('\n')
				_, err = file.Seek(int64(len(row1)), io.SeekStart)
				errExit(err)
				scanner := bufio.NewScanner(file)
				scanner.Split(bufio.ScanLines)
				var valArray[5] float64
				for scanner.Scan(){
								word := scanner.Text()
								lisVal := strings.Split(word, ",")
								for i:=1; i<len(lisVal); i++ {
												var val float64
												val, err := strconv.ParseFloat(lisVal[i], 64 )
												errExit(err)
												valArray[i-1] = val
								}
							  rating(lisVal[0],valArray[0],valArray[1],valArray[2],valArray[3],valArray[4])
						}
				file.Close()

}

func main(){
				var fileName string
				fmt.Print("Enter file name: ")
				fmt.Scan(&fileName)
				readFile(fileName)

}
