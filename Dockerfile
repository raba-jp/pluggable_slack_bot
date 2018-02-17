FROM alpine

COPY bin/maguro_linux /maguro
CMD ["./maguro"]
