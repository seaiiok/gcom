package gxml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type XmlElement struct {
	StartElementSpace string
	StartElementLocal string
	StartAttr         []attr
	TagText           string
	EndElementSpace   string
	EndElementLocal   string
}

type attr struct {
	Space string
	Local string
	Value string
}

func XmlAsRow(xmlFile string) ([]XmlElement, error) {

	xmls := make([]XmlElement, 0)
	xmle := XmlElement{}
	xmle.StartAttr = make([]attr, 0)
	att := attr{}

	f, err := os.OpenFile(xmlFile, os.O_RDONLY, 0744)
	if err != nil {
		return nil, err
	}
	decoder := xml.NewDecoder(f)

	for {
		token, err := decoder.Token() //读取一个标签或者文本内容
		if err == io.EOF {
			return xmls, nil
		}
		if err != nil {
			return nil, err
		}
		switch tp := token.(type) { //读取的TOKEN可以是以下三种类型：StartElement起始标签，EndElement结束标签，CharData文本内容
		case xml.StartElement:
			se := xml.StartElement(tp) //强制类型转换
			xmle.StartElementSpace = stringTrimSpaceAndPrefixBom(se.Name.Space)
			xmle.StartElementLocal = stringTrimSpaceAndPrefixBom(se.Name.Local)

			for _, v := range se.Attr {
				att.Space = stringTrimSpaceAndPrefixBom(v.Name.Space)
				att.Local = stringTrimSpaceAndPrefixBom(v.Name.Local)
				att.Value = stringTrimSpaceAndPrefixBom(v.Value)
				xmle.StartAttr = append(xmle.StartAttr, att)
			}

		case xml.EndElement:
			ee := xml.EndElement(tp)
			xmle.EndElementSpace = stringTrimSpaceAndPrefixBom(ee.Name.Space)
			xmle.EndElementLocal = stringTrimSpaceAndPrefixBom(ee.Name.Local)

		case xml.CharData: //文本数据，注意一个结束标签和另一个起始标签之间可能有空格
			cd := xml.CharData(tp)
			if data := stringTrimSpaceAndPrefixBom(string(cd)); len(data) != 0 {
				xmle.TagText = data
			}
		default:
			fmt.Println("exec --- default")
		}
		xmls = append(xmls, xmle)
	}
}

func stringTrimSpaceAndPrefixBom(v string) string {
	return strings.TrimSpace(strings.TrimPrefix(v, string([]byte{239, 187, 191})))
}
