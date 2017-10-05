package core

import (
	"fmt"
	"github.com/GoCollaborate/artifacts/task"
	"github.com/GoCollaborate/wrappers/taskHelper"
	"net/http"
)

func ExampleJobHandler01(w http.ResponseWriter, r *http.Request) *task.Job {
	job := task.MakeJob()
	job.Tasks(&task.Task{task.SHORT,
		task.BASE, "exampleFunc",
		[]task.Countable{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
		[]task.Countable{0},
		task.NewTaskContext(struct{}{}), 0})
	// map once, reduce once, repeat
	job.Stacks("core.ExampleTask.SimpleMapper",
		"core.ExampleTask.SimpleReducer", "core.ExampleTask.SimpleMapper",
		"core.ExampleTask.SimpleReducer")
	return job
}

func ExampleFunc(source *[]task.Countable,
	result *[]task.Countable,
	context *task.TaskContext) chan bool {
	out := make(chan bool)
	// deal with passed in request
	go func() {
		fmt.Println("Example Task Executed...")
		var total int
		for _, n := range *source {
			total += n.(int)
		}
		*result = append(*result, total)
		out <- true
	}()
	return out
}

type SimpleMapper int

func (m *SimpleMapper) Map(inmaps map[int]*task.Task) (map[int]*task.Task, error) {
	return taskHelper.Slice(inmaps, 3), nil
}

type SimpleReducer int

func (r *SimpleReducer) Reduce(maps map[int]*task.Task) (map[int]*task.Task, error) {
	var sum int
	for _, s := range maps {
		for _, r := range (*s).Result {
			sum += r.(int)
		}
	}
	fmt.Printf("The sum of numbers is: %v \n\n", sum)
	fmt.Printf("The task set is: %v", maps)
	return maps, nil
}
