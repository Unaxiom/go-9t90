## Go client for 9t90 URL shortener

## Install package
```bash
go get github.com:Unaxiom/go-9t90
```

```go
func testGRPC() {
	<-time.After(time.Second * 3)
	conn, err := grpc.Dial("localhost"+constants.AppPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	// Code removed for brevity
	ctx := context.Background()

	md := metadata.New(map[string]string{"auth_token": ""})
	ctx = metadata.NewOutgoingContext(ctx, md)

	client := sdk.NewURLsServiceClient(conn)
	url, err := client.ViewByCode(ctx, &sdk.SearchKeyRequest{
		SearchKey: "Ukp9aciZUg4M",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n\n%#v", url)
}

```