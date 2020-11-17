/**
 * Created by cks
 * Date: 2020-11-17
 * Time: 11:24
 */
package task

type Task struct {
}

func New() *Task{
	return new(Task)
}

func (task *Task) Run() {
	
}