package main

import(
				"os"
				"fmt"
				"strings"
				"bufio"
				"log"
				"math/rand"
				"strconv"
				"crypto/sha1"

)
func errExit(err error){
				if err != nil{
								log.Fatal(err)
				}
}

func openF(fname string)[]string{
		file, err := os.Open(fname)
		errExit(err)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var lisWord []string
		lisWord = make([]string,0)
		for scanner.Scan(){
						line := scanner.Text()
						line = strings.TrimSpace(line)
						words := strings.Split(line, " ")
						for i:=0; i<len(words); i++{
										lisWord = append(lisWord, words[i])
						}
		}
		return lisWord
}

func genWords(lisWord []string)string{
				var words string
					genWord:= rand.Intn(len(lisWord))
					words = strings.Title(lisWord[genWord])
					for true{
						genWord = rand.Intn(len(lisWord))
						value := strings.Title(lisWord[genWord])
						if words != value{
										words += value
										break
						}
		}
return words
}

func main(){
				args := os.Args
				if len(args)<3{
								fmt.Println("Usage:... ./passwordGen file seed")
					os.Exit(3)
				}
				seed,err:= strconv.ParseInt(args[2],10,64)
				errExit(err)
				rand.Seed(seed)
				lisWord := openF(args[1]) 
				words := genWords(lisWord)
				num := rand.Intn(98765-12345)+12345
				numbers:= strconv.Itoa(num)
				var runes [8]string = [8]string{"!","@","#","$","%","^","&","*"}  
				randRune := runes[rand.Intn(len(runes))]
				password := words+numbers+randRune
				hashVal := sha1.Sum([]byte(password))
				fmt.Printf("password: %s\n",password)
				fmt.Printf("SH1 password hash: %x\n",hashVal)
} 
