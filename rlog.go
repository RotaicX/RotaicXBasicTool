package RotaicxBasicTool

import (
	"fmt"
	"os"
	"time"
)

type rlog struct {
	nowTime string

	SaveToFile bool
	// The accompanying parameters of "SaveToFile"
	SavePath     string
	SaveFileName string
}

var Rlog rlog

func (r *rlog) Println(printData interface{}) error {
	printPackStr := "\u001B[33m[%s] [INFO]: %s\u001B[0m\n"
	return r.basicPrintf(printPackStr, printData)
}

func (r *rlog) Errorln(printData interface{}) error {
	printPackStr := "\u001B[32m[%s] [INFO]: %s\u001B[0m\n"
	return r.basicPrintf(printPackStr, printData)
}

func (r *rlog) basicPrintf(printPackStr string, printData interface{}) error {
	r.getNowTime()
	printLogData := fmt.Sprintf(printPackStr, r.nowTime, printData.(string))

	if r.SaveToFile {
		return r.saveToFile(printLogData)
	}
	fmt.Printf(printLogData)

	return nil
}

func (r *rlog) saveToFile(saveData string) error {
	if r.SavePath == "" || r.SaveFileName == "" {
		return fmt.Errorf("some or all of the parameters required for this operation are not legal")
	}

	fullPath := fmt.Sprintf("%s/%s", r.SavePath, r.SaveFileName)
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	checkError(err)

	file.Write([]byte(saveData))

	defer checkError(file.Close())
	return err
}

func (r *rlog) getNowTime() {
	r.nowTime = time.Now().Format("2006-01-02 15:04:05")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
