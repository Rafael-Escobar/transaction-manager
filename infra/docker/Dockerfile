# using the super-lean distroless image
FROM gcr.io/distroless/base:latest

ENV PORT=8080

# Change to the app directory to maintain the same root reference as in development.
WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY ./transaction-manager /app/transaction-manager

# Copy configuration file from the builder stage.
COPY ./app/config/config.yml /app/config.yml

# Default starting supplier api in verbose mode
ENTRYPOINT ["/app/transaction-manager"]
CMD ["--verbose", "--config", "/app/config.yml", "serve"]