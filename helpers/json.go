package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"github.com/cheggaaa/pb/v3"
	"sync"
)

type Data struct{
	Type string `json:"type"`
	ProjectId string `json:"project_id"`
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey string `json:"private_key"`
	ClientEmail string `json:"client_email"`
	ClientId string `json:"client_id"`
	AuthUri string `json:"auth_uri"`
	TokenUri string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl string `json:"client_x509_cert_url"`
}

func Getfiles(pathh string) int{
	d, err := os.Open(pathh)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	i := 0
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".json" {
				i++
			}
		}

	}
	return i
}

func JsonAdd(pathh string, grpmail string){
	d, err := os.Open(pathh)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	count := Getfiles(pathh)
	serv, err := GetService()
	if err != nil {
		log.Fatalf("Unable to create service, Error is : %v", err)
	}
	bar := pb.StartNew(count)
	var wg sync.WaitGroup
	wg.Add(count)
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".json" {
				go func(sem) {
					defer wg.Done()
					// fmt.Println(file.Name())
					var data Data
					dir := path.Join(pathh, file.Name())
					//fmt.Println(dir)
					d, err := ioutil.ReadFile(dir)
					if err != nil {
						fmt.Print(err)
					}
					//fmt.Println(d)
					err = json.Unmarshal(d, &data)
					if err != nil {
						fmt.Println("error:", err)
					}
					Addemails(data.ClientEmail, grpmail, serv)
					bar.Increment()
				}()
			}
		}
	}
	wg.Wait()
	bar.Finish()
}