package util


import (
	
	"os"
	"path/filepath"
	"fmt"
	"github.com/beevik/etree"
	"goodliving/rf-env/pkg/setting"
)

func ReadSource(env string) (map[string]interface{}, error) {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	abspath := filepath.Join(pwd, setting.EnvFileName)
	
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(abspath); err != nil {
		fmt.Printf("文件读取失败，请检查文件%s是否存在", abspath)
	}

	root := doc.SelectElement("environment")
	// 使用interface{}初始化一个一维映射，interface{} 可以代表任意类型
	// interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	env_info := make(map[string]interface{})
	
	for _, info := range root.SelectElements(env) {
		for _, c := range info.ChildElements() {
			env_info[c.Tag] = readTypeValue(info, c.Tag)
		}
	}
	return env_info, err
}


func readTypeValue(info *etree.Element, choose_type string) (map[string]interface{}) {
	result := make(map[string]interface{})
	type_choice := "./" + choose_type + "/*"

	for _, e := range info.FindElements(type_choice) {
		result[e.Tag] = e.Text()
	}
	return result
}
















