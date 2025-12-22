# REDIS STREAM Example

## How to Deploy

#### Git pull / clone
```bash
git pull origin main
```
#### Build image
```bash
sudo docker compose build 
```
#### Deploy
```bash
sudo docker compose up --no-deps -d 
```

## Sender redis-cli

```bash
 XADD eventName MAXLEN ~ 1000 * status "1" description "desc" action "status" 
```

## Sender with go

```go
errs := redis.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "eventName",
		MaxLen: 5000,
		Values: map[string]interface{}{
			"status": 1,
			"description":     time.Now().UTC(),
			"action":     "status",
		},
	}).Err()
	if errs != nil {
		fmt.Println(errs)
	}

```







