package utils

import (
	"bufio"
	"fmt"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

// 余弦相似度
func CalCosSim(f1, f2 []float64) float64 {
	l1 := len(f1)
	l2 := len(f2)

	if l1 != l2 {
		fmt.Printf("两个向量长度不一致，f1的长度是:%v, f2的长度是：%v", l1, l2)
		return -1
	}
	var rr float64 = 0.0
	var f1r float64 = 0.0
	var f2r float64 = 0.0
	for i := 0; i < l1; i++ {
		rr += f1[i] * f2[i]
		f1r += f1[i] * f1[i]
		f2r += f2[i] * f2[i]

	}
	var rs float64 = rr / (math.Sqrt(f1r) * math.Sqrt(f2r))
	return rs
}

// 带解析，带换行
func Printfln(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

// []string，删除单个元素。
func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// []string去重
func RemoveDup(qqEmailList []string) []string {
	return removeDup(qqEmailList)
}

// []string去重
func removeDup(qqEmailList []string) []string {
	qqEString := strings.Join(qqEmailList, "")
	newQEmailList := qqEmailList
	for i, dataStr := range qqEmailList {
		qqe_count := strings.Count(qqEString, dataStr)
		if qqe_count > 1 {
			newQEmailList = Remove(newQEmailList, i)
			if i == len(newQEmailList) {
				break
			}
			i--
		}
	}
	return newQEmailList
}

// 查看数据类型。
func GetType(value interface{}) string {
	return fmt.Sprintf("%T", value)

}

// 外部接口，验证权限。
func Run(password string, cmd string) { //外部使用。
	if password == "123456" {
		run(cmd)
	} else {
		fmt.Println("密码错误！")
	}

}

// 实际干活
func run(cmd string) {
	exec.Command(cmd).Run() //执行命令
}

// 读取所有内容→字符串，
// 如果错→defaultA参字符串。
func LoadStringFromFile1(fileNameA string, defaultA string) string {
	fileT, err := os.Open(fileNameA) //打开文件
	if err != nil {
		return defaultA
	}
	defer fileT.Close()                        // 关闭
	fileContentT, err := ioutil.ReadAll(fileT) //一次读取所有的字节
	if err != nil {
		return defaultA
	}
	return string(fileContentT) // 将字节转换为字符串

}

// 更简单的读取完整文本的方法
// 从文件中读取所有内容→字符串
// 错→defaultA
func LoadStringFromFile(fileNameA string, defaultA string) string {
	//fileContentT, err := ioutil.ReadFile(fileNameA) //
	fileContentT, err := os.ReadFile(fileNameA) // 1.6版本之后可以用os.ReadFile
	if err != nil {
		return defaultA
	}
	return string(fileContentT) //
}

// 读取所有内容→字符串列表
func LoadStringListFromFile(fileNameA string, defaultA string) []string {
	fileContentT, err := os.ReadFile(fileNameA) // 1.6版本之后可以用os.ReadFile
	if err != nil {
		return []string{defaultA}
	}
	return strings.Split(string(fileContentT), "，")
}

// 从文件读取行→指定行数（limitA）
// 错→
func LoadLimitFromFile(fileNameA string, limitA int) string {
	fileT, errT := os.Open(fileNameA) // 打开文件
	if errT != nil {
		return "文件打开错误：" + errT.Error()
	}
	defer fileT.Close()     // 关闭
	var buf strings.Builder // 自+
	reader := bufio.NewReader(fileT)
	limitT := 0 //行数
	for true {
		strT, err := reader.ReadString('\n') //行读取。
		if err != nil {
			//最后一行加入
			buf.WriteString(strT) // 加入
			break
		}
		buf.WriteString(strT) // 加入
		limitT++

		if limitT >= limitA {
			break
		}
	}
	return buf.String() // 返回字符串
}

// 将字符串写入文本文件。
// 1.创建新文件，2、缓冲读写3、保存前刷新缓存。
func SaveStringToFile(strA string, fileA string) string {
	fileT, errT := os.Create(fileA) // 创建新文件
	if errT != nil {
		return "创建文件错误信息：" + errT.Error()

	}
	defer fileT.Close()               // 函结前关闭文件
	writerT := bufio.NewWriter(fileT) //创建写入者
	writerT.WriteString(strA)         // 写入
	writerT.Flush()                   // 更新
	return ""

}

// 判断文件或目录
// FileExists判断文件或目录存在
// os.Stat(fileNameA)：这是一个系统调用，它返回一个描述文件 fileNameA 的os.FileInfo 类型值。
// 如果调用成功，os.FileInfo 将提供如名字，大小，修改时间等文件详细信息。如果文件不存在或者发生了其它错误，则返回一个错误。
// return errT == nil || os.IsExist(errT):
// 最后，这段代码返回的是一个布尔值，如果为真(true)则表示文件存在，如果为假(false)则表示文件不存在。
// || 是逻辑或操作，只要满足 errT == nil 或者 os.IsExist(errT) 中的任何一个条件，结果就是 true。
func FileExists(fileNameA string) bool {
	_, errT := os.Stat(fileNameA)
	return errT == nil || os.IsExist(errT)
	//没出错，或者存在。
	//	只要满足其中的任何一个条件，结果就是 true。
}

// 判断路径名是否是目录
func IsDirectory(dirNameA string) bool {
	f, err := os.Open(dirNameA) // 打开文件
	if err != nil {
		return false
	}
	defer f.Close()

	fi, err := f.Stat() //返回os.FileInfo 类型的值可以获取到文件的名字、大小、修改时间等元信息。
	if err != nil {
		return false
	}
	//mode := fi.Mode() 这行代码调用了 os.FileInfo 类型的 Mode 方法，这个方法返回的是 os.FileMode 类型的值，它描述了文件的模式和权限位。
	//mode.IsDir() 这行代码调用了 os.FileMode 类型的 IsDir 方法，这个方法返回一个布尔值，如果文件是一个目录，那么返回 true，否则返回 false。
	if mode := fi.Mode(); mode.IsDir() { //
		return true
	} else {
		return false
	}
}

// 删除文件
func DelFile(fileNameA string) {
	if !FileExists(fileNameA) {
		Printfln("文件%v不存在", fileNameA)
	}
	errT := os.Remove(fileNameA) // 删除文件
	if errT != nil {
		Printfln("删除文件时发生了错误：%v", errT.Error())
		return
	}
	Printfln("已成功删除文件%v。", fileNameA)
}

// 删除文件夹
func DelDir(dirNameA string) {
	if !FileExists(dirNameA) {
		Printfln("文件夹%v不存在", dirNameA)
		return
	}
	errT := os.RemoveAll(dirNameA) // 删除文件
	if errT != nil {
		Printfln("删除文件夹时发生了错误：%v", errT.Error())
		return
	}
	Printfln("已成功删除文件夹%v。", dirNameA)
}

// 创建文件夹
func CreateDir(dirNameA string) {
	errT := os.Mkdir(dirNameA, 0777)
	var isDir = false
	if errT != nil {
		if !FileExists(dirNameA) {
			Printfln("创建目录时发生错误：%v", errT.Error())
			return
		}
		isDir = true

	}
	if !isDir {
		Printfln("已成功创建目录%v", dirNameA)
	} else {
		Printfln("目录%v已存在", dirNameA)
	}

}

// 创建新目录和新文件
func CreateDirAndFile(dirNameA string, fileNameA string) {
	//权限模式0777是一个八进制的数，在UNIX或者类UNIX系统中，
	//表示这个目录的所有者(user)、所属的组(group)和其他所有人(others)都有读（4）、写（2）和执行（1）的权限。
	//每个数字是这三个权限的对应的和。所以，数字7(4+2+1)就表示有读、写和执行权限。
	errT := os.Mkdir(dirNameA, 0777)
	var isDir = false
	if errT != nil {
		if !FileExists(dirNameA) {
			Printfln("创建目录时发生错误：%v", errT.Error())
			return
		}
		isDir = true

	}
	if !isDir {
		Printfln("已成功创建目录%v", dirNameA)
	} else {
		Printfln("目录%v已存在", dirNameA)
	}

	fileT, errT := os.OpenFile(fileNameA, os.O_CREATE, 0666)
	//	os.Open和os.OpenFile函数都是用来打开文件的，但他们之间的主要区别在于使用的权限和模式。
	//os.Open(name string) (*os.File, error) 函数是以只读模式打开一个名为 name 的文件。这对于你仅仅想读取文件内容的场景很有用。
	//它同样是打开一个名为 name 的文件，但是你可以通过 flag 和 perm 参数指定文件的打开模式和权限。
	//打开模式 flag 可以是如下一项或多项的组合：
	//os.O_RDONLY（只读），os.O_WRONLY（只写），os.O_RDWR（读写），os.O_APPEND（追加），os.O_CREATE（如果不存在则创建新文件）等。
	//你可以按位或操作（|）来组合使用这些值。
	//权限 perm 则用于设置新创建的文件（如果使用了os.O_CREATE 选项）的权限，是一个 os.FileMode 类型的值。
	//所以在处理文件读取和操作上，os.OpenFile因为其更多的可选项提供了更多灵活性，而os.Open更常用在只读取文件的简单场景中。
	if errT != nil {
		Printfln("创建文件时发生错误：%v", errT.Error())
		return
	}
	defer fileT.Close()
}

// 移动文件 或者 改名
func MoveFile(oldFileNameA string, newFileNameA string) {
	errT := os.Rename(oldFileNameA, newFileNameA) //改名
	if errT != nil {
		Printfln("移动文件时发生了错误：%v", errT.Error())
		return
	}
	Printfln("已成功移动文件%v到%v。", oldFileNameA, newFileNameA)
}

// 获取文件大小
func GetFileSize(fileNameA string) int64 {
	fileInfoT, errT := os.Stat(fileNameA) // 获取文件的状态信息
	if errT != nil {
		Printfln("获取文件信息时发生错误：%v", errT.Error())
		return -1
	}
	return fileInfoT.Size() //字节内容大小。
}

// 文件拷贝
// old →new
func CopyFile(oldFileNameA string, newFileNameA string) {
	os.MkdirAll(newFileNameA, 0755) // 无目录自动创建目录。
	os.RemoveAll(newFileNameA)      //删除最里层的目录
	oldFileT, errT1 := os.Open(oldFileNameA)
	if errT1 != nil {
		Printfln("打开%v文件中发生错误：%v", oldFileNameA, errT1.Error())
		return
	}
	defer oldFileT.Close()
	newFileT, errT2 := os.OpenFile(newFileNameA, os.O_CREATE|os.O_RDWR, 0666)
	//0666是在UNIX和类UNIX系统（包括Linux和OS X）中表示文件权限的一个八进制数。
	//在这个系统中，每个文件的权限包括三个部分：用户（User）权限、组（Group）权限和其他（Other）用户权限。
	//每一部分可以包含读（Read）、写（Write）和执行（Execute）三种权限。
	//八进制数0666可以分解为三个数字，分别表示用户、组和其他用户的权限。
	//在这个案例中，每个都是6。
	//八进制数6等于二进制的110，代表读（值为4）和写（值为2）权限。
	//所以，0666说明了文件所有者、所属组和所有其他用户都有这个文件的读写权限。
	//读和写的功能。7的话就是有执行权限。
	if errT2 != nil {
		Printfln("打开%v文件中发生错误：%v", newFileNameA, errT2.Error())
		return
	}
	defer newFileT.Close()

	bufT := make([]byte, 5) // 一次传输 //内存缓冲区
	for {
		countT, err := oldFileT.Read(bufT) //一次传输 //老的获取数据。
		if err != nil {
			if err == io.EOF {
				break
			} //到结尾了就退出循环
			Printfln("从源文件中读取数据时发生错误：%v", err.Error())
			return
		}
		_, errT3 := newFileT.Write(bufT[:countT]) // 写到新的里面
		if errT3 != nil {
			Printfln("从源文件中写入数据时发生错误：%v", errT3.Error())
			return
		}

	}
	//

}

// 文本文件编码转换:GB18030→UTF-8
// 转换GB18030编码的字节切片为UTF-8编码
// 需引用golang.org/x/text/encoding/simplifiedchinese
func ConverBytesFromGB18030ToUTF8(srcA []byte) []byte {
	bufT := make([]byte, len(srcA)*4) // GB:4个字节
	transformer := simplifiedchinese.GB18030.NewDecoder()
	countT, _, errT := transformer.Transform(bufT, srcA, true)
	//go
	//nDst, nSrc, err := t.Transform(dst, src, atEOF)
	//nDst 是转换输入后写入 dst 的字节数
	//nSrc 是从 src 中读取的字节数
	//err 是转换过程中可能出现的错误
	if errT != nil {
		return nil
	}
	return bufT[:countT]
}

// 需引用golang.org/x/text/encoding/simplifiedchinese
func ConverBytesFromUTF8ToGB18030(srcA []byte) []byte {
	bufT := make([]byte, len(srcA)*1) // GB:4个字节
	transformer := simplifiedchinese.GB18030.NewDecoder()
	countT, _, errT := transformer.Transform(bufT, srcA, true)
	//go
	//nDst, nSrc, err := t.Transform(dst, src, atEOF)
	//nDst 是转换输入后写入 dst 的字节数
	//nSrc 是从 src 中读取的字节数
	//err 是转换过程中可能出现的错误
	if errT != nil {
		return nil
	}
	return bufT[:countT]
}

// 查看编码类型
// go get -u github.com/saintfish/chardet
// 如果你在中国，由于众所周知的原因，可能无法直接访问 GitHub，这时候需要设置 Go 的代理。你可以使用 GOPROXY 环境变量来设置，例如：
// bash export GOPROXY=https://goproxy.cn,direct
func GetEncodingType(fileNameA string) string {
	data, err := os.ReadFile(fileNameA)
	if err != nil {
		log.Fatal(err)
	}

	// 创建文本检测器
	detector := chardet.NewTextDetector()

	// 检测字节切片的编码
	result, err := detector.DetectBest(data)
	if err != nil {
		log.Fatal(err)
	}

	// 打印检测到的编码
	return result.Charset

}

// GB18030→UTF-8 文件转换
func ConverFileFromGB18030ToUTF8(fileNameA string) {
	myStr, errT := ioutil.ReadFile(fileNameA) // 读取文件
	if errT != nil {
		Printfln("读取文件%v时发生错误：%v", fileNameA, errT.Error())
		return
	}
	newStr := ConverBytesFromGB18030ToUTF8(myStr) // 文本转换编码
	newFileT, errT2 := os.OpenFile(fileNameA, os.O_CREATE|os.O_RDWR, 0666)
	if errT2 != nil {
		Printfln("打开文件%v时发生错误：%v", fileNameA, errT2.Error())
		return
	}
	defer newFileT.Close()
	newFileT.Write(newStr) // 写入UTF8编码
}

// []string 排序
// 文本排序前:[abc rst def 123]
// 文本排序后:[123 abc def rst]
const Ascending = 1   //升序
const descending = -1 //降序

func SortStringList(strList []string, flag int) {
	switch flag {
	case 1:
		//	升序
		sort.Sort(sort.StringSlice(strList))
		break
	case -1:
		//	降序
		sort.Sort(sort.Reverse(sort.StringSlice(strList)))
		break

	}

}

type CharSet string

const (
	UTF8    = CharSet("UTF-8")
	GB18030 = CharSet("GB18030")
)

// 二进制转换
// []byte → string // 中文问题
func ConverByteString(byte []byte, charset CharSet) string {
	var str string
	switch charset {
	case charset:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// 返回随机数密码
// countT 密码长度
func GetRandomString(countT int) string {
	rand.Seed(time.Now().Unix()) // 初始化伪随机序列
	//	准备码表
	baseNT := "0123456789"
	baseT := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + baseNT

	//	初始化字符串
	resultT := ""

	//循环逐个字符生成
	for i := 0; i < countT; i++ {
		idxT := rand.Intn(len(baseT)) // 索引号
		resultT += baseT[idxT:(idxT + 1)]
	}
	return resultT
}

// 字符串 转换成 []byte
// string →[]byte
func ConverStringToByteList(strName string) []byte {
	return []byte(strName)
}

// 判断是文件夹还是文件
func IsDirectoryOrFile(pathName string) string {
	if FileExists(pathName) {
		//存在
		result := IsDirectory(pathName) // 判断是否是文件夹
		if result {
			return "文件夹"
		} else {
			return "文件"
		}
	} else {
		return "路径不存在"
	}
}

// 写入文件//覆盖
func InFile(fileName string, inStr string) int {
	os.MkdirAll(fileName, 0755)      // 无目录自动创建目录。
	os.RemoveAll(fileName)           //删除最里层的目录
	file, err := os.Create(fileName) //创建文件

	if err != nil {
		fmt.Println("打开文件出错，", err.Error())
		return -1 //失败
	}
	defer file.Close()          // 函数结束前关闭文件
	wt := bufio.NewWriter(file) // 创建写入
	wt.WriteString(inStr)       // 写入字符串
	wt.Flush()                  // 更新。
	return 1                    // 成功

}

// 统计文件行数。
func CountFileLines(filePath string) int {
	//统计文件行数。
	fileT, errT := os.Open(filePath)
	if errT != nil {
		fmt.Println("打开文件出错")
		return -1
	}
	defer fileT.Close()
	var i = 0 //行数

	reader := bufio.NewReader(fileT) //读取器
	for {
		//byteList, isPrefix, err := reader.ReadLine() // 每行
		_, isPrefix, err := reader.ReadLine() // 每行 // 统计行不需要数据。
		if err == io.EOF {
			break
		}
		if !isPrefix {
			//读取完一行。
			i++ //行数加一
		}
		//time.Sleep(time.Second * 3)
	}
	return i

}

// 用于标志是否初始化过随机数种子的变量
var ifRandomizedG = false

// Random初始化随机数种子，不会重复操作
func Randomize() {
	if !ifRandomizedG {
		rand.Seed(time.Now().Unix())
		ifRandomizedG = true
	}
}

// GenerateRandomString生成一个可定制的随机字符串
func GenerateRandomString(minCharCountA, maxCharCountA int, hasUpperA, hasLowerA, hasDigitA, hasSpecialCharA, hasSpaceA, hasOtherChars bool) string {
	Randomize()

	if minCharCountA <= 0 {
		return ""
	}
	if maxCharCountA <= 0 {
		return ""
	}

	if minCharCountA > maxCharCountA {
		return ""
	}

	countT := minCharCountA + rand.Intn(maxCharCountA+1-minCharCountA)
	baseT := ""
	if hasUpperA {
		baseT += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if hasLowerA {
		baseT += "abcdefghijklmnopqrstuvwxyz"
	}
	if hasDigitA {
		baseT += "0123456789"
	}
	if hasSpecialCharA {
		baseT += "!@#$%^&*-=[]{}."
	}

	if hasSpaceA {
		baseT += " "
	}

	if hasOtherChars {
		baseT += "/\\:*\"<>|(),+?;"
	}
	rStrT := ""
	var idxT int
	for i := 0; i < countT; i++ {
		idxT = rand.Intn(len(baseT))
		rStrT += baseT[idxT:(idxT + 1)]
	}
	return rStrT

}
