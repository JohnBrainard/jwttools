# JWT Tools

Command line tools for generating and debugging JWT tokens

You can:

* Create new tokens from scratch
* Create new tokens from presets
* Display information about tokens

## Token Generation

### From scratch

`jwttools generate [-iss issuer] [-aud audience] [-sub subject] [-key secret] [-exp 1h]`

* `-iss` issuer
* `-aud` audience
* `-sub` subject
* `-exp` token expiration in the form of a duration value (eg. `1h15m30s`)
* `-key` secret
* `-preset` name of preset to use

### From Presets

You can create token presets in your `config.json` file.

`$HOME/.jwttools/config.json`

```json
{
  "presets": {
    "my-token": {
      "info":"My generic token",
      "claims": {
        "iss": "my-issuer",
        "sub": "my-subject",
        "aud": "classrooms-consumer"
      },
      "key": "secret",
      "expires": "1h"
    }
  }
}
```

* `my-token`: Name of the preset which is passed to the `jwttools` command
* `info`: short, descriptive text about the token. Displayed with the `jwttools presets` command
* `claims`: Raw JSON used in the token's payload
	* `iss`: standard JWT 'issuer' field
	* `sub`: standard JWT 'subject' field
	* `aud`: standard JWT 'audience' field
* `key`: secret used to sign the generated token. Can be overriden by the `-key` argument to most commands
* `expires`: expiration in the format of `"2h15m30s"`. Can be overidden by the `-exp` argument to most commands

## Init command

Executing the `init` command will create the `$HOME/.jwttools/config` file if it doesn't already exist.
Executing it repeatedly will not overwrite an already existing `config.json` file.

`jwttools init`

## Info command

The `info` command parses and displays the token fields.

`jwttools info -token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWRpZW5jZSIsImV4cCI6MTUxMDcwODA3MywiaXNzIjoiaXNzdWVyIiwic3ViIjoic3ViamVjdCJ9.MokQk4jORaloT3whhaMl0VjeZ6Q8Oas3UzpxjQiA_jg`

```
Expiration: 2017-11-14 18:07:53 -0700 MST
Claims:
  {
    "aud": "audience",
    "exp": 1510708073,
    "iss": "issuer",
    "sub": "subject"
  }
```

## Presets command

The `presets` command displays a summary of your presets

`jwttools presets [-keys] [-preset preset-name]`

* `-keys` include the secret in the summary list
* `-preset` only print the summary of the `-preset` argument
* `-verbose` print extra token information

## Edit command

The `edit` command allows you to use your favorite editor to more conveniently create or edit token
presets.

`jwttools edit -preset preset-name`

`jwttools` will use the `EDITOR` environment variable to determine which editor to use.

*Note* The edit command will create a backup of the current config file before saving changes to it.
