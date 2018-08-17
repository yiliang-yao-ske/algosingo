package main

import (
    "fmt"
    "bufio"
    "os"
)


func main(){

//    rp, _ := io.Pipe()
//    //fmt.Println(rp)
//
//    data := make([]byte, 64)
//    n, err := rp.Read(data)
//
//    if nil != err {
//        fmt.Println(err)
//    }
//
//    fmt.Println("data", n , string(data))

    reader := bufio.NewReader(os.Stdin)


    data := make([]byte, 64)

    n, err := reader.Read(data)

    if nil != err {
        fmt.Println(err)
    }

    fmt.Println("data", n , string(data))

    fmt.Printf("%s\n", string(data[6]))

}

