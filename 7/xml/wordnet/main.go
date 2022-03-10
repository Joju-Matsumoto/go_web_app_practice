package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type LexicalEntry struct {
	XMLName xml.Name `xml:"LexicalEntry"`
	Id      string   `xml:"id,attr"`
	Lemma   Lemma    `xml:"Lemma"`
	Senses  []Sense  `xml:"Sense"`
}

type Lemma struct {
	Word string `xml:"writtenForm,attr"`
	Pos  string `xml:"partOfSpeech,attr"`
}

type Sense struct {
	Id     string `xml:"id,attr"`
	Synset string `xml:"synset,attr"`
}

type Lexicon struct {
	XMLName  xml.Name       `xml:"Lexicon"`
	Label    string         `xml:"label,attr"`
	Language string         `xml:"language,attr"`
	Owner    string         `xml:"owner,attr"`
	Version  string         `xml:"version,attr"`
	Entries  []LexicalEntry `xml:"LexicalEntry"`
}

type LexicalResource struct {
	XMLName   xml.Name    `xml:"LexicalResource"`
	Lexicon   Lexicon     `xml:"Lexicon"`
	SenseAxis []SenseAxis `xml:"SenseAxes>SenseAxis"`
}

type SenseAxis struct {
	Id      string   `xml:"id,attr"`
	Rel     string   `xml:"relType,attr"`
	Targets []Target `xml:"Target"`
}

type Target struct {
	Id string `xml:"ID,attr"`
}

func readFile(filename string) []byte {
	xmlData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return xmlData
}

func writeFile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	var lr LexicalResource
	err := xml.Unmarshal(readFile("xml/wordnet/jpn_wn_lmf.xml"), &lr)
	if err != nil {
		fmt.Println("Error Unmarshaling XML data:", err)
		return
	}

	entries := &lr.Lexicon.Entries

	fmt.Println("Words:", len(*entries))

	// xmlData, err := xml.MarshalIndent(lr, "", "    ")
	// if err != nil {
	// 	fmt.Println("Error Marshaling XML data:", xmlData)
	// 	return
	// }
	// writeFile("xml/wordnet/copy.xml", append([]byte(xml.Header), xmlData...))

	// var lex Lexicon

	// xmlFile, err := os.Open("xml/wordnet/sample.xml")
	// if err != nil {
	// 	fmt.Println("Error opening XML file:", err)
	// 	return
	// }
	// defer xmlFile.Close()

	// decoder := xml.NewDecoder(xmlFile)
	// for {
	// 	t, err := decoder.Token()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error decoding XML into tokens:", err)
	// 		return
	// 	}

	// 	switch se := t.(type) {
	// 	case xml.StartElement:
	// 		fmt.Println(se.Attr)
	// 		if se.Name.Local == "Lexicon" {
	// 			decoder.DecodeElement(&lex, &se)
	// 		}
	// 	}
	// }

	// fmt.Println(lex)
}
