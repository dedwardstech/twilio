# Twilio
This package just contains methods for Twilio SMS currently. Feel free to submit PRs for 
other Twilio products. Also feel free to submit suggestions!

### SMS
The SMS package contains 4 important structs.


##### Client
The Client struct is responsible for building requests for
a particular account/number combination.

Creating a new client

````go
client, err := sms.NewClient(&sms.ClientOptions{
    AccountSid: "xxxxxx",
    AuthToken: "xxxxxx",
    Number: "+18008675309"
})
````

The client does not send SMS messages. It exists purely to
build the http Requests. Sending is delegated back to the consumer.

Building a request
```go
msgReq, err := client.NewRequest("+18009999999", "Hello World")
if err != nil {
	// handle error
}

httpClient := &http.Client{}
resp, err := httpClient.Do(msgReq)
if err != nil {
	// handle http error
}
...
```

##### Parser
The parser is responsible for parsing incoming Twilio
messages from a Twilio webhook.

Creating a new parser

````go
parser := sms.NewMessageParser()
msg := parser.FromRequest(r)
````
