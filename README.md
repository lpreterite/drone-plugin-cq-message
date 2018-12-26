# drone-plugin-cq-message

push message by cqhttp in drone plugin

## Build

```sh
go build
```

## Docker

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o release/linux/amd64/drone-plugin-cq-message
docker build --rm -t lpreterite/drone-plugin-cq-message .
```

## Usage

```sh
docker run --rm \
  -e PLUGIN_CQ_HOST=http://cqhttp \
  -e PLUGIN_CQ_ACTION=send_msg \
  -e PLUGIN_CQ_TOKEN=kSLuTF2GC2Q4q4ugm3 \
  -e PLUGIN_CQ_QUERY='{"group_id": "734751943", "message": "[CQ:at,qq=53421639] hello!!"}' \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  lpreterite/drone-plugin-cq-message
```