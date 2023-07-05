### Demo generate otp service using golang:

#### Generate OTP ( Should apply rate limit for this service)

```
Route: http://localhost:8080
Method: POST
Response: {
  "otp": "453402",
  "timestamp": 1688528760 // timestamp of the generated otp
}
Description: Return OTP with expired after 30 seconds
```

#### Verify OTP

```
Route: http://localhost:8080/:otp
Method: GET
Response: {
  "result": false, // return false after 30 seconds
  "timestamp": 1688528791
}

```
