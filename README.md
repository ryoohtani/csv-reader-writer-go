# csv-reader-writer-go
A Go-based project for CSV file handling and output generation. It includes features to read input CSV files, perform data transformations, and generate output files. Docker is used for easy deployment. Aimed at learning Go and Docker, this project provides hands-on experience in CSV manipulation, containerization, and application development.

**環境構築**

*dockerコマンド*

* ビルドコマンド
```
docker-compose build
```
* 環境の立ち上げ
```
docker-compose up -d
```
* コンテナにアクセス(app_goはコンテナ名)
```
docker exec -it app_go  /bin/bash   
```

*Pythonのコマンド*
* GOの実行コマンド
```
go run ファイル名.go
```