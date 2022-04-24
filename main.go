package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

const targetDevice = "Levolor"

// https://community.smartthings.com/t/make-motorized-levolor-blinds-smarter/168656/54
var unlockService = mustParse("ab529100-5310-483a-b4d3-7f1eaa8134a0")
var unlockCharacteristic = mustParse("ab529101-5310-483a-b4d3-7f1eaa8134a0")
var raiseLowerService = mustParse("ab529200-5310-483a-b4d3-7f1eaa8134a0")
var raiseLowerCharacteristic = mustParse("ab529201-5310-483a-b4d3-7f1eaa8134a0")

func mustParse(uuid string) bluetooth.UUID {
	u, err := bluetooth.ParseUUID(uuid)
	if err != nil {
		panic(err)
	}
	return u
}

type Command byte

const (
	Open  Command = 0x01
	Close Command = 0x02
)

func main() {
	server := NewServer(adapter)

	err := server.Start()
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/send", server.SendCommandEndpoint)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("error starting http server:", err)
		os.Exit(1)
	}
}

type Server struct {
	adapter *bluetooth.Adapter
}

func NewServer(adapter *bluetooth.Adapter) *Server {
	s := &Server{adapter: adapter}
	return s
}

func (s *Server) Start() error {
	retries := 10
	enabled := false
	for i := 0; i < retries; i++ {
		err := s.adapter.Enable()
		if err != nil {
			fmt.Printf("Attempt %d/%d failed to enable adapter: %v\n", i, retries, err)
			time.Sleep(time.Second * 30)
			continue
		} else {
			enabled = true
			break
		}
	}
	if !enabled {
		return fmt.Errorf("failed to enable adapter after %d attempts", retries)
	}
	return nil
}

func (s *Server) SendCommandEndpoint(w http.ResponseWriter, req *http.Request) {
	group, err := strconv.ParseUint(req.FormValue("group"), 10, 8)
	if err != nil {
		fmt.Println("Failed to parse group:", err)
		http.Error(w, "Failed to parse group", http.StatusBadRequest)
		return
	}

	command, err := strconv.ParseUint(req.FormValue("command"), 10, 8)
	if err != nil {
		fmt.Println("Failed to parse command:", err)
		http.Error(w, "Failed to parse command", http.StatusBadRequest)
		return
	}

	err = s.SendCommand(req.Context(), byte(group), Command(command))
	if err != nil {
		fmt.Println("Failed to send command:", err)
		http.Error(w, "Failed to send command", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Sent command")
}

func (s *Server) SendCommand(ctx context.Context, group byte, command Command) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		<-ctx.Done()
		adapter.StopScan()
	}()

	done := false
	// Start scanning.
	fmt.Println("scanning...")
	return adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if done {
			return
		}

		if device.LocalName() == targetDevice {
			err := s.sendCommandToAddress(device.Address, group, command)
			if err != nil {
				fmt.Println("failed to send command:", err)
			} else {
				done = true
				cancel()
			}
		}
	})
}

func (s *Server) sendCommandToAddress(addr bluetooth.Addresser, group byte, command Command) error {
	dev, err := s.adapter.Connect(addr, bluetooth.ConnectionParams{})
	if err != nil {
		return fmt.Errorf("failed to connect to device: %w", err)
	}
	defer dev.Disconnect()

	fmt.Println("Connected to device:", addr.String())

	err = s.writeToCharacteristic(dev, unlockService, unlockCharacteristic, []byte{})
	if err != nil {
		return fmt.Errorf("failed to unlock device: %w", err)
	}
	err = s.writeToCharacteristic(dev, raiseLowerService, raiseLowerCharacteristic,
		[]byte{0xF1, byte(command), group})
	if err != nil {
		return fmt.Errorf("failed to send command: %w", err)
	}

	fmt.Println("Sent command to device:", addr.String())
	return nil
}

func (s *Server) writeToCharacteristic(dev *bluetooth.Device, service, char bluetooth.UUID, data []byte) error {
	services, err := dev.DiscoverServices([]bluetooth.UUID{service})
	if err != nil {
		return err
	}

	if len(services) != 1 {
		return fmt.Errorf("expected 1 service, got %d", len(services))
	}

	chars, err := services[0].DiscoverCharacteristics([]bluetooth.UUID{char})
	if err != nil {
		return err
	}

	if len(chars) != 1 {
		return fmt.Errorf("expected 1 characteristic, got %d", len(chars))
	}

	_, err = chars[0].WriteWithoutResponse(data)
	return err
}
