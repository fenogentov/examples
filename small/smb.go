package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	"github.com/hirochachacha/go-smb2"
)

func main() {
	l := getSaveFilesList()
	fmt.Println(l)

	smb := "172.16.10.84:445"

	conn, err := net.Dial("tcp", smb)
	if err != nil {
		fmt.Println("\nconnect", smb, "error", err)
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "user",
			Password: "12345",
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		fmt.Println("\nconnect", smb, "error", err)
	}
	defer s.Logoff()

	smbFS, err := s.Mount("rotabb")
	if err != nil {
		fmt.Println("\nconnect", smb, "error", err)
	}
	defer smbFS.Umount()

	
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("io.Copy", srcFile, "->", dstFile, "error", err)
		return nil
	}
}

func getSaveFilesList() map[string]os.FileInfo {
	files := make(map[string]os.FileInfo)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".go" {
			files[path] = info
		}

		return nil

	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return files
}
