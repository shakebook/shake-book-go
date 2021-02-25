function genProto {
  DOMAIN=$1
  PROTO_PATH=./service/${DOMAIN}/proto
  GO_OUT_PATH=./service/${DOMAIN}/proto/api/v1
  WEB_PATH=${HOME}/web-src/shake-book-manager
  DART_PATH=${HOME}/flutter-project/mokong/lib/api/v1
  PROTOC_GEN_DART=${HOME}/Development/flutter/.pub-cache/bin/protoc-gen-dart 

  protoc -I=${PROTO_PATH} \
    -I ${HOME}/go/pkg/mod \
    -I ${HOME}/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.4.1 \
    --go_out=plugins=grpc,paths=source_relative:${GO_OUT_PATH} ${PROTO_PATH}/${DOMAIN}.proto 
  
  protoc -I=${PROTO_PATH} \
    -I ${HOME}/go/pkg/mod \
    -I ${HOME}/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.4.1 \
    --dart_out=${DART_PATH} --plugin=protoc-gen-dart=${PROTOC_GEN_DART} ${PROTO_PATH}/${DOMAIN}.proto
  
  if [ $2 ];then
    return;
  fi
  protoc -I=${PROTO_PATH} \
    -I ${HOME}/go/pkg/mod \
    -I ${HOME}/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.4.1 \
    --validate_out="lang=go",paths=source_relative:${GO_OUT_PATH} \
    --grpc-gateway_opt allow_delete_body=true\
    --grpc-gateway_out=paths=source_relative,grpc_api_configuration=${PROTO_PATH}/${DOMAIN}.yaml:${GO_OUT_PATH} ${PROTO_PATH}/${DOMAIN}.proto 

  PROTOBUF_JS_BIN=${WEB_PATH}/node_modules/protobufjs/bin
  WEB_PROTO_DIR=${WEB_PATH}/src/service/api/v1

  ${PROTOBUF_JS_BIN}/pbjs -t static -w es6 ${PROTO_PATH}/${DOMAIN}.proto --no-create --no-encode --no-decode --no-verify --no-delimited --force-number -o ${WEB_PROTO_DIR}/${DOMAIN}_pb_tmp.js
  echo 'import * as $protobuf from "protobufjs";\n' > ${WEB_PROTO_DIR}/${DOMAIN}_pb.js
  cat ${WEB_PROTO_DIR}/${DOMAIN}_pb_tmp.js >> ${WEB_PROTO_DIR}/${DOMAIN}_pb.js
  rm ${WEB_PROTO_DIR}/${DOMAIN}_pb_tmp.js
  ${PROTOBUF_JS_BIN}/pbts -o ${WEB_PROTO_DIR}/${DOMAIN}_pb.d.ts ${WEB_PROTO_DIR}/${DOMAIN}_pb.js

}

genProto account
genProto auth 1
genProto file
genProto manager