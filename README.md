# Task Tracker CLI

## 簡介
此專案源自於 roadmap [Task Tracker](https://roadmap.sh/projects/task-tracker) project

## 安裝
需要先安裝 [Go](https://go.dev/)


```bash
git clone https://github.com/y3933y3933/task-tracker-cli.git
cd task-tracker-cli
```

## 啟動

```bash
go run .
```

## 功能

### 新增
`task-cli add <description>`

### 更新
`task-cli update <id> <description>`

### 刪除
`task-cli delete <id>`

### 顯示
```
task-cli list
task-cli list done
task-cli list todo
task-cli list in-progress
```

### 更改狀態
```
task-cli mark-in-progress
task-cil mark-done
```