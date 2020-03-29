package handlers

import (
	"context"
	"fmt"
	"grpc_server/source/rpc"
	"io"
	"log"
	"time"
)

type ExampleHandlers struct {
}

func (c *ExampleHandlers) Send(ctx context.Context, in *rpc.Req) (*rpc.Resp, error) {
	time.Sleep(5 * time.Second)
	//fmt.Println("msg data: ", in.Data)
	return &rpc.Resp{Code: 0, Data: "success111."}, nil
}

func (c *ExampleHandlers) List(in *rpc.Req, stream rpc.TestService_ListServer) error {
	fmt.Println("list msg data: ", in.Data)
	for n := 0; n <= 100; n++ {
		tmp := fmt.Sprintf("%s:%d", in.Data, n)
		err := stream.Send(&rpc.Resp{Code: 0, Data: tmp})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ExampleHandlers) Record(stream rpc.TestService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&rpc.Resp{Code: 0, Data: "Record success"})
		}
		if err != nil {
			return err
		}
		log.Printf("Record stream.Recv req.Data: %s", r.Data)
	}
	return nil
}

func (c *ExampleHandlers) Route(stream rpc.TestService_RouteServer) error {
	n := 0
	for {
		//recv
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Route Recv rep.Data: %s", r.Data)
		n++

		//send
		tmp := fmt.Sprintf("%s:%d", r.Data, n)
		err = stream.Send(&rpc.Resp{Code: 0, Data: tmp})
		if err != nil {
			return err
		}
	}
	return nil
}
