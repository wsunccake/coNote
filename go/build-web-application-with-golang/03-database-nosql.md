# database - nosql

---

## content

- [nosql](#nosql)
  - [go-redis - redis](#go-redis---redis)
  - [mongo-driver - mongo](#mongo-driver---mongo)

---

## nosql

### go-redis - redis

```go
package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	Host     = "127.0.0.1"
	Port     = 6379
	Password = ""
	Db       = 0
)

func main() {
	Addr := fmt.Sprintf("%s:%d", Host, Port)
	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       Db,
	})

	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)
}
```

```bash
#!/bin/bash
set_data() {
  docker exec -i redis redis-cli << EOF
SET foo ABC
SET bar XYZ
GET bar
EOF
}

set_data

go build .
./demo
```

### mongo-driver - mongo

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Userinfo struct {
	Username   string
	Department string
	Created    time.Time
}

const (
	UserName     string = "root"
	Password     string = "password"
	Addr         string = "127.0.0.1"
	Port         int    = 27017
	Database     string = "foo"
	MaxLifetime  int    = 5000
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func main() {

	// conn := fmt.Sprintf("mongodb://%s:%s@%s:%d/?timeoutMS=%d", UserName, Password, Addr, Port, MaxLifetime)
	conn := fmt.Sprintf("mongodb://%s:%d/?timeoutMS=%d", Addr, Port, MaxLifetime)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(conn).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	coll := client.Database(Database).Collection("userinfo")

	// create
	userinfo := Userinfo{"slene", "RD", time.Now()}
	_, err = coll.InsertOne(context.TODO(), userinfo)

	userinfos := []interface{}{
		Userinfo{"alex", "sales", time.Now()},
		Userinfo{"frank", "IT", time.Now()},
	}
	_, err = coll.InsertMany(context.TODO(), userinfos)

	// query
	cursor, err := coll.Find(context.TODO(), bson.M{"department": "RD"})
	if err != nil {
		panic(err)
	}

	var results []Userinfo
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}

	// delete
	_, err = coll.DeleteMany(context.TODO(), bson.M{"department": "RD"})
	if err != nil {
		panic(err)
	}
}
```

```bash
set_data() {
  docker exec -i mongo mongosh << EOF
help
show databases
show dbs
show collections
show tables
use foo
db.userinfo.insert({"username": "slene", "department": "RD", "created": new Date()})
db.userinfo.find({})
db.userinfo.deleteMany({})
EOF
}

###
### main
###

go build .
./demo

set_data
```
