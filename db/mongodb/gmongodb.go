/*
 * go4api - an api testing tool written in Go
 * Created by: Ping Zhu 2019.09
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package gmongodb

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/Aysnine/go4api/cmd"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClients map[string]*mongo.Client

type MongoDBExec struct {
	TargetMongoDB    string
	Database         string
	Collection       string
	Cmd              string
	Filter           interface{}
	FindFilter       interface{}
	UpdateFilter     interface{}
	CmdResults       interface{}
	CmdAffectedCount int
}

func InitMongoDBConnection() {
	MongoDBClients = make(map[string]*mongo.Client)

	mongs := cmd.GetMongoDBConfig()

	// master only
	for k, v := range mongs {
		ip := v.Ip
		port := v.Port
		user := v.UserName

		pw := ""
		pwV := v.Password
		pwV = strings.Replace(pwV, "${", "", -1)
		pwV = strings.Replace(pwV, "}", "", -1)
		if len(pwV) > 0 {
			pw = os.Getenv(pwV)
		}

		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + user + ":" + pw + "@" + ip + ":" + fmt.Sprint(port)))
		if err != nil {
			panic(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
		err = client.Connect(ctx)

		if err != nil {
			panic(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			panic(err)
		}

		key := strings.ToLower(k)
		MongoDBClients[key] = client
	}
}

func Run(cmdStr string) (int, interface{}, string) {
	var err error
	mongoExecStatus := ""

	tDb := "master"

	// if strings.HasPrefix(cmdStr, "db.") {
	// example: cmdStr => "dbname, db.getCollection('collectionName').deleteMany({'name': 'value'});"
	// example: cmdStr => "dbname, db.getCollection('collectionName').updateOne({'_id': ObjectId('xxx')}, {$set: {'key': 'value'}});"

	sl := strings.SplitN(cmdStr, ",", 2)

	db := strings.TrimSpace(sl[0])
	// content in first () is collection name
	pos1 := strings.Index(sl[1], "(")
	pos2 := strings.Index(sl[1], ")")

	coll := sl[1][pos1+2 : pos2-1]
	// content in second () is filter / updatefilter, to be safe, use second "(" and last ")"
	sR := strings.Replace(sl[1], "(", "R", 1)
	pos3 := strings.Index(sR, "(")

	pos4 := strings.LastIndex(sl[1], ")")

	filter := sl[1][pos3+1 : pos4]
	cmd := sl[1][pos2+2 : pos3]

	//
	mongoDBExec := &MongoDBExec{
		TargetMongoDB: tDb,
		Database:      db,
		Collection:    coll,
		Cmd:           cmd,
		Filter:        filter,
		FindFilter:    "",
		UpdateFilter:  "",
		CmdResults:    "",
	}

	err = mongoDBExec.Do()
	if err == nil {
		mongoExecStatus = "cmdSuccess"
	} else {
		mongoExecStatus = "cmdFailed"
	}

	return mongoDBExec.CmdAffectedCount, mongoDBExec.CmdResults, mongoExecStatus
}

func (mongoDBExec *MongoDBExec) Do() error {
	client := MongoDBClients[mongoDBExec.TargetMongoDB]

	var err error
	var res interface{}
	var v *mongo.SingleResult

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	//
	switch strings.ToUpper(mongoDBExec.Cmd) {
	case "FINDONE":
		collection := client.Database(mongoDBExec.Database).Collection(mongoDBExec.Collection)

		// if has filter specified, i.e. contains ":"
		if strings.Count(mongoDBExec.Filter.(string), ":") > 0 {
			sl := strings.Split(mongoDBExec.Filter.(string), ",")
			findFilter := sl[0]

			findFilterKey, findFilterValue := getFindFilterKeyValue(findFilter)
			filter := getFindFilterBson(findFilterKey, findFilterValue)

			v = collection.FindOne(ctx, filter)

			r, err := v.DecodeBytes()

			if err != nil {
				// panic(err)
				fmt.Println("!! Warning, mongo: no documents in result. ", err)
			} else {
				res = r.String()

				mongoDBExec.CmdAffectedCount = -1
				mongoDBExec.CmdResults = res
			}
		} else {
			filter := bson.D{{}}
			v = collection.FindOne(ctx, filter)

			r, err := v.DecodeBytes()

			if err != nil {
				// panic(err)
				fmt.Println("!! Warning, mongo: no documents in result")
			} else {
				res = r.String()

				mongoDBExec.CmdAffectedCount = -1
				mongoDBExec.CmdResults = res
			}
		}
		// fmt.Println(">>>>>>>> res: ", res)
	case "UPDATEONE":
		collection := client.Database(mongoDBExec.Database).Collection(mongoDBExec.Collection)

		//
		sl := strings.Split(mongoDBExec.Filter.(string), ",")
		findFilter := sl[0]
		findFilterKey, findFilterValue := getFindFilterKeyValue(findFilter)
		//
		updateFilter := sl[1]
		updateFilterKey, updateFilterValue := getUpdateFilterKeyValue(updateFilter)

		//
		res, err = collection.UpdateOne(ctx,
			bson.D{{findFilterKey, findFilterValue}},
			bson.M{"$set": bson.M{updateFilterKey: updateFilterValue}},
		)

		if err != nil {
			panic(err)
		}

		if err == nil {
			mongoDBExec.CmdAffectedCount = 1
			mongoDBExec.CmdResults = res
		}
	case "DELETEMANY":
		collection := client.Database(mongoDBExec.Database).Collection(mongoDBExec.Collection)

		// if has filter specified, i.e. contains ":"
		if strings.Count(mongoDBExec.Filter.(string), ":") > 0 {
			sl := strings.Split(mongoDBExec.Filter.(string), ",")
			findFilter := sl[0]

			findFilterKey, findFilterValue := getFindFilterKeyValue(findFilter)
			filter := getFindFilterBson(findFilterKey, findFilterValue)

			res, err = collection.DeleteMany(ctx, filter)

			if err != nil {
				panic(err)
			}

			if err == nil {
				mongoDBExec.CmdAffectedCount = -1
				mongoDBExec.CmdResults = res
			}
		} else {
			filter := bson.D{{}}
			res, err = collection.DeleteMany(ctx, filter)

			if err != nil {
				panic(err)
			}

			if err == nil {
				mongoDBExec.CmdAffectedCount = -1
				mongoDBExec.CmdResults = res
			}
		}
	default:
		mongoDBExec.CmdAffectedCount = -1
		fmt.Println("!! Warning, Command ", mongoDBExec.Cmd, " is not supported currently, will enhance it later")
	}

	return err
}

func getFindFilterBson(findFilterKey string, findFilterValue interface{}) interface{} {
	var filter = bson.D{{}}

	switch reflect.TypeOf(findFilterValue).Kind().String() {
	case "string":
		ff := findFilterValue.(string)
		ff = ff[2 : len(ff)-1]

		// if is reg expression
		if strings.HasPrefix(ff, "/") {
			filter = bson.D{{findFilterKey, primitive.Regex{Pattern: ff, Options: ""}}}
		} else {
			filter = bson.D{{findFilterKey, findFilterValue}}
		}
	default:
		filter = bson.D{{findFilterKey, findFilterValue}}
	}

	return filter
}

func getFindFilterKeyValue(findFilter string) (string, interface{}) {
	findFilterKey := strings.TrimSpace(strings.Split(findFilter, ":")[0])
	findFilterKey = findFilterKey[2 : len(findFilterKey)-1]

	findFilterValue := strings.TrimSpace(strings.Split(findFilter, ":")[1])
	findFilterValue = findFilterValue[0 : len(findFilterValue)-1]

	// fmt.Println("!>>>>>>>>: ", findFilter, findFilterKey, findFilterValue)

	var obj interface{}
	if strings.HasPrefix(findFilterValue, "ObjectId") {
		findFilterValue = findFilterValue[9:]
		findFilterValue = findFilterValue[1 : len(findFilterValue)-2]
		obj, _ = primitive.ObjectIDFromHex(findFilterValue)
	} else {
		findFilterValue = findFilterValue[1 : len(findFilterValue)-1]
		obj = findFilterValue
	}

	return findFilterKey, obj
}

func getUpdateFilterKeyValue(updateFilter string) (string, interface{}) {
	kvSlice := strings.Split(updateFilter, ":")

	updateFilterKey := strings.TrimSpace(kvSlice[1])
	updateFilterKey = updateFilterKey[2 : len(updateFilterKey)-1]

	updateFilterValue := strings.TrimSpace(kvSlice[2])
	updateFilterValue = updateFilterValue[1 : len(updateFilterValue)-3]

	return updateFilterKey, updateFilterValue
}
