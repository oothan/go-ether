package main

import (
	"fmt"
	"time"
)

func main() {
	layoutISO := "2006-01-02"
	tt := time.Now()
	fmt.Println( tt.Format(layoutISO))
	//t, _ := time.Parse(layoutISO, )
	//createdAt := fmt.Sprintf("date(created_at) = date(\"%v\")", t.Format(layoutISO))

}
