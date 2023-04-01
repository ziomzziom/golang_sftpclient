package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func main() {
	// Get SFTP credentials
	username := getSftpUsername()
	password := getSftpPassword()
	serverIP := getServerIP()

	// Connect to server
	conn, err := ssh.Dial("tcp", serverIP+":22", &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.PublicKeysCallback(agent.NewClient),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()

	// Open SFTP session
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println("Failed to open SFTP session:", err)
		return
	}
	defer sftpClient.Close()

	// Implement rate limiting and IP blocking
	var (
		mutex    sync.Mutex
		attempts = make(map[string]int)
	)

	for {
		// Get file to upload
		filename := getFilename()
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Failed to open file:", err)
			continue
		}
		defer file.Close()

		// Check if IP is blocked
		ip := getIPAddress()
		mutex.Lock()
		if attempts[ip] >= 3 {
			fmt.Println("IP", ip, "is blocked")
			mutex.Unlock()
			time.Sleep(1 * time.Minute)
			continue
		}
		mutex.Unlock()

		// Upload file
		remoteFile, err := sftpClient.Create(filename)
		if err != nil {
			fmt.Println("Failed to create remote file:", err)
			continue
		}
		defer remoteFile.Close()

		_, err = remoteFile.ReadFrom(file)
		if err != nil {
			fmt.Println("Failed to upload file:", err)
			continue
		}

		// Increment attempts and sleep to implement rate limiting
		mutex.Lock()
		attempts[ip]++
		mutex.Unlock()
		time.Sleep(10 * time.Second)

		fmt.Println("File", filename, "uploaded successfully")
	}
}

func getSftpUsername() string {
	fmt.Print("Enter SFTP username: ")
	var username string
	fmt.Scanln(&username)
	return username
}

func getSftpPassword() string {
	fmt.Print("Enter SFTP password: ")
	var password string
	fmt.Scanln(&password)
	return password
}

func getServerIP() string {
	fmt.Print("Enter server IP address: ")
	var serverIP string
	fmt.Scanln(&serverIP)
	return serverIP
}

func getFilename() string {
	fmt.Print("Enter filename to upload: ")
	var filename string
	fmt.Scanln(&filename)
	return filename
}

func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Failed to get IP address:", err)
		return ""
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}

	return ""
}
