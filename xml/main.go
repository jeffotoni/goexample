package main

import (
  "fmt"

  "github.com/clbanning/mxj"
  charset "golang.org/x/net/html/charset"
)

var xmlData = []byte(`<?xml version="1.0" encoding="ISO-8859-1"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
  xmlns:ws="http://ws.service.soa.mw.s3wf.com.br"
  xmlns:mes="http://message.ws.connector.soa.mw.s3wf.com.br">
  <soapenv:Header/>
  <soapenv:Body>
    <ws:executeSyncResponse>
      <ws:message xmlns:ws="http://ws.service.soa.mw.s3wf.com.br"
        xmlns:mes="http://message.ws.connector.soa.mw.s3wf.com.br">
        <mes:payload>&lt;pfe-msg&gt;&lt;data&gt;&lt;recharge-number&gt;0&lt;/recharge-number&gt;&lt;/data&gt;&lt;/xpy-msg&gt;</mes:payload>
        <mes:property>
          <mes:name>ClientName</mes:name>
          <mes:value/>
        </mes:property>
        <mes:property>
          <mes:name>ServiceName</mes:name>
          <mes:value>S3WF1</mes:value>
        </mes:property>
        <mes:property>
          <mes:name>ResultCode</mes:name>
          <mes:value>1</mes:value>
        </mes:property>
        <mes:property>
          <mes:name>RemoteId</mes:name>
          <mes:value>X393399X9X9X</mes:value>
        </mes:property>
        <mes:property>
          <mes:name>Version</mes:name>
          <mes:value>1</mes:value>
        </mes:property>
      </ws:message>
    </ws:executeSyncResponse>
  </soapenv:Body>
</soapenv:Envelope>`)

func main() {

  mxj.XmlCharsetReader = charset.NewReaderLabel
  m, err := mxj.NewMapXml(xmlData)
  if err != nil {
    fmt.Println("merr:", err.Error())
    return
  }

  fmt.Println(m)
}
