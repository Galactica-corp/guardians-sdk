# Galactica Guardian CLI

The Galactica Guardian CLI is a powerful command-line interface designed for guardians in the Galactica blockchain
ecosystem. Guardians play a pivotal role in verifying and managing user data on the blockchain. The CLI provides
essential tools for tasks such as certificate issuance, revocation, renewal, and more.

## Installation

To install the Galactica Guardian CLI, run the following command:

```shell
go install github.com/galactica-corp/guardians-sdk/v3/cmd/galactica-guardian@latest
```

Ensure that your Go environment is properly configured before running the installation command.

## Usage

Once installed, you can use the Galactica Guardian CLI to perform various tasks. The CLI provides several commands to
manage certificates, interact with the blockchain, and handle key operations. Each command supports the `--help` flag
for detailed information and usage guidelines.

The CLI **requires** the [Merkle Proof Service](https://github.com/Galactica-corp/merkle-proof-service) for issuing
and revoking Zero Knowledge Certificates. Ensure that this service is running and accessible to perform these operations.

### Commands Overview

* `generateEdDSAKeyPair`: Generate EdDSA key pairs for managing zero knowledge certificates.
* `createZKCert`: Create a Zero Knowledge Certificate (ZKCert) for on-chain verification.
* `issueZKCert`: Issue a ZKCert to the Galactica blockchain registry.
* `revokeZKCert`: Revoke a ZKCert from the Galactica blockchain registry.
* `renewZKCert`: Renew a ZKCert with an updated expiration date.
* `encryptZKCert`: Encrypt a ZKCert with a holder's encryption key.
* `printEdDSAPublicKey`: Print EdDSA public key for a given private key.

## License

This project is licensed under the GNU General Public License v3.0 (GPL-3.0). See the [LICENSE](LICENSE) file for
details.

Happy guarding in the Galactica ecosystem!
