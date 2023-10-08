Unescapes URL and redirects to URL in path. Useful for Firefox [keywords feature](https://support.mozilla.org/en-US/kb/how-search-from-address-bar) with Wayback Machine:
```
http://100.100.100.100:8080/https://web.archive.org/web/*/%s
```

By definition vulnerable as [open redirect](https://cwe.mitre.org/data/definitions/601.html)

Client is responsible for infinite redirects.
