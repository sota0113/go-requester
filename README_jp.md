# go-requester, just keep requesting to some endpoint.

英語版は[こちら](/README_jp).
Docker image [IMAGE](https://hub.docker.com/r/sota0113/go-requester)
あるエンドポイントに対してリクエストを投げ続けます。
デフォルトのエンドポイントは、"http://localhost/list"です。 
エンドポイントは以下の環境変数を設定することで変更可能です。
  
```
RQ_HOST		#宛先ホスト名かIPアドレスを設定してください。デフォルトの値は"localhost"です。
RQ_PROTOCOL	#リクエストで使用するプロトコルを設定してください。デフォルトの値は"http"です。
RQ_PATH		#リクエストのパスを設定してください。デフォルトの値は"/list"です。
RQ_PORT		#リクエストのポート番号を指定してください。デフォルトの値は"80"です。
```
  
Dockerで動かす場合は、dockerfileの環境変数を構成してください。
Kubernetesの場合は、本リポジトリのkubernetes/deployment.yamlを構成してください。

アプリケーションを実行すると以下のようなログが出力されます。
```
# 誤ったエンドポイントが指定されている場合
requesting http://wronghost:30080/hoge.
Get Error in requesting http://wronghost:30080/hoge.
Get Error in requesting http://wronghost:30080/hoge.
Get Error in requesting http://wronghost:30080/hoge.
...

#正しいエンドポイントが指定されている場合
# This is a senario of requesting to right endpoint.
requesting http://righthost:30080/fuga.
<result of your request>
<result of your request>
<result of your request>
...
```

各リクエストは1秒ごとに実行されます。
現時点で宛先のホットリロードには対応していません。
