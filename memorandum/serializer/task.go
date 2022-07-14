package serializer

import "todo_list/model"

type Task struct {
	ID        uint   `json:"id"example:"1"`
	Title     string `json:"title"example:""`
	Content   string `json:"content"example:""`
	View      uint64 `json:"view"example:""`
	Status    string `json:"status"example:""`
	CreatedAt int64  `json:"created_at"example:""`
	StartTime int64  `json:"start_time"example:""`
	EndTime   int64  `json:"end_time"example:""`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:      item.ID,
		Title:   item.Title,
		Content: item.Content,
		//View:      item.View(),
		Status:    item.Status,
		CreatedAt: item.CreatedAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}

}

func BuildTaskAll(item []model.Task) []Task {
	var tasklist []Task
	for _, v := range item {
		tasklist = append(tasklist, BuildTask(v))
	}
	return tasklist

}
