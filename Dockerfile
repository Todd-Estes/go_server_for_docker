FROM debian:stable-slim
# COPY source destination
COPY go_server_for_docker /bin/go_server_for_docker
ENV PORT :8000
CMD ["/bin/go_server_for_docker"]