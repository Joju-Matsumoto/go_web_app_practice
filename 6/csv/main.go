package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Joju-Matsumoto/go_web_app_practice/6/memory/models"
)

func main() {
	// ファイル作成
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	posts := []models.Post{
		{Id: 1, Content: "Hello", Author: "joju"},
		{Id: 2, Content: "World", Author: "joju"},
		{Id: 3, Content: "Goalng", Author: "matsumoto"},
		{Id: 4, Content: "Web App", Author: "matsumoto"},
	}

	// csv.Writerで書き込み
	writer := csv.NewWriter(csvFile)
	for _, post := range posts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	// ファイルオープン
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// csv.Readerで読み込み
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // 行内に全てのフィールドが揃っていないくても許容する
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var redPosts []models.Post
	for _, record := range records {
		id, _ := strconv.ParseInt(record[0], 0, 0)
		post := models.Post{Id: int(id), Content: record[1], Author: record[2]}
		redPosts = append(redPosts, post)
	}
	fmt.Println(redPosts)
}
