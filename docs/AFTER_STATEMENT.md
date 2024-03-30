```sql
DROP TRIGGER IF EXISTS after_update_public_rooms ON public.rooms;
DROP FUNCTION IF EXISTS after_update_public_rooms();

CREATE FUNCTION after_update_public_rooms()
RETURNS TRIGGER AS $$
BEGIN
  PERFORM pg_notify('rooms:updated', NEW.id::text);
  RETURN NULL;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER after_update_public_rooms
AFTER UPDATE ON public.rooms
FOR EACH ROW EXECUTE PROCEDURE after_update_public_rooms();
COMMENT ON TRIGGER after_update_public_rooms ON public.rooms IS 'trigger to perform notify with id';
```

Q: как бы сделать так, чтобы отправлялась одна нотификация на изменение сразу нескольких рядов в таблице, т.е. по запросу: `UPDATE rooms SET kind = 'PUBLIC_GROUP' where id IN (2, 4)`

***

Ваш текущий код отправляет уведомление для каждого обновленного ряда. Если вы хотите отправить одно уведомление для нескольких обновленных рядов, вам нужно будет изменить подход. 

Один из способов - использовать временную таблицу (которая автоматически удаляется при завершении сессии) для хранения обновленных `ID`, а затем использовать триггер `AFTER STATEMENT`, чтобы отправить уведомление после выполнения запроса `UPDATE`. Вот пример кода:

```sql
DROP TRIGGER IF EXISTS after_update_public_rooms ON public.rooms;
DROP FUNCTION IF EXISTS after_update_public_rooms();

CREATE FUNCTION after_update_public_rooms()
RETURNS TRIGGER AS $$
BEGIN
  CREATE TEMP TABLE IF NOT EXISTS updated_ids (id INT);
  INSERT INTO updated_ids VALUES (NEW.id);
  RETURN NULL;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER after_update_public_rooms
AFTER UPDATE ON public.rooms
FOR EACH ROW EXECUTE PROCEDURE after_update_public_rooms();
COMMENT ON TRIGGER after_update_public_rooms ON public.rooms IS 'trigger to collect ids';

DROP TRIGGER IF EXISTS notify_updated_rooms ON public.rooms;
DROP FUNCTION IF EXISTS notify_updated_rooms();

CREATE FUNCTION notify_updated_rooms()
RETURNS TRIGGER AS $$
DECLARE
  ids TEXT;
BEGIN
  -- LOCK TABLE updated_ids IN EXCLUSIVE MODE;
  SELECT string_agg(id::text, ', ') INTO ids FROM updated_ids;
  PERFORM pg_notify('rooms:updated', ids);
  DELETE FROM updated_ids WHERE id = ANY(string_to_array(ids, ', ')::int[]);
  -- TRUNCATE updated_ids;
  RETURN NULL;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER notify_updated_rooms
AFTER UPDATE ON public.rooms
FOR EACH STATEMENT EXECUTE PROCEDURE notify_updated_rooms();
COMMENT ON TRIGGER notify_updated_rooms ON public.rooms IS 'trigger to perform notify with id';
```

Этот код создает временную таблицу `updated_ids` для хранения обновленных `ID`. Триггер `after_update_public_rooms` добавляет `ID` каждого обновленного ряда в эту таблицу. Триггер `notify_updated_rooms` вызывается после каждого запроса `UPDATE` и отправляет уведомление с `ID` всех обновленных рядов. Затем он очищает временную таблицу для следующего запроса `UPDATE`.

Обратите внимание, что этот код предназначен для демонстрации и может потребовать дополнительной настройки или оптимизации для вашего конкретного случая. Пожалуйста, тестируйте его в безопасной среде перед использованием в производственной среде.

