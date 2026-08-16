package main

import (
	"fmt"
	"runtime"
	"time"

	"golang.org/x/sys/unix"
)

func printThreadID() {
	for {
		tid := unix.Gettid()
		fmt.Printf("Before LockOSThread: TID=%d\n", tid)

		runtime.LockOSThread()

		tid = unix.Gettid()
		fmt.Printf("After LockOSThread: TID=%d\n", tid)

		time.Sleep(2 * time.Second)
	}
}

func threadID() {
	for i := 0; i < 10; i++ {
		tid := unix.Gettid()
		fmt.Printf("Unlocked | iteration=%d | TID=%d\n", i, tid)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("---- Locking Thread ----")

	runtime.LockOSThread()

	for i := 0; i < 10; i++ {
		tid := unix.Gettid()
		fmt.Printf("LOCKED | Iteration=%d | TID=%d\n", i, tid)

		time.Sleep(100 *time.Millisecond)
	}
}

func main() {
	go threadID()

	for {
		time.Sleep(time.Second)
	}
}
