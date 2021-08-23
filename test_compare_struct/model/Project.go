package model

import "reflect"

type Project struct {
	Name string
}

func (project *Project) IsEmpty() bool {
	return reflect.DeepEqual(project, &Project{})
}
