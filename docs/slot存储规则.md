## Slot的存储规则

除了64位参数(double,long)外，其他所有类型都只占用1个slot

### long的slot中存储规则

```go
slots[index].Num = int32(val)  //索引n存储低32位
slots[index+1].Num = int32(val >> 32) //索引n+1存储高32位

slots[index].Type = utils.SlotLong
slots[index+1].Type = utils.SlotLong
```

### long的slot中存储规则

```go 
bits := math.Float64bits(val)  //将值变成bit类型
slots[index].Num = int32(val)  //索引n存储低32位
slots[index+1].Num = int32(val >> 32) //索引n+1存储高32位

l.slots[index].Type = utils.SlotDouble
l.slots[index+1].Type = utils.SlotDouble
```