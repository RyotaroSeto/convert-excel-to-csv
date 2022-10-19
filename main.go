package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func main() {
	r := gin.Default()
	r.POST("/:filename", func(c *gin.Context) {
		fn := c.Param("filename")

		if err := ValidateParam(fn); err != nil {
			log.Println(err)
			c.JSON(404, createErrResponse("400", "100"))
			return
		}

		if _, err := os.Stat("./file/" + fn + ".xlsx"); err != nil {
			log.Println(err)
			c.JSON(404, createErrResponse("400", "100"))
			return
		}

		if err := ExcelToCSV(os.Stdout, "./file/"+fn+".xlsx", 0); err != nil {
			log.Println(err)
			c.JSON(404, createErrResponse("400", "100"))
			return

		}
		c.JSON(200, gin.H{"message": "ExcelからCSVファイルにエクスポート成功"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, createErrResponse("404", "102"))
	})
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func ValidateParam(fn string) error {
	if fn == "" {
		return fmt.Errorf("パラメータが不正です。")
	}
	return nil
}

func ExcelToCSV(w io.Writer, path string, sheetIndex int) error {
	excel, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	rows, err := excel.Rows(excel.GetSheetName(sheetIndex))
	if err != nil {
		return err
	}
	csvw := csv.NewWriter(w)
	defer csvw.Flush()
	csvFile, err := os.Create("./output/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	for rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			return err
		}
		if err := csvw.Write(cols); err != nil {
			return err
		}
		writer.Write(cols)
	}
	writer.Flush()
	return nil
}
