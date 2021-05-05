package service

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"pricingengine"
	"testing"
	"time"
)

func TestSpeed_test(t *testing.T) {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", "http://localhost:3000/generate_pricing", timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		//Bring up the server
		log.Info("Starting the server for test....")
		go Start()
	}
	//Set up the body
	req := pricingengine.GeneratePricingRequest{
		Dob:            "2001-01-01",
		InsuranceGroup: 35,
		LicenseSince:   "2012-08-01",
	}
	reqBody, _ := json.Marshal(req)
	//Set up chi and start it
	startTime := time.Now()
	numOfTests := 3000
	for i := 0; i <= 3000; i++ {
		_, err := http.Post("http://localhost:3000/generate_pricing", "application/json", bytes.NewBuffer(reqBody))
		assert.Nil(t, err)
	}
	endTime := time.Now()
	log.Infof("total time taken for %v requests is %v", numOfTests, endTime.Sub(startTime).Seconds())
	assert.True(t, endTime.Sub(startTime).Seconds() < 10)
}
