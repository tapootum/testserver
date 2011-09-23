package main

import "fmt"
import "io"
import "os"
import "strconv"

var wait = make (chan int)
var count int
func CopyFile(num int) {
	numq := strconv.Itoa(num)


	//fmt.Println(num)
	file := "/Volumes/tapootum-1/"+"tapootum" + numq +".mov"
	//file :="tapootum" + numq + ".mov"
	oFile := "../123.mov"


	dst, _ := os.Create(file)
        src, _ := os.Open(oFile)

        if _, err := io.Copy(dst, src); err != nil {
                fmt.Printf("Error: %s\n", err)
        } else {
		count++
                fmt.Printf("Copy successful %d\n",num)
        }

	wait <- 1
}


func copy1(l int){
        //x:=1
	for i:=1 ; i<=l ; i++{
       		go CopyFile(i)
	fmt.Println("###############################################",i)
        }
	for i:=1 ; i<=l ; i++{
		<- wait
	}
}


func copy5(){
        x:=1
//for i:=1 ; i<=n ; i++{

                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)

        <- wait
        <- wait
        <- wait
        <- wait
        <- wait
//        }
}


func copy10(){
	x:=1
//for i:=1 ; i<=n ; i++{

                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        x++
                go CopyFile(x)
                        

        <- wait
        <- wait
        <- wait
        <- wait
        <- wait
        <- wait
        <- wait
        <- wait
        <- wait
        <- wait
//	}
}

func main() {
	times1 , _ , _ := os.Time()
	 if len(os.Args) == 1 {
                fmt.Println("Don't Copy File")
        }else {

        s := os.Args[1]
        n , _ := strconv.Atoi(s)

	copy1(n)
/*
	switch n {
	case  1 : copy1()
	case  5 : copy5()
	case  10: copy10()
	}
*/
}
	times2 , _ , _ := os.Time()

	tt := times2-times1

	fmt.Println(count)
	//m := tt/60
	//s := tt%60 
	fmt.Printf(".............%d s...........\n",tt)
}
