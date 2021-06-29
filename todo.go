package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item Struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

// note: upper case type will be visible outside of the package
type List []item

func (l *List) Add(task string) {
	t := item {
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: timeTime{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l

	// Error handler
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	// Add to a slice of everything before the target index, everything after
	// thereby removing the target
	*l = append(ls[:i - 1], ls[i:]...)
}