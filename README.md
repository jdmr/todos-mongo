# TODOS

A list of things that need to be done.

Create with `curl`:

```bash
curl -d '{"name": "walk the dog", "description": "walk the dog", "priority": {"id": "3", "name": "High"}, "status": {"id": "1", "name": "Not Started"}, "user": {"id": "test", "name": "Test User"}}' -H "Content-Type: application/json" -X POST http://localhost:8080/todos
```

List with `curl`:

```bash
curl http://localhost:8080/todos
```
