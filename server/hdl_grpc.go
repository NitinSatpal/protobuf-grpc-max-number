package main

import (
	"io"
	"log"
	"net"

	"max/pb"

	"google.golang.org/grpc"
)

type grpcServiceServer struct{}

func (*grpcServiceServer) Message(stream pb.Service_MessageServer) error {
	sess := globals.sessionStore.Create(stream)

	defer func() {
		sess.closeGrpc()
		sess.cleanUp()
		log.Println("grpc exited", sess.sid)
	}()

	go sess.writeGrpcLoop()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
			// return nil
		}
		if err != nil {
			return err
		}

		sess.dispatch(in)
	}

	return nil
}

func (sess *Session) writeGrpcLoop() {
	defer func() {
		sess.closeGrpc() // exit MessageLoop
	}()

	for {
		select {
		case msg, ok := <-sess.send:
			if !ok {
				// channel closed
				return
			}
			if err := grpcWrite(sess, msg); err != nil {
				log.Println("g.writeGrpcLoop", err)
				return
			}
		case msg := <-sess.stop:
			// Shutdown requested, don't care if the message is delivered
			if msg != nil {
				grpcWrite(sess, msg)
			}
			return

		}
	}
}

func grpcWrite(sess *Session, msg interface{}) error {
	out := sess.grpcnode
	if out != nil {
		// Will panic if msg is not of *pbx.ServerMessage type. This is an intentional panic.
		return out.Send(msg.(*pb.ServerMessage))
	}
	return nil
}

func serveGrpc(addr string) (*grpc.Server, error) {
	if addr == "" {
		return nil, nil
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()
	pb.RegisterServiceServer(srv, &grpcServiceServer{})
	log.Printf("gRPC server is registered at [%s]", addr)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Println("gRPC server failed:", err)
		}
	}()

	return srv, nil
}
