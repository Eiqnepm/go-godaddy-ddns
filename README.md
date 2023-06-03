# go-godaddy-ddns

<details>
  <summary>Docker Compose example</summary>

```yaml
version: "3"

services:
  go-godaddy-ddns:
    environment:
      API_KEY: "3mM44UchDmQvBg_CRNcUUaamzsizKH4HquNz4"
      API_SECRET: "HL78KHU3r853gHz8LSGVgz"
      DOMAIN: "example.com"
    image: "eiqnepm/go-godaddy-ddns:latest"
    restart: "unless-stopped"
```

</details>

## Environment variables

| Variable     | Default | Description                         |
| ------------ | ------- | ----------------------------------- |
| `API_KEY`    |         | GoDaddy API key                     |
| `API_SECRET` |         | GoDaddy API secret                  |
| `DOMAIN`     |         | Domain name                         |
| `NAME`       | `@`     | Record name                         |
| `IPV4`       | `true`  | Update IPv4 IP using an A record    |
| `IPV6`       | `false` | Update IPv6 IP using an AAAA record |
| `TIMEOUT`    | `300`   | Seconds between each IP check       |
