package corntask

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/tiandidoushidali/dex-dodo/backend/app/sync/internal/service/corntask/etherum"
	"github.com/tiandidoushidali/dex-dodo/backend/app/sync/internal/service/corntask/solana"
	"github.com/zeromicro/go-zero/core/threading"
)

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

func (task *Task) process() {
	fmt.Print("in process")

	c := cron.New(
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),
			cron.DelayIfStillRunning(cron.DefaultLogger),
		),
		cron.WithParser(cron.NewParser(cron.SecondOptional)),
	)
	etherum.RegisterTask(c)
	solana.RegisterTask(c)

	c.Start()
}

func (task *Task) Run() {
	threading.GoSafe(task.process)
}
