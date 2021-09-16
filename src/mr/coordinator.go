package mr

import "log"
import "net"
import "os"
import "net/rpc"
import "net/http"


type Coordinator struct {
	// Your definitions here.

}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {   // 通信操作,变量改变
	reply.Y = args.X + 1
	return nil
}


//
// start a thread that listens for RPCs from worker.go  意思和tcp 差不多?（listen ）
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.

// main/mrcoordinator.go expects mr/coordinator.go to implement a Done() method that returns true 
//when the MapReduce job is completely finished; at that point, mrcoordinator.go will exit.
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.


	return ret
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
// The coordinator can't reliably distinguish between crashed workers, workers that are alive but have stalled for some reason, 
//and workers that are executing but too slowly to be useful. The best you can do is have the coordinator wait for some amount of time, 
//and then give up and re-issue the task to a different worker. 
//For this lab, have the coordinator wait for ten seconds; after that the coordinator should assume the worker has died (of course, it might not have).


func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.


	c.server()
	return &c
}
