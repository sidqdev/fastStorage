
# FastStorage for telegram bots or another projects

You can set and get data quickly from database(postgres or redis)




## Usage/Examples

### For postgres you should create database(for example in Django)
```python3
class Storage(models.Model):
    chat_id = models.BigIntegerField()
    user_id = models.BigIntegerField()
    json_data = models.JSONField(null=True)
    stage = models.TextField(null=True)
    class Meta:
        db_table = "storage_data
```

```golang
    storage.SetPostgresConfig("modtest", "modtest", "modtest") // For postgres user/pass/db 
    storage.SetRedisConfig(0, "", "name") // For redis dbIndex/pass/dbname(any)

    //STAGES
    stage, err := storage.GetStage(chat_id, user_id)
    err := storage.SetStage(chat_id, user_id, "STAGE")

    //DATA
    err := storage.GetData(chat_id, user_id, &struct)
    err := storage.SetData(chat_id, user_id, struct)

```

