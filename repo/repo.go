package repo

import "fmt"
import model "github.com/roshniashok/crud/model"

type CallRecords []model.CallRecord

var Callrecords CallRecords
var currentId int32

func RepoFindRecord(id int32) model.CallRecord {
	for _, t := range Callrecords {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return model.CallRecord{}
}

func RepoCreateRecord(t model.CallRecord) model.CallRecord {
	currentId += 1
	t.Id = currentId
	Callrecords = append(Callrecords, t)
	return t
}

func RepoDestroyRecord(id int32) error {
	for i, t := range Callrecords {
		if t.Id == id {
			Callrecords = append(Callrecords[:i], Callrecords[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find CallRecord with id of %d to delete", id)
}
