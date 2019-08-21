package workers

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

var workerMap = make(map[string]Worker)

//TODO bind queue and consumer
func GetConnections() []ConnInfo {
	conns := ParseConnInfo()
	return conns
}

func GetEnableWorkers() []WorkerInfo {
	result := []WorkerInfo{}
	workItems := ParseWorkerInfo()
	if len(workItems) > 0 {
		for _, value := range workItems {
			if value.Enable == true {
				result = append(result, value)
			}
		}
	}

	return result
}

func InitManager() {
	connList := GetConnections()
	if len(connList) == 0 {
		panic(errors.New("connect not aval"))
	}

	workInfos := GetEnableWorkers()
	workStructMap := GetWorkerStructMap()
	for _, item := range workInfos {
		if workStructMap[item.Name] != nil {
			log.Printf("%s", workStructMap[item.Name])
			structPtr := reflect.New(reflect.TypeOf(workStructMap[item.Name]))
			s := structPtr.Elem()
			if s.Kind() == reflect.Struct {
				for k, v := range item.Payload {
					log.Printf("Value：%s", v.(string))
					field := s.FieldByName(k)
					if field.IsValid() && field.CanSet() {
						field.SetString(v.(string))
					}
				}
			}
			w, _ := NewWorker(connList[0], item, s.Interface().(Handler))
			workerMap[item.Name] = w
		}
	}
}

func Start() {
	for name, woker := range workerMap {
		fmt.Printf("woker ['%s'] start working", name)
		woker.Run()
	}
}

func Pause(name string) {

}
