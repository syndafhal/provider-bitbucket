FROM alpine:3.20.3


RUN apk --no-cache add ca-certificates bash curl unzip


ARG TARGETOS
ARG TARGETARCH


ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider-bitbucket


USER root

# Setup Terraform installation
ARG TERRAFORM_VERSION="1.5.7"
RUN curl -fsSL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip -o /tmp/terraform.zip \
    && unzip /tmp/terraform.zip -d /usr/local/bin/ \
    && chmod +x /usr/local/bin/terraform \
    && rm /tmp/terraform.zip

EXPOSE 8080

ENTRYPOINT ["provider-bitbucket"]