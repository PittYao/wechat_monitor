build:
	docker build -t fanyaoyao12138/wechat_monitor:latest .
tag:
	docker tag wechat_monitor fanyaoyao12138/wechat_monitor
push:
	docker push  fanyaoyao12138/wechat_monitor:latest

all:
	make tag
	make build
	make push

