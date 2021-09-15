package mr
//都是标准库
import "fmt"
import "log"
import "net/rpc"  
import "hash/fnv"


//
// Map functions return a slice of KeyValue.
//
type KeyValue struct {
	Key   string
	Value string
}

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}


//
// main/mrworker.go calls this function.
//
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {  //worker接受map和reduce函数。
	//map 实现 可参考 steal some code from mrsequential.go 
	//for reading Map input files, for sorting intermedate key/value pairs between the Map and Reduce, and for storing Reduce output in files.
	// intermediate files is mr-X-Y, where X is the Map task number, and Y is the reduce task number. (得到的中间文件保存命名)
	// The worker's map task code will need a way to store intermediate key/value pairs in files in a way that can be correctly read back during reduce tasks. 
	// One possibility is to use Go's encoding/json package. To write key/value pairs to a JSON file:保存格式
	
	// reduce 实现
	// Workers will sometimes need to wait,One possibility is for workers to periodically ask the coordinator for work, sleeping with time.Sleep() between each request
	// 暂时不懂： Another possibility is for the relevant RPC handler in the coordinator to have a loop that waits, either with time.Sleep() or sync.Cond. Go runs the handler for each RPC in its own thread, so the fact that one handler is waiting won't prevent the coordinator from processing other RPCs.
	// Your worker implementation here.

	// uncomment to send the Example RPC to the coordinator.（hint1修改）
	// CallExample()

}

//
// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
//
func CallExample() {
// rpc 中定义的结构体
	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	call("Coordinator.Example", &args, &reply)

	// reply.Y should be 100.
	fmt.Printf("reply.Y %v\n", reply.Y)
}

//
// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")  // 都是标准库
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
