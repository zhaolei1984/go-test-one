package gochan

import "fmt"

type JobInterface interface {
	StartJob()
	ProcessJob() error
	StopJob() error
}

type jobStruct struct {
	jobDoChan chan int
	jobDoneChan chan int
	maxNum int
}


func NewJob(num int) JobInterface  {
	return &jobStruct{
		jobDoChan:make(chan int, 10),
		jobDoneChan:make(chan int),
		maxNum:num,
	}
}

func DoJob()  {
	job:=NewJob(200)
	go job.StartJob()
	job.ProcessJob()
	job.StopJob()
}

func (j *jobStruct)StartJob () {
	for {
		select {
		case <-j.jobDoneChan:
			break
		case num:=<-j.jobDoChan:
			fmt.Println(num)
		}
	}
	fmt.Println("game over!")
}

func (j *jobStruct)ProcessJob () error {
	for i := 0; i < j.maxNum; i++ {
		j.jobDoChan <- i
	}
	return nil
}

func (j *jobStruct)StopJob () error{
	j.jobDoneChan <- 1
	return nil
}