package main

import "fmt"

func main() {
	// var m1 map[string]string
	// fmt.Println("m1 length:", len(m1))
	// m5 := map[string]string{
	// 	"key1": "value1",
	// 	"key2": "value2",
	// }
	// fmt.Println("m5 length:", len(m5))
	// fmt.Println("m5:", m5["key1"])
	// var m2 = make(map[string]string, 8)
	// m2["key1"] = "张三"
	// fmt.Println("m2 length:", len(m2))
	// // 遍历map集合main
	// for key, value := range m2 {
	// 	fmt.Println("iterate map, m[", key, "] =", value)
	// }
	// // 使用内置函数删除指定的key
	// _, exist_key1 := m2["key1"]
	// fmt.Println("before delete, exist 10: ", exist_key1)
	// delete(m2, "key1")
	// _, exist_key1 = m2["key1"]
	// fmt.Println("after delete, exist 10: ", exist_key1)

	siteMap := make(map[string]string, 4) /*创建集合 */
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"
	name, ok := siteMap["Wiki"] /*如果确定是真实的,则存在,否则不存在 */

	fmt.Println(ok)
	fmt.Println(name)

}
