package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	client := blogpb.NewBlogServiceClient(conn)

	gserver := gin.Default()

	gserver.POST("/create/",func(ctx gin.Context){

		request:=&blogpb.CreateBlogReq{ID:1,
									AuthorID:"Auth001",
									Content:"My first grpc",
									Title:"First gRPC"
								}

		if resp,err:=client.CreateBlog(ctx,request); err==nil{
			ctx.JSON(http.StatusOK,gin.H{
				"result": fmt.Sprint(CreateBlogRes.result)
			})
		}else{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		}
	})

	gserver.GET("/read/:id/",func(ctx gin.Context){

			id,err:=strconv.ParseUint(ctx.Param("id"),10,64)
			if err!=nil{
				ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
				return
			}
			request:=&blogpb.ReadBlogReq{ID:int64(id)}

			if response,err:= client.ReadBlog(ctx,request);err==nil{
				ctx.JSON(http.StatusOK,gin.H{
					"result": fmt.Sprint(ReadBlogRes.result)
				})
			}else{
				ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			}
	})


	gserver.PUT("/update/id",func(ctx gin.Context){
		id,err:=strconv.ParseUint(ctx.Param("id"),10,64)
			if err!=nil{
				ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
				return
			}
			request:=&blogpb.UpdateBlogReq{ID:int64(id)}
			if response,err:= client.UpdateBlog(ctx,request);err==nil{
				ctx.JSON(http.StatusOK,gin.H{
					"result": fmt.Sprint(UpdateBlogRes.result)
				})
			}else{
				ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			}

	})

	gserver.DELETE("/delete/id",func(ctx gin.Context){
		id,err:=strconv.ParseUint(ctx.Param("id"),10,64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		request:=&blogpb.DeleteBlogReq{ID:int64(id)}
		if response,err:= client.DeleteBlog(ctx,request);err==nil{
			ctx.JSON(http.StatusOK,gin.H{
				"result": fmt.Sprint(DeleteBlogRes.result)
			})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		}


	})

	if err := gserver.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
