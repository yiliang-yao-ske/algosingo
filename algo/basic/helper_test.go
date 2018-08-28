//package helper
package main

import (
    "testing"
    "reflect"
    //"basic.helper"
)

func TestAverage(t *testing.T) {
    v := getheight(12)
    if reflect.TypeOf(v).String() != "int" {
        t.Error("Excepted 1.5, got ", v)
    }

}


// go test -v helper.go helper_test.go
