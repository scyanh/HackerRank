package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

var numRequestsToMake int
var numConcurrentRequests int

func init() {
	flag.IntVar(&numRequestsToMake, "total-requests", 30, "total # of requests to make")
	flag.IntVar(&numConcurrentRequests, "concurrent-requests", 30, "pool size, request concurrency")
}

type Counter struct {
	count int
}

func (c *Counter) Inc() {
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

func (c *Counter) Set(v int) {
	c.count = v
}

type ValuedCounter struct {
	mu    *sync.Mutex
	count int
}

func (c *ValuedCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func (c *ValuedCounter) Set(v int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count = v
}

func TestExplicitRace(t *testing.T) {
	flag.Parse()

	reqCount := Counter{}

	go func() {
		http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			value := reqCount.Value()
			fmt.Printf("handling request: %d\n", value)
			time.Sleep(1 * time.Nanosecond)
			reqCount.Set(value + 1)
			fmt.Fprintln(w, "Hello, client")
		}))
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	requestsChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(numConcurrentRequests)

	// start a pool of 100 workers all making requests
	for i := 0; i < numConcurrentRequests; i++ {
		go func() {
			defer wg.Done()
			for range requestsChan {
				res, err := http.Get("http://localhost:8080/")
				if err != nil {
					t.Fatal(err)
				}
				_, err = ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					t.Error(err)
				}
			}
		}()
	}

	go func() {
		for i := 0; i < numRequestsToMake; i++ {
			requestsChan <- i
		}
		close(requestsChan)
	}()

	wg.Wait()

	fmt.Printf("Num Requests TO Make: %d\n", numRequestsToMake)
	fmt.Printf("Num Handled: %d\n", reqCount.Value())
	if numRequestsToMake != reqCount.Value() {
		t.Errorf("expected %d requests: received %d", numRequestsToMake, reqCount.Value())
	}
}

/*func TestWhyDoesntWork2(t *testing.T) {
	var mapMutex = sync.RWMutex{}
	data:=make(map[int]string)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(w *sync.WaitGroup) {
		defer w.Done()
		mapMutex.Lock()
		data[0] = "First entry"
		mapMutex.Unlock()
	}(&wg)

	wg.Add(1)
	go func(w *sync.WaitGroup) {
		defer w.Done()
		mapMutex.Lock()
		data[1] = "Second entry"
		mapMutex.Unlock()
	}(&wg)

	wg.Add(1)
	go func(w *sync.WaitGroup) {
		defer w.Done()
		time.Sleep(1 * time.Second)
		fmt.Println(data[0], data[1])
	}(&wg)

	wg.Wait()
}*/

/*func Test1SortAscending(t *testing.T) {
	arr := []int{2, 1, 3}
	expected := []int{1, 2, 3}
	t.Logf("initial arr= %d", arr)

	ascArr:=SortAscending(arr)
	t.Logf("sort asc= %d", ascArr)
	if !reflect.DeepEqual(ascArr, expected) {
		t.Errorf("expected sort asc %d but got %d", expected, ascArr)
	}else{
		t.Logf("sorting ok")
	}

}

func Test2SortAscending(t *testing.T) {
	arr := []int{12, 6, 3}
	expected := []int{3, 6, 12}
	t.Logf("initial arr= %d", arr)

	ascArr:=SortAscending(arr)
	t.Logf("sort asc= %d", ascArr)
	if !reflect.DeepEqual(ascArr, expected) {
		t.Errorf("expected sort asc %d but got %d", expected, ascArr)
	}else{
		t.Logf("sorting ok")
	}
}

func Test3SortAscending(t *testing.T) {
	arr := []int{6, 3, 9}
	expected := []int{3, 6, 9}
	t.Logf("initial arr= %d", arr)

	ascArr:=SortAscending(arr)
	t.Logf("sort asc= %d", ascArr)
	if !reflect.DeepEqual(ascArr, expected) {
		t.Errorf("expected sort asc %d but got %d", expected, ascArr)
	}else{
		t.Logf("sorting ok")
	}
}*/

