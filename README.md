# go-logging
Simple module to proper logging to elasticsearch

ENV
| KEY                 	| DEFAULT                 	| REQUIRED             	| DESCRIPTION                                                                   	|
|---------------------	|-------------------------	|----------------------	|-------------------------------------------------------------------------------	|
| LOGGING_HOST        	| https://es2.jubelio.com 	| no                   	| ES Host                                                                       	|
| LOGGING_INDICE      	| logs-app-default        	| no                   	| Indice name to save logs                                                      	|
| LOGGING_APIKEY      	|                         	| yes, username/apikey 	| APIKEY to authenticate to ES, must have permisssion to readANDwrite to indice 	|
| LOGGING_USERNAME    	|                         	| yes, username/apikey 	| Username to authenticate to ES, if APIKEY exist, will use APIKEY instead      	|
| LOGGING_PASSWORD    	|                         	| yes, username/apikey 	| Password to authenticate to ES, if APIKEY exist, will use APIKEY instead      	|
| LOGGING_LEVEL       	| INFO                    	| no                   	| options: FATAL, ERROR, WARN, INFO, DEBUG, TRACE                               	|
| LOGGING_ACTIVE      	| false                   	| no                   	| when set to false, will not send logs to ES                                   	|
| LOGGING_STDOUT      	| true                    	| no                   	| when set to false, will not log to stdout                                     	|
| LOGGING_SERVICENAME 	| go-logging              	| no                   	| service name                                                                  	|
