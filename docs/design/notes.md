# Things to address

* Should a user be able to specify DGNs in a config file
  * For example Spyder has their own "custom" DGN and other vendors might do the same
* A user SHOULD be able to specify instance names in a config file - including overriding existing ones
* Maybe bitfields should be packed into a single byte vs split out

#
# MQTTT
# 

Initial idea is to send only changes to MQTTT. We don't want to use up a lot of bandwidth
sending mostly unchanged messages. We'll also need some means to optionally filter out messages.

For a message structure we can have something like :

ID (Eg - MikesMotorhome)
RVC
DGNName
InstanceName
JSON Payload (Getter Methods and values)


