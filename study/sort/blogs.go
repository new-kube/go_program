package main
 
import (
    "sort"
    "fmt"
)
 
/*
Go 的排序思路和 C 和 C++ 有些差别。
1.C 默认是对数组进行排序
2.C++ 是对一个序列进行排序
3.Go 待排序的可以是任何对象， 虽然很多情况下是一个 slice (分片， 类似于数组)，或是包含 slice 的一个对象。
 
排序(接口)的三个要素：
1.待排序元素个数 n ；
2.第 i 和第 j 个元素的比较函数 cmp ；
3.第 i 和 第 j 个元素的交换 swap ；
乍一看条件 3 是多余的， c 和 c++ 都不提供 swap 。
 
c 的 qsort 的用法：
 
qsort(data, n, sizeof(int), cmp_int); data 是起始地址， n 是元素个数， sizeof(int) 是每个元素的大小， cmp_int 是一个比较两个 int 的函数。
 
c++ 的 sort 的用法：
 
sort(data, data+n, cmp_int); data 是第一个元素的位置， data+n 是最后一个元素的下一个位置， cmp_int 是比较函数。
*/
 
/*
基本类型排序
*/
/*
1.升序排序
说明：对于int、float64、string数组/分片的排序，
     go分别提供sort.Ints()、sort.Float64s()、sort.Strings()函数（默认从小->大排序）
*/
func upSort(){
    intList := [] int {, , , , , , , , , }
    float8List := [] float64 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
    stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
 
    sort.Ints(intList)
    sort.Float64s(float8List)
    sort.Strings(stringList)
 
    fmt.Printf("%v\n%v\n%v\n",intList,float8List,stringList)
    /*
    打印结果：
    [0 1 2 3 4 5 6 7 8 9]
    [3.14 4.2 5.9 10 12.3 27.81828 31.4 50.4 99.9]
    [a b c d f i w x y z]
    */
}
 
/*
2.降序排序
*/
func downSort(){
    intList := [] int {, , , , , , , , , }
    float8List := [] float64 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
    stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
    sort.Sort(sort.Reverse(sort.IntSlice(intList)))
    sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
    sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
    fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
    /*
    打印结果：
    [9 8 7 6 5 4 3 2 1 0]
    [99.9 50.4 31.4 27.81828 12.3 10 5.9 4.2 3.14]
    [z y x w i f d c b a]
    */
}
 
/*
3.深度理解排序
sort 包中有一个 sort.Interface 接口，该接口有三个方法 Len() 、 Less(i,j) 和 Swap(i,j) 。
通用排序函数 sort.Sort 可以排序任何实现了 sort.Inferface 接口的对象(变量)。
对于 [] int 、[] float64 和 [] string 除了使用特殊指定的函数外，
还可以使用改装过的类型 IntSclice 、 Float64Slice 和 StringSlice ，
然后直接调用它们对应的 Sort() 方法；因为这三种类型也实现了 sort.Interface 接口，
所以可以通过 sort.Reverse 来转换这三种类型的 Interface.Less 方法来实现逆向排序，
这就是前面最后一个排序的使用。
下面使用了一个自定义(用户定义)的 Reverse 结构体， 而不是 sort.Reverse 函数， 来实现逆向排序。
*/
 
// 自定义的 Reverse 类型
type Reverse struct {
    sort.Interface    //这样，Reverse可以接纳任何实现了sort.Interface的对象
}
 
// Reverse 只是将其中的 Inferface.Less 的顺序对调了一下
func (r Reverse) Less(i, j int) bool {
    return r.Interface.Less(j, i)
}
 
//自定义排序
func selfDefineSort(){
    ints := []int{, , , , , }
 
    sort.Ints(ints)                                     // 特殊排序函数，升序
    fmt.Println("after sort by Ints:\t", ints)
 
    doubles := []float64{2.3, 3.2, 6.7, 10.9, 5.4, 1.8}
 
    sort.Float64s(doubles)
    fmt.Println("after sort by Float64s:\t", doubles)   // [1.8 2.3 3.2 5.4 6.7 10.9]
 
    strings := []string{"hello", "good", "students", "morning", "people", "world"}
    sort.Strings(strings)
    fmt.Println("after sort by Strings:\t", strings)    // [good hello mornig people students world]
 
    ipos := sort.SearchInts(ints, -)    // int 搜索
    fmt.Printf("pos of 5 is %d th\n", ipos)
 
    dpos := sort.SearchFloat64s(doubles, 20.1)    // float64 搜索
    fmt.Printf("pos of 5.0 is %d th\n", dpos)
 
    fmt.Printf("doubles is asc ? %v\n", sort.Float64sAreSorted(doubles))
 
    doubles = []float64{3.5, 4.2, 8.9, 100.98, 20.14, 79.32}
    // sort.Sort(sort.Float64Slice(doubles))    // float64 排序方法 2
    // fmt.Println("after sort by Sort:\t", doubles)    // [3.5 4.2 8.9 20.14 79.32 100.98]
    (sort.Float64Slice(doubles)).Sort()         // float64 排序方法 3
    fmt.Println("after sort by Sort:\t", doubles)       // [3.5 4.2 8.9 20.14 79.32 100.98]
 
    sort.Sort(Reverse{sort.Float64Slice(doubles)})    // float64 逆序排序
    fmt.Println("after sort by Reversed Sort:\t", doubles)      // [100.98 79.32 20.14 8.9 4.2 3.5]
 
}
 
/*
4.结构体类型排序
结构体类型的排序是通过使用 sort.Sort(slice) 实现的， 只要 slice 实现了 sort.Interface 的三个方法就可以。
*/
 
/*
(1)模拟IntSlice排序
缺点：
根据 Age 排序需要重新定义 PersonSlice 方法，绑定 Len 、 Less 和 Swap 方法，
如果需要根据 Name 排序， 又需要重新写三个函数； 如果结构体有 4 个字段，有四种类型的排序，那么就要写 3 × 4 = 12 个方法，
即使有一些完全是多余的;
根据不同的标准 Age 或是 Name，真正不同的体现在 Less 方法上，所以可以将 Less 抽象出来，
每种排序的 Less 让其变成动态的.见（2）
*/
type Person struct {
    Name string
    Age int
}
//按照Person.Age从大-》小排序（PersonSlice是person[]的模版）
type PersonSlice [] Person
 
//重写len()方法
func(a PersonSlice) Len() int{
    return len(a)
}
//重写Swap()方法
func (a PersonSlice) Swap(i,j int){
    a[i],a[j]=a[j],a[i]
}
//重写Less()方法
func (a PersonSlice) Less(i,j int ) bool{
    return a[j].Age < a[i].Age
}
func IntSliceSort(){
    people:=[] Person{
        {"zhang san", },
        {"li si", },
        {"wang wu", },
        {"zhao liu", },
    }
    fmt.Println(people) //[{zhang san 12} {li si 30} {wang wu 52} {zhao liu 26}]
 
    sort.Sort(PersonSlice(people)) //按照 Age 的逆序排序
    fmt.Println(people) //[{wang wu 52} {li si 30} {zhao liu 26} {zhang san 12}]
 
    sort.Sort(sort.Reverse(PersonSlice(people))) //按照 Age 的升序排序
    fmt.Println(people)//[{zhang san 12} {zhao liu 26} {li si 30} {wang wu 52}]
}
 
/*
(2)封装成 Wrapper
*/
type Person2 struct {
    Name string
    Age int
}
 
type PersonWrapper2 struct { //注意此处
    people [] Person2
    by func(p,q * Person2) bool
}
 
func (pw PersonWrapper2) Len() int  {//重写len()方法
    return len(pw.people)
}
func (pw PersonWrapper2) Swap(i,j int){ //重写Swap()方法
    pw.people[i],pw.people[j]=pw.people[j],pw.people[i]
}
func (pw PersonWrapper2) Less(i,j int) bool{//重写Less()方法
    return pw.by(&pw.people[i], &pw.people[j])
}
func wrapperSort(){
    people := [] Person2{
        {"zhang san", },
        {"li si", },
        {"wang wu", },
        {"zhao liu", },
    }
 
    fmt.Println(people)
 
    sort.Sort(PersonWrapper2{people, func (p, q *Person2) bool {
        return q.Age < p.Age    // Age 递减排序
    }})
    fmt.Println(people)
 
    sort.Sort(PersonWrapper2{people, func (p, q *Person2) bool {
        return p.Name < q.Name    // Name 递增排序
    }})
    fmt.Println(people)
    /*
    执行结果：
    [{zhang san 12} {li si 30} {wang wu 52} {zhao liu 26}]
    [{wang wu 52} {li si 30} {zhao liu 26} {zhang san 12}]
    [{li si 30} {wang wu 52} {zhang san 12} {zhao liu 26}]
    */
}
 
/*
(3)进一步封装
*/
type Person3 struct {
    Name string
    Age  int
}
 
type PersonWrapper3 struct {
    people [] Person3
    by func(p, q * Person3) bool
}
 
type SortBy func(p, q *Person3) bool
 
func (pw PersonWrapper3) Len() int {         // 重写 Len() 方法
    return len(pw.people)
}
func (pw PersonWrapper3) Swap(i, j int){     // 重写 Swap() 方法
    pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
func (pw PersonWrapper3) Less(i, j int) bool {    // 重写 Less() 方法
    return pw.by(&pw.people[i], &pw.people[j])
}
 
// 封装成 SortPerson 方法
func SortPerson(people [] Person3, by SortBy){
    sort.Sort(PersonWrapper3{people, by})
}
 
func wrapperSorts(){
    people := [] Person3{
        {"zhang san", },
        {"li si", },
        {"wang wu", },
        {"zhao liu", },
    }
    fmt.Println(people)
 
    sort.Sort(PersonWrapper3{people, func (p, q *Person3) bool {
        return q.Age < p.Age    // Age 递减排序
    }})
    fmt.Println(people)
 
    SortPerson(people, func (p, q *Person3) bool {
        return p.Name < q.Name    // Name 递增排序
    })
    fmt.Println(people)
    /*
    运行结果：
    [{zhang san 12} {li si 30} {wang wu 52} {zhao liu 26}]
    [{wang wu 52} {li si 30} {zhao liu 26} {zhang san 12}]
    [{li si 30} {wang wu 52} {zhang san 12} {zhao liu 26}]
    */
}
 
func main(){
    fmt.Println("升序==========")
    upSort()   //升序
    fmt.Println("降序==========")
    downSort()
    fmt.Println("模拟IntSlice排序==========")
    IntSliceSort()
    fmt.Println("封装成wrapper排序==========适合项目应用")
    wrapperSort()
    fmt.Println("更深层封装成wrapper排序==========适合项目应用")
    wrapperSorts()
}
