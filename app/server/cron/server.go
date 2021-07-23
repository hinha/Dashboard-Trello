package cron_server

import (
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

const AppName = "trelloProject"

type Server struct {
	Logger *log.Entry
	worker *work.WorkerPool
	pool   *redis.Pool
}

var batchTime = "0 */2 * * * *"

func New(logger *log.Entry) *Server {
	s := &Server{
		Logger: logger,
	}

	s.pool = pool()

	worker := work.NewWorkerPool(struct{}{}, 3, AppName, s.pool)

	fh := trelloJobHandler{logger: logger}
	worker.PeriodicallyEnqueue(batchTime, fh.name()).JobWithOptions(fh.name(), work.JobOptions{
		Priority: 1,
		MaxFails: 3,
	}, func(c *work.Job) error {
		err := fh.job(c)
		if err != nil {
			logger.Error(err)
		}
		return nil
	}).Middleware(fh.middleware)

	s.worker = worker
	return s
}

func (s *Server) Start() {
	s.worker.Start()
}

func (s *Server) Stop() {
	s.worker.Stop()
	s.deleteNameSpace()
}

func (s *Server) deleteNameSpace() {
	conn := s.pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	iter := 0
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", AppName+":*"))
		if err != nil {
			panic(err)
		}

		iter, _ = redis.Int(arr[0], nil)
		key, _ := redis.Strings(arr[1], nil)

		for _, k := range key {
			_, err := conn.Do("DEL", k)
			if err != nil {
				break
			}
		}

		if iter == 0 {
			break
		}
	}

}
