package copyfile

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Test() {
	start := time.Now()
	// wg := &sync.WaitGroup{}
	// workerPool := 3
	url := "https://odinxclitestbucket.s3.us-east-1.amazonaws.com/fuk7IZtW3xIIy6WyzH0k58zkKs.mp4?response-content-disposition=inline&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEO7%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJGMEQCIHAQ%2Bzok8FVkSj4MnjKvNiBnHhTW88yr94vjodQtj6xJAiBybjv36UhG1UAoyVmZjHZU%2Fk58Flgqg0CWvYWFMhtjByrtAgiH%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAAaDDExMjA5MDYwOTg0NSIMeDJAFUfl%2BsdMCAyOKsECwyAAbs6Ced6kpNeOzblxN9%2BYicIQ2JJfCDwYpH6wVVYBztXmIZSWSG80nvbgZmzTclwSMnta40LoNI1LVHVfeFm8ICliovV9OS7fKv0vkmEdjHiPO%2FuDfal1MxZejjYlbw%2FChKOUAkRTI7NF%2Bqdtn9OVdiRkm0P7B450znGhJhYBPGtmAeNcEKS01TrbDwskE7JJU3ZlLJQZYt4P5jLUrx6ZEgKxZNu%2ByKY3X2aceodkL5u86bKlhGyUlVQVfDHuUZs1eATTsdXphOAFQghsB%2Fsa2S2Wd4MB94I%2F48SnvZssNQGjWrX1QhAZkB%2Fr5M9iNbAkX1KQYixLRo3AdrWAjORR3%2BGJp4T8ypf1S9ob5GCZhLF0%2FPpV4BTwAlDxSeoWUoJCih9nT1rVA2qUsM%2B%2BgSldBwyAv%2F2Hnh%2BGyYiDJ59%2BMISTqrMGOrQCCBXjqAFOB3qLHy8EmRhlz4ZTIajMRL2YPXZZhnAaXz7weDjFjD%2BJshgnhZUFhsz3JxkArzd38H2uKvcS65ej3FAxmmLfeWr1POhbFs2Kkh%2FlGpX6biPi8lT5IxmA4ZgIc4mZGOKqn56yLwYhjCMH9qfpJGjN2DxtiSUhgoNOti5Am%2Fycxjur8rEjmhitNPzGI2PDJwCu7SgZq3iq82m%2BOAYGg%2BhvPp%2Bv2EkeJHklMq2BQxZxFVVT1pJ6ceX%2FfMAOytrqbFqaJGrJP51nUOrvRg5rV0q7ZwNUkKmvGMhxsxUwMEHeCq3AqlZiD%2BPQmhK02CTE5AriCgFfUyYDXH7y9%2B9ou%2BPKgEHo3%2BLtBveRP1iVTlG1YMQ%2FCEkU5qGE7i%2BFhXl22BxgIn%2FcxFWQ3MiJNRZXJ0g%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20240613T055540Z&X-Amz-SignedHeaders=host&X-Amz-Expires=7200&X-Amz-Credential=ASIARUGI67C22ANVRAI5%2F20240613%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=4f374a6867300697092eddf070678bd3dabb8339b32e7b8294ff8b91944fe467"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Header.Get("Accept-Ranges"))
	fmt.Println(resp.ContentLength)
	// contentLength := resp.ContentLength
	if resp.Header.Get("Accept-Ranges") != "" {
		// chunkSize := float32(contentLength) / float32(workerPool)
		// currChunk := chunkSize
		// for i := 0; i < 1; i++ {
		// 	wg.Add(1)
		// 	start := i * int(currChunk)
		// 	end := (i + 1) * int(currChunk)

		handleConnection(&url)

		// }
		// wg.Wait()
		fmt.Printf("took %v \n", time.Since(start))
	} else {
		fmt.Println("Can't perform concurrent download...")
	}
}

func handleConnection(url *string) {
	// defer wg.Done()
	file, err := os.Create(fmt.Sprintf("%d.mp4", 5))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var client *http.Client = &http.Client{}
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Range", "bytes=3614605-")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("content length", resp.ContentLength)
	written, err := io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(written)

}
