package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      int      `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	// Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   int    `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlfile, err := os.Open("xml/simple/simple.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlfile.Close()

	xmlData, err := ioutil.ReadAll(xmlfile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)

	saveData, err := xml.MarshalIndent(post, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling XML data:", err)
		return
	}

	// saveFile, err := os.Create("xml/simple/copy.xml")
	// if err != nil {
	// 	fmt.Println("Error opening XML file:", err)
	// 	return
	// }
	// defer saveFile.Close()

	// _, err = saveFile.WriteString(xml.Header + string(saveData))
	// if err != nil {
	// 	fmt.Println("Error writing XML data:", err)
	// 	return
	// }

	err = ioutil.WriteFile("xml/simple/copy.xml", append([]byte(xml.Header), saveData...), 0644)
	if err != nil {
		fmt.Println("Error WriteFile:", err)
	}
}
