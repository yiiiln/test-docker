##基底image 這裡指定的是golang官方的最新image
#FROM golang:latest
##指定專案資料夾 也就是整個專案的根目錄
#WORKDIR /Users/h10/go/src/TestDockerV1
##將專案內所有程式碼複製到docker根目錄
#COPY . .
##按照go.mod安裝專案所需套件
#RUN go mod download
##將專案編譯成可執行檔
#RUN go build -o main .
##指定專案對外port號 記得要和專案內所指定的port號一樣
#EXPOSE 8080
##執行已編譯好的2進位檔
#CMD ["./main"]

FROM golang:latest

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]