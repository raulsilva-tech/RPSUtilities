package usecase

import (
	"strconv"

	"github.com/IOTechSystems/onvif"
	"github.com/use-go/onvif/device"
)

type RebootUseCase struct{}

func NewRebootUseCase() *RebootUseCase {
	return &RebootUseCase{}
}

func (uc *RebootUseCase) Execute(host string, port int, user, password string) error {

	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    host + ":" + strconv.Itoa(port),
		Username: user,
		Password: password,
		AuthMode: onvif.UsernameTokenAuth, //digest , both . usernametokenauth
	})

	if err != nil {
		//logIt(fmt.Sprintf("NewDevice: Host %s, retornou o erro: %s", host, err.Error()))
		return err
	}

	_, err = dev.CallMethod(device.SystemReboot{})

	if err != nil {
		// log.Println(err)
		//logIt(fmt.Sprintf("System Reboot: Host %s, retornou o erro: %s", host, err.Error()))
		return err
		//} else {
		// fmt.Println(gosoap.SoapMessage(readResponse(getResponse)).StringIndent())
		// logIt(fmt.Sprintf("System Reboot: Host %s", host))
		//fmt.Println("success")
	}

	return nil
}
