package grpc_clinet

// import (
// 	"log"

// 	ps "EcoTrack/UserServis/protos"


// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// type Clients interface {
// 	Post() ps.PostServiceClient
// }

// type ServiceMeneger struct {
// 	postService ps.PostServiceClient
// }

// func New() *ServiceMeneger {

// 	postConn, err := grpc.NewClient(":7080", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Println("error while connecting to post service : ", err)
// 	}

// 	postService := ps.NewPostServiceClient(postConn)

// 	return &ServiceMeneger{
// 		postService: postService,
// 	}

// }

// func (s *ServiceMeneger) Post() ps.PostServiceClient {
// 	return s.postService
// }
