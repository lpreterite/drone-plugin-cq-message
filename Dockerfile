FROM plugins/base:multiarch

ADD release/linux/amd64/drone-plugin-cq-message /bin/
ENTRYPOINT ["/bin/drone-plugin-cq-message"]