### 增删改以后请访问
~~https://dev.talentpool.co.jp/i18no/rebuild 进行重塑~~

移除rebuild重塑ID序列放在refresh里面

https://dev.talentpool.co.jp/i18no/refresh 进行缓存刷新

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

### 修改下拉列表翻译分类 (json全字段提交)
URL:https://dev.talentpool.co.jp/uitranslate/update/selectcategorylist

Method: POST

```
{
    "id":11,
    "classname":"CompanyType",
    "tag":"companyType"
}
```
RESPONSE
```
{
    "msg": "",
    "status": 0
}
```
#### 说明
status 0 正常 1 错误

### 根据下拉下来列表分类ID获取该分类下面的翻译内容列表
URL: http://dev.talentpool.co.jp/uitranslate/get/selectcategorylist/classid/:id

Method: Get

Example:

Get http://dev.talentpool.co.jp/uitranslate/get/selectcategorylist/classid/5

Response
```
{
    "data": [
        {
            "id": 1,
            "english": "Full time",
            "japanese": "正社員"
        },
        {
            "id": 2,
            "english": "Freelance",
            "japanese": "フリーランス"
        },
        {
            "id": 3,
            "english": "Contract",
            "japanese": "業務委託"
        },
        {
            "id": 4,
            "english": "Dispatched",
            "japanese": "派遣"
        }
    ],
    "msg": "",
    "status": 0
}
```
#### 说明
status 0 正常 1 错误

### 修改该下拉分类下面的某个翻译内容
URL: http://dev.talentpool.co.jp/uitranslate/update/selectcategorybyclassid

Method: POST
```
{
    "id":1,
    "english":"Full Time",
    "japanese":"正社員",
    "classid":5
}
```
Response
```
{
    "msg": "ok",
    "status": 0
}
```
#### 说明
classid 分类id

### 新增该下拉分类下面的某个翻译内容
URL: https://dev.talentpool.co.jp/uitranslate/add/selectcategorybyclassid

Method: POST
```
{
    "english":"Full Time",
    "japanese":"正社員",
    "classid":5
}
```
Response
```
{
    "msg": "ok",
    "status": 0
}
```
### 删除该下拉分类下面的某个翻译内容
URL: https://dev.talentpool.co.jp/uitranslate/delete/selectcategorybyclassid

Method: POST
```
{
    "id":8,
    "classid":5
}
```
Response
```
{
    "msg": "ok",
    "status": 0
}
```

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
classname 分类名称 这部分的classname禁止修改，仅可以修改tag 修改的时候请原样提交

tag 代表可以通过 http://dev.talentpool.co.jp/i18nUI/{classname}.json 访问到json格式翻译

tag 代表可以通过 http://dev.talentpool.co.jp/i18nUI/{classname}_{tag}.json 访问到json格式翻译

### 修改界面翻译分类 (json全字段提交)
URL:https://dev.talentpool.co.jp/uitranslate/get/uicategorylist

METHOD: POST

```
{
    "id":100,
    "classname":"s1",
    "tag":"home"
}
```
RESPONSE
```
{
    "msg": "",
    "status": 0
}
```
#### 说明
classname 请原样提交

status 0 正常 1 错误

### 根据界面翻译分类ID获取该分类下面的翻译内容列表
URL: http://dev.talentpool.co.jp/uitranslate/get/uicategorylist/classid/:id

Method: Get

Example:

Get http://dev.talentpool.co.jp/uitranslate/get/uicategorylist/classid/100

Response
```
{
    "data": [
        {
            "id": 376,
            "transkey": "jobseeker",
            "english": "Jobseeker",
            "japanese": "求職者",
            "classid": 100
        },
        {
            "id": 377,
            "transkey": "email",
            "english": "Email",
            "japanese": "Email",
            "classid": 100
        },
        {
            "id": 378,
            "transkey": "enteremail",
            "english": "Enter email",
            "japanese": "Enter email",
            "classid": 100
        },
        {
            "id": 379,
            "transkey": "password",
            "english": "Password",
            "japanese": "Password",
            "classid": 100
        },
        {
            "id": 380,
            "transkey": "enterpassword",
            "english": "Enter password",
            "japanese": "Enter password",
            "classid": 100
        },
        {
            "id": 381,
            "transkey": "confirmpassword",
            "english": "Confirm Password",
            "japanese": "Confirm Password",
            "classid": 100
        },
        {
            "id": 382,
            "transkey": "enterpasswordagain",
            "english": "Enter password again",
            "japanese": "もう一度パスワードを入力",
            "classid": 100
        },
        {
            "id": 383,
            "transkey": "signup",
            "english": "Sign Up",
            "japanese": "Sign Up",
            "classid": 100
        },
        {
            "id": 384,
            "transkey": "alreadyhaveanaccount",
            "english": "Already have an account",
            "japanese": "Already have an account",
            "classid": 100
        },
        {
            "id": 385,
            "transkey": "signinhere",
            "english": "Sign in here",
            "japanese": "Sign in here",
            "classid": 100
        },
        {
            "id": 386,
            "transkey": "agree",
            "english": "I agree to the Terms of Service and Privacy Policy.",
            "japanese": "I agree to the Terms of Service and Privacy Policy",
            "classid": 100
        },
        {
            "id": 389,
            "transkey": "test",
            "english": "Test",
            "japanese": "Test",
            "classid": 100
        }
    ],
    "msg": "",
    "status": 0
}
```
#### 说明
```
// 这里的id的值为100的时候代表分类是s1,前面已经修改该tag的备注为"home
// 使用日语语言访问：https://dev.talentpool.co.jp/i18nUI/s1_home.json可以获得
{
    "home": {
        "agree": "I agree to the Terms of Service and Privacy Policy",
        "alreadyhaveanaccount": "Already have an account",
        "confirmpassword": "Confirm Password",
        "email": "Email",
        "enteremail": "Enter email",
        "enterpassword": "Enter password",
        "enterpasswordagain": "もう一度パスワードを入力",
        "jobseeker": "求職者",
        "password": "Password",
        "signinhere": "Sign in here",
        "signup": "Sign Up",
        "test": "Test"
    }
}
```
#### 说明
transkey: 这里的transkey 代表json翻译文件里面的指定翻译字段名，比如transkey为agree

status 0 正常 1 错误



### 修改某个界面翻译分类下的某个内容
URL: http://dev.talentpool.co.jp/uitranslate/update/uicategorybyclassid

Method: POST
```
{
    "id":381,
    "transkey":"confirmpassword",
    "english":"Confirm Password",
    "japanese":"パスワードを認証する",
    "classid":100
}
```
Response
```
{
    "msg": "ok",
    "status": 0
}
```

#### 说明
transkey: 这里的transkey 代表json翻译文件里面的指定翻译字段名，比如transkey为agree

status 0 正常 1 错误

### 新增某个界面翻译分类下的某个内容
URL: http://dev.talentpool.co.jp/uitranslate/add/uicategorybyclassid

Method: POST
```
{
    "transkey":"new",
    "english":"New",
    "japanese":"新しい",
    "classid":100
}
```
Response
```
{
    "msg": "ok",
    "status": 0
}
```
#### 说明
transkey: 这里的transkey 代表json翻译文件里面的指定翻译字段名，比如transkey为agree

status 0 正常 1 错误

### 删除某个界面翻译分类下的某个内容
URL: http://dev.talentpool.co.jp/uitranslate/delete/uicategorybyclassid

Method: POST
```
{
    "id":390,
    "classid":100
}
```
Response
```
{
    "msg": "ok",
    "status": 0
}
```

### 获取职位分类 JobCategory 一级列表
URL: http://dev.talentpool.co.jp/uitranslate/get/jobcategoryclass

Method: Get

Response
```
{
    "data": [
        {
            "id": 1,
            "english": "UI Design",
            "japanese": "デザイナー",
            "parentid": 0
        },
        {
            "id": 2,
            "english": "Product Management",
            "japanese": "プロダクトマネジャー",
            "parentid": 0
        },
        {
            "id": 3,
            "english": "Consulting",
            "japanese": "コンサル",
            "parentid": 0
        },
        {
            "id": 4,
            "english": "Front-end",
            "japanese": "フロントエンド",
            "parentid": 0
        },
        {
            "id": 5,
            "english": "Back-end",
            "japanese": "バックエンド",
            "parentid": 0
        },

        ...
    ],
    "msg": "",
    "status": 0
}
```

### 根据一级职位分类列表ID获取二级列表
URL: http://dev.talentpool.co.jp/uitranslate/get/jobcategorysubclass/:id

Method: Get

Example: GET http://dev.talentpool.co.jp/uitranslate/get/jobcategorysubclass/1

RESPONE
```
{
    "data": [
        {
            "id": 17,
            "english": "Figma",
            "japanese": "Figma",
            "parentid": 1
        },
        {
            "id": 18,
            "english": "Axure",
            "japanese": "Axure",
            "parentid": 1
        },
        {
            "id": 19,
            "english": "Adobe Photoshop",
            "japanese": "Adobe Photoshop",
            "parentid": 1
        },
        {
            "id": 20,
            "english": "Adobe Illustrator",
            "japanese": "Adobe Illustrator",
            "parentid": 1
        },
        {
            "id": 21,
            "english": "Adobe XD",
            "japanese": "Adobe XD",
            "parentid": 1
        }
    ],
    "msg": "",
    "status": 0
}
```
#### 说明
parentid 上级分类ID

### 更新职位一级分类
URL: http://dev.talentpool.co.jp/uitranslate/update/jobcategoryclass

Method: POST
```
{
    "id": 1,
    "english": "UI Design",
    "japanese": "デザイナー"
}
```
RESPONSE
```
{
    "msg": "ok",
    "status": 0
}
```

### 添加一级分类
URL: http://dev.talentpool.co.jp/uitranslate/add/jobcategoryclass

Method: POST
```
{
    "english": "New",
    "japanese": "New"
}
```
RESPONSE
```
{
    "msg": "ok",
    "status": 0
}
```
### 删除一级分类 必须保证该分类下没有二级分类

URL: http://dev.talentpool.co.jp/uitranslate/delete/jobcategoryclass

Method: POST

```
{
    "id":51
}
```
RESPONSE
```
{
    "msg": "ok",
    "status": 0
}
```
### 更新职位二级分类
URL: http://dev.talentpool.co.jp/uitranslate/update/jobcategorysubclass

Method: POST
```
{
    "id": 17,
    "english": "Figma",
    "japanese": "Figma",
    "parentid": 1
}
```
RESPONSE
```
{
    "msg": "ok",
    "status": 0
}
```
### 添加二级分类
URL: http://dev.talentpool.co.jp/uitranslate/add/jobcategorysubclass

Method: POST
```
{
    "english": "New",
    "japanese": "New",
    "parentid":1
}
```
RESPONSE
```
{
    "msg": "ok",
    "status": 0
}
```
### 删除二级分类

URL: http://dev.talentpool.co.jp/uitranslate/delete/jobcategorysubclass

Method: POST

```
{
    "id":52
}
```
RESPONSE
```
{
    "msg": "ok",
    "status": 0
}
```

### 翻译 —— 英文转日文
URL: http://dev.talentpool.co.jp/uitranslate/en2ja

Method: POST
```
{
    "text":"hello world"
}
```
RESPONSE
```
{
    "data": {
        "translations": [
            {
                "detected_source_language": "EN",
                "text": "ハローワールド"
            }
        ]
    },
    "msg": "ok",
    "status": 0
}
```
#### 说明
text 翻译结果

### 翻译 —— 日文转英文
URL: http://dev.talentpool.co.jp/uitranslate/ja2en

Method: POST
```
{
    "text":"ハローワールド"
}
```
RESPONSE
```
{
    "data": {
        "translations": [
            {
                "detected_source_language": "JA",
                "text": "Hello World"
            }
        ]
    },
    "msg": "ok",
    "status": 0
}
```
#### 说明
text 翻译结果
