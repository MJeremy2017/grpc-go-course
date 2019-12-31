package main

import (
	"context"
	"fmt"
	"github.com/grpc-go-course/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var collection *mongo.Collection

type server struct {
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (*server) CreateBlog(ctx context.Context, request *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := request.Blog
	data := blogItem{
		AuthorID: blog.AuthorId,
		Content:  blog.Content,
		Title:    blog.Title,
	}

	result, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal err: [%v]", err),
		)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID) // cast to object id
	if !ok {
		return nil, status.Error(codes.Internal, "ID fetching err")
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.AuthorId,
			Content:  blog.Content,
			Title:    blog.Title,
		},
	}, nil

}

func (*server) ReadBlog(ctx context.Context, request *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	blog_id := request.BlogId
	oid, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Cannot parse id")
	}

	data := &blogItem{}
	filter := bson.D{{"_id", oid}}
	result := collection.FindOne(context.Background(), filter)

	if err := result.Decode(data); err != nil {
		return nil, err
	}

	response := &blogpb.ReadBlogResponse{
		Blog: dataToBlog(data),
	}
	return response, nil
}

func (*server) UpdateBlog(ctx context.Context, request *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {

	blog := request.Blog
	oid, err := primitive.ObjectIDFromHex(blog.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Cannot parse id")
	}

	// find the blog
	data := &blogItem{}
	filter := bson.D{{"_id", oid}}
	result := collection.FindOne(context.Background(), filter)

	if err := result.Decode(data); err != nil {
		return nil, err
	}

	// update content
	data.Title = blog.Title
	data.AuthorID = blog.AuthorId
	data.Content = blog.Content

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Error(codes.Internal, "Cannot update")
	}

	return &blogpb.UpdateBlogResponse{
		Blog: dataToBlog(data),
	}, nil

}

func (*server) DeleteBlog(ctx context.Context, request *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	blogId := request.BlogId
	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Cannot parse id")
	}

	filter := bson.D{{"_id", oid}}
	deleteRes, deleteErr := collection.DeleteOne(context.Background(), filter)
	if deleteErr != nil {
		return nil, deleteErr
	}

	if deleteRes.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "id not found")
	}

	return &blogpb.DeleteBlogResponse{BlogId: blogId}, nil

}

func (*server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &blogItem{}
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("cannot decode [%v]", err))
		}

		resp := &blogpb.ListBlogResponse{Blog: dataToBlog(data)}
		sendErr := stream.Send(resp)
		if sendErr != nil {
			return status.Errorf(codes.Internal, "err sending [%v]", sendErr)
		}
	}

	if cursor.Err() != nil {
		return status.Errorf(codes.Internal, "do not know err [%v]", cursor.Err())
	}

	return nil

}

func dataToBlog(data *blogItem) *blogpb.Blog {
	return &blogpb.Blog{
		Id: data.ID.Hex(),
		Title: data.Title,
		AuthorId: data.AuthorID,
		Content: data.Content,
	}
}


func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Inside server ...")

	// connect to mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// connect to database
	collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // open a port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})
	reflection.Register(s)  // enable reflection on the service

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// a channel waiting for signal size 1
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping server")
	s.Stop()
	lis.Close()
	fmt.Println("Stoping MongoDB")
	client.Disconnect(context.TODO())
	fmt.Println("Sever stopped")

}
