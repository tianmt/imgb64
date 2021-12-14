package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	help        bool
	b2i, i2b    bool
	image_path  string
	base64_str  string
	base64_path string
	target_path string
)

const IMG_BASE64_STR_HEAD = "data:image/jpg;base64,"

var (
	HELP_INFO              = errors.New("PRINT HELP INFO.")
	PARAM_ERROR            = errors.New("PARAM ERROR.")
	PARAM_EMPTY_ERROR      = errors.New("PARAM EMPTY ERROR.")
	FILE_OPEN_ERROR        = errors.New("FILE OPEN ERROR.")
	FILE_READ_ERROR        = errors.New("FILE READ ERROR.")
	FILE_WRITE_ERROR       = errors.New("FILE WRITE ERROR.")
	FILE_NOT_EXISTS_ERROR  = errors.New("FILE NOT EXISTS.")
	PATH_IS_NOT_FILE_ERROR = errors.New("PATH IS NOT FILE.")
)

func init() {
	flag.BoolVar(&help, "h", false, "help info.")
	flag.BoolVar(&i2b, "i", false, "image to base64 string.")
	flag.BoolVar(&b2i, "b", false, "base64 string to image.")
	flag.StringVar(&image_path, "ip", "", "image path.")
	flag.StringVar(&base64_str, "bs", "", "base64 string.")
	flag.StringVar(&base64_path, "bp", "", "base64 path.")
	flag.StringVar(&target_path, "tp", "", "target path.")

	flag.Usage = usage
}

func preprocessing() error {
	if help {
		return HELP_INFO
	}

	if i2b && !b2i {
		if err := ckFile(image_path); err != nil {
			return err
		}
	} else if b2i && !i2b {
		if (base64_str == "" && base64_path == "") || target_path == "" {
			return PARAM_EMPTY_ERROR
		}
		if base64_path != "" {
			return ckFile(base64_path)
		}
	} else {
		return PARAM_ERROR
	}

	return nil
}

func ckFile(file_full_path string) error {
	st, err := os.Stat(file_full_path)
	if err != nil {
		return FILE_NOT_EXISTS_ERROR
	}
	if st.IsDir() {
		return PATH_IS_NOT_FILE_ERROR
	}

	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: imgb64 [-bih] [-ip image_path] [-bs base64_string] [-bp base64_path] [-tp target_path] \nOptions:\n")
	flag.PrintDefaults()
}

func image2Base64Str() error {
	img, err := ioutil.ReadFile(image_path)
	if err != nil {
		return FILE_READ_ERROR
	}

	base64_str := base64.StdEncoding.EncodeToString(img)

	if target_path != "" {
		if err := ioutil.WriteFile(target_path, []byte(IMG_BASE64_STR_HEAD+base64_str), os.ModePerm); err != nil {
			return FILE_WRITE_ERROR
		}
	}

	return nil
}

func base64Str2Image() error {
	b64_str := ""
	if base64_str != "" {
		b64_str = base64_str
	} else if base64_path != "" {
		data, err := ioutil.ReadFile(base64_path)
		if err != nil {
			return FILE_READ_ERROR
		}
		b64_str = string(data)
	}

	if strings.HasPrefix(b64_str, IMG_BASE64_STR_HEAD) {
		b64_str = strings.Replace(b64_str, IMG_BASE64_STR_HEAD, "", 1)
	}

	img_data, _ := base64.StdEncoding.DecodeString(b64_str)
	img_file, err := os.OpenFile(target_path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return FILE_OPEN_ERROR
	}
	defer img_file.Close()

	if _, err := img_file.Write(img_data); err != nil {
		return FILE_WRITE_ERROR
	}

	return nil
}

func main() {
	flag.Parse()

	// 参数检查
	if err := preprocessing(); err != nil {
		fmt.Println(err.Error())
		usage()
		return
	}

	// 逻辑处理
	if i2b {
		// image 转 base64
		if err := image2Base64Str(); err != nil {
			fmt.Println(err.Error())
		}
	}
	if b2i {
		// base64 转image
		if err := base64Str2Image(); err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Done...")
	return
}
