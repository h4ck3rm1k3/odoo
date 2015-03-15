package main;
// lifted from http://rosettacode.org/wiki/Walk_a_directory/Recursively#Go
import (
	"fmt"
	"os"
	"path/filepath"
	"encoding2/json"
//	"encoding/xml"
	"io/ioutil"
	"reflect"
	"strings"
)


type Load struct {}

type Manifest struct {
	Application bool
	Author interface{} // string or []string
	AutoInstall bool `json:"auto_install"`
	Active bool `json:"active"`
	Bootstrap bool `json:"bootstrap"`
	Category string
	Url string
	Complexity string
	Certificate interface{} // string or bool
	Data []string // array of file name
	Demo []string // array of file names
	DemoXML []string `json:"demo_xml"`
	Depends []string
	CSS []string
	Images []string
	Description string
	ExternalDependancies	map[string][]string `json:"external_dependencies"`
	//ExternalDependancies	map[string][]interface{} `json:"external_dependancies"`
	//ExternalDependancies	map[string]interface{} `json:"external_dependancies"`
	//ExternalDependancies	interface{} `json:"external_dependancies"`
	Icon string // relative file name
	InitXML []string `json:"init_xml"`
	Installable interface {}//bool or string
	License string
	ModulePath string `json:"module_path"`
	Maintainer string `json:"maintainer"`
	Name string
	PostLoad *Load `json:"post_load"`
	Qweb []string
	Sequence interface{}//int or string
	Summary string
	Test []string
	UpdateXML []string `json:"update_xml"`
	Version string
	Web bool
	Website string
}

func DecodeDict(file []byte) {

	var data map[string]interface {}
	err := json.Unmarshal(file, &data)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	VisitJson(data)
}

func VisitJson(data map[string]interface {} ){
	for k, v := range data {
		fmt.Printf("Key:%s:\n", k)
		//fmt.Printf("pair:%s\t%s\n", k, v)

		switch t := v.(type) {
		case int:
			fmt.Printf("\tInteger: %v\n", t)
		case float64:
			fmt.Printf("\tFloat64: %v\n", t)
		case string:
			fmt.Printf("\tString: %v\n", t)
		case bool:
			fmt.Printf("\tString: %v\n", t)
		case nil:
			fmt.Printf("\tNil: %v\n", t)

		case map[string]interface{}:
			for i,n := range t {
				fmt.Printf("\t\tKV: %v\n", i)
				switch u := n.(type) {
				case []interface {}:
					for j,m := range u { // string
						fmt.Printf("\t\tItem: %v= %v\n", j, m)
					}
				}

			}

		case []interface {}:
			for i,n := range t {
				fmt.Printf("\t\tItem: %v= %v\n", i, n)
			}
		default:
			var r = reflect.TypeOf(t)
			fmt.Printf("\tOther:%v\n", r)
		}
	}

}

func Emit(m Manifest) {
	//for j,m := range m.ExternalDependancies{		fmt.Printf("Raw DEPS:%s %s:\n", j, m)	}	

	output, err2 := json.MarshalIndent(m, "  ", "    ")
	if err2 != nil {
	 	fmt.Printf("error: %v\n", err2)
	}
	fmt.Printf("JSON:")
	os.Stdout.Write(output)
}

// func EmitXML(m Manifest) {

// 	output, err2 := xml.MarshalIndent(m, "  ", "    ")
// 	if err2 != nil {
// 	 	fmt.Printf("error: %v\n", err2)
// 	}
// 	fmt.Printf("XML:")
// 	os.Stdout.Write(output)
// }

func VisitPackageContents(fn string)  {
	//fmt.Printf("Read: %v\n", fn)
	file, e := ioutil.ReadFile(fn)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	list := strings.Split(filepath.Dir(fn),"/")
	ftype := list[len(list)-1]

	var m Manifest
	err1 := json.Unmarshal(file, &m)
	if err1 != nil {
		fmt.Printf("error1:%v data:%s\n", err1, file)
		return
	}
	//fmt.Printf("Raw:%v:\n", m)

	
	if m.ExternalDependancies != nil {
		//fmt.Printf("Raw DEPS:%s:\n", m.ExternalDependancies)
	// 	var r = reflect.TypeOf(m.ExternalDependancies)
	// 	fmt.Printf("\tOther:%v\n", r)
	}
	
	for _,v := range m.Data {

		//fmt.Printf("data: %v %v\n", x, v)
		VisitPackageDataFile(ftype, m.ModulePath + "/"+ v)
	}

	//Emit(m);
	//EmitXML(m);

	//DecodeDict(file)
}

func VisitPackage(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if !!fi.IsDir() {
		return nil // not a file.  ignore.
	}
	matched, err := filepath.Match("__openerp__.json", fi.Name())
	if err != nil {
		fmt.Println(err) // malformed pattern
		return err       // this is fatal.
	}
	if matched {
				
		VisitPackageContents(fp)
	}
	return nil
}

func main() {
	filepath.Walk(".", VisitPackage)
	//VisitPackageContents("addons/auth_ldap/__openerp__.json")
}
