FROM debian:stable-slim
# COPY source destination
COPY go_server_for_docker /bin/go_server_for_docker
CMD ["/bin/go_server_for_docker"]