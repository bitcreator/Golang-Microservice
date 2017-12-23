FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY micro /
CMD ["/micro"]