package vilin_parser

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type TestErrorParse struct {
	DataFilePath string
	TestErr      string
}

func TestErrorMainParse(t *testing.T) {
	testingCase := []TestErrorParse{
		{DataFilePath: "dummyPath",
			TestErr: "can`t open file in dataFromFile(). Error is: open dummyPath: no such file or directory in scannerFromFile() in Parse()"},
		{DataFilePath: "data_tests_source/data.dummyext",
			TestErr: "func Parse can handle 'csv' and 'prn' files, extension .dummyext is not implemented"},
	}

	for number, val := range testingCase {
		resErr := Parse(val.DataFilePath)
		if !strings.Contains(resErr.Error(), val.TestErr) {
			t.Errorf("unexpected error in case [%v] \n Expected: %v \n Got: %v", number, val.TestErr, resErr)
		}
	}
}

type TestParse struct {
	DataFilePath  string
	HtmlFilePatch string

	TestHtml string
}

func TestMainParse(t *testing.T) {
	testingCase := []TestParse{
		{DataFilePath: "data_source/data.prn",
			HtmlFilePatch: "result_html/prn_table.html",
			TestHtml:      "<head>\n    <meta charset=\"UTF-8\"/>\n    <title></title>\n</head>\n<table>\n\t<tr>\n\t\t<td>First name</td>\n\t\t<td>Address</td>\n\t\t<td>Postcode</td>\n\t\t<td>Mobile</td>\n\t\t<td>Limit</td>\n\t\t<td>Birthday</td>\n\t</tr>\n\t<tr>\n\t\t<td>Oliver</td>\n\t\t<td>Via Archimede, 103-91</td>\n\t\t<td>2343aa</td>\n\t\t<td>000 1119381</td>\n\t\t<td>6000000</td>\n\t\t<td>19570101</td>\n\t</tr>\n\t<tr>\n\t\t<td>Harry</td>\n\t\t<td>Leonardo da Vinci 1</td>\n\t\t<td>4532 AA</td>\n\t\t<td>010 1118986</td>\n\t\t<td>10433301</td>\n\t\t<td>19751203</td>\n\t</tr>\n\t<tr>\n\t\t<td>Jack</td>\n\t\t<td>Via Rocco Chinnici 4d</td>\n\t\t<td>3423 ba</td>\n\t\t<td>0313-111475</td>\n\t\t<td>93543</td>\n\t\t<td>19740604</td>\n\t</tr>\n\t<tr>\n\t\t<td>Noah</td>\n\t\t<td>Via Giannetti, 4-32</td>\n\t\t<td>2340 CC</td>\n\t\t<td>28932222</td>\n\t\t<td>34</td>\n\t\t<td>19940906</td>\n\t</tr>\n\t<tr>\n\t\t<td>Charlie</td>\n\t\t<td>Via Aldo Moro, 7</td>\n\t\t<td>3209 DD</td>\n\t\t<td>30-34563332</td>\n\t\t<td>4531</td>\n\t\t<td>19981107</td>\n\t</tr>\n\t<tr>\n\t\t<td>Mia</td>\n\t\t<td>Via Due Giugno, 12-1</td>\n\t\t<td>4220 EE</td>\n\t\t<td>43433344329</td>\n\t\t<td>9087</td>\n\t\t<td>19700515</td>\n\t</tr>\n\t<tr>\n\t\t<td>Lily</td>\n\t\t<td>Arcisstraße 21</td>\n\t\t<td>12343</td>\n\t\t<td>&#43;44 728 343434</td>\n\t\t<td>765599</td>\n\t\t<td>19971003</td>\n\t</tr>\n</table>",
		},
		{DataFilePath: "./data_source/data.csv",
			HtmlFilePatch: "result_html/csv_table.html",
			TestHtml:      "<head>\n    <meta charset=\"UTF-8\"/>\n    <title></title>\n</head>\n<table>\n\t<tr>\n\t\t<td>Name</td>\n\t\t<td>Address</td>\n\t\t<td>Postcode</td>\n\t\t<td>Mobile</td>\n\t\t<td>Limit</td>\n\t\t<td>Birthday</td>\n\t</tr>\n\t<tr>\n\t\t<td>&#34;Oliver, El&#34;</td>\n\t\t<td>&#34;Via Archimede, 103-91&#34;</td>\n\t\t<td>2343aa</td>\n\t\t<td>000 1119381</td>\n\t\t<td>6000000</td>\n\t\t<td>01/01/1999</td>\n\t</tr>\n\t<tr>\n\t\t<td>&#34;Harry&#34;</td>\n\t\t<td>Leonardo da Vinci 1</td>\n\t\t<td>4532 AA</td>\n\t\t<td>010 1118986</td>\n\t\t<td>343434</td>\n\t\t<td>31/12/1965</td>\n\t</tr>\n\t<tr>\n\t\t<td>&#34;Jack&#34;</td>\n\t\t<td>&#34;Via Rocco Chinnici 4d&#34;</td>\n\t\t<td>3423 ba</td>\n\t\t<td>0313-111475</td>\n\t\t<td>22</td>\n\t\t<td>05/04/1984</td>\n\t</tr>\n\t<tr>\n\t\t<td>&#34;Noah&#34;</td>\n\t\t<td>&#34;Via Giannetti, 4-32&#34;</td>\n\t\t<td>2340 CC</td>\n\t\t<td>28932222</td>\n\t\t<td>434</td>\n\t\t<td>03/10/1964</td>\n\t</tr>\n\t<tr>\n\t\t<td>&#34;Charlie&#34;</td>\n\t\t<td>&#34;Via Aldo Moro, 7&#34;</td>\n\t\t<td>3209 DD</td>\n\t\t<td>30-34563332</td>\n\t\t<td>343.8</td>\n\t\t<td>04/10/1954</td>\n\t</tr>\n\t<tr>\n\t\t<td>&#34;Mia&#34;</td>\n\t\t<td>&#34;Via Due Giugno, 12-1&#34;</td>\n\t\t<td>4220 EE</td>\n\t\t<td>43433344329</td>\n\t\t<td>6343.6</td>\n\t\t<td>10/08/1980</td>\n\t</tr>\n\t<tr>\n\t\t<td>&#34;Lilly&#34;</td>\n\t\t<td>Arcisstraße 21</td>\n\t\t<td>12343</td>\n\t\t<td>&#43;44 728 343434</td>\n\t\t<td>34342.3</td>\n\t\t<td>20/10/1997</td>\n\t</tr>\n</table>"},
	}

	for number, val := range testingCase {
		resErr := Parse(val.DataFilePath)
		if resErr != nil {
			t.Errorf("unexpected error in case [%v] \n Got: %v", number, resErr)
		}
		file, err := os.Open(val.HtmlFilePatch)
		if err != nil {
			t.Errorf("error open file %v in case [%v]", val.HtmlFilePatch, number)
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("error ReadAll file %v in case [%v]", val.HtmlFilePatch, number)
		}

		if !(string(fileBytes) == val.TestHtml) {
			t.Errorf("wrong recorded html in case[%v] \n Expected: %v \n Got: %v", number, val.TestHtml, string(fileBytes))
		}
	}
}
