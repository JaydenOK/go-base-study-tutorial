package pool

import (
	"fmt"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	pool := New(2)
	defer Stop()

	if pool == nil {
		t.Fail()
	}
}

func TestRemoveWorker(t *testing.T) {
	pool := New(2)
	defer Stop()

	worker := func(i int) {}

	for i := 1; i <= 2; i++ {
		if err := AddWorker(worker); err != nil {
			t.Fatal(err)
		}
	}

	if err := RemoveWorker(worker); err != nil {
		t.Fatal(err)
	}

	if WorkersNum() != 1 {
		t.Fatal("should have one worker left")
	}

	if err := Delegate(1); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidWorker(t *testing.T) {
	pool := New(2)
	defer Stop()

	if AddWorker("worker") == nil {
		t.Fail()
	}
}

func TestWorkers(t *testing.T) {
	delegateWorkToWorkers(t, 2, 3, 3) // same workers as jobs
	delegateWorkToWorkers(t, 2, 3, 2) // less workers then jobs
	delegateWorkToWorkers(t, 2, 3, 5) // more workers than jobs
}

func delegateWorkToWorkers(t *testing.T, poolSize int, jobsAmount int, workersAmount int) {
	var wg sync.WaitGroup

	pool := New(poolSize)
	out := make(chan int, jobsAmount)
	worker := func(i int) {
		defer wg.Done()
		out <- i
	}

	for i := 1; i <= workersAmount; i++ {
		if err := AddWorker(worker); err != nil {
			t.Fatal(err)
		}
	}

	wg.Add(jobsAmount)

	for i := 1; i <= jobsAmount; i++ {
		if err := Delegate(i); err != nil {
			t.Fatal(err)
		}
	}

	go func() {
		wg.Wait()
		close(out)
		Stop()
	}()

	sum := 0
	for n := range out {
		sum += n
	}

	if sum == 0 {
		fmt.Printf("Delegating job %d to %d workers failed", jobsAmount, workersAmount)
		t.Fail()
	}
}
