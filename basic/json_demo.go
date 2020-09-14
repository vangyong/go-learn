package basic

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

// json str 转map
func JsonToMap(jsonStr string) map[string]interface{} {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("=====json str 转map======")
		fmt.Println(dat)
		fmt.Println(dat["host"])
	}
	return dat
}

// json str 转struct
func JsonToStruct(jsonStr string) Config {
	var config Config
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(config)
		fmt.Println(config.Host)
	}
	return config
}

// struct转json str
func StructToJson(config Config) string {
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("================struct 到json str==")
		fmt.Println(string(b))
		return string(b)
	}
	return ""
}


// map转json str
func MapToJson(dat map[string]interface{}) string {
	fmt.Println("===========map 到json str==============")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)
	return ""
}


// array 到 json str
func ArrayToJson(array []string) string {
	lang, err := json.Marshal(array)
	if err == nil {
		fmt.Println("========array 到 json str==")
		fmt.Println(string(lang))
	}
	return string(lang)
}

// array 到 json str
func JsonToArray(jsonStr string) []string {
	var array []string
	if err := json.Unmarshal([]byte(jsonStr), &array); err == nil {
		fmt.Println("================json 到 []string==")
		fmt.Println(array)
	}
	return array
}