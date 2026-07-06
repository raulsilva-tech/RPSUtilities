package usecase

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/icholy/digest"
)

type Output struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type FBIGetHashFingerprintUseCase struct{}

func NewFBIGetHashFingerprintUseCase() *FBIGetHashFingerprintUseCase {
	return &FBIGetHashFingerprintUseCase{}
}

func (uc *FBIGetHashFingerprintUseCase) Execute(host string, port int, user, password, url string) (Output, error) {

	//requere para usuario colocar o dedo no leitor
	output, err := digestRequest(host, port, user, password, url)
	if output.Code != 200 || err != nil {
		return output, err
	}

	client := &http.Client{
		Transport: &digest.Transport{
			Username: user,
			Password: password,
		},
		Timeout: time.Second * 0,
	}

	var base64Fingerprint string = ""

	// http://192.168.1.201/cgi-bin/eventManager.cgi?action=attach&codes=[All]&heartbeat=5
	eventManagerURL := fmt.Sprintf("http://%s:%s/cgi-bin/eventManager.cgi?action=attach&codes=[All]&heartbeat=5", host, strconv.Itoa(port))
	// fmt.Println("URL: ", eventManagerURL)
	// Faça a requisição GET
	resp, err := client.Get(eventManagerURL)
	if err != nil {
		return Output{0, err.Error()}, err
	}
	defer resp.Body.Close()

	// fmt.Println("Conectado! Aguardando eventos...")

	scanner := bufio.NewScanner(resp.Body)
	// emptyLines := 0
	// event := make(map[string]string)

	captured := false

outer:
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(">>", line)

		// if strings.TrimSpace(line) == "" {
		// 	emptyLines++
		// 	if emptyLines == 2 {
		// 		// fmt.Println("Evento completo:", event)
		// 		event = make(map[string]string)
		// 		emptyLines = 0
		// 	}
		// 	continue
		// }

		// emptyLines = 0

		// if line == "Code=_FingerPrintCollect_;action=Stop;index=0" {
		// 	base64Fingerprint = "timeout"
		// 	break
		// }

		// Se linha tem formato "key: value"
		if parts := strings.SplitN(line, ": ", 2); len(parts) == 2 {

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			// fmt.Println("key: ", key, ", value: ", value)

			switch key {

			case "\"CaptureTimes\"":
				// fmt.Printf("value: '%s'\n", value)
				if value[0] == '2' {
					captured = true
				}

			case "\"FingerprintData\"":
				if len(value) < 10 {
					base64Fingerprint = "invalid"

				} else {
					//retirando caracteres '"' e '\'
					base64Fingerprint = strings.ReplaceAll(value[1:len(value)-2], "\\", "")
					//fmt.Println("size: ", len(base64Fingerprint))
					break outer
				}
			case "\"Status\"":
				if value == "\"Off\"" {
					// requireFingerprint(host, port, user, password, url)
					// log.Println("OFF")
					break outer
				}
				// case "\"ErrorCode\"":
				// 	if value == "285933616" {
				// 		base64Fingerprint = "invalid"
				// 		// fmt.Println("Essa digital já está cadastrada para outro usuário!")
				// 		break outer
				// 	}

			}
			if len(base64Fingerprint) == 1080 {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return Output{0, err.Error()}, err
	}
	if !captured {
		base64Fingerprint = "timeout"
	}

	return Output{http.StatusOK, base64Fingerprint}, nil
}

func digestRequest(host string, port int, user, password, url string) (Output, error) {
	requireURL := fmt.Sprintf("http://%s:%s%s", host, strconv.Itoa(port), url)

	//implementação para capturar a hash da digital do usuario
	client := &http.Client{
		Transport: &digest.Transport{
			Username: user,
			Password: password,
		},
		Timeout: time.Second * 0,
	}

	// Faça a requisição GET
	// log.Println("URL: ", requireURL)
	resp, err := client.Get(requireURL)

	if err != nil {
		return Output{0, err.Error()}, err
	}
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return Output{resp.StatusCode, string(body)}, err
	}
	// log.Println("URL: ", requireURL, " CHAMADA")

	defer resp.Body.Close()
	return Output{resp.StatusCode, "OK"}, nil
}
