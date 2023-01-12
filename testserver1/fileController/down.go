package down

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func downPart(wg *sync.WaitGroup, name, url string, client http.Client, start, end int) {
	defer wg.Done()
	part, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer part.Close()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = part.Write(body)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func StratDu(url string) {

	client := http.Client{}

	res, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	urlSplit := strings.Split(url, "/")
	filename := urlSplit[len(urlSplit)-1]

	if res.Header.Get("Accept-Ranges") != "bytes" {
		fmt.Println("unable to download file in multipart")
		return
	}

	cntLen, err := strconv.Atoi(res.Header.Get("Content-Length"))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("cntLen error")
	}

	amDown := 5 * 1024 * 1024
	tot := cntLen / amDown

	wg := sync.WaitGroup{}

	for i := 0; i <= tot; i++ {
		wg.Add(1)
		fmt.Println("Dowmpart:  ", i)
		name := fmt.Sprintf(filename+"%d", i)
		start := i * amDown
		if i == tot {
			end := cntLen - (i-1)*amDown
			//(i + 1) * amDown
			go downPart(&wg, name, url, client, start, end)
		} else {
			end := (i + 1) * amDown
			go downPart(&wg, name, url, client, start, end)
		}

	}

	wg.Wait()

}
