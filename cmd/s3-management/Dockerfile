FROM jrei/systemd-centos:7

COPY S3Management.viper.yaml /etc/S3Management/S3Management.viper.yaml
COPY s3-management /usr/local/bin/s3-management

CMD ["s3-management", "run"]
