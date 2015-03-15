/*
read the data files from odoo
*/
package main;
//import "encoding2/xml"
import "encoding/xml"
import "fmt"
import "io/ioutil"
import "os"
import "strings"
import "path/filepath"

type MenuItem struct {
	Name string `xml:"name,attr"`;
	Id string `xml:"id,attr"`;
	Groups string `xml:"groups,attr"`;
	Sequence string `xml:"sequence,attr"`;
}

type Button  struct {
	Name string `xml:"string,attr"`;
}

type HeaderField struct { 
	Name string `xml:"name,attr"`;
	Widget string `xml:"widget,attr"`;
	NoLabel string `xml:"nolabel,attr"`;
}

type Header  struct {
	Button []Button `xml:"button,attr"`;
	Field  []HeaderField `xml:"field"`;
}

type Form  struct {
	Name string `xml:"string,attr"`;
}

type Field struct { 
	Name string `xml:"name,attr"`;
	Body string `xml:",chardata"`
	Form Form `xml:"form"`
	Header Header `xml:"header"`
}

type Record struct {
	Id string `xml:"id,attr"`;
	Model string `xml:"model,attr"`;
	Field []Field `xml:"field"`;
}

type DataObj struct {
	Menuitem []MenuItem `xml:"menuitem"`;
	Record []Record `xml:"record"`;
}

type DataFile struct { // openerp
	Data DataObj `xml:"data"`
}

func decode_csv(name string, fn string, file []byte) {
	//fmt.Printf("csv fn:%v\n", fn)
}

func decode_yml(name string, fn string, file []byte) {
	//fmt.Printf("yaml fn:%v\n", fn)
}

func decode_xml_views(fn string, file []byte) {
	//fmt.Printf("view\n")
}

func decode_xml_security(fn string, file []byte) {
	//fmt.Printf("sec\n")
}

func decode_xml_report(fn string, file []byte) {
	//fmt.Printf("report\n")
}

func decode_xml_data(fn string, file []byte) {
	fmt.Printf("xml fn:%v\n", fn)
	//fmt.Printf("data\n")
	var m DataFile
	//var m map[string] interface {}
	//fmt.Printf("xml fn:%v\n", fn)
	err := xml.Unmarshal(file, &m)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	fmt.Printf("data:%v\n", m)
}

func decode_xml_wizard(fn string, file []byte) {
	//fmt.Printf("wizard\n")
}

func decode_xml_res(fn string, file []byte) {
	//Resource
}

func decode_xml_module(fn string, file []byte) {
}

func decode_xml_project(fn string, file []byte) {
}

func decode_xml_chart(fn string, file []byte) {
}

func decode_xml_edi(fn string, file []byte) {
}

func decode_xml_ir(fn string, file []byte) {
	//Information Repository
	//The Information Repository is used to store data needed by OpenERP to know how to work as an application - to define menus, windows, views, wizards, database tables, etc.
	//https://www.odoo.com/forum/help-1/question/what-mean-ir-ui-view-60174
}

func decode_xml_workflow(fn string, file []byte) {
}


func decode_xml(name string , fn string, file []byte) {

	var dir = filepath.Dir(fn)
	list := strings.Split(dir,"/")
	ftype := list[len(list)-1]

	switch {
	case ftype == "views":
		decode_xml_views(fn, file);

	case ftype == "wizard":
		decode_xml_wizard(fn, file);

	case ftype == "res":
		decode_xml_res(fn, file);

	case ftype == "module":
		decode_xml_module(fn, file);

	case ftype == "report":
		decode_xml_report(fn, file);

	case ftype == "project":
		decode_xml_project(fn, file);

	case ftype == "chart":
		decode_xml_chart(fn, file);

	case ftype == "workflow":
		decode_xml_workflow(fn, file);

	case ftype == "edi":
		decode_xml_edi(fn, file);

	case ftype == "ir":
		decode_xml_ir(fn, file);

	case ftype == "data":
		decode_xml_data(fn, file);

	case ftype == "security":
		decode_xml_security(fn, file);
	default:
		if name == ftype {
			decode_xml_data(fn, file);
		} else {
			//fmt.Printf("dir :%v\n", dir)
			fmt.Printf("unknown :%v %v\n",name, ftype)
		}
		
	}		

	//	for _, v := range list {		fmt.Printf("path :%v\n", v)	}

}

func VisitPackageDataFile(name string, fn string)  {
	//fmt.Printf("Read: %v\n", fn)
	file, e := ioutil.ReadFile(fn)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var ext = filepath.Ext(fn)	
	if ext == ".csv" {
		decode_csv(name, fn, file)
	} else if ext == ".xml" {
		decode_xml(name, fn, file)
	} else if ext == ".yml" {
		decode_yml(name, fn, file)
	} else {
		fmt.Printf("ext:%v\n", ext)
	}

}
