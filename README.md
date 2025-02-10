# 居民身份证校验

## 接口
URL:
```
/id?id_number=11010219840406970X
```
其中11010219840406970X为[维基百科](https://zh.wikipedia.org/wiki/%E5%85%AC%E6%B0%91%E8%BA%AB%E4%BB%BD%E5%8F%B7%E7%A0%81)上的样例身份证
返回值:
```json
{
    "code": 10000,
    "data": {
        "age": 40,
        "birthday": "1984-04-06",
        "gender": "女",
        "is_adult": true,
        "zodiac": "鼠",
        "constellation": "白羊座",
        "address": {
            "province": "北京市",
            "city": "市辖区",
            "area": "西城区"
        }
    },
    "msg": "success"
}
```
身份证不合法时的返回:
```json
{
    "code": 10001,
    "data": null,
    "msg": "非法的身份证号"
}
```

## 数据源
数据源来自 https://github.com/modood/Administrative-divisions-of-China
