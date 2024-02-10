package main

import "fmt"
import "bufio"
import "os"

func ReadFile() string{
	var str string
	f,_:=os.Open("input.txt")
	defer f.Close()
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		str+=scanner.Text()
	}
	return str
}

func CheckFloor(s string){
	var floor int
	for i,c:=range(s){
		if c=='('{
			floor++
		}else{
			floor--
		}
	}
	fmt.Println("Floor ->",floor)
}

func main(){
	lettura:=ReadFile()
	CheckFloor(lettura)
}