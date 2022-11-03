#!/bin/sh
docker run -p 8081:8081\
	-e https_proxy="172.17.0.1:7890"\
	grfw_server
