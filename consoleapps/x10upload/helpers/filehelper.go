package x10upload

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tk "github.com/eaciit/toolkit"
)

func MoveFile(From string, To string) {
	args := []string{"mv", From, To}
	args0 := strings.Join(args, " ")

	if err := exec.Command("/bin/sh", "-c", args0).Run(); err != nil {
		tk.Printf("Error: %#v\n", err.Error())
	}
}

func CopyFile(From string, To string) {
	args := []string{"cp", From, To}
	args0 := strings.Join(args, " ")

	if err := exec.Command("/bin/sh", "-c", args0).Run(); err != nil {
		tk.Printf("Error: %#v\n", err.Error())
	}
}

func UrlReplacer(Url string) string {
	formattedstring := strings.Replace(Url, " ", "\\ ", -1)
	return formattedstring
}

func ProcessFile(inbox string, process string, failed string, success string, reporttype string, webapps string) {

	inboxfolder, _ := ioutil.ReadDir(inbox)
	if len(inboxfolder) > 0 {
		for _, f := range inboxfolder {
			err := ConvertPdfToXml(inbox, process, f.Name())
			if err != nil {
				tk.Println(err.Error())
			}
			filename := strings.TrimRight(f.Name(), ".pdf")
			xmlfilename := filename + ".xml"

			DeleteFile(".png", process)
			DeleteFile(".jpg", process)

			ExtractPdfDataCibilReport(process, process, f.Name(), reporttype, xmlfilename, inbox, success, failed, webapps)

		}
	}

}

func DeleteFile(ext string, folder string) {
	procinbox, _ := ioutil.ReadDir(folder)
	for _, files := range procinbox {
		if filepath.Ext(files.Name()) == ext {
			os.RemoveAll(folder + "/" + files.Name())
		}
	}
}
