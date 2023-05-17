## JSON Web Token(JWT)

> JSON Web Token (JWT) is an open standard ([RFC 7519](https://tools.ietf.org/html/rfc7519)) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed. JWTs can be signed using a secret (with the HMAC algorithm) or a public/private key pair using RSA or ECDSA.

> Although JWTs can be encrypted to also provide secrecy between parties, we wirll focus on signed tokens. Signed tokens can verify the integrity of the claims contained within it, while encrypted tokens hide those claims from other parties. When tokens are signed using public/private key pairs, the signature also certifies that only the party holding the private key is the one that signed it.

Read [here](https://jwt.io/introduction/) for more details.

JSON Web Token is dot-separated string:

```
[header - base64].[payload - base64].[signature]
```

### Head

`Base64Url` encoded.

> The header typically consists of two parts: the type of the token, which is JWT, and the signing algorithm being used, such as HMAC SHA256 or RSA.

```
{
  "alg": "HS256",
  "typ": "JWT"
}
```

### Payload

`Base64Url` encoded. Pass parameters such as issuer (e.g. google.com), audience(e.g. client to consume token), and expiration date.

> The second part of the token is the payload, which contains the claims. Claims are statements about an entity (typically, the user) and additional data. There are three types of claims: registered, public, and private claims.

Registered claims are a set of predefined claims recommended to implement. See registered claims [here](https://tools.ietf.org/html/rfc7519#section-4.1).

```
{
  "sub": "1234567890"
  "aud": "example.com",
  "iss": "google.com",
}
```

### Signature

Encrypted using algorithm such as SHA256 and _private key_. _private key_ needs to be random and complex, and it should be securely stored on server.

> The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is.

> To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that.

```
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```

### Generate token

> The output is three Base64-URL strings separated by dots that can be easily passed in HTML and HTTP environments, while being more compact when compared to XML-based standards such as SAML.

To generate token, use `base64` to encode `[header]` and `[payload]`(`map[string]string` in this project).

To generate signature, encrypt `[header]` and `[payload]` together using secret key.

For validation, decode `[header]` and `[payload]` and combine them to create new signature. If this matches with the signature in token, token is valid. Make sure expiration date is not expired.

### SECURITY NOTICE

In systems using HS256 algorithm, we sign with private key and use public key(usually text-based PEM format) as verification key.

In systems using HMAC, verification key will be server's secret signing key(HMAC use the same key for signing and verifying).

As security practice, it is recommended to verify algorithm in header of token.

Read [here](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/) for more details.
