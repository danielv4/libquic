// go build -o libquic.dll -buildmode=c-shared libquic.go
// go build -buildmode=c-archive libquic.go
package main



import (
	"C"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
    "fmt"
	"math/big"
	quic "github.com/lucas-clemente/quic-go"
)


func generateTLSConfig() *tls.Config {

	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}



var i quic.Listener
var g quic.Session
var m quic.Stream


//export quic_listen
func quic_listen(url string) {

	listener, err := quic.ListenAddr(url, generateTLSConfig(), nil)
	if err != nil {

	}
	
	fmt.Printf("%T \n", listener)
	fmt.Printf("%T \n", quic.ListenAddr)
	
	i = listener
}


//export quic_accept
func quic_accept() {

	sess, err := i.Accept(context.Background())
	if err != nil {

	}
	
	fmt.Printf("%T \n", sess)
	fmt.Printf("%T \n", i.Accept)
	
	g = sess
}


//export quic_accept_stream
func quic_accept_stream() {

	stream, err := g.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%T \n", stream)
	fmt.Printf("%T \n", g.AcceptStream)
	
	m = stream
}


//export quic_stream_write
func quic_stream_write(write_bytes []byte) C.int {

	n, err := m.Write(write_bytes)
	if err != nil {
		return C.int(-1)
	}

	return C.int(n)
}


//export quic_stream_read
func quic_stream_read(read_bytes []byte) C.int {

	n, err := m.Read(read_bytes)
	if err != nil {
		return C.int(-1)
	}
	
	return C.int(n)
}


//export quic_open_stream_sync
func quic_open_stream_sync(url string) {

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(url, tlsConf, nil)
	if err != nil {

	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {

	}
	
	fmt.Printf("%T \n", stream)
	fmt.Printf("%T \n", session.OpenStreamSync)

	stream = m
}


func main() {

	// quic_listen("127.0.0.1:4433")
	
	// quic_accept()
	
	// quic_accept_stream()
	
	// quic_stream_write(byte[], 1024)
	
	// quic_stream_read(byte[], 1024)
	
	// quic_open_stream_sync("127.0.0.1:4433")
}







