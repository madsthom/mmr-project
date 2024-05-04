# mmr-project

MMR Project

## Setup GORM and Atlas

Run the following command to install Atlas:

```bash
curl -sSf https://atlasgo.sh | sh
```

Add migration:

```bash
atlas migrate diff --env gorm
```

Apply migration:

```bash
atlas migrate apply --dir "file://db/migrations" --url postgres://root:root123\!@localhost:3306/foosball
```
