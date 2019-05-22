package mqtt

import (
	"goc/logface"
	"time"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
)

var log = logface.New(logface.InfoLevel)

type Cfg struct {
	Topic       string
	IsSubscribe bool
	Qos         int
	Broker      string
	Username    string
	Password    string
	ClientID    string
	AliveTime   int
	Timeout     int
	IsClean     bool
	Will        string `json:"Will,omitempty"`
	ConCb       func() `json:"ConCb,omitempty"`
	DisconCb    func() `json:"DisconCb,omitempty"`
}

type Handle struct {
	client pahoMqtt.Client
	option *pahoMqtt.ClientOptions
	token  *pahoMqtt.Token
	chRecv chan (string)
	conf   Cfg
}

func New(cfg Cfg) *Handle {

	m := &Handle{chRecv: make(chan string, 10)}
	m.conf = cfg
	log.Debug("mqtt cfg:[%+v]", m.conf)
	m.option = pahoMqtt.NewClientOptions()

	m.option.SetAutoReconnect(true)
	m.option.AddBroker(m.conf.Broker)
	m.option.SetClientID(m.conf.ClientID)
	m.option.SetPassword(m.option.Password)
	m.option.SetUsername(m.option.Username)
	m.option.SetCleanSession(m.conf.IsClean)
	m.option.SetConnectTimeout(time.Second * time.Duration(m.conf.Timeout))
	m.option.SetKeepAlive(time.Second * time.Duration(m.conf.AliveTime))

	if len(m.conf.Will) > 0 {
		m.option.SetWill(m.conf.Topic, m.conf.Will, 1, false)
	}

	m.option.SetOnConnectHandler(func(c pahoMqtt.Client) {
		if cfg.ConCb != nil {
			cfg.ConCb()
		}
	})

	m.option.SetConnectionLostHandler(func(c pahoMqtt.Client, err error) {
		if cfg.DisconCb != nil {
			cfg.DisconCb()
		}
	})

	m.client = pahoMqtt.NewClient(m.option)
	log.Debug("mqtt client init success")
	return m
}

func (p *Handle) Connect() error {
	if token := p.client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("mqtt connect fail")
	}

	var recvHandle = func(client pahoMqtt.Client, msg pahoMqtt.Message) {
		log.Trace("mqtt recv payload [%s]", string(msg.Payload()))
		p.chRecv <- string(msg.Payload())
	}

	if p.conf.IsSubscribe {
		if token := p.client.Subscribe(p.conf.Topic, byte(p.conf.Qos), recvHandle); token.Wait() && token.Error() != nil {
			log.Panic("mqtt subscribe fail")
		}
	}

	log.Debug("mqtt connect success")
	return nil
}

func (p *Handle) DisConnect() error {
	if p.conf.IsSubscribe {
		if token := p.client.Unsubscribe(p.conf.Topic); token.Wait() && token.Error() != nil {
			log.Panic("mqtt unsubscribe fail")
		}
	}
	p.client.Disconnect(250)
	log.Debug("mqtt disconnect success")
	return nil
}

func (p *Handle) Send(str string) error {
	if token := p.client.Publish(p.conf.Topic, byte(p.conf.Qos), false, str); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	log.Debug("mqtt send success")
	return nil
}

func (p *Handle) Recv() string {
	return <-p.chRecv
}
