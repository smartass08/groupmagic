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
	"golang.org/x/sync/semaphore"
	"google.golang.org/api/admin/directory/v1"
	"sync"
	"context"
	"time"
)

var wg sync.WaitGroup
const LIMIT int64 = 30

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
		log.Fatal(err)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".json" {
				count++
			}
		}
	}
	return count
}

func AddEmailByFile(filePath string, groupMail string, srv *admin.Service) bool {
	fmt.Printf("[AddEmail]: %s\n",filePath)
	defer wg.Done()
	var data Data
	d, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return false
	}
	err = json.Unmarshal(d, &data)
	if err != nil {
		log.Println(err)
		return false
	}
	Addemails(data.ClientEmail, groupMail, srv)
	return true
}

func JsonAdd(dirPath string, groupMail string){
	d, err := os.Open(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	files, err := d.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	count := Getfiles(dirPath)
	sem := semaphore.NewWeighted(LIMIT)
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	service,err := GetService()
	if err != nil {
		log.Fatalf("Unable to create service, Error is : %v", err)
	}
	bar := pb.StartNew(count)
	for _,file := range files{
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".json" {
				filePath := path.Join(dirPath, file.Name())
				go EmailGoroutine(sem,ctx,filePath,groupMail,service,bar)
				wg.Add(1)
			}
		}
	}
	wg.Wait()
	bar.Finish()
}

func EmailGoroutine(sem *semaphore.Weighted, ctx context.Context, filePath string, groupMail string, service *admin.Service, bar *pb.ProgressBar) {
	sem.Acquire(ctx,1)
	AddEmailByFile(filePath,groupMail,service)
	bar.Increment()
	sem.Release(1)
}
