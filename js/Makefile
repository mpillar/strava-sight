BROWSERIFY = ./node_modules/.bin/browserify
WATCHIFY = ./node_modules/.bin/watchify

APP_JS = ./index.js
BUNDLE_DIR = .
BUNDLE_JS = ${BUNDLE_DIR}/bundle.js
BUNDLE_MIN_JS = ${BUNDLE_DIR}/bundle.min.js
BROWSERIFY_PARAMS = -t [ babelify --presets [ es2015 react ] ]

start:
	./start.sh

install:
	npm install

bundle:
	${BROWSERIFY} ${BROWSERIFY_PARAMS} ${APP_JS} -o ${BUNDLE_JS}

watch:
	${WATCHIFY} -v ${BROWSERIFY_PARAMS} ${APP_JS} -o ${BUNDLE_JS}

clean:
	rm -rf node_modules
	rm -f ${BUNDLE_JS}
