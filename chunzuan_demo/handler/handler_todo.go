package handler

import (
	"golang.org/x/net/context"

	"git.rd.rijin.com/chunzuan_demo/infrastructure"
	"git.rd.rijin.com/chunzuan_demo/models"
	"git.rd.rijin.com/chunzuan_demo/proto"
)

type Todo struct {
}

func (p *Todo) AddTask(ctx context.Context, req *service_todo.AddTaskReq, resp *service_todo.Null) (err error) {
	task := &models.Task{
		Id:          req.Id,
		OwnerId:     req.OwnerId,
		Title:       req.Title,
		Description: req.Description,
	}

	if _, err = infrastructure.Engine.Insert(task); err != nil {
		return err
	}
	return err
}
