FROM golang:1.15.6-alpine

ENV TYG_ENTERPRISE_URL="https://replace.me.with.github.or.internal.url"
ENV TYG_TAGS='{{paths: ["README.md"], tags: ["edits-readme"]}}'
ENV TYG_GIT_API_KEY="REPLACE_ME"
ENV TYG_WEBHOOK_SECRET_KEY="REPLACE_ME"

# for no reason
EXPOSE 8080

WORKDIR /app
COPY ./ /app

RUN rm -rf .git && \
    go install .

FROM alpine:3.12.3
COPY --from=0 /go/bin/tag-your-git /usr/local/bin/

CMD ["tag-your-git"]