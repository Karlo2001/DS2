package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	pb "DS2/courses_pb"

	"google.golang.org/grpc"
)

var (
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of courses")
	port       = flag.Int("port", 10000, "The server port")
)

type courseServer struct {
	pb.UnimplementedCoursesServer
	savedCourses []*pb.Course
}

func (s *courseServer) Get(ctx context.Context, in *pb.CourseRequest) (*pb.Course, error) {
	for _, course := range s.savedCourses {
		if course.Id == in.Id {
			return course, nil
		}
	}
	// No course was found, return nil
	return nil, nil
}

func (s *courseServer) loadCourses(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to load default courses: %v", err)
		}
	} else {
		data = exampleData
	}
	if err := json.Unmarshal(data, &s.savedCourses); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func NewServer() *courseServer {
	s := &courseServer{}
	s.loadCourses(*jsonDBFile)
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCoursesServer(grpcServer, NewServer())
	grpcServer.Serve(lis)
}

var exampleData = []byte(`[{
	"id": 1,
	"name": "First course",
	"satisfaction_rating": 1.430
}, {
	"id": 2,
	"name": "Second course",
	"satisfaction_rating": 2.672
}, {
	"id": 3,
	"name": "Third course",
	"satisfaction_rating": 3.258
}]`)
