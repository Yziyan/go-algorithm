// @Author: Ciusyan 12/5/23

package other

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongo(t *testing.T) {
	ctx := context.TODO()

	// 设置 MongoDB 连接
	clientOptions := options.Client().ApplyURI("mongodb://zhiyan:1234@localhost:27017")

	// 获取 MongoDB 客户端
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	require.NoError(t, err)

	t.Log("Connected to MongoDB!")

	// 获取集合
	collection := client.Database("articledb").Collection("myCollection")

	// 插入文档
	doc := bson.D{{Key: "name", Value: "志颜"}, {Key: "age", Value: "23131231"}}
	insertResult, err := collection.InsertOne(ctx, doc)
	require.NoError(t, err)
	t.Log("Inserted a single document: ", insertResult.InsertedID)
	//
	//// 更新文档
	//filter := bson.D{{Key: "name", Value: "Zhiyan"}}
	//update := bson.D{
	//	{Key: "$set", Value: bson.D{
	//		{Key: "age", Value: "2000"},
	//	}},
	//}
	//updateResult, err := collection.UpdateOne(ctx, filter, update)
	//
	//require.NoError(t, err)
	//t.Logf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	//
	//// 查询文档
	//var result bson.M
	//err = collection.FindOne(ctx, filter).Decode(&result)
	//require.NoError(t, err)
	//
	//t.Log("Found a single document: ", result)
	//
	//// 删除文档
	//deleteResult, err := collection.DeleteOne(ctx, filter)
	//require.NoError(t, err)
	//
	//t.Logf("Deleted %v documents in the mycollection collection\n", deleteResult.DeletedCount)
	//
	//// 关闭连接
	//err = client.Disconnect(ctx)
	//require.NoError(t, err)
	//
	//t.Log("Connection to MongoDB closed.")
}
