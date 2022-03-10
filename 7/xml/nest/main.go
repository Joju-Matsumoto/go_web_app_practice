package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       int       `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Comments []Comment `xml:"comments>comment"`
	// Xml      string    `xml:",innerxml"`
}

type Author struct {
	Id   int    `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      int    `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlData, err := ioutil.ReadFile("xml/nest/nest.xml")
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	var post Post
	err = xml.Unmarshal(xmlData, &post)
	if err != nil {
		fmt.Println("Error unmarshaling XML data:", err)
		return
	}

	fmt.Println(post)

	saveData, err := xml.MarshalIndent(post, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling XML data:", err)
		return
	}
	fmt.Println(string(saveData))

	err = ioutil.WriteFile("xml/nest/copy.xml", append([]byte(xml.Header), saveData...), 0644)
	if err != nil {
		fmt.Println("Error save XML data:", err)
		return
	}
}
