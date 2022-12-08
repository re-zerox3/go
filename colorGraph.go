package main


import (
				"fmt"
				"os"
				"strings"
				"bufio"
				"log"
				"sort"
)


type Graph struct{
				vertices []*Vertex
}


type Vertex struct{
				color string
				num int
				adjacent []*Vertex
}


func (g *Graph)addVertex(k string){
				if containVertex(g.vertices, k) == false{
						g.vertices= append(g.vertices,&Vertex{color:k})
				}
}

func (g *Graph)addEdge(from, to string){
				fromV := g.getVertex(from)
				toV :=g.getVertex(to)
				state := g.containEdge(fromV, toV)
				if state==false{
						fromV.adjacent=append(fromV.adjacent, toV)
				}
}

func (g *Graph)containEdge(fromV, toV *Vertex)bool{
				for  v:=0; v<len(fromV.adjacent); v++{
							if  fromV.adjacent[v] == toV{			
									return true
							}
			}
		return false
}

func (g *Graph)getVertex(k string)*Vertex{
				for i, v := range g.vertices{
								if v.color==k {
												return g.vertices[i]
								}
				}
				return nil
}

func containVertex(s []*Vertex, k string)bool{
				for _, v:=range s{
								if k == v.color{
												return true
								}
				}
				return false
}

func errExit(err error){
				if err != nil{
								log.Fatal(err)
				}
}

func (g *Graph)openFile(fName string)[][]string{
			 nodes:= make([][]string,0)
				file, err := os.Open(fName)
				errExit(err)
				scanner := bufio.NewScanner(file)
				scanner.Split(bufio.ScanLines)
				for scanner.Scan(){
								line := scanner.Text()
								words :=strings.Split(line,"->")
								nodes = append(nodes, words)
								for i:=0; i<len(words); i++{
												words[i] = strings.TrimSpace(words[i])
											g.addVertex(words[i])
								}
				}
				return nodes
}

func (g *Graph)getGraph()map[string][]string{
			color := make(map[string][]string)
				for _,v:=range g.vertices{
							//	fmt.Println("valll: ", v.color)
								adjVal:=make([]string,0)
					for _, v := range v.adjacent{
								//	fmt.Println(v.color)
												adjVal = append(adjVal, v.color)
								}
								color[v.color] = adjVal
				}
			return color
}

func (g *Graph)combo(){
				colors :=g.getGraph()
				combo := make(map[string]int)
				for k,v :=range colors{
								if len(v)!=0{
									keySplit := strings.Split(k,"_")
									key :=keySplit[0]
									for i:=0;i<len(v);i++ {
												val := strings.Split(v[i],"_")
												_, ok := combo[key+"->"+val[0]]
												if ok {
																combo[key+"->"+val[0]] +=1
												}else{
																combo[key+"->"+val[0]] = 1
												}
									}

								}
				}
			
				g.sortOutput(combo)
}

func (g *Graph)sortOutput(combo1 map[string]int){

				colors := g.getGraph()
				keys:=make([]string,0) 
								fmt.Println("-------------Just Graph---------------")
								
								for k := range colors {
												keys = append(keys, k)

								}
								sort.Strings(keys)
								for _, k := range keys{
												fmt.Println(k)
												sort.Strings(colors[k])
												colorVal:= colors[k]
												for i:=0; i<len(colorVal); i++{
																fmt.Printf("\t==> %s\n",colorVal[i])
												}
								}
					keys = make([]string,0)
					fmt.Println("\n=============COMBOS=================")
					for k:= range combo1{
									keys = append(keys, k)
					}
					sort.Strings(keys)
					for i:=0; i<len(keys); i++{
								fmt.Printf("\t%s = %d\n", keys[i],combo1[keys[i]])
					}
	
}

func main(){
				args := os.Args
				if len(args)<2{
								fmt.Println("Usage:... ./colorGraph fileName")
								os.Exit(3)
				}

				fileName := args[1]
				colorGraph := &Graph{}
				
				edges := colorGraph.openFile(fileName)

				for i:=0; i<len(edges); i++{
								colorGraph.addEdge(edges[i][0],edges[i][1])
	 			
				}
				colorGraph.combo()
			
}
