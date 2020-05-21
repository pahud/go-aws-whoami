# go-aws-whoami

`pahud/aws-whoami` simple implementation in Go.

# Usage

```sh
# docker run it
docker run -v $HOME/.aws:/root/.aws -p 8080:8080 pahud/go-aws-whoami:latest

# curl it
curl 0:8080/whoami
```

<details>
<summary> 👉Response </summary>
{
  "Account": "112233445566",
  "Arn": "arn:aws:iam::112233445566:user/pahud",
  "UserId": "AIDAJVHX3XBRH4E4UGWWK"
}
</details>