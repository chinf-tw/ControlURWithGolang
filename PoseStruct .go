package main

import (
	"encoding/json"
	"os"
)

type Pose struct {
	PoseName string
	PoseData []float64
}

type Poses []Pose

func (ps *Poses) AddPosesToJsonFile(fileName string) {
	data, err := json.Marshal(ps)
	check(err)
	filePath := "./" + fileName + ".json"
	f, err := os.Create(filePath)
	check(err)

	defer f.Close()
	_, err = f.Write(data)
	check(err)
}

func (ps *Poses) AddPose(p Pose) {
	*ps = append(*ps, p)
}

func (p Pose) newPose(name string, data []float64) Pose {
	return Pose{PoseName: name, PoseData: data}
}
