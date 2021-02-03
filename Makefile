.PHONY: proto

#ARCH
ARCH="`uname -s`"
LINUX="Linux"
DARWIN="Darwin"
#Main build target
all:build

build:
	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		echo $(ARCH); \
		echo "protoc --proto_path=. --doc_out=./doc --doc_opt=html,index.html --go_out=plugins=grpc:. */*.proto;";\
#		sed -i 's/dm-gitlab.bolo.me\/hubpd\/proto/./' common/common.proto;\
		protoc --proto_path=. --doc_out=./doc --doc_opt=html,index.html --go_out=plugins=grpc:. */*.proto;\
#		sed -i 's/.\/common/dm-gitlab.bolo.me\/hubpd\/proto\/common/' common/common.proto;\
#		protoc --proto_path=. --doc_out=./doc --doc_opt=html,index.html --go_out=plugins=grpc:. */*.proto;\
	elif [ $(ARCH) = $(DARWIN) ]; \
	then \
		echo $(ARCH); \
		echo "protoc --proto_path=. --doc_out=./doc --doc_opt=html,index.html --go_out=plugins=grpc:. */*.proto;";\
		sed -i '' 's/dm-gitlab.bolo.me\/hubpd\/proto/./' common/common.proto;\
		protoc --proto_path=. --doc_out=./doc --doc_opt=html,index.html --go_out=plugins=grpc:. */*.proto;\
		sed -i '' 's/.\/common/dm-gitlab.bolo.me\/hubpd\/proto\/common/' common/common.proto;\
		protoc --proto_path=. --doc_out=./doc --doc_opt=html,index.html --go_out=plugins=grpc:. */*.proto;\
	else \
		echo "ARCH unknow, error found!"; \
	fi
	rm -rf dm-gitlab.bolo.me
	#bash mockgen.sh

