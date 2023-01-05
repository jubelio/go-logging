# go-logging
Simple module to proper logging to elasticsearch

ENV
```
### ES host | default: https://es2.jubelio.com
LOGGING_HOST=

### Indice name to save logs | default: logs-app-default
LOGGING_INDICE=

### APIKEY to authenticate to ES, must have permisssion to readANDwrite to indice | default: none
LOGGING_APIKEY=

### Username to authenticate to ES, if APIKEY exist, will use APIKEY instead | default: none
LOGGING_USERNAME=

### Password to authenticate to ES, if APIKEY exist, will use APIKEY instead | default: none
LOGGING_PASSWORD
```
