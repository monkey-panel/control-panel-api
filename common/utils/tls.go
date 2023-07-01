package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/fs"
	"log"
	"math/big"
	"net"
	"time"
)

type Certificate struct {
	ServerKey string `json:"server_tls_key"`
	ServerPem string `json:"server_tls_public"`
}

func SummonCert() {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Organization: []string{"console-panel-api"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		SubjectKeyId:          []byte{1, 2, 3, 4, 5},
		BasicConstraintsValid: true,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	caKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"Monkey-Cat"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	cert.DNSNames = append(cert.DNSNames, "127.0.0.1")
	cert.IPAddresses = append(cert.IPAddresses, net.ParseIP("0.0.0.0"))

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	pub := &priv.PublicKey
	privPm := priv
	if caKey != nil {
		privPm = caKey
	}
	ca_b, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, privPm)
	if err != nil {
		log.Println("create failed", err)
		return
	}
	certificate := &pem.Block{
		Type:    "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   ca_b,
	}
	ca_b64 := pem.EncodeToMemory(certificate)
	AutoWriteFile("data/server.pem", ca_b64, fs.ModePerm)

	privateKey := &pem.Block{
		Type:    "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   x509.MarshalPKCS1PrivateKey(priv),
	}
	priv_b64 := pem.EncodeToMemory(privateKey)
	AutoWriteFile("data/server.key", priv_b64, fs.ModePerm)
}
