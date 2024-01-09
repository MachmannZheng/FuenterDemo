package synctest

import (
	"fmt"
	"sync"
)

type User struct {
	depositAmout float32
	sync.Mutex
}

var BALANCE float32 = 12000

func SynctestRaceCondition() {
	fmt.Println("-- execute sync test Race Condition function --")
	var wg sync.WaitGroup
	user := User{depositAmout: 1200}
	for i := 0; i < 10; i++ {
		wg.Add(3)
		go user.deposit(&wg)
		go user.deposit(&wg)
		go user.deposit(&wg)
	}
}

func (u *User) deposit(wg *sync.WaitGroup) {
	u.Lock()
	BALANCE += u.depositAmout
	u.Unlock()
	wg.Done()
}
