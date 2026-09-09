package usecase

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/icholy/digest"
)

type FBIGetFaceUseCase struct{}

func NewFBIGetFaceUseCase() *FBIGetFaceUseCase {
	return &FBIGetFaceUseCase{}
}

func (uc *FBIGetFaceUseCase) Execute(host string, port int, user, password, url string, timeout int) (Output, error) {

	//LIGA CAPTURA
	go func() {
		time.Sleep(1000 * time.Millisecond)
		//requere usuario cadastrar face
		// http://192.168.15.4/cgi-bin/accessControl.cgi?action=captureCmd&type=1&heartbeat=5&timeout=30
		if url == "" {
			url = "/cgi-bin/accessControl.cgi?action=captureCmd&type=1&heartbeat=5&timeout=" + fmt.Sprintf("%d", timeout)
		}
		_, err := digestRequest(host, port, user, password, url)
		if err != nil {
			return
		}

	}()

	fullUrl := fmt.Sprintf("http://%s:%d/cgi-bin/snapManager.cgi?action=attachFileProc&Flags[0]=Event&Events=[CitizenPictureCompare]", host, port)
	resp, err := openStream(fullUrl, user, password, timeout)
	if err != nil {
		return Output{-1, err.Error()}, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return Output{resp.StatusCode, "Erro ao interpretar Content-Type: " + err.Error()}, err
	}

	boundary, ok := params["boundary"]
	if !ok {
		_, _ = io.ReadAll(resp.Body)

		// fmt.Println("Status:", resp.Status)
		// fmt.Println("Headers:", resp.Header)
		// fmt.Println("Body:")
		// fmt.Println(string(b))

		return Output{
			resp.StatusCode,
			"Boundary não encontrado",
		}, err

		// return Output{resp.StatusCode, "Boundary não encontrado no Content-Type"}, err
	}

	//PROCESSAR SNAPMANAGER PARA OBTER A IMAGEM DA CAPTURA
	reader := multipart.NewReader(resp.Body, boundary)

	var msg string

	for {
		part, err := reader.NextPart()
		if err != nil {
			if errors.Is(err, io.EOF) {
				msg = "Stream encerrado"
				break
			}

			return Output{-1, err.Error()}, err
		}

		// fmt.Println("Next Part OK")

		if finished, err := handlePart(part); finished {
			// fmt.Println("Imagem capturada")

			if err != nil {
				return Output{-1, err.Error()}, err
			}

			return Output{
				resp.StatusCode,
				"Imagem capturada",
			}, nil
		}
	}

	return Output{0, msg}, nil
}

const maxJPEGSize = 10 * 1024 * 1024 // 10 MB

func handlePart(part *multipart.Part) (bool, error) {
	contentType := part.Header.Get("Content-Type")

	// Não é imagem: descarta essa part e continua o stream
	if contentType != "image/jpeg" {
		part.Close()
		return false, nil
	}

	// fmt.Println("header de imagem!")

	file, err := os.Create("fbi_snapshot.jpg")
	if err != nil {
		return true, err
	}
	defer file.Close()

	// fmt.Println("arquivo criado!")

	err = readJPEG(part, file)
	if err != nil {
		// fmt.Println(err.Error())
		return true, err
	}

	// fmt.Println("Imagem salva!")

	return true, nil
}

func readJPEG(r io.Reader, w io.Writer) error {
	buf := make([]byte, 8192)

	var last byte
	var total int64

	for {
		n, err := r.Read(buf)

		if n > 0 {
			data := buf[:n]
			end := n
			foundEnd := false

			// Procura o marcador JPEG EOI: FF D9
			for i := 0; i < n; i++ {
				if last == 0xFF && data[i] == 0xD9 {
					end = i + 1
					foundEnd = true
					break
				}

				last = data[i]
			}

			// Proteção contra imagem muito grande
			if total+int64(end) > int64(maxJPEGSize) {
				return fmt.Errorf(
					"imagem excedeu o tamanho máximo de %d MB",
					maxJPEGSize/(1024*1024),
				)
			}

			// Escreve somente até o FFD9
			if _, err := w.Write(data[:end]); err != nil {
				return err
			}

			total += int64(end)

			// Encontrou o final real do JPEG
			if foundEnd {
				return nil
			}
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}
	}
}

func openStream(url, user, password string, timeout int) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "multipart/x-mixed-replace")

	client := &http.Client{
		Transport: &digest.Transport{
			Username: user,
			Password: password,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(timeout) * time.Second,
					KeepAlive: time.Duration(timeout) * time.Second,
				}).DialContext,
				DisableKeepAlives: false,
			},
		},
		Timeout: time.Duration(timeout) * time.Second, // stream contínuo
	}

	return client.Do(req)
}
