// https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph

func loadTLSCredentials() (credentials.TransportCredentials, error) {
    // Load server's certificate and private key
    serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
    if err != nil {
        return nil, err
    }

    // Create the credentials and return it
    config := &tls.Config{
        Certificates: []tls.Certificate{serverCert},
        ClientAuth:   tls.NoClientCert,
    }

    return credentials.NewTLS(config), nil
}

func main() {
    ...

    tlsCredentials, err := loadTLSCredentials()
    if err != nil {
        log.Fatal("cannot load TLS credentials: ", err)
    }

    grpcServer := grpc.NewServer(
        grpc.Creds(tlsCredentials)
    )

   
}



func loadTLSCredentialsClient() (credentials.TransportCredentials, error) {
    // Load certificate of the CA who signed server's certificate
    pemServerCA, err := ioutil.ReadFile("self.crt")
    if err != nil {
        return nil, err
    }

    certPool := x509.NewCertPool()
    if !certPool.AppendCertsFromPEM(pemServerCA) {
        return nil, fmt.Errorf("failed to add server CA's certificate")
    }

    // Create the credentials and return it
    config := &tls.Config{
        RootCAs:      certPool,
    }

    return credentials.NewTLS(config), nil
}
