package main
import(
	"fmt"
	//"net/rpc"
	"os"
	"net/rpc"
)

type arg struct {
	Key string
	Value string
}
func main() {
	pair:=new(arg)
	hello:=""
	address:="localhost:2905"
	client,err:=rpc.Dial("tcp",address)
if err==nil{
	if len(os.Args)>=2  {
		if os.Args[1]=="-p" || os.Args[1]=="-P" && len(os.Args)>=4 {
			pair.Key=os.Args[2]
			pair.Value=os.Args[3]
			err:=client.Call("Server.AddKey",pair,&hello)
			if err!=nil {
				fmt.Println(err)
				os.Exit(9)
			}
			if hello=="err"{
				os.Exit(10)
			}
			fmt.Println(hello)
		}else
		if os.Args[1]=="-d" || os.Args[1]=="-D" && len(os.Args)>=3{
			pair.Key=os.Args[2]
			pair.Value=os.Args[3]
			err:=client.Call("Server.Delete",pair,&hello)
			if err!=nil {
				fmt.Println(err)
				os.Exit(9)
			}
			if hello=="err" {
				os.Exit(10)
			}
		}else
		if os.Args[1]=="-g" || os.Args[1]=="-G" && len(os.Args)>=3 {
			pair.Key=os.Args[2]
			pair.Value=os.Args[3]
			err:=client.Call("Server.GetKey",pair,&hello)
			if err!=nil {
				fmt.Println(err)
				os.Exit(9)
			}
			if hello=="err" {
				os.Exit(10)
			}
			fmt.Println(hello)

		}else
		if os.Args[1]=="-u" || os.Args[1]=="-U"&& len(os.Args)>=4 {
				pair.Key=os.Args[2]
				pair.Value=os.Args[3]
				err:=client.Call("Server.Update",pair,&hello)
				if err!=nil {
					fmt.Println(err)
					os.Exit(9)
				}
				if hello=="err" {
					os.Exit(10)
				}
				fmt.Println(hello)
			}else{
				fmt.Println("------------Option Not found------------------")
				fmt.Println("ADD KEY 		-p Key 		value")
				fmt.Println("Update KEY 		-u Key 		Newvalue")
				fmt.Println("Delete KEY 		-d Key")
				fmt.Println("Inverse Map 		-i Key")
				fmt.Println("Get Key 		-g Key")
			}
	}else{
	fmt.Println("------------Option------------------")
	fmt.Println("ADD KEY 		-p Key 		value")
	fmt.Println("Update KEY 		-u Key 		Newvalue")
	fmt.Println("Delete KEY 		-d Key")
	fmt.Println("Inverse Map 		-i Key")
	fmt.Println("Get Key 		-g Key")
}
}else{
	fmt.Println(err)

}



}


