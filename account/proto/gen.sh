
protoc -I=. \
  -I ${HOME}/go/pkg/mod \
  -I ${HOME}/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.4.1 \
  --go_out=plugins=grpc,paths=source_relative:api/v1 \
  --validate_out="lang=go",paths=source_relative:api/v1 \
  --grpc-gateway_out=paths=source_relative,grpc_api_configuration=account.yaml:api/v1 account.proto


PROTOBUF_JS_BIN=$HOME/web-src/shakebook_manager/node_modules/protobufjs/bin
WEB_PROTO_DIR=$HOME/web-src/shakebook_manager/src/service/api/v1

$PROTOBUF_JS_BIN/pbjs -t static -w es6 account.proto --no-create --no-encode --no-decode --no-verify --no-delimited --force-number -o $WEB_PROTO_DIR/account_pb_tmp.js
echo 'import * as $protobuf from "protobufjs";\n' > $WEB_PROTO_DIR/account_pb.js
cat $WEB_PROTO_DIR/account_pb_tmp.js >> $WEB_PROTO_DIR/account_pb.js
rm $WEB_PROTO_DIR/account_pb_tmp.js
${PROTOBUF_JS_BIN}/pbts -o $WEB_PROTO_DIR/account_pb.d.ts $WEB_PROTO_DIR/account_pb.js

# protoc \
#   -I=. \
#   -I ${HOME}/go/pkg/mod \
#   -I ${HOME}/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.4.1 \
#   --validate_out="lang=go",paths=source_relative:gen/go \
#   account.proto