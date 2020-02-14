package main

import (
	"github.com/linkedin/goavro"
	"log"
)

func main()  {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	codec, err := goavro.NewCodec(`
        {
          "type": "record",
          "name": "redis_RtDataBlock",
          "fields" : [
            {"name": "sName", "type": "string"},
            {"name": "sValue", "type": "string"},
            {"name": "sQuality", "type": "string","default":"0"},
            {"name": "sTagType", "type": "string"},
            {"name": "sTimeStamp", "type": "string"},
            {"name": "sParam", "type": "string","default":""}
          ]
        }`)
	if err != nil {
		log.Println(err)
	}

	// NOTE: May omit fields when using default value
	textual := []byte(`{"sName":"Power","sValue":"0.20000","sQuality":"0","sTagType":"15","sTimeStamp":"2019-12-19 18:16:04"}`)
	// Convert textual Avro data (in Avro JSON format) to native Go form
	native, _, err := codec.NativeFromTextual(textual)//从JSON字符串转化为Go原生数据
	if err != nil {
		log.Println(err)
	}
	log.Printf("%T",native)//native的类型： map[string]interface {}
	v := native.(map[string]interface {})
	log.Println(v["sName"])
	// Convert native Go form to binary Avro data
	binary, err := codec.BinaryFromNative(nil, native)//序列化为二进制
	if err != nil {
		log.Println(err)
	} else {
		log.Println(binary)
		log.Println(string(binary))
	}

	nData := make(map[string]interface {})
	nData["sName"] = "Power"
	nData["sValue"] = "0.20000"
	nData["sTagType"] = "15"
	nData["sTimeStamp"] = "2019-12-19 18:16:04"
	nData["sQuality"] = "0"
	binary, err = codec.BinaryFromNative(nil, nData)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(binary)
		log.Println(string(binary))
	}
}