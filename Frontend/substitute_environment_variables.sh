#!/usr/bin/env sh

find '/usr/share/nginx/html/js' -name '*.js*' -exec sed -i 's|BASE_CONTENT_URL_PLACEHOLDER|'${BASE_CONTENT_URL}'|g' {} \;
find '/usr/share/nginx/html/js' -name '*.js*' -exec sed -i 's|BASE_AUTH_URL_PLACEHOLDER|'${BASE_AUTH_URL}'|g' {} \;
nginx -g "daemon off;"