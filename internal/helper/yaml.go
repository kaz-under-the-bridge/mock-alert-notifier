package helper

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

//func UnmarshalYamlFromFile(file *os.File, model *interface{}) error {
//	//decoder := yaml.NewDecoder(file)
//	//err := decoder.Decode(model)
//	//if err != nil {
//	//	return err
//	//}
//	//fmt.Println(model)
//	//return nil
//	data, _ := io.ReadAll(file)
//	fmt.Println(string(data))
//	err := yaml.Unmarshal(data, model)
//	if err != nil {
//		return err
//	}
//	pp.Println(model)
//	return nil
//}
//
//func UnmarshalYamlFromBytes(bytes []byte, i interface{}) (interface{}, error) {
//	err := yaml.Unmarshal(bytes, &i)
//	if err != nil {
//		return nil, err
//	}
//	return i, nil
//}

func UnmarshalYaml[T any](data []byte, v *T) error {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
