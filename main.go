package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	urlverifier "github.com/davidmytton/url-verifier"
)

func DownloadFile(filep string, uploadDir string,url string) (string, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "an error occred -",err
	}
	defer resp.Body.Close()

	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil{
	  return "", err
	}

	// Create destination file path
	filePa := filepath.Join(uploadDir, filep)
  
	// Create the file
	out, err := os.Create(filePa)
	file := out.Name()
	if err != nil {
		return "an error occred -",err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("an error occred -",err)
	}
	return file,nil 
}

func Url(url string) (string, error){
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	verifier := urlverifier.NewVerifier()
	ret, err := verifier.Verify(url)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	  
	fmt.Printf("Result: %+v\n", ret)

	imageNameWithExt := path.Base(url)
	fmt.Println("imageNameWithExt--",imageNameWithExt)

	parts := strings.Split(imageNameWithExt, "?")
	fmt.Println("parts", parts)

	if len(parts) > 0 {
		parts[1] = ""
	}

	newURL := strings.Join(parts, "")
	fmt.Println("new Url", newURL)
	
	fileNameFilterd := filepath.Ext(newURL)
	fmt.Println(fileNameFilterd)
	if !allowedExtensions[fileNameFilterd]{
		fmt.Printf("unsupported file extension: %s", newURL)
	}

	if ret.IsURL {
		file, erro := DownloadFile(newURL, "uploads",url)
		if erro != nil {
			fmt.Println("failed to download the file :", erro)
		}

		fmt.Println("file succefully downloaded,", file)
		example := "filesource/" + file 
		return example, erro
	}

	if !ret.IsURL {
		err := errors.New("invalied url")
		return "",err
	}

	return "and error occred", err
}

func main() {
	_, err := Url("url -of-the-image")
	if err != nil {
		fmt.Println("an error occred", err)
	}
	// return url, nil
}