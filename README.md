# mmr-project
MMR Project


## Setup GORM and Atlas
Run the following command to install Atlas: 
```
curl -sSf https://atlasgo.sh | sh
```

Add migration: 
```
atlas migrate diff --env gorm 
```

Apply migration:
```
atlas migrate apply --dir "file://db/migrations" --url mysql://root:root123\!@localhost:3306/foosball

```