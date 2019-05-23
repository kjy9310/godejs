rm -rf ./deploy/*
cd api && go build && mv ./api ../deploy/server && cd ..
cd react-web && npm run build && cp -R ./build ../deploy && cd ..
