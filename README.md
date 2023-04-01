<h2>GoLang SFTP</h2>
<p>The code in sftpclient.go implements an SFTP client in Go language. It connects to an SFTP server using SSH, authenticates with a username and password, and then opens an SFTP session to upload a file to the server.</p>

<p>The script prompts the user to enter their SFTP username, password, server IP address, and filename to upload. It then establishes a connection to the server, opens an SFTP session, and creates a remote file on the server with the same filename as the local file.</p>

<p>The script implements rate limiting and IP blocking to prevent excessive upload attempts from the same IP address. If an IP address attempts to upload more than three times within a certain time period, the script blocks that IP address for a minute before allowing any further attempts.</p>

<p>Overall, the code provides a basic SFTP client implementation with some additional security features to prevent abuse and protect the server.</p>


<ul>
<li>Clone the repository:</li>
</ul>

```
git clone https://github.com/ziomzziom/golang.git
```

<ul>
<li>Navigate to the src directory:</li>
</ul>

```
cd golang/src/
```

<ul>
<li>Build the script:</li>
</ul>

```
go build sftpclient.go
```

<ul>
<li>Run the script:</li>
</ul>

```
./sftpclient
```

<ul>
<li>The script will prompt you to enter the SFTP credentials and the server IP address. Once connected to the server, you can enter the filename to upload. The script will then upload the file to the server using SFTP.</li>
</ul>
