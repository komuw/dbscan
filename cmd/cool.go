package main

import (
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/komuw/sir/pkg"
	"github.com/pkg/errors"
)

func priSecForward(requestBytes []byte, remoteAddr string, reqResp *sir.RequestsResponse) {
	backendConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		err = errors.Wrapf(err, "dial failed for address %s of backend %v", remoteAddr, reqResp.Backend)
		log.Fatalf("%+v", err)
	}
	defer backendConn.Close()
	err = backendConn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		err = errors.Wrapf(err, "unable to set backendConn deadline of backend %v", reqResp.Backend)
		log.Fatalf("%+v", err)
	}
	log.Printf("frontend connected to backend %v(%v)", reqResp.Backend, remoteAddr)

	_, err = backendConn.Write(requestBytes)
	if err != nil {
		err = errors.Wrapf(err, "backendConn.Write of backend %v failrd", reqResp.Backend)
		log.Fatalf("%+v", err)
	}
	reqResp.HandleRequest(requestBytes)
	log.Printf("we sent request to backend %v(%v) \n %v", reqResp.Backend, remoteAddr, string(requestBytes))

	responseBytes, err := ioutil.ReadAll(backendConn)
	if err != nil {
		err = errors.Wrapf(err, "unable to read & log response of backend %v(%v)", reqResp.Backend, remoteAddr)
		log.Fatalf("%+v", err)
	}
	reqResp.HandleResponse(responseBytes)
	log.Printf("we got response from backend %v(%v) \n %v", reqResp.Backend, remoteAddr, string(responseBytes))

	// requestBuf := new(bytes.Buffer)
	// responseBuf := new(bytes.Buffer)
	// ch := make(chan bool)

	// // forward data from client to server
	// go func() {
	// 	tee := io.TeeReader(frontendConn, requestBuf)
	// 	io.Copy(backendConn, tee)
	// 	ch <- true
	// }()

	// // forward data from server to client
	// go func() {
	// 	tee := io.TeeReader(backendConn, responseBuf)
	// 	io.Copy(frontendConn, tee)
	// 	ch <- true
	// }()

	// <-ch
	// <-ch
	// //////////////////////////////////// LOG REQUEST  & RESPONSE ////////////////////////
	// requestBytes, err := ioutil.ReadAll(requestBuf)
	// if err != nil {
	// 	err = errors.Wrap(err, "unable to read & log request")
	// 	log.Fatalf("%+v", err)
	// }
	// requestBytes = bytes.Trim(requestBytes, sir.NulByte)
	// reqResp.HandleRequest(requestBytes)
	// log.Printf("we sent request to backend %v \n %v", reqResp.Backend, string(requestBytes))

	// responseBytes, err := ioutil.ReadAll(responseBuf)
	// if err != nil {
	// 	err = errors.Wrapf(err, "unable to read & log response of backend %v", reqResp.Backend)
	// 	log.Fatalf("%+v", err)
	// }
	// reqResp.HandleResponse(responseBytes)
	// log.Printf("we got response from backend %v \n %v", reqResp.Backend, string(responseBytes))
	// //////////////////////////////////// LOG REQUEST  & RESPONSE ////////////////////////

	// reqResp.L.Lock()
	// reqResp.NoOfAllRequests++
	// reqResp.NoOfAllResponses++
	// log.Printf("lengthOfLargestRequest for backend %v %v", reqResp.Backend, reqResp.LengthOfLargestRequest)
	// log.Printf("lengthOfLargestResponse for backend %v %v", reqResp.Backend, reqResp.LengthOfLargestResponse)
	// reqResp.L.Unlock()
}
