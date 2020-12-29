package mongo

import (
	"context"
	"github.com/fire-g/mark-go-db/db"
	"github.com/fire-g/mark-go-util/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	Config *db.DatabaseConfig
	err    error
	Client *mongo.Client
)

func InitMongo() *mongo.Client {
	Client, err = mongoInit(Config)
	if err != nil {
		logger.Error.Print("MongodbConfig init():\t", err)
	}
	return Client
}

//Mongodb连接池初始化
func mongoInit(config *db.DatabaseConfig) (client *mongo.Client, e error) {
	// context是go的标准库包，不是很清楚这个包的作用，文档上面也没有写很清楚，知道的朋友希望指点
	//一下
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// 构建mongo连接可选属性配置
	opt := new(options.ClientOptions)
	// 设置最大连接的数量
	opt = opt.SetMaxPoolSize(uint64(10))
	// 设置连接超时时间 5000 毫秒
	du, _ := time.ParseDuration("5000")
	opt = opt.SetConnectTimeout(du)
	// 设置连接的空闲时间 毫秒
	mt, _ := time.ParseDuration("5000")
	opt = opt.SetMaxConnIdleTime(mt)
	// 开启驱动
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://"+config.Username+":"+config.Password+"@"+config.Uri), opt)
	if err != nil {
		e = err
		return nil, err
	}
	// 注意，在这一步才开始正式连接mongo
	e = client.Ping(ctx, readpref.Primary())
	return client, e
}
