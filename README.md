# Chivato
## Telegram webhook
* Read updates  
https://api.telegram.org/bot\<api\>/getUpdates

* Remove WebHook  
https://api.telegram.org/bot\<api\>/deleteWebhook  
* Get Webhook info  
https://api.telegram.org/bot\<api\>/getWebhookInfo  

* Llamadas varias  
curl -F "url=$WEBHOOK"  https://api.telegram.org/bot\<api\>/setWebhook  
curl -s -X POST "https://api.telegram.org/bot\<api\>/sendMessage" -d chat_id=$ID -d text="$1"    
curl -s -X POST "https://api.telegram.org/bot\<api\>/sendMessage" -d chat_id=$ID2 -d text="$1"    
curl -F "url=$WEBHOOK"  https://api.telegram.org/bot\<api\>/setWebhook .

## Configure env Linux/mac
go env -w GO111MODULE="on"  
go env -w GOPATH=$(dirname $(dirname $(pwd)))
