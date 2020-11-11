package main

import (
	"context"
	"log"
	"net"
	blogpb "path/CrudWithgRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type blogItem struct {
	ID       int    `json:"id"`
	AuthorID string `json:"author_id"`
	Content  string `json:"content"`
	Title    string `json:"title"`
}

type BlogServiceServer struct {
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterAddServiceServer(srv, &server{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *BlogServiceServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {

	db, err := db.getconn()
	if err != nil {
		log.Println("Something Went Wrong")
	}

	userid := req.GetId()
	authorid := req.GetAuthorId()
	usertitle := req.GetTitle()
	usercontent := req.GetContent()

	result, err := db.Query("insert into new_table values(?,?,?,?)", userid, authorid, usertitle, usercontent)
	if err != nil {
		panic(err.Error())
	}
	response := &blogpb.CreateBlogRes{Result: result}
	defer db.Close()
	return response, nil
}

func (s *BlogServiceServer) ReadBlog(ctx context.Context, req *blogpb.ReadBlogReq) (*blogpb.ReadBlogRes, error) {

	db, err := db.getconn()
	if err != nil {
		log.Println("Something Went Wrong")
	}
	userid := req.GetId()

	result := "SELECT * FROM new_table WHERE id=?"
	row := db.QueryRow(result, userid)
	blogr := blogItem{}
	rek, _ := row.Scan(&blogr.ID, &blogr.AuthorID, &blogr.Title, &blogr.Content)
	response := &blogpb.ReadBlogRes{Result: rek}
	defer db.Close()
	return response, nil

}

func (s *BlogServiceServer) UpdateBlog(ctx context.Context, req *UpdateBlogReq) (*blogpb.UpdateBlogRes, error) {

	poi := blogItem{}
	db, err := db.getconn()
	defer db.Close()
	if err != nil {
		log.Println("Something Went Wrong")
	}
	userid := req.GetId()
	query := "UPDATE new_blog SET author_id = ?, content = ?, title = ? WHERE id = ?"
	row := db.QueryRow(query, userid)
	if err != nil {
		panic(err.Error())
	}
	data, _ := row.Scan(poi.AuthorID, poi.Content, poi.Title)
	response := &blogpb.UpdateBlogRes{Result: data}
	defer db.Close()
	return response, nil
}

func (s *BlogServiceServer) DeleteBlog(ctx context.Context, req *DeleteBlogReq) (*blogpb.DeleteBlogRes, error) {

	poi := blogItem{}
	db, err := db.getconn()
	defer db.Close()
	if err != nil {
		log.Println("Something Went Wrong")
	}
	userid := req.GetId()
	query := "DELETE FROM new_blog WHERE id = ?"
	row := db.QueryRow(query, userid)
	if err != nil {
		log.Println("Something Went Wrong")
	}
	response := &blogpb.DeleteBlogRes{Result: "Successfully deleted"}
	return response, nil
}
