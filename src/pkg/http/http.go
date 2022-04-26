package http

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func CallRemote() error {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl -d "-i /home/app/function/assets/video1-tom_fisk_pexels_id5210841.mp4"  127.0.0.1:8080/function/pymemo

	//bodyContent := "-i /home/app/function/assets/video1-tom_fisk_pexels_id5210841.mp4"
	//body := strings.NewReader(bodyContent)
	req, err := http.NewRequest("POST", "http://localhost:8080/function/pymemo", nil)

	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Println(err)
		return err
	}

	print(string(respDump))
	defer resp.Body.Close()

	return nil
}
