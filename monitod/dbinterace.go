package monitod

import (
	"fmt"
	_ "log"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)
//if this is available through the whole package
//then i only need to import the struct
//keep it here for now, this represents the document

const (
	uri = "127.0.0.1:27017"
	monitoDb = "monito"
	cpuStats = "cpu_stats"
	networkStats = "net_stats"
	memStats = "mem_stats"
)

//make this exportable
func DbFactory(connectionInfo *mgo.Collection) (err error) {

	fmt.Println("whatever")
	session, err := mgo.Dial(uri)

	if err !=nil {
		return err
	}

	defer session.Close()
	connectionInfo = session.DB(monitoDb).C(cpuStats)

	if err != nil {
		return err
	}

	return nil
}
