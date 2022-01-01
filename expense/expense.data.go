package expense

import (
	"sync"
)


var expenseMap = struct {
	sync.RWMutex
	m map[int]Expense
}{m: make(map[int]Expense)}
