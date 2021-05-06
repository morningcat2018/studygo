package book_manage

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/**
init() 方法不能手动调用
只有在其他包中导入当前包时自动调用
没有入参 没有出参
*/
func init() {
	fmt.Println("初始化图书档案")
	readFileToArray()
	printBookList(bookList)
}

type Book struct {
	bookName    string
	bookCode    string
	author      string
	publishYear int
}

func (book *Book) GetBookString() string {
	var str = book.bookCode + "\t" + book.bookName + "\t" + book.author + "\t" + strconv.Itoa(book.publishYear)
	return str + "\n"
}
func (book *Book) PrintBook() {
	fmt.Printf("编号：%s\t", book.bookCode)
	fmt.Printf("书名：%s\t", book.bookName)
	fmt.Printf("作者：%s\t", book.author)
	fmt.Printf("出版时间：%d\n", book.publishYear)
}

var (
	bookList = list.New()
	fileName = "book-list.dat"
)

func DataPanel() {
	// readFileToArray()
	for {
		fmt.Print(getPanelString())
		fmt.Print("请输入指令：")
		var commond int
		fmt.Scanln(&commond)
		if commond == 0 {
			break
		} else if commond == 1 {
			inputBookInfo()
		} else if commond == 2 {
			queryBook()
		} else if commond == 3 {
			deleteBookByCode()
		} else if commond == 4 {
			printBookList(bookList)
		}
	}
}

func getPanelString() string {
	var content = "********************\n"
	content += "\t1. 录取新书\n"
	content += "\t2. 查找书籍\n"
	content += "\t3. 删除数据\n"
	content += "\t4. 打印数据\n"
	content += "\t0. 退出\n"
	content += "********************\n"
	return content
}

func newBook(code, name, author string, publishYear int) *Book {
	// You can safely return a pointer to local variable as a local variable will survive the scope of the function.
	p := Book{bookName: name, bookCode: code, author: author, publishYear: publishYear}
	// p.author = author // 也可赋值
	return &p
}

func printBookList(books *list.List) {
	for e := books.Front(); e != nil; e = e.Next() {
		book := e.Value.(*Book)
		book.PrintBook()
		//fmt.Println("------")
	}
}

func addToList(book *Book) {
	bookList.PushBack(book)
}

func inputBookInfo() {
	var code, name, author string
	var year int
	fmt.Print("请输入书籍编码：")
	fmt.Scanln(&code)
	fmt.Print("请输入书名：")
	fmt.Scanln(&name)
	fmt.Print("请输入作者：")
	fmt.Scanln(&author)
	fmt.Print("请输入出版时间：")
	fmt.Scanln(&year)
	newBook := newBook(code, name, author, year)
	addToList(newBook)
	writeListToFile()
	//PrintBook(book_2)
}

func queryBook() {
	var name string
	queryBookList := list.New()
	fmt.Print("请输入书名：")
	fmt.Scanln(&name)
	for e := bookList.Front(); e != nil; e = e.Next() {
		book := e.Value.(*Book)
		if book.bookName == name {
			queryBookList.PushBack(book)
		}
	}
	if queryBookList.Len() > 0 {
		fmt.Println("查找到的书籍为：")
		printBookList(queryBookList)
	}
}

func deleteBookByCode() {
	var code string
	var deleteCount = 0
	fmt.Print("请输入书籍编码：")
	fmt.Scanln(&code)
	for e := bookList.Front(); e != nil; e = e.Next() {
		book := e.Value.(*Book)
		if book.bookCode == code {
			bookList.Remove(e)
			deleteCount++
		}
	}
	if deleteCount > 0 {
		writeListToFile()
		fmt.Printf("删除了%d本书籍\n", deleteCount)
	} else {
		fmt.Println("未查到需要删除的书籍")
	}
}

func readFileToArray() {
	inputFile, inputError := os.Open(fileName)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n")
		return // exit the function on error
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if inputString != "\n" && len(inputString) > 0 {
			bookInfo := strings.Split(inputString, "\t")
			year := strings.Replace(bookInfo[3], "\r", "", -1)
			year = strings.Replace(year, "\n", "", -1)
			publishYear, err := strconv.Atoi(year)
			if err != nil {
				fmt.Printf("时间数据格式不正确\n")
				return // exit the function on error
			}
			book := newBook(bookInfo[0], bookInfo[1], bookInfo[2], publishYear)
			addToList(book)
		}
		if readerError == io.EOF {
			return
		}
	}
}

func writeListToFile() {
	outputFile, outputError := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	for e := bookList.Front(); e != nil; e = e.Next() {
		book := e.Value.(*Book)
		outputWriter.WriteString(book.GetBookString())
	}
	outputWriter.Flush()
}
