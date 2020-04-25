# jsonlogparser

Log format in `nginx.conf`:
```text
log_format json '{ "@timestamp": "$time_iso8601", '
                  '"remote_addr": "$remote_addr", '
                  '"cookie_bar": "$cookie_bar", '
                  '"set_cookie": "$sent_http_set_cookie", '
                  '"body_bytes_sent": "$body_bytes_sent", '
                  '"status": "$status", '
                  '"request": "$request", '
                  '"url": "$uri", '
                  '"request_method": "$request_method", '
                  '"upstream": "$upstream_addr", '
                  '"response_time": "$temprt", '
                  '"http_referrer": "$http_referer", '
                  '"http_user_agent": "$http_user_agent" }';
```
