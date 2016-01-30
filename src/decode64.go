// http://golang.org/pkg/io/ioutil/
// http://golang.org/pkg/encoding/xml/
package main

import (
	"bufio"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type envelop struct {
	XMLName xml.Name `xml:"SigEnvelope"`
	// First letter must be capital. Cannot use `content`
	Objects objct
}

type objct struct {
	XMLName xml.Name `xml:"Object"`
	Content string   `xml:",chardata"`
}

func main() {
	argsWithoutProg := os.Args[1:]
	arg := argsWithoutProg[0]

	xmlFile, err := os.Open(arg)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	e := envelop{}
	//xmlContent, _ := ioutil.ReadFile("Castle.xml")
	//xmlContent, _ := ioutil.ReadAll(xmlFile)

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&e)

	//err = xml.Unmarshal(xmlContent, &d)
	if err != nil {
		panic(err)
	}
	//	fmt.Println("XMLName:", e.XMLName)
	//	fmt.Println("Objects XMLName:", e.Objects.XMLName)
	//fmt.Println("Objects:", e.Objects.Content)
	data, err := base64.StdEncoding.DecodeString(e.Objects.Content)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//fmt.Printf("%q\n", data)
	//fmt.Println(data)
	f, err := os.Create(arg + ".xml")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	n4, err := w.Write(data)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
