FROM --platform=linux node:18.7-alpine AS cosmos-faucet

ENV COSMJS_VERSION=0.29.3

RUN npm install @cosmjs/faucet@${COSMJS_VERSION} --global --production

ENV FAUCET_CONCURRENCY=2
ENV FAUCET_PORT=4500
ENV FAUCET_GAS_PRICE=0.001stake
ENV FAUCET_GAS_LIMIT=200000
# Prepared keys for determinism
ENV FAUCET_MNEMONIC="vanish hover hammer plate oppose runway loan silent sponsor mansion connect rug total stereo gloom jaguar chalk gallery skin valley city tattoo water rubber"
ENV FAUCET_ADDRESS_PREFIX=cosmos
ENV FAUCET_TOKENS="stake, token"
ENV FAUCET_CREDIT_AMOUNT_STAKE=100
ENV FAUCET_CREDIT_AMOUNT_TOKEN=100
ENV FAUCET_COOLDOWN_TIME=0

EXPOSE 4500

ENTRYPOINT [ "cosmos-faucet" ]