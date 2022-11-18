package uploads

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skr/models"
	"github.com/szyhf/go-excel"
)

func UploadFile(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		log.Println(file.Filename)
		err := c.SaveUploadedFile(file, "./saved/"+file.Filename)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(c.PostForm("key"))
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	ReadFile(files[0].Filename)

}

// defined a struct
type Standard struct {
	// use field name as default column name
	ID int
	// column means to map the column name
	EmpCode string `xlsx:"column(Emp Code)"`
	// you can map a column into more than one field
	EmpCodePtr *string `xlsx:"column(Emp Code)"`
	// omit `column` if only want to map to column name, it's equal to `column(Emp Code)`
	EmpName string `xlsx:"Emp Name"`

	EmpNamePtr *string `xlsx:"column(Emp Name)"`

	Branch string `xlsx:"Branch"`

	BranchPtr *string `xlsx:"column(Branch)"`

	Role string `xlsx:"Role"`

	RolePtr *string `xlsx:"column(Role)"`

	MobileNumber string `xlsx:"Mobile number"`

	MobileNumberPtr *string `xlsx:"column(Mobile number)"`

	Emailid string `xlsx:"Emailid"`

	EmailidPtr *string `xlsx:"column(Emailid)"`

	IDType string `xlsx:"ID type"`

	IDTypePtr *string `xlsx:"column(ID type)"`

	Temp *Temp `xlsx:"column(UnmarshalString)"`
	// support default encoding of json
	//	TempEncoding *TempEncoding `xlsx:"column(UnmarshalString);encoding(json)"`
	// use '-' to ignore.
	Ignored string `xlsx:"-"`
}

// func (this Standard) GetXLSXSheetName() string {
// 	return "Some other sheet name if need"
// }

type Temp struct {
	Foo string
}

func (this Standard) GetXLSXSheetName() string {
	return "Some other sheet name if need"
}

// self define a unmarshal interface to unmarshal string.
func (this *Temp) UnmarshalBinary(d []byte) error {
	return json.Unmarshal(d, this)
}

func ReadFile(fileName string) {
	fmt.Println("=file Name==" + fileName)

	conn := excel.NewConnecter()
	errp := conn.Open("./saved/" + fileName)
	if errp != nil {
		panic(errp)
	}
	defer conn.Close()

	// Generate an new reader of a sheet
	// sheetNamer: if sheetNamer is string, will use sheet as sheet name.
	//             if sheetNamer is int, will i'th sheet in the workbook, be careful the hidden sheet is counted. i ∈ [1,+inf]
	//             if sheetNamer is a object implements `GetXLSXSheetName()string`, the return value will be used.
	//             otherwise, will use sheetNamer as struct and reflect for it's name.
	// 			   if sheetNamer is a slice, the type of element will be used to infer like before.
	rd, err := conn.NewReader("Sheet3")
	if err != nil {
		panic(err)
	}
	defer rd.Close()

	// for rd.Next() {
	// 	var s Standard
	// 	// Read a row into a struct.
	// 	err := rd.Read(&s)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%+v", s)
	// }

	//Read all is also supported.
	var stdList []Standard
	err = rd.ReadAll(&stdList)
	if err != nil {
		panic(err)
		//return
	}
	//fmt.Printf("%+v", stdList)

	for _, val := range stdList {

		//fmt.Println("===Sunil=====", val.Emailid)

		c := models.UserExcel{}
		c.Branch = val.Branch
		c.Empcode = *val.EmpCodePtr
		c.Role = val.Role
		c.Moblienumber = val.MobileNumber
		c.Emailid = val.Emailid
		c.Usertpe = val.IDType
		cr1, errorsave := c.SaveData()

		fmt.Println(cr1)
		if errorsave != nil {
			//return &models.UserExcel{}, errorsave
		}

	}

}
