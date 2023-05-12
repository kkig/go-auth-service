## JSON Web Token(JWT)

JSON Web Token is dot-separated string:

```
[header - base64].[payload - base64].[signature]
```

`[header]`: Base64 encoded.
`[payload]`: Base64 encoded. Pass parameters such as issuer (e.g. google.com), audience(e.g. client to consume token), and expiration date.
`[signature]`: Encrypted using algorithm such as SHA256 and _private key_. _private key_ needs to be random and complex, and it should be securely stored on server.

To generate signature, encrypt `[header]` and `[payload]` together using secret key.

For validation, decode `[header]` and `[payload]` and combine them to create new signature. If this matches with the signature in token, token is valid. Make sure expiration date is not expired.
