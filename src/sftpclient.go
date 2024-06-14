package main

import (
    "fmt"
    "io"
    "os"
    "golang.org/x/crypto/ssh"
    "github.com/pkg/sftp"
    "log"
)

func main() {
    var (
        username string
        password string
        server   string
        port     string
        filePath string
        action   string
    )

    fmt.Print("Enter username: ")
    fmt.Scan(&username)
    fmt.Print("Enter password: ")
    fmt.Scan(&password)
    fmt.Print("Enter server address: ")
    fmt.Scan(&server)
    fmt.Print("Enter server port: ")
    fmt.Scan(&port)
    fmt.Print("Enter file path: ")
    fmt.Scan(&filePath)
    fmt.Print("Upload or Download (u/d): ")
    fmt.Scan(&action)

    // Configure SSH client
    config := &ssh.ClientConfig{
        User: username,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    // Connect to the server
    conn, err := ssh.Dial("tcp", server+":"+port, config)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // Create new SFTP client
    client, err := sftp.NewClient(conn)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    if action == "u" {
        // Upload file
        uploadFile(client, filePath)
    } else if action == "d" {
        // Download file
        downloadFile(client, filePath)
    } else {
        fmt.Println("Invalid action")
    }
}

func uploadFile(client *sftp.Client, filePath string) {
    srcFile, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer srcFile.Close()

    dstFile, err := client.Create(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("File uploaded successfully.")
}

func downloadFile(client *sftp.Client, filePath string) {
    srcFile, err := client.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer srcFile.Close()

    dstFile, err := os.Create(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("File downloaded successfully.")
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
