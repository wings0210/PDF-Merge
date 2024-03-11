package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"github.com/unidoc/unipdf/v3/common/license"
	"os"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	// 读取所有要合并的PDF文件
	err := license.SetMeteredKey("37e25811aa085b9dc046fcd99e62a6cd4afc49c2dfbc2dd7e834cd7edae8f2a7")
	if err != nil {
		fmt.Printf("无法设置许可证: %s\n", err.Error())
		os.Exit(1)
	}

	//pdfFiles := []string{"/home/autolabor/work/test_go/1.pdf", "/home/autolabor/work/test_go/2.pdf"}
	// 创建一个新的PDF文档
	//outFile := "/home/autolabor/work/test_go/merged.pdf"

	pdfFiles_1, _, err := dlgs.File("选择文件", "", false)
	if err != nil {
		fmt.Printf("无法打开文件选择对话框: %s\n", err.Error())
		os.Exit(1)
	}

	// 检查用户是否取消了选择
	if pdfFiles_1 == "" {
		fmt.Println("用户取消了文件选择")
		os.Exit(0)
	}
	pdfFiles_2, _, err := dlgs.File("选择文件", "", false)
	if err != nil {
		fmt.Printf("无法打开文件选择对话框: %s\n", err.Error())
		os.Exit(1)
	}

	// 检查用户是否取消了选择
	if pdfFiles_2 == "" {
		fmt.Println("用户取消了文件选择")
		os.Exit(0)
	}
	pdfFiles := []string{pdfFiles_1, pdfFiles_2}

	outFile_temp, _, err := dlgs.File("选择文件", "", true)
	if err != nil {
		fmt.Printf("无法打开文件选择对话框: %s\n", err.Error())
		os.Exit(1)
	}
	// 检查用户是否取消了选择
	if outFile_temp == "" {
		fmt.Println("用户取消了文件选择")
		os.Exit(0)
	}
	outFile := outFile_temp + "/merged.pdf"
	outPdf := creator.New()

	// 遍历每个PDF文件，将其内容添加到新的PDF文档中
	for _, file := range pdfFiles {
		err := mergePDF(file, outPdf)
		if err != nil {
			fmt.Printf("无法合并PDF文件 '%s': %s\n", file, err.Error())
			os.Exit(1)
		}
	}

	// 保存合并后的PDF文档
	err = outPdf.WriteToFile(outFile)
	if err != nil {
		fmt.Printf("无法保存合并后的PDF文件: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("PDF文件成功合并为 '%s'\n", outFile)
}

// 合并PDF文件的函数
func mergePDF(file string, outPdf *creator.Creator) error {
	// 打开PDF文件
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// 读取PDF文件内容
	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return err
	}

	// 获取PDF文件的页数
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	// 遍历每页并将其添加到输出PDF文档中
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return err
		}
		outPdf.AddPage(page)
	}

	return nil
}
