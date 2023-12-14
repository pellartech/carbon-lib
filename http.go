package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EventMessage struct {
	Event     string                 `json:"event"`
	Timestamp string                 `json:"timestamp"`
	Extra     map[string]interface{} `json:"extra,omitempty"`
}

type SignedEventRequest struct {
	Data      EventMessage `json:"data"`
	Signature string       `json:"signature"`
}

func assembleMessage(eventName string, extra map[string]interface{}) EventMessage {
	return EventMessage{
		Event:     eventName,
		Timestamp: time.Now().Format(time.RFC3339),
		Extra:     extra,
	}
}

func preparePayload(privateKey string, msg EventMessage) (SignedEventRequest, error) {
	messageBytes, err := json.Marshal(msg)
	if err != nil {
		return SignedEventRequest{}, fmt.Errorf("error marshalling message: %v", err)
	}

	signature, err := SignMessage(privateKey, string(messageBytes))
	if err != nil {
		return SignedEventRequest{}, fmt.Errorf("error signing message: %v", err)
	}

	return SignedEventRequest{
		Data:      msg,
		Signature: signature,
	}, nil
}

func sendHTTPRequest(baseURL, publicKey string, payload SignedEventRequest) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling payload: %v", err)
	}

	requestURL := fmt.Sprintf("%s/accounts/%s/earn", baseURL, publicKey)
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response: %s", resp.Status)
	}

	return nil
}

func SendSignedEvent(baseURL string, privateKey string, eventName string, extra map[string]interface{}) error {
	publicKey, err := PublicKeyFromPrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("error deriving public key: %v", err)
	}

	msg := assembleMessage(eventName, extra)
	payload, err := preparePayload(privateKey, msg)
	if err != nil {
		return err
	}

	return sendHTTPRequest(baseURL, publicKey, payload)
}
