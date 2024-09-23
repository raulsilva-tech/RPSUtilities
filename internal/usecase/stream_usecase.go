package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/base"
)

type StreamUseCase struct{}

func NewStreamUseCase() *StreamUseCase {
	return &StreamUseCase{}
}

var (
	ErrNoMediaReturned = errors.New("no media returned")
	ErrRequestTimedOut = errors.New("request timed out")
)

func (uc *StreamUseCase) Execute(url string) error {

	// fmt.Println("Check: " + url)
	c := gortsplib.Client{}

	// u, err := base.ParseURL("rtsp://192.168.1.25:554/user=&password=&channel=1&stream=0.sdp")
	u, err := base.ParseURL(url)
	if err != nil {
		log.Println(url + ": " + fmt.Sprintf("%T\n", err) + " | " + err.Error())
		return err
	}

	err = c.Start(u.Scheme, u.Host)
	if err != nil {
		log.Println(url + ": " + fmt.Sprintf("%T\n", err) + " | " + err.Error())
		return err
	}
	defer c.Close()

	desc, _, err := c.Describe(u)
	if err != nil {
		log.Println(url + ": " + fmt.Sprintf("%T\n", err) + " | " + err.Error())
		return err
	}

	if len(desc.Medias) <= 0 {
		log.Println(url + ": medias vazio")
		return ErrNoMediaReturned
	}

	log.Println("success")
	return nil

}
