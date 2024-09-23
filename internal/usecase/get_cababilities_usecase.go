package usecase

import (
	"strconv"

	"github.com/IOTechSystems/onvif"
	"github.com/use-go/onvif/device"
)

type GetCapabilitiesUseCase struct{}

func NewGetCapabilitiesUseCase() *GetCapabilitiesUseCase {
	return &GetCapabilitiesUseCase{}
}

func (uc *GetCapabilitiesUseCase) Execute(host string, port int, user, password string) error {

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

	getCapabilities := device.GetCapabilities{Category: "All"}
	_, err = dev.CallMethod(getCapabilities)

	if err != nil {
		// log.Println(err)
		//logIt(fmt.Sprintf("Get Capabilities: Host %s, retornou o erro: %s", host, err.Error()))
		return err
		//} else {
		// fmt.Println(gosoap.SoapMessage(readResponse(getResponse)).StringIndent())
		//fmt.Println("success")
	}

	return nil
}
