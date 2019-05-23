rm -rf ./deploy/*
cd api && go build && mv ./api ../deploy && cd ..
cd react-web && npm run build && cp -R ./build ../deploy && cd ..
