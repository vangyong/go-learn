//package utils
//
//import "fmt"
//import "io/ioutil"
//import "goshark"
//import "golibwireshark"
//
////打印抓包数据
//func PrintPcap(filePath string) {
//	//	file := "E:\testfile\2019-04-23 14_12_45-HTTP.pcap"
//	file := filePath
//	d := goshark.NewDecoder()
//	if err := d.DecodeStart(file); err != nil {
//		log.Println("Decode start fail:", err)
//		return
//	}
//	defer d.DecodeEnd()
//
//	f, err := d.NextPacket()
//	if err != nil {
//		log.Println("Get packet fail:", err)
//		return
//	}
//
//	key := "igmp.maddr"
//	value, ok := f.Iskey(key)
//	if ok {
//		fmt.Printf("key: %s\nvalue: %s\n", key, value)
//	}
//}
