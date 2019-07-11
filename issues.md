1. provider.go do not have comments
2. client.go tries to return single entity instead of slice
3. if response contains just string generated client is not compiled
```yaml
      responses:
        '200':
          description: password has been changed successful
          content:
            application/json:
              schema:
                type: string
```
4. 