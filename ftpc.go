package main

import "fmt"
import "os"
import "io/ioutil"
import "./ftp"
import "strconv"
import "runtime"

var wait = make (chan int)
var count int
func ftpc(num int) {

	numq := strconv.Itoa(num)
	//fmt.Println(num)

	//########  Store File ###########

	file := "Alfresco/User Homes/tapootum/"+"tapootum" + numq +".pdf"
	//file :="test/tapootum" + numq + ".avi"

	//#########  Read File  ############

	rFile := "test/tapootum" + numq + ".pdf"



    // new ftp
    ftp := new(ftp.FTP)
    // set debug, default false
    ftp.Debug = true
    // connect
    ftp.Connect("192.168.1.48", 21)
    // login
    ftp.Login("admin", "tapootum")
    // login failure
    if ftp.Code == 530 {
        fmt.Println("error: login failure")
        os.Exit(-1)
    }else {
    // pwd
    ftp.Pwd()
    fmt.Println("code:", ftp.Code, ", message:", ftp.Message)
    // mkdir new dir
    //ftp.Mkd("/smallfish")
    // stor new file
    ftp.Request("TYPE I")
    b, _ := ioutil.ReadFile(rFile)
    ftp.Stor(file, b)
    // quit
    ftp.Quit()

	fmt.Printf("FTP Send File complete ################################# %d\n",num)
	count++
	}
     //### channel ####
	wait <- 1

}

func ftprun(l int) {
	for i:=1 ; i<=l ; i++{
		go ftpc(i)
		fmt.Printf("############################################################ %d \n",i)
	}
	for i:= 1 ; i<=l ; i++{
		<- wait
	}
}




func main(){
	
	runtime.GOMAXPROCS(4)
	times1 , _ , _ := os.Time()
	if len(os.Args) == 1 {
	 fmt.Println("Don't Send File")
	}else {

        s := os.Args[1]
        n , _ := strconv.Atoi(s)

	ftprun(n)

/*	switch n {
	case  1 : ftprun(1)
	case  5 : ftprun(5)
	case  10: ftprun(10)
	case  100: ftprun(100)
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