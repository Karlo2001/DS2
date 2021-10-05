package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "DS2/courses_pb"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func printCourse(client pb.CoursesClient, req *pb.CourseRequest) {
	log.Printf("Getting course for id %d", req.Id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	course, err := client.Get(ctx, req)
	if err != nil {
		log.Printf("%v.Get(_) = _, %v: ", client, err)
		fmt.Printf("No such id exists. Please enter valid id\n")
		return
	}
	log.Println(course)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCoursesClient(conn)

	var in string
	var id int32
	for {
		fmt.Printf("\nEnter course id (1-3): ")
		fmt.Scanln(&in)
		i, err := strconv.ParseInt(in, 10, 32)
		if err != nil {
			fmt.Printf("Couldn't read integer\n")
			continue
		}
		id = int32(i)
		printCourse(client, &pb.CourseRequest{Id: id})
	}
}
