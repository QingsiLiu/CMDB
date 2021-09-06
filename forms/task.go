package forms

type TaskForm struct {
	ID           int    `form:"id"`
	Name         string `form:"name"`
	StartTime    string `form:"start_time"`
	CompleteTime string `form:"complete_time"`
	DeadlineTime string `form:"deadline_time"`
	Status       string `form:"status"`
	Content      string `form:"content"`
	User         int    `form:"user"`
}

func (t *TaskForm) StatusInt() int {
	switch t.Status {
	case "新建":
		return 0
	case "开始":
		return 1
	case "暂停":
		return 2
	case "完成":
		return 3
	}
	return 4
}
