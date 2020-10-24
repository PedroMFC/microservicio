// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package main

import (
	//"fmt"
	//"time"

	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func main() {

	// create redis connection pool
redisPool := &redis.Pool{
	Dial: func() (redis.Conn, error) {
		  c, err := redis.DialURL("redis://")
		  if err != nil {
			  return nil, err
		  }
		  return c, err
	  },
  }
  
  // initialize celery client
  cli, _ := gocelery.NewCeleryClient(
	  gocelery.NewRedisBroker(redisPool),
	  &gocelery.RedisCeleryBackend{Pool: redisPool},
	  5, // number of workers
  )
  
  // task
  add := func(a, b int) int {
	  return a + b
  }
  
  // register task
  cli.Register("worker.add", add)
  
  // start workers (non-blocking call)
  cli.StartWorker()
  
  // wait for client request
  //time.Sleep(10 * time.Second)
  
  <-make(chan int)

  // stop workers gracefully (blocking call)
  cli.StopWorker()
}
