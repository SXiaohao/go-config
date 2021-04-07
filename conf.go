/**
 * @Author: sxiaohao
 * @Description:
 * @File:  conf
 * @Version: 1.0.0
 * @Date: 2021/4/6 下午7:21
 */

package goConfig

import (
	"io/ioutil"
	"strings"
)

type Config struct {
	filename string
	data     map[string]map[string]string
}

func NewConfig(filename string) (Config, error) {

	config := Config{
		filename: filename,
		data:     make(map[string]map[string]string),
	}

	//文件读取到缓冲区
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	//解析内容
	err = config.parseContent(buf)

	if err != nil {
		panic(err)
	}

	return config, nil
}


//获取值
func (config Config) GetString(key string) string {

	index := strings.LastIndex(key, ".")

	nodeKey := key[:index]
	key = key[index+1:]


	return config.data[nodeKey][key]
}

//解析内容
func (config *Config) parseContent(buf []byte) error {
	//记录当前节点
	var node string

	//建立临时Map存储数据
	tempMap := make(map[string]string)

	//换行符分割内容
	content := strings.Split(string(buf), "\n")

	for _, ctn := range content {

		//去掉前后空格
		ctn = strings.Trim(ctn, " ")

		//检测有注释的跳过
		if strings.HasPrefix(ctn, "#") {
			continue
		} else if ctn == "" { //检测空行跳过
			continue
		} else if strings.HasPrefix(ctn, "[") && strings.HasSuffix(ctn, "]") { //获取节点
			ctn = strings.Trim(ctn, "[")
			ctn = strings.Trim(ctn, "]")
			node = ctn
		} else { //获取节点下的数据
			//取出等号左右两边数据
			str := strings.Split(ctn, "=")
			//去除数据前后的空格
			str[0] = strings.Trim(str[0], " ")
			str[1] = strings.Trim(str[1], " ")

			tempMap[str[0]] = str[1]

		}

	}
	//将临时数据赋值给当前节点
	config.data[node] = tempMap
	return nil
}
