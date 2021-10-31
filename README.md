# auditlog-sample-gorm-echo

Sample implementation for audit logs with echo and gorm.

You can implement audit logs which have an application user name using like the following codes.

- Add user identifier to echo context in Auth middleware.

main.go

```
		ctx := context.WithValue(c.Request().Context(), "UserName", username)
		r := c.Request().WithContext(ctx)
		c.SetRequest(r)

		return next(c)
```

- Use custom logger with gorm

main.go

```
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: auditlog.Default.LogMode(logger.Info), // for Audit logs
	})
```

- Make custom logger with the context

auditlog/auditlog.go

```
	pri := fmt.Sprintf(logPrefix, ctx.Value("UserName").(string))
```

- Use db with context

main.go

```
	ctxdb := db.WithContext(c.Request().Context())
```
