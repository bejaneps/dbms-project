// Package ssl generates an x509 certificate and private key for tls(ssl) connections
package ssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func init() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serial, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}

	subject := pkix.Name{
		Country:      []string{"North Cyprus"},
		Organization: []string{"European University of Lefke"},
		CommonName:   "EUL",
	}

	templ := x509.Certificate{
		SerialNumber: serial,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := x509.CreateCertificate(rand.Reader, &templ, &templ, pk.Public, pk)
	if err != nil {
		log.Fatal(err)
	}

	certFile, err := os.Create("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	certFile.Close()

	err = pem.Encode(certFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})
	if err != nil {
		log.Fatal(err)
	}

	keyFile, err := os.Create("key.pem")
	if err != nil {
		log.Fatal(err)
	}

	err = pem.Encode(keyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk),
	})
	if err != nil {
		log.Fatal(err)
	}
}
