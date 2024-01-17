package helper

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
)

var testEmailYamlData = []byte(`
- name: test_name
  subject: test_subject
  body: test_body`)

// 異なるパッケージのテストがある&数が少ないのでTestMainでpre/post処理を書かずに個別に書く
//func TestMain(m *testing.M) {
//	m.Run()
//}

func getTestFile(t *testing.T) (file *os.File, closeFunc func()) {
	file, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}

	closeFunc = func() {
		file.Close()
		os.Remove(file.Name())
	}

	if _, err = file.Write(testEmailYamlData); err != nil {
		t.Fatal(err)
	}

	data, _ := io.ReadAll(file)
	fmt.Println(string(data))

	return file, closeFunc
}

//func TestUnmarshalYamlFromFile(t *testing.T) {
//	file, closeFunc := getTestFile(t)
//	defer closeFunc()
//
//	email := model.Email{}
//	var i interface{} = email
//
//	data, _ := io.ReadAll(file)
//	fmt.Println(string(data))
//	UnmarshalYamlFromFile(file, &i)
//
//	_, ok := i.(model.Email)
//	if !ok {
//		t.Fatal("failed to cast interface{} to model.Email")
//	}
//
//}
//
//func TestUnmarshalYamlFromBytes(t *testing.T) {
//	email := model.Email{}
//	var i interface{} = email
//
//	done, err := UnmarshalYamlFromBytes(testEmailYamlData, &i)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	pp.Println(done)
//
//	//mail, ok := i.(*model.Email)
//	//if !ok {
//	//	t.Fatal("failed to cast interface{} to model.Email")
//	//}
//
//	//_ := done.(*model.Email)
//}
//

func TestUnmarshal(t *testing.T) {
	email := model.Email{}
	if err := UnmarshalYaml(testEmailYamlData, &email); err != nil {
		t.Fatal(err)
	}

	fmt.Println(email)
}
