package resource

import (
	"io"
	"mp-projeto/shared/logger"
	port_shared "mp-projeto/shared/port"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Gateway struct {
	Interval     int
	IntervalInc  int
	Url          string
	log          port_shared.ILogger
	RegisterName string
	Scheme       string
	Host         string
	Port         string
	scIp         string
	scPort       string
	scName       string
}

func NewGateway(scheme, host, port string, interval, intervalInc int, Logger ...port_shared.ILogger) *Gateway {

	Gateway := &Gateway{
		Interval:    interval,
		IntervalInc: intervalInc,
		log:         nil,
		Scheme:      scheme,
		Host:        host,
		Port:        port,
	}

	if len(Logger) > 0 {
		Gateway.log = Logger[0]
	}
	return Gateway
}

func (o *Gateway) ReqId() string {

	return uuid.New().String()
}

func (o *Gateway) GetUrl() string {

	return o.Scheme + "://" + o.Host + ":" + o.Port + "/"
}

func (o *Gateway) logger() port_shared.ILogger {

	if o.log != nil {
		return o.log
	}
	return nil
}

func (o *Gateway) sendMessage(method, hUrl string, body ...io.Reader) (*http.Response, error) {

	var b io.Reader
	reqid := o.ReqId()
	client := &http.Client{}

	if len(body) > 0 {
		b = body[0]
	}

	req, err := http.NewRequest(method, hUrl, b)

	if err != nil {
		return nil, err
	}
	req.Header.Set("reqid", reqid)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *Gateway) register(ip, port, name string) {

	o.scIp = ip
	o.scPort = port
	o.scName = name
	interval := o.Interval
	TimeTicker := time.NewTicker(time.Second * time.Duration(interval))
	TickerChannel := make(chan bool)

	for {
		select {
		case <-TickerChannel:
			return
		case <-TimeTicker.C:
			response, err := o.sendMessage(http.MethodGet, o.GetUrl()+"register?ip="+ip+"&port="+port+"&name="+name)
			if err != nil {
				interval = interval + o.IntervalInc
				o.logger().Error(err.Error())
			} else {
				response.Body.Close()
				o.sendPulse()
				TickerChannel <- true
			}
		}
	}
}

func (o *Gateway) deregister(ip, port, name string) {

	response, err := o.sendMessage(http.MethodGet, o.GetUrl()+"unregistry?ip="+ip+"&port="+port+"&name="+name)

	if err != nil {
		o.logger().Error(err.Error())
	} else {
		body, _ := io.ReadAll(response.Body)
		response.Body.Close()
		o.logger().SetExtraPart(logger.FromPart, "gateway-res").Info(string(body))
	}
}

func (o *Gateway) sendPulse() {

	interval := o.Interval
	TimeTicker := time.NewTicker(time.Second * time.Duration(interval))
	TickerChannel := make(chan bool)

	for {
		select {
		case <-TickerChannel:
			return
		case <-TimeTicker.C:
			response, err := o.sendMessage(http.MethodGet, o.GetUrl()+"heartbeat")
			if err != nil {
				interval = interval + o.IntervalInc
				o.logger().Error(err.Error())
				o.register(o.scIp, o.scPort, o.scName)
				TickerChannel <- true
			} else {
				response.Body.Close()
				interval = o.Interval
			}
			TimeTicker = time.NewTicker(time.Second * time.Duration(interval))
		}
	}
}
