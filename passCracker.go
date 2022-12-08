package main

import(
				"fmt"
			//	"encoding/hex"
				"os"
				"log"
				"bufio"
				"strings"
				"crypto/sha1"
				"strconv"
)

func errExit(err error){
				if err != nil{
								log.Fatal(err)
				}
}

func openFile(fname string)[]string{
				lisWord := make([]string,0)
				file, err := os.Open(fname)
				errExit(err)
				scanner := bufio.NewScanner(file)
				scanner.Split(bufio.ScanLines)
				for scanner.Scan(){
								word := scanner.Text()
								word = strings.TrimSpace(word)
								lisWord = append(lisWord, word)
				}
				return lisWord
}

func workerG(id int,n int, work int,mutte int, ch chan bool, list []string, lPass *[]string){
				var end int
				start := id *work
				if (id +1)==n{
								end = start + work + mutte
				}else{
								end = start + work
				}
				
				for i := start; i< end; i++{
								for j:=0; j<len(list); j++{
												*lPass = append(*lPass,list[i]+list[j])
							  }
				}
				ch <- true
}

func findMatch(lisPass []string){

				found := false
				for i:=0; i<len(lisPass); i++{
							hash:=fmt.Sprintf("%x",sha1.Sum([]byte(lisPass[i])))
							if hash==os.Args[3]{
											found = true
											fmt.Printf("found: %s %s\n",lisPass[i],hash)
											break
						}		
				}
				if found ==false{
								fmt.Printf("Password not found after %d word combinations\n",len(lisPass))
				}
				
}




func main(){
				if len(os.Args) <4{
								fmt.Println("Usage:... ./passCracker fileName N SHA-1_hash")
								os.Exit(3)
				}			

				numG, err0 := strconv.Atoi(os.Args[2])
				errExit(err0)

				if numG> 100 || numG<1 {
								fmt.Println("N must be >=1 and <=100")
								os.Exit(3)
				}

				listW :=	openFile(os.Args[1])
				if len(listW)< numG{
								fmt.Println("Number of Job must be greater than N(workers)")
								os.Exit(0)
				}

				work := len(listW) / numG
				mutte := len(listW) % numG
				ch := make(chan bool)
				lisPass := make([]string,0)
			
				for i :=0; i<numG; i++{
								go workerG(i,numG, work,mutte, ch,listW, &lisPass)
				}
				
				for i:=0; i<numG; i++{
								<-ch
				}

				findMatch(lisPass)

}
