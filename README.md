### 获取全部下拉列表翻译分类
URL：https://dev.talentpool.co.jp/uitranslate/get/selectcategorylist

Method：Get

Response:
```
{
    "data": [
        {
            "id": 5,
            "classname": "JobType",
            "tag": "jobtype"
        },
        {
            "id": 6,
            "classname": "Languages",
            "tag": "languages"
        },
        {
            "id": 7,
            "classname": "WorkStyle",
            "tag": "workStyle"
        },
        {
            "id": 8,
            "classname": "Education",
            "tag": "education"
        },
        {
            "id": 9,
            "classname": "Salary",
            "tag": "salary"
        },
        {
            "id": 10,
            "classname": "Industry",
            "tag": "industry"
        },
        {
            "id": 11,
            "classname": "Industry",
            "tag": "companyType"
        },
        {
            "id": 12,
            "classname": "Country",
            "tag": "country"
        },
        {
            "id": 13,
            "classname": "JapaneseCity",
            "tag": "japanList"
        }
    ],
    "msg": "",
    "status": 0
}
```
#### 说明
classname 分类名称

tag 代表可以通过 https://dev.talentpool.co.jp/i18n/{tag}.json 访问到json格式翻译

status 0 正常 1 错误

### 获取界面翻译分类 目前显示10个预留了几十个
URL：https://dev.talentpool.co.jp/uitranslate/get/uicategorylist

Method：Get

Response:
```
{
    "data": [
        {
            "id": 100,
            "classname": "s1",
            "tag": ""
        },
        {
            "id": 101,
            "classname": "s2",
            "tag": ""
        },
        {
            "id": 102,
            "classname": "s3",
            "tag": ""
        },
        {
            "id": 103,
            "classname": "s4",
            "tag": ""
        },
        {
            "id": 104,
            "classname": "s5",
            "tag": ""
        },
        {
            "id": 105,
            "classname": "s6",
            "tag": ""
        },
        {
            "id": 106,
            "classname": "s7",
            "tag": ""
        },
        {
            "id": 107,
            "classname": "s8",
            "tag": ""
        },
        {
            "id": 108,
            "classname": "s9",
            "tag": ""
        },
        {
            "id": 109,
            "classname": "s10",
            "tag": ""
        }
    ],
    "msg": "",
    "status": 0
}
```
#### 说明
classname 分类名称 这部分的classname禁止修改，仅可以修改tag 修改的时候不要传classname

tag 代表可以通过 http://dev.talentpool.co.jp/i18nUI/{classname}.json 访问到json格式翻译

tag 代表可以通过 http://dev.talentpool.co.jp/i18nUI/{classname}_{tag}.json 访问到json格式翻译

status 0 正常 1 错误
