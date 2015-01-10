package monitod

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
//if this is available through the whole package
//then i only need to import the struct
//keep it here for now, this represents the document
type individualCpuStats struct {
	user      uint64
	nice      uint64
	system    uint64
	idle      uint64
	iowait    uint64
}
 
const (
	uri = "127.0.0.1:27017" 
	monitodb = "monito"
	cpustats = "cpu_stats"
)
//make this exportable
func DbFactory() string connectionInfo {
	session, err := mgo.Dial(uri)
	if err !=nil {
		panic(err)
	}
	defer session.Close()
	dbSession := session.DB(monitodb).C(cpustats)
	if err != nil {
		log.Fatal(err)
	}
}
