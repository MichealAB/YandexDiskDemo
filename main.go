package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	url2 "net/url"
	"os"
	"path/filepath"
)

type URlforUpload struct {
	OperationId string `json:"operation_id"`
	Href        string `json:"href"`
	Method      string `json:"method"`
	Templated   bool   `json:"templated"`
}

type URLforDown struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

func main() {
	GetURLForUpload()
	GetUrlForDownload()
}

func GetURLForUpload() {
	YandexDiskToken := os.Getenv("TOKEN")
	FileName := "красивый лес.jpeg"
	url := "https://cloud-api.yandex.net/v1/disk/resources/upload?path=/" + FileName
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", YandexDiskToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	var GetHref URlforUpload
	_ = json.Unmarshal(body, &GetHref)
	HrefForUpload := GetHref.Href
	fmt.Println(HrefForUpload)
	UpLoadFileOnDisk(HrefForUpload)
}
func UpLoadFileOnDisk(urlHref string) {
	filePath := "C:/Users/Михаил/Desktop/qwer.jpg"
	url := urlHref
	method := "PUT"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(filePath)
	if errFile1 != nil {
		log.Fatalln("File not found", errFile1)
	}
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base("/C:/Users/Михаил/Desktop/qwer.jpg"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func GetUrlForDownload() {
	YandexDiskToken := os.Getenv("TOKEN")
	DiskFilePath := "/красивый лес.jpeg"
	url := "https://cloud-api.yandex.net/v1/disk/resources/download?path=" + url2.QueryEscape(DiskFilePath)
	fmt.Println(url)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", YandexDiskToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	var GetHrefDown URLforDown
	_ = json.Unmarshal(body, &GetHrefDown)
	HrefForDownload := GetHrefDown.Href
	fmt.Println(HrefForDownload)
	DownloadOnPC(HrefForDownload)
}

func DownloadOnPC(urlHref string) {
	fmt.Println(urlHref)
	url := urlHref
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "yandexuid=369435711663506784")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err1 := os.WriteFile("./codejpeg.txt", body, 0666)
	if err != nil {
		fmt.Println("Не удалось записать в файл", err1)
	}
}
