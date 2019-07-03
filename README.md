# go-requester, just keep requesting to some endpoint.

For Japanese, see [日本語](/README_jp.md).
Docker image [IMAGE](https://hub.docker.com/r/sota0113/go-requester)
Just keep requesting to a custom endpoint.  
Default endpoint is "http://localhost/list".  
The endpoint cloud be changed by setting environment variable as below.
  
```
RQ_HOST		#Set hostname or IP address. Default value is "localhost".
RQ_PROTOCOL	#Set request protocol. Default value is "http".
RQ_PATH		#Set request path. Default value is "list".
RQ_PORT		#Set request port. Default value is "3030".
```
  
If you run this app on Docker, edit dockerfile and configure each environemnt.
If on Kubernetes, edit kubernetes/deployment.yaml on this repo.

This app print logs to stdout as below.
```
# This is a senario of requesting to wrong endpoint.
requesting http://wronghost:30080/hoge.
Get Error in requesting http://wronghost:30080/hoge.
Get Error in requesting http://wronghost:30080/hoge.
Get Error in requesting http://wronghost:30080/hoge.
...

# This is a senario of requesting to right endpoint.
requesting http://wronghost:30080/hoge.
<result of your request>
<result of your request>
<result of your request>
...
```

Each request is executed every 1 second.
Hot reload of request endpoint is not supported.
