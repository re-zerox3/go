package hw11

import (
  "fmt"
  "log"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

type Student struct{
	id int
	name string
	age int
	gpa float64
}

func check(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

//CODE GOES HERE!
func Insert(db *sql.DB,name string,age int,gpa float64){
	
	db, err0 := sql.Open("sqlite3","./hw11.sqlite")
	check(err0)
	defer db.Close()
	tx, _:=db.Begin()
	cmd := "INSERT INTO students"+"(name,age, gpa)"+"VALUES"+"(?,?,?)"
	stmt,err1 := tx.Prepare(cmd)
	defer stmt.Close()
	check(err1)
	stmt.Exec(name, age, gpa)
	tx.Commit()
}

func SelectAll(db *sql.DB)[]Student{

	db, err0 := sql.Open("sqlite3","./hw11.sqlite")
	check(err0)
	defer db.Close()
	rows, err1 := db.Query("SELECT * FROM students")
	check(err1)
	defer rows.Close()
	studentL := make([]Student,0)
	for  rows.Next(){
		var id int
		var name string
		var age int
		var gpa float64
		var student Student
		err2:= rows.Scan(&id,&name,&age,&gpa)
		check(err2)
		student=Student{id:id,name:name,age:age,gpa:gpa}
		studentL = append(studentL,student)
	}
	return studentL
}

func Select(db *sql.DB, column string, operator rune, value interface{})[]Student{
	db, err0 := sql.Open("sqlite3","./hw11.sqlite")
	check(err0)
	defer db.Close()
	synm := fmt.Sprintf("%c", operator)
	rows, err1:= db.Query("SELECT * FROM students WHERE "+ column + synm +"$3", value)
	check(err1)
	defer rows.Close()
	studentL := make([]Student,0)
	for rows.Next(){
		var id int
		var name string
		var age int
		var gpa float64
		var student Student
		rows.Scan(&id,&name,&age,&gpa)
		student = Student{id:id, name:name, age:age, gpa:gpa}
		studentL = append(studentL,student)
	}
	return studentL
}

func (s *Student)Print(){
	fmt.Printf("%d %s %d %.1f\n",s.id, s.name, s.age, s.gpa)
}

func (s *Student)Delete(db *sql.DB){
	db, err := sql.Open("sqlite3","./hw11.sqlite")
	check(err)
	defer db.Close()
	cmd := "DELETE FROM students WHERE id = ?"
	db.Exec(cmd,s.id)
}

func (s *Student)Update(db *sql.DB, column string, value interface{}){
	db, err := sql.Open("sqlite3","./hw11.sqlite")
	check(err)
	defer db.Close()
	cmd:="UPDATE students "+"SET "+ column +" = ? WHERE id = ?"
	db.Exec(cmd,value, s.id)
}
