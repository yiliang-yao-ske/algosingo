package main


import (
    "fmt"
    "reflect"
    "strings"
)



func findLongestSubstring(s string)(find_s string){

    var i int = 1

    n := len(s)
    fmt.Println(n)
    var currlen int

    var st int = 0

    var maxlen int = 0

    var start int

    pos := make(map[string]int)

    sArr := strings.Fields(s)
    //fmt.Println(sArr)
    //fmt.Println(s[0])
    //fmt.Println(string(s[0]))

    pos[string(s[0])] = 0
    //fmt.Println(pos)


    for i = 1; i < n; i++ {
        _, is_exist := pos[string(s[i])]
        //fmt.Println(is_exist)
        //fmt.Println(pos)
        if !is_exist {
            pos[string(s[i])] = i
        }else{
            if pos[string(s[i])] >= st {
                currlen = i - st

                if maxlen < currlen {
                    maxlen = currlen
                    start = st
                }
                st = pos[string(s[i])] + 1
            }

            pos[string(s[i])] = i
        }

    }

    if maxlen < (i - st) {
        maxlen = i - st
        start = st
    }

    //fmt.Println(start)
    //fmt.Println(maxlen)
    return s[start:start+maxlen]

}


func main(){

    s := "GEEKSFORGEEKS"

    find_s := findLongestSubstring(s)

    fmt.Println(s, reflect.TypeOf(s))
    fmt.Println(find_s)
}
