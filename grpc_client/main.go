package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_client/source/rpc"
	"grpc_client/source/utils"
	"io"
	"log"
	"time"
)

func send(_addr string, _func string) {
	log.Println("connect grpc server :", _addr)
	conn, err := grpc.Dial(_addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("connect grpc server faild：", err)
	}
	defer conn.Close()
	client := rpc.NewTestServiceClient(conn)

	// send
	if _func == "send" {
		// data := new(any.Any)
		// data.Value = []byte(`{"aa":"bb"}`)
		resp, err := client.Send(context.Background(), &rpc.Req{Data: "hello"})
		if err != nil {
			log.Fatalf("could not greet: %v", err.Error())
		}
		log.Printf("send get resp.Code: %d , resp.Data: %s", resp.Code, resp.Data)
	}

	// list
	if _func == "list" {
		stream, err := client.List(context.Background(), &rpc.Req{Data: "hello list"})
		if err != nil {
			log.Fatalf("list error: %v", err.Error())
		}
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("list recv error: %v", err.Error())
			}
			log.Printf("list resp: resp.Code: %d, resp.Data: %s", resp.Code, resp.Data)
		}
	}

	//record
	if _func == "record" {
		stream2, err := client.Record(context.Background())
		if err != nil {
			log.Fatalf("record error: %v", err.Error())
		}
		for n := 0; n < 10; n++ {
			tmp := fmt.Sprintf("record:%d", n)
			err := stream2.Send(&rpc.Req{Data: tmp})
			if err != nil {
				log.Fatalf("record send error: %v", err.Error())
			}
		}
		resp2, err := stream2.CloseAndRecv()
		if err != nil {
			log.Fatalf("record close error: %v", err.Error())
		}
		log.Printf("record resp: resp.Code: %d, resp.Data: %s", resp2.Code, resp2.Data)
	}
	//route
	if _func == "route" {
		stream3, err := client.Route(context.Background())
		if err != nil {
			log.Fatalf("route error: %v", err.Error())
		}

		for n := 0; n <= 10; n++ {
			err = stream3.Send(&rpc.Req{Data: "route"})
			if err != nil {
				log.Fatalf("route send error: %v", err.Error())
			}

			resp3, err := stream3.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("route resv error: %v", err.Error())
			}
			log.Printf("route resp: resp.Code: %d, resp.Data: %s", resp3.Code, resp3.Data)
			time.Sleep(time.Second)
		}
		stream3.CloseSend()
	}
}

func main() {
	utils.Init()
	gwAddr := utils.GetConf().Server.Addr
	gwFunc := utils.GetConf().Server.Func
	send(gwAddr, gwFunc)
}
