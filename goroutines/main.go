package main

import (
	"fmt"
	"time"
)

func main() {
	done:=make(chan bool)
   go greet("world")
   go greet("Ali")

   go SleepAwake(done)
   go greet("Rukhsar")
   <-done
   
}
func greet(msg string)  {
	fmt.Println("Hello", msg)
}

func SleepAwake(doneChane chan bool)  {
	fmt.Println("Sleeping for 5 seconds...")
	 time.Sleep(5 * time.Second)
	 fmt.Println("Awake now")
	 doneChane <- true
}