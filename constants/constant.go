package constants

//const RFC3339     = "2006-01-02T15:04:05Z07:00" //系统常量
// const TimeFormat = time.RFC3339
const TimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"

// 订单变更交换机名（事件）及队列名
const EventOrderChange string = "OrderChangeEx"
const QueueOrderChange string = "OrderChangeQueue"

// 包裹
const QueuePackageChange string = "PackageChangeEx"
const EventPackageChange string = "PackageChangeQueue"

// 库存
const EventStockChange string = "StockChangeEx"
const QueueStockChange string = "StockChangeQueue"
