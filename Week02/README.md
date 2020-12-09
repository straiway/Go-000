## 作业题目
我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么，应该怎么做请写出代码？

## 代码实现
```go
// dao层
package dao

import (
    "github.com/pkg/errors"
    "sql"
)

var ErrNotFound = errors.New("data not found")

func GetUserByID(userID uint64) (*User, error) {

    ...

    var user User
    err := stmt.QueryRow().Scan(&user.ID, &user.name)
    if err == sql.ErrNoRows {
        return nil, errors.Wrapf(ErrNotFound, "user id not found: %d", userID)
    }
    if err != nil {
        return nil, errors.Wrap(err, "get user error")
    }
    
    return user, nil
}
```

```go
// service层
package service

import (
	"dao"
    "github.com/pkg/errors"
)

func GetUserDetail(userID uint64) (*UserDetail, error) {
    user, err := dao.GetUserByID(userID)
    if err != nil {
        return nil, errors.WithMessagef(err, "get user detail by id error:%d", userID)
    }
    
    // 其他用户信息获取
    ...
    
}
```

```go
// api层
package api

import (
    "service"
)

func GetUserProfile(userID uint64) {
    err := service.GetUserDetail(userID)
    if err != nil {
        // 统一记录错误日志
        log.Error(err)
        return
    }
    
    ...

}
```

## 参考答案
```go
// dao层
 return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))

```

```go
// biz层
if errors.Is(err, code.NotFound} {

}
```