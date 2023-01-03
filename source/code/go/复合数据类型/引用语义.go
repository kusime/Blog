package main
import "fmt"
type user struct{
	name string
	ID int
}

func main() {
	// ! 这里测试 map ，slice ，array ，Struct
	testMap := map[string]int{"ming":111}
	testSlice := []int{1,2}
	testArray := [2]int{1,2}
	testStruct := user{name:"ming",ID:0}
	fmt.Println(testMap,testArray,testSlice,testStruct)
	mapTest(testMap)
	arrayTest(testArray)
	sliceTest(testSlice)
	fmt.Println(testMap,testArray,testSlice,testStruct)


	/* 
map[ming:111] [1 2] [1 2] {ming 0}
map[ming:666] [1 2] [666 2] {ming 0}
	*/
	// ! 可以看见除了数组，其他默认都是引用语义
	// ! 所以之前把数组切片和map归纳到那边是没毛病的
}

func mapTest(TmpMap map[string]int)  {
	TmpMap["ming"]=666
}

func sliceTest(TmpSlice []int)  {
	TmpSlice[0]=666
}

func arrayTest(TmpArray [2]int){
	TmpArray[0]=666
}

func arrayStruct(TmpStruct user){
	TmpStruct.name = "cloud"
	TmpStruct.ID = 666
}