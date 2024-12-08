## 11.1 Private and Public Keys
1. We hash the msg
   - `"I love you" -> hash -> "fseefqedsjcnkjnca"`
2. Generate Key Pair|
   - `KeyPair (private Key, Public Key)`
3. Sign the Hash
   - `("fseefqedsjcnkjnca" + private Key) -> 'Signature'`
 
- Then, If somebody asked you to prove that this is your signature, What do you do?
- How are you going to prove that you are the one who signed tha hash.
- -> Verifying by the Public Key!

4. Verify
    - `("fseefqedsjcnkjnca" + 'Signature' + public Key) -> True/false`
