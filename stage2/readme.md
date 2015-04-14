# decode_json
JSON Decode challenge

# how-to

## Challenge2

### Faulty Sensors

You have a bunch of sensors in your smart home. Each sensor gives you exactly 4 values. You usually poll them using http GET. Some sensors started to fail, and did not send all the values anymore. Identify the failed sensors in your request to avoid a disaster!

### GET http://forgot_to_update_hostname/stage1/data.json
```bash
> GET /stage1/data.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 10:25:30 GMT
< Content-Length: 109
< Content-Type: text/plain; charset=utf-8
<

{
   "sensors" : [
      {
         "id" : 2346,
         "humidity" : 40,
         "temperature" : 25
      },
      {
         "id" : 2871,
         "humidity" : 40,
         "temperature" : 25
      },
      {
         "have_smoke" : false,
         "id" : 22,
         "humidity" : 40,
         "temperature" : 25
      },
      {
         "have_smoke" : false,
         "id" : 234,
         "humidity" : 40,
         "temperature" : 25
      },
      {
         "have_smoke" : false,
         "id" : 39,
         "humidity" : 40,
         "temperature" : 25
      }
   ]
}
```

### POST http://forgot_to_update_hostname/stage1/submit.json
```bash
> POST /stage1/submit.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
> Content-type: application/json
> Content-Length: 78
>

{"team":"YOURTEAMNAME", "faulty":[2346,2871]}

< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 08:14:12 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<

{"ok":true}
```
