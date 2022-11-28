# reed
Remote Envelope Encryption Dictionary

This is a simple gRPC service which uses AWS KMS service to encrypt and decrypt whatever content you send it. To make itself useful
as a service to be used for handling PII, it also acts a dictionary of sorts for sensitive data. Essentially, consumming services
will call it to encrypt a piece of data, and it will store sufficient unidentifiable metadata for performing decryption at a later
stage when requested.

It uses AWS KMS, but employs the Envelope Encryption technique recommended by Amazon in order to be able to keep sensitive data encrypted
while at the same time not having to rely on a central encryption key and at the same time helps to avoid storing the plain version of the
encryptions keys used (in case of a leak).

This service is strongly opinionated in the sense that it limits the radius of key exposure by enforcing usage of a new enncryption key per
data entry encryption request. Evidently, it is also opinionated due to the use of AWS KMS.
