package inter

import (
	"fmt"
	"strconv"
	"sync"
)

func StartThis() {
	//var numbers []int32
	/*numbers := []int32{6, 2, 4, 10}
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	minDiff := numbers[0] - numbers[1]
	if minDiff < 0 {
		minDiff = minDiff * -1
	}
	fmt.Println("minDiff=", minDiff)
	for i := range numbers {
		//fmt.Println(numbers[i])
		if i < len(numbers)-1 {
			diff := numbers[i+1] - numbers[i]
			if diff == minDiff {
				fmt.Println(numbers[i], numbers[i+1])
			}
		}
	}*/
}

func Prob1() {
	/*fmt.Println("in Prob1")
	// Declaring some variables
	var inpu1 int
	var input2 int

	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scanf("%d", &inpu1)
	fmt.Scanf("%d", &input2)

	fmt.Printf("The 1= %d 2= %d ", inpu1, input2)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")

	text := ""
	if t, _ := reader.ReadString('\n'); true {
		text = strings.Trim(t, " \n")
	}
	fmt.Println("text=", text)

	split := strings.Split(text, " ")

	var arr1 []int
	for _, el := range split {
		myInt, _ := strconv.Atoi(el)
		arr1 = append(arr1, myInt)
	}

	fmt.Println("split", split)
	fmt.Println("arr1", arr1)*/

}

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (n *NumericInput) Add(r rune) {
	digit := ""
	for i := 48; i <= 57; i++ {
		if int(r) == i {
			digit = strconv.Itoa(int(r) - 48)

			n.input += digit
			fmt.Println("digit=", digit)
		}
	}
}

func (n *NumericInput) GetValue() string {
	return n.input
}

func V1() {
	fmt.Println("in Prob1")
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}

const maxWorkers = 4

var workersInfo map[int]int
var mapMutex = sync.RWMutex{}

type job struct {
	name string
	days int
}

func doWork(id int, j job) {
	mapMutex.Lock()
	if _, ok := workersInfo[id]; ok {
		workersInfo[id] += 1
	} else {
		workersInfo[id] = 1
	}
	mapMutex.Unlock()

	//fmt.Printf("worker%d: started %s, working for %d\n", id, j.name, j.days)
	fmt.Println(getLog(j.name, j.days))
	//fmt.Printf("worker%d: completed %s!\n", id, j.name)
}

func V2() {
	workersInfo = make(map[int]int)
	//map[string]int = { "Robert" : 30, "John" : 475, "Elly" : 1022, "Bob" : 99 }
	data := make(map[string]int)
	data["Robert"] = 30
	data["John"] = 475
	data["Elly"] = 1022
	data["Bob"] = 99

	//ConvertDays(data)

	// channel for jobs
	jobs := make(chan job)

	// start workers
	wg := &sync.WaitGroup{}
	wg.Add(maxWorkers)
	for i := 1; i <= maxWorkers; i++ {
		go func(i int) {
			defer wg.Done()

			for j := range jobs {
				doWork(i, j)
			}
		}(i)
	}

	// add jobs
	for key, val := range data {
		name := fmt.Sprintf("%s", key)
		//fmt.Printf("adding: %s %s\n", name, val)
		jobs <- job{name, val}
	}
	close(jobs)

	// wait for workers to complete
	wg.Wait()

	fmt.Printf("\nInfo:\n")
	fmt.Printf("Workers Count: %d\n", len(workersInfo))
	for key, val := range workersInfo {
		fmt.Printf("Worker#%d -> %d elements processed\n", key, val)
	}

}

func getLog(name string, days int) string {
	weeks := days / 7
	days = days % 7

	switch days {
	case 0:
		return fmt.Sprintf("%s ha trabajando en la empresa %d semanas", name, weeks)
	case 1:
		return fmt.Sprintf("%s ha trabajando en la empresa %d semanas y %d día", name, weeks, days)
	default:
		return fmt.Sprintf("%s ha trabajando en la empresa %d semanas y %d días", name, weeks, days)
	}
}


